package DB

import "github.com/jinzhu/gorm"

var db gorm.DB

type GormController struct {
	Txn *gorm.DB
}

func InitDB() {
	//var(
	//	driver, url string
	//	found bool
	//)
	//
	//if driver, found = echo.
	//설정파일 구성할 방법 찾아라
}