package model

import (
	"github.com/CloudSilk/pkg/db"
	"github.com/CloudSilk/pkg/db/mysql"
	"github.com/CloudSilk/pkg/db/sqlite"
)

var dbClient db.DBClientInterface

// Init Init
func Init(connStr string, debug bool) {
	dbClient = mysql.NewMysql(connStr, debug)
	if debug {
		AutoMigrate()
	}
}

func InitSqlite(database string, debug bool) {
	dbClient = sqlite.NewSqlite2("", "", database, "", debug)
	if debug {
		AutoMigrate()
	}
}

func InitDB(client db.DBClientInterface, debug bool) {
	dbClient = client
	if debug {
		AutoMigrate()
	}
}

// AutoMigrate 自动生成表
func AutoMigrate() {
	dbClient.DB().AutoMigrate(&Metadata{}, &MetadataField{}, &Page{}, &PageToolBar{}, &PageField{}, &PageButton{}, &Template{},
		&Service{}, &CodeFile{}, &ServiceFunctional{}, &Cell{}, &CellMarkup{}, &CellAttrs{}, &CellConnecting{}, &Form{}, &FormVersion{}, &FileTemplate{},
		&FunctionalTemplate{}, &SystemObject{})
}
