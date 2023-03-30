package component

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tiyee/gokit/pkg/consts"
	"time"
)

var RDb *sql.DB
var WDb *sql.DB

const MaxIdle int = 20
const MaxOpen int = 20
const RDsn = consts.RDsn
const WDsn = consts.WDsn

func InitMysql() error {
	if db, err := sql.Open("mysql", RDsn); err == nil {
		db.SetConnMaxLifetime(time.Second * 20)
		db.SetMaxIdleConns(MaxIdle)
		db.SetMaxOpenConns(MaxOpen)
		if err := db.Ping(); err == nil {
			RDb = db
		} else {
			return err
		}
	} else {
		return err
	}
	if db, err := sql.Open("mysql", WDsn); err == nil {
		db.SetConnMaxLifetime(time.Second * 20)
		db.SetMaxIdleConns(MaxIdle)
		db.SetMaxOpenConns(MaxOpen)
		if err := db.Ping(); err == nil {
			WDb = db
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}
