package database

import (
	"fmt"
	"log"

	"github.com/seregant/decision-making-go/config"

	"github.com/jinzhu/gorm"
)

func DbConnect() *gorm.DB {
	var conf = config.SetConfig()
	var addr = conf.User + ":" + conf.Pass + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.DbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", addr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type Pertanyaan struct {
	ID         int    `gorm:"primary_key:yes;column:pertanyaan_id;auto_increment:yes;type:int(5)"`
	Pertanyaan string `gorm:"column:pertanyaan_string;type:varchar(300)"`
}

func (Pertanyaan) TableName() string {
	return "pertanyaan"
}

type Pengetahuan struct {
	ID   int    `gorm:"primary_key:yes;column:pengetahuan_id;auto_increment:yes;type:int(5)"`
	Nama string `gorm:"column:pengetahuan_nama;type:char(140)"`
}

func (Pengetahuan) TableName() string {
	return "pengetahuan"
}

type Decisions struct {
	ID_Decision int    `gorm:"type:int(5);column:decision_id;auto_increment:yes;primary_key:yes"`
	Decision    string `gorm:"type:char(100);colum:decision_name"`
}

func (Decisions) TableName() string {
	return "decisions"
}

type Pengetahuan_value struct {
	ID_Pertanyaan  int `gorm:"type:int(5);column:pertanyaan_id"`
	ID_Pengetahuan int `gorm:"type:int(5);column:pengetahuan_id"`
	ID_Decision    int `gorm:"type:int(5);column:decision_id"`
	Yes            int `gorm:"type:int(5);column:is_yes"`
}

func (Pengetahuan_value) TableName() string {
	return "pengetahuan_value"
}

func CreateTable() {
	var db = DbConnect()
	defer db.Close()

	fmt.Println("Creating tables...")

	db.AutoMigrate(&Pertanyaan{})
	db.AutoMigrate(&Pengetahuan{})
	db.AutoMigrate(&Decisions{})
	db.AutoMigrate(&Pengetahuan_value{})
}
