package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	"github.com/justinas/nosurf"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"main.go/handler"
	"main.go/storage/postgres"
)

//go:embed assets
var assetFiles embed.FS

//go:embed migrations
var migrationFiles embed.FS

func main() {

	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	
	p := config.GetInt("server.port")
	fmt.Println("Running on port : ", p)

	decoder := form.NewDecoder()

	postGresStore,err := postgres.NewPostgresStorage(config)
	if err != nil {
		log.Fatalln(err)
	}
	goose.SetBaseFS(migrationFiles)
	if err := goose.SetDialect("postgres"); err != nil {
        log.Fatalln(err)
    }
	
    if err := goose.Up(postGresStore.DB.DB, "migrations"); err != nil {
        log.Fatalln(err)
    }

	lifeTime := config.GetDuration("session.lifeTime")
	idleTime := config.GetDuration("session.idleTime")
	sessionManager := scs.New()
	sessionManager.Lifetime = lifeTime * time.Hour
	sessionManager.IdleTimeout = idleTime * time.Minute
	sessionManager.Cookie.Name = "web-session"
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Secure = true
	sessionManager.Store = NewSQLXStore(postGresStore.DB)

	var assetFS = fs.FS(assetFiles)
	staticFiles, err1 := fs.Sub(assetFS, "assets/templates/images")
	if err1 != nil {
		log.Fatal(err)
	}

	templateFiles, err := fs.Sub(assetFS, "assets/templates")
	if err != nil {
		log.Fatal(err)
	}
	
	chi := handler.NewHandler(sessionManager,decoder,postGresStore,staticFiles,templateFiles)
	newChi := nosurf.New(chi)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), newChi); err != nil {
		log.Fatalln(err)
	}
}
