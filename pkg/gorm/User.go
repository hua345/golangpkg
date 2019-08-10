package gorm

import "time"

//Column	指定列名
//Type	指定列数据类型
//PRIMARY_KEY	将列指定为主键
//UNIQUE	将列指定为唯一
//DEFAULT	指定列默认值
//PRECISION	指定列精度
//NOT NULL	将列指定为非 NULL
//AUTO_INCREMENT	指定列是否为自增类型
//INDEX	创建具有或不带名称的索引, 如果多个索引同名则创建复合索引
//UNIQUE_INDEX	唯一索引
//-	忽略此字段
type User struct {
	// GORM 默认会使用名为ID的字段作为表的主键。
	ID          int64  `gorm:"type:bigint;PRIMARY_KEY;not null"`
	Name        string `gorm:"type:varchar(128)"`
	UserMobile  string `gorm:"type:varchar(16);INDEX;not null"`
	Age         int
	CreatedTime time.Time
	CreatedBy   int64
	UpdatedTime time.Time
	UpdatedBy   int64
}

func (User) TableName() string {
	return "myUser"
}
