package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mysql struct {
	UserName      string `json:"username"`
	Password      string `json:"password"`
	Database      string `json:"database"`
	Address       string `json:"address"`
	Parameters    string `json:"parameters"`
	MaxIdle       int    `json:"maxidle"`
	MaxOpen       int    `json:"maxopen"`
	Debug         bool   `json:"debug"`
	MigrationsDir string `json:"migrationsdir"`
}

func (m *Mysql) String() string {
	return fmt.Sprintf("%s:%s@%s/%s?%s", m.UserName, m.Password, m.Address, m.Database, m.Parameters)
}

func (m *Mysql) New() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", m.String())
	if err != nil {
		return nil, err
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	db.DB().SetConnMaxLifetime(-1) //make connection reuseable forever
	db.DB().SetMaxIdleConns(m.MaxIdle)
	db.DB().SetMaxOpenConns(m.MaxOpen)
	db.LogMode(m.Debug)
	return db, nil
}
