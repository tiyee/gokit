package uploader

import (
	"github.com/tiyee/gokit/pkg/component/redis"
	"github.com/tiyee/gokit/pkg/controllers/uploader"
	"github.com/tiyee/gokit/pkg/engine"
	"github.com/tiyee/gokit/pkg/repository/pic_repo"
	"io"
	"net/http"
	"strconv"
	"time"
)

type RetResult struct {
	UploadID string                `json:"upload_id"`
	Status   int8                  `json:"status"`
	Url      string                `json:"url"`
	Chunks   []*uploader.ChunkETag `json:"chunks"`
}

func Init(c *engine.Context) {
	size := c.QueryInt("size", 0)
	chunkSize := c.QueryInt("chunk_size", 0)
	digest := c.Query("digest")

	if size*chunkSize < 1 || len(digest) != 32 {

		c.AjaxError(122, "数据不合法", digest)
		return
	}
	chunk, err := io.ReadAll(c.Request().Body)
	if err != nil || len(chunk) < 1 {
		c.AjaxError(123, "数据不合法", len(chunk))
		return
	}
	if row, err := pic_repo.GetUrlBySha(digest); err == nil {
		if row.Size == size {
			c.AjaxSuccess("ok", RetResult{
				Status: uploader.FulfilledStatus,
				Url:    row.Url,
				Chunks: []*uploader.ChunkETag{},
			})
			return
		}

	}
	mime := http.DetectContentType(chunk)
	ext := uploader.Ext(mime)
	if ext == "" {
		c.AjaxError(127, "文件类型不合法", size)
		return
	}
	//fmt.Println(string(digest), size, chunkSize, http.DetectContentType(chunk))
	id := uint64(time.Now().Unix() - 1599613749)
	objectName := strconv.FormatInt(int64(id), 10) + "." + ext
	var dir string
	if size < 5*1024*1024 {
		dir = "u"
	} else if size > 10*1024*1024 {
		dir = "b"
	} else {
		dir = "m"
	}
	fn := func(m *uploader.Meta) {
		m.ChunkSize = int(chunkSize)
		m.Object = dir + "/" + objectName
		m.Digest = digest
		m.Size = size
		m.ContentType = mime
	}
	upload, err := uploader.FromDigest(c.Ctx(), uploader.NewRedisClient(c.Ctx(), redis.RedisClient), uploader.OSS, digest, fn)
	//upload, err := uploader.NewMeta(c, uploader.OSS, fn)
	if err != nil {
		c.AjaxError(125, err.Error(), digest)
		return
	}
	status := upload.Touch()
	if status == uploader.FulfilledStatus {

		c.AjaxSuccess("ok", RetResult{
			Status: upload.Status,
			Url:    upload.Url(),
			Chunks: []*uploader.ChunkETag{},
		})
		return
	}
	if status == uploader.PendingStatus {
		c.AjaxSuccess("ok", RetResult{
			UploadID: upload.UploadID,
			Status:   upload.Status,
			Url:      "",
			Chunks:   upload.Pending(),
		})
		return
	}
	err = upload.Init()
	if err != nil {
		c.AjaxError(126, err.Error(), 1)
		return
	}
	c.AjaxSuccess("ok", RetResult{
		UploadID: upload.UploadID,
		Status:   upload.Status,
		Url:      "",
		Chunks:   upload.Pending(),
	})
	return
}
