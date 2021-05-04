package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"test/Model"
	"test/config"
)

var db *gorm.DB

type GormController struct {
	Txn *gorm.DB
}

func InitDB() {
	var (
		driver, url string
	)

	//Configuration
	configuration := config.GetConfiguration()
	driver = configuration.DBtype
	url = configuration.DBUrl

	var dberr error
	db, dberr = gorm.Open(driver, url)
	if dberr != nil {
		log.Fatal(dberr)
		log.Fatalf(url, driver)
	}

	//db.LogMode(true)
	migrate()
}

//테이블 생성
func migrate () {
	db.AutoMigrate(&Model.Push{})
}

//Begin a transaction
func (c GormController) Begin() {
	c.Txn = db.Begin()
}

//Commit the transaction
func (c *GormController) Commit () {
	if c.Txn != nil {
		c.Txn.Commit()
		c.Txn = nil
	}
}

//DBManager return gorm (Connection Done)
func DBManager() *gorm.DB {
	return db
}