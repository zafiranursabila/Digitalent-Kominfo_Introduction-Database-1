package main

import (
	"database/sql"
	"fmt"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-Database-1/sql-generic/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

func main(){
	cfg,err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db,err := connect(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

//database.InsertCustomer(database.Customer{
	//	FirstName:"Zafira",
	//	LastName:"Nur Sabila",
	//	NpwpId:"id-16",
	//	Age:20,
	//	CustomerType:"Premium",
	//	Street:"Str",
	//	City:"Bekasi",
	//	State:"Indo",
	//	ZipCode:"55555",
	//	PhoneNumber:"0812345",
	//},db)

	//database.GetCustomers(db)
	database.DeleteCustomer(16,db)
	database.UpdateCustomer(30,16,db)
}


func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}


func connect(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",cfg.User,cfg.Password,cfg.Host,cfg.Port,cfg.DbName,cfg.Config))
	if err != nil {
		return nil, err
	}

	log.Println("db successfully connected")
	return db, nil
}