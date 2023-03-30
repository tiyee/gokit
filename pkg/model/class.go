package model

import (
	"github.com/tiyee/gokit/pkg/component/idgen"
	"github.com/tiyee/gokit/pkg/helps"
	"github.com/tiyee/gokit/pkg/vo"
	"time"
)

type Class struct {
	Id        int64  `json:"id"`         //id
	Name      string `json:"name"`       //标题
	Teachers  string `json:"teachers"`   //老师uid，多个逗号隔开
	Catalog   int8   `json:"catalog"`    //课程分类
	Tag       int8   `json:"tag"`        //1:公开课,2:专题课,3:系统课,4:答疑,5:一对一
	Creator   int64  `json:"creator"`    //创建者ID
	Thumb     string `json:"thumb"`      //缩略图
	Summary   string `json:"summary"`    //简介
	StartTime int64  `json:"start_time"` //课程起始时间
	EndTime   int64  `json:"end_time"`   //课程结束时间
	Quota     int    `json:"quota"`      //限额
	Volume    int    `json:"volume"`     //已售数目
	LessonNum int    `json:"lesson_num"` //课时数
	Introduce string `json:"introduce"`  //课程介绍图片地址
	State     int8   `json:"state"`      //0:待上线，1:上线，99：下线
	CTime     int64  `json:"ctime"`      //创建时间
	MTime     int64  `json:"mtime"`      //修改时间
	Price     int    `json:"price"`
}

func (c *Class) TableName() string {
	return "t_class"
}
func (c *Class) Pk() string {
	return "id"
}
func (c *Class) Format() *vo.Class {
	return &vo.Class{
		ClassID:      idgen.FromId(c.Id).Encode(),
		ClassName:    c.Name,
		StartTime:    c.StartTime,
		EndTime:      c.EndTime,
		EndDate:      time.Unix(c.EndTime, 0).Format("2006-01-02"),
		TotalLessons: c.LessonNum,
		Duration:     helps.Duration2(c.StartTime, c.EndTime),
		MTime:        c.MTime,
		CTime:        c.CTime,
		Price:        c.Price,
		Tag:          c.Tag,
	}
}
