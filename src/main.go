package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	html "html/template"
	text "text/template"

	"github.com/fiorix/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

var (
	VERSION = "tip"
	APPNAME = "%name%"

	// Templates
	HTML *html.Template
	TEXT *text.Template
)

func main() {
	configFile := flag.String("c", "%name%.conf", "")
	flag.Usage = func() {
		fmt.Println("Usage: %name% [-c %name%.conf]")
		os.Exit(1)
	}
	flag.Parse()

	var err error
	config, err := loadConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Parse templates.
	HTML = html.Must(html.ParseGlob(config.TemplatesDir + "/*.html"))
	TEXT = text.Must(text.ParseGlob(config.TemplatesDir + "/*.txt"))

	// Set up databases.
	rc := redis.New(config.DB.Redis)
	db, err := sql.Open("mysql", config.DB.MySQL)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s %s \n", APPNAME, VERSION)

	// Start HTTP server.
	s := new(httpServer)
	s.init(config, rc, db)
	go s.ListenAndServe()
	go s.ListenAndServeTLS()

	// Sleep forever.
	select {}
}

func openLog(filename string) *os.File {
	f, err := os.OpenFile(
		filename,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Fatal(err)
	}
	return f
}
