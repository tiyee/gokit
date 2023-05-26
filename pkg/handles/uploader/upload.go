package uploader

import (
	"github.com/tiyee/gokit/pkg/component/redis"
	"github.com/tiyee/gokit/pkg/controllers/uploader"
	"github.com/tiyee/gokit/pkg/engine"
	"io"
)

type uploadRet struct {
	ETag     string `json:"etag"`
	Index    int    `json:"index"`
	UploadID string `json:"upload_id"`
}

func Upload(c *engine.Context) {
	index := c.QueryInt("index", -1)
	uploadID := c.Query("upload_id")

	if index < 0 || len(uploadID) < 1 {
		c.AjaxError(128, "数据不合法", uploadID)
		return
	}
	chunk, err := io.ReadAll(c.Request().Body)
	if err != nil || len(chunk) < 1 {
		c.AjaxError(123, "数据不合法", len(chunk))
		return
	}
	upload, err := uploader.FromUploadID(c.Ctx(), uploader.NewRedisClient(c.Ctx(), redis.RedisClient), uploadID)
	if err != nil {
		c.AjaxError(123, err.Error(), index)
		return
	}
	chunkPart, err := upload.UploadPart(int(index), chunk)
	if err != nil {
		c.AjaxError(123, err.Error(), index)
		return
	}
	data := uploadRet{
		ETag:     chunkPart.ETag,
		Index:    int(index),
		UploadID: uploadID,
	}
	c.AjaxSuccess("ok", data)
	return

}
