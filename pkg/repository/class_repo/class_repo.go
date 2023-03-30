package class_repo

import (
	"github.com/tiyee/gokit/pkg/component/orm"
	"github.com/tiyee/gokit/pkg/model"
)

func Class(classID int64) (*model.Class, error) {
	return orm.NewORM(&model.Class{}).Row("id=?", []interface{}{classID})
}
func ClassLimit(where string, values []interface{}, offset, page int64) ([]*model.Class, error) {
	return orm.NewORM(&model.Class{}).Limit(where, values, offset, page, "id desc")
}
func ClassCount(where string, values []interface{}) (int64, error) {
	return orm.NewORM(&model.Class{}).Count(where, values)
}
func ClassUpdateByPk(cls *model.Class, data map[string]interface{}, pk int64) (int64, error) {
	return orm.NewORM(cls).UpdateByPk(data, pk)
}
func ClassSave(cls *model.Class) (int64, error) {
	return orm.NewORM(cls).Save()
}
func ClassUpdate(cls *model.Class) (int64, error) {
	return orm.NewORM(cls).Update(cls.Id)
}
