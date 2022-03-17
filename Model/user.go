package Model

import (
	uuid "github.com/satori/go.uuid"
	"go-study/Library/Gorm"
	"go-study/Library/MyTime"
	"gorm.io/gorm"
)

//表名
var tableName = "user"

//
// User
// @Description: 表字段结构体
//
type User struct {
	ID        uint64       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Uuid      uuid.UUID    `gorm:"column:uuid;not null;default:'';uniqueIndex:user_uuid_uindex;comment:'全局唯一标识'" json:"uuid"`
	Name      string       `gorm:"column:name;;not null;default:'';comment:'用户名'" json:"name"`
	CreatedAt *MyTime.Time `gorm:"column:created_at;index:user_created_at_index;<-:create" json:"created_at"`
	UpdatedAt *MyTime.Time `gorm:"column:updated_at;index:user_updated_at_index;<-:update" json:"updated_at"`
	DeletedAt *MyTime.Time `gorm:"column:deleted_at;" json:"-"`
}

// UserModel
// @Description: 初始化model
// @return *gorm.DB
//
func UserModel() *gorm.DB {
	return Gorm.Mysql.Table(tableName)
}
