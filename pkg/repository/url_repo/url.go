package url_repo

import (
	"github.com/tiyee/gokit/pkg/component/orm"
	"github.com/tiyee/gokit/pkg/models"
)

func GetUrl(id int) (*models.Url, error) {
	return orm.New(&models.Url{}).Row("id = ?", []any{id})
}
func GetUrlBySha(sha string) (*models.Url, error) {
	return orm.New(&models.Url{}).Row("sha = ?", []any{sha})
}
func Save(row *models.Url) (int64, error) {
	return orm.New(row).Save()
}
