package pic_repo

import (
	"github.com/tiyee/gokit/pkg/component/orm"
	"github.com/tiyee/gokit/pkg/models"
)

func GetPic(id int) (*models.Pic, error) {
	return orm.New(&models.Pic{}).Row("id = ?", []any{id})
}
func GetUrlBySha(sha string) (*models.Pic, error) {
	return orm.New(&models.Pic{}).Row("sha = ?", []any{sha})
}
func Save(row *models.Pic) (int64, error) {
	return orm.New(row).Save()
}
