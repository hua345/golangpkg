package gorm

type Book struct {
	Id       int64  `gorm:"type:bigint;PRIMARY_KEY;not null"`
	BookName string `gorm:"type:varchar(128)"`
	Price    string `json:"price"`
	BookDesc string `gorm:"type:varchar(128)"`
}

func (Book) TableName() string {
	return "book"
}
