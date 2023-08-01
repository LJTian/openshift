package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SendDb(data []byte) (err error) {

	fmt.Printf("DBSend data is [%s]\n", data)

	var log Logs

	if err = json.Unmarshal(data, &log); err != nil {
		fmt.Println(err)
		return
	}

	db := db.Create(&log)
	if db.Error == nil {
		fmt.Println(db.Error)
		return errors.New("插入数据库失败")
	}

	return
}

func connect(dsn string) (err error) {

	fmt.Printf("进行数据库链接，链接地址[%s]\n", dsn)
	//dsn := "root:123456@tcp(81.70.17.60:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("链接数据库失败:[%s]\n", err)
		return
	}
	fmt.Printf("链接数据库成功~\n")
	return
}

func StartDB(dsn string) (err error) {

	if err = connect(dsn); err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&Logs{})
	return
}
