package mysql

type Book struct {
	Id       int64  `json:"id"`
	BookName string `json:"bookName"`
	Price    string `json:"price"`
	BookDesc string `json:"bookDesc"`
}
