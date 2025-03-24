package main

import (
	"fmt"
	"log"

	"github.com/thinkgos/sqlcipher-gorm"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlcipher.Open("gorm.db?_pragma_key=2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99&_pragma_cipher_page_size=4096"))
	if err != nil {
		panic(err)
	}

	err = db.Exec("create table foo (id integer not null primary key, name text)").Error
	if err != nil {
		log.Printf("%q\n", err)
		return
	}

	for i := 0; i < 100; i++ {
		err = db.Exec("insert into foo(id, name) values(?, ?)", i, fmt.Sprintf("こんにちわ世界%03d", i)).Error
		if err != nil {
			log.Fatal(err)
		}
	}
	type foo struct {
		Id   int
		Name string
	}

	var rows []foo

	err = db.Raw("select id, name from foo").Find(&rows).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range rows {
		log.Println(v)
	}
	db.Exec("drop table foo")
}
