package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/madasatya6/go-native/applications/config"
)

var DB Databases

var MySQLDsn = fmt.Sprintf("%v:%v@/%v", "root", "", "go-ecommerce")
var MySQLDockerDsn = fmt.Sprintf("%v:%v@tcp(%s)/%v", "root", "lampupijar77", "mysql-network", "go-ecommerce") //docker
var postgreeDsn = "dbname=ecommerce user=postgres password=lampupijar77 host=localhost sslmode=disable"

type Methods interface{
	SetMysql(dsn string)
	SetPostgre(dsn string)
	TestPing(dbNames []string)
}

type Databases struct{
	MySQL *sql.DB
	Postgre *sql.DB
}

func (d *Databases) SetMysql(dsn string) {
	db, _ := MySQL(dsn)
	d.MySQL = db 
}

func (d *Databases) SetPostgre(dsn string) {
	db, _ := Postgres(dsn)
	d.Postgre = db 
}

func (d *Databases) TestPing(dbNames []string) {
	var err error
	for i:=0; i < len(dbNames); i++ {
		if dbNames[i] == "mysql" {
			err = d.MySQL.Ping()
			if err != nil {
				log.Fatal(err.Error())
			}
			//set global variable
			config.MySQL = d.MySQL

		} else if dbNames[i] == "postgree" {
			err = d.Postgre.Ping()
			if err != nil {
				log.Fatal(err.Error())
			}
			//set global variable
			config.Postgree = d.Postgre
		}
	}

	if err == nil {
		log.Println("Successfully connected with database")
	}
}

func Init(data map[string]interface{}) *Databases {
	var method Methods
	method = &DB 
	method.SetMysql(MySQLDsn)
	method.SetPostgre(postgreeDsn)
	method.TestPing([]string{
		"mysql",
	})
	return &DB
}

func MySQL(config string) (*sql.DB, error) {
	db, err := sql.Open("mysql", config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Postgres(config string) (*sql.DB, error) {
	db, err := sql.Open("postgres", config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

