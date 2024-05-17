package mysql

import (
	"fmt"
	"log"
	"time"
	"todo-list/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Task struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

var DB *gorm.DB

func NewClient() *gorm.DB {
	mysqlUser := config.MainConfig.MySQL.User
	mysqlPassword := config.MainConfig.MySQL.Password
	mysqlDatabase := config.MainConfig.MySQL.DataBase
	dataSourceName := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlDatabase)

	// Из-за особенностей докера сделаем 5 попыток подключиться к mysql
	var err error
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open("mysql", dataSourceName)
		if err == nil {
			err = DB.DB().Ping()
			if err == nil {
				log.Println("Клиент Mysql успешно создан")
				DB.AutoMigrate(&Task{}) //Создадим таблицу Task
				return DB
			}
		}
		log.Println("Не удалось подключиться к базе данных, ждем 5 секунд и повторяем попытку...")
		time.Sleep(5 * time.Second)
	}
	log.Fatal("Не удалось подключиться к базе данных после 5 попыток: ", err)
	return nil
}
