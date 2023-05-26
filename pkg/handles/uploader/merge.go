package uploader

import (
	"encoding/json"
	"github.com/tiyee/gokit/pkg/component/redis"
	"github.com/tiyee/gokit/pkg/controllers/uploader"
	"github.com/tiyee/gokit/pkg/engine"
	"github.com/tiyee/gokit/pkg/models"
	"github.com/tiyee/gokit/pkg/repository/pic_repo"
	"time"
)

func Merge(c *engine.Context) {
	digest := c.Query("digest")
	uploadID := c.Query("upload_id")
	var partials []uploader.ChunkETag

	if err := json.NewDecoder(c.Request().Body).Decode(&partials); err != nil {
		c.AjaxError(129, err.Error(), uploadID)
		return
	}
	cache := uploader.NewRedisClient(c.Ctx(), redis.RedisClient)
	meta, err := uploader.FromUploadID(c.Ctx(), cache, uploadID)
	if err != nil {
		c.AjaxError(130, err.Error(), uploadID)
		return
	}
	if meta.Digest != string(digest) {
		c.AjaxError(129, "digest do not match", meta.Digest)
		return
	}
	err = meta.Merge(partials)
	if err != nil {
		c.AjaxError(131, err.Error(), uploadID)
		return
	}
	row := models.Pic{
		Id:     0,
		OpenId: "",
		Url:    meta.Url(),
		Size:   meta.Size,
		Sha:    digest,
		Ip:     c.Request().Header.Get("X-Forward-For"),
		Ua:     c.Request().UserAgent(),
		Ctime:  time.Now().Unix(),
	}
	if n, err := pic_repo.Save(&row); err == nil {
		row.Id = int64(n)
		data := &RetResult{
			UploadID: meta.UploadID,
			Status:   meta.Status,
			Url:      meta.Url(),
			Chunks:   meta.Pending(),
		}
		c.AjaxSuccess("ok", data)
	}

}
