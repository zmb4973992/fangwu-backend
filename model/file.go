package model

import (
	"errors"

	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/gorm"
)

type File struct {
	Base
	Name         string
	SizeKb       float64 `gorm:"index;"`
	BusinessType *string `gorm:"index;"`
	BusinessId   *int64  `gorm:"index;"`
	Sort         int     `gorm:"index;"`
}

// TableName 修改数据库的表名
func (f File) TableName() string {
	return "file"
}

func (b *File) BeforeCreate(tx *gorm.DB) (err error) {
	b.Id = idgen.NextId()
	if b.Id == 0 {
		return errors.New("生成id失败")
	}
	return nil
}
