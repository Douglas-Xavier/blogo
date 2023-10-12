package data

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Datasource interface {
	GetDB() *sql.DB
}

type ActualDBConn struct {
	db *sql.DB
}

var (
	host     string
	port     int
	user     string
	password string
	dbname   string
	instance Datasource
)

// ToDO: usar variaveis de ambiente ao inves de valores hardcoded
func init() {
	host = "0.0.0.0"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "blogo"
}

func CreateDataSource() *ActualDBConn {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}
	//esses valores seriam maiores em aplicações reais, coloquei valores baixos pois tenho preguiça
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Minute * 2)
	println("connect")
	return &ActualDBConn{db}
}

func (p *ActualDBConn) GetDB() *sql.DB {
	return p.db
}

func GetDBInstance() Datasource {
	if instance == nil {
		instance = CreateDataSource()
	}

	return instance
}
