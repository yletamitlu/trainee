package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	. "github.com/yletamitlu/trainee/internal/mwares"
	statDelivery "github.com/yletamitlu/trainee/internal/stat/delivery"
	statRepos "github.com/yletamitlu/trainee/internal/stat/repository"
	statUcase "github.com/yletamitlu/trainee/internal/stat/usecase"
	"log"
	"os"
)

func main() {
	conn, err := sqlx.Connect("pgx",
		"postgres://"+os.Getenv("DB_USER")+":statuser@localhost:5432/"+os.Getenv("DB_NAME"))
	if err != nil {
		logrus.Info(err)
	}

	conn.SetMaxOpenConns(8)
	conn.SetMaxIdleConns(8)

	if err := conn.Ping(); err != nil {
		logrus.Info(err)
	}
	defer conn.Close()

	router := fasthttprouter.New()

	statR := statRepos.NewStatRepos(conn)
	statU := statUcase.NewStatUcase(statR)
	statD := statDelivery.NewStatDelivery(statU)
	statD.Configure(router)

	fmt.Printf("Server started...")
	log.Fatal(fasthttp.ListenAndServe(":80", Use(router.Handler, PanicRecovering, SetHeaders, AccessLog)))
}
