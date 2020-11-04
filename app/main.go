package main

import (
	blacklistUsecase "bitbucket.org/fabribraguev/api-toolbox/app/application/usecase/blacklist"
	blacklistRepository "bitbucket.org/fabribraguev/api-toolbox/app/infrastructure/data/db/mysql/blacklist"
	v1 "bitbucket.org/fabribraguev/api-toolbox/app/infrastructure/web/rest/v1"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"os"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "MYSQL56_USER"
	mysql_users_password = "MYSQL56_PASSWORD"
	mysql_users_host     = "MYSQL56_HOST"
	mysql_users_port     = "MYSQL56_PORT"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host) + ":" + os.Getenv(mysql_users_port)
	schema   = "apolo"
)

func main() {

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORS())

	_ = v1.NewHealthCheck(e)

	client := createMysqlClient()
	blackListMysqlRepository := blacklistRepository.NewBlackListMysqlRepository(client)
	blackListUsecase := blacklistUsecase.NewBlackListUsecase(blackListMysqlRepository)

	_ = v1.NewBlackListHandler(e,blackListUsecase)

	e.Logger.Fatal(e.Start(":8080"))

}

func createMysqlClient() *sql.DB {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema)

	client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	//client.SetMaxOpenConns(10)

	if err = client.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("database successfully configured")

	return client
}

