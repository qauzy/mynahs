package model

import (
	"github.com/zzpu/mynahs/data"
	"github.com/zzpu/mynahs/internal/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	ID        uint             `gorm:"column:id;type:int(11) unsigned AUTO_INCREMENT;not null;primarykey" json:"id"`
	CreatedAt utils.FormatDate `gorm:"column:created_at;type:timestamp;<-:create" json:"created_at"`
	UpdatedAt utils.FormatDate `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

func (model *BaseModel) DB() *gorm.DB {
	return DB()
}

type ContainsDeleteBaseModel struct {
	BaseModel
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index" json:"-"`
}

func DB() *gorm.DB {
	return data.MysqlDB
}
