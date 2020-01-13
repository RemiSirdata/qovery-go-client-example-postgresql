package main

import (
	"flag"
	"fmt"
	"github.com/Qovery/qovery-go-client"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
)

var (
	configurationFilename = flag.String("config-filename", "../../test_files/local_configuration.json", "Qovery configuration filename")
	databaseName          = flag.String("dbname", "my-pql", "")
	bind                  = flag.String("bind", ":8080", "Http port")
)

func main() {
	flag.Parse()

	app := iris.New()
	app.Get("/databases", func(ctx iris.Context) {
		printDbStatus(ctx)
	})
	app.Run(iris.Addr(*bind))
}

func printDbStatus(ctx iris.Context) {
	qv, err := qovery.New(configurationFilename)
	if err != nil {
		ctx.Writef("fail to init qv client: %s", err.Error())
		return
	}
	databases := qv.GetDatabaseConfigurations()
	if len(databases) == 0 {
		ctx.Writef("no databases found")
		return
	}
	for _, dbConf := range qv.GetDatabaseConfigurations() {
		ctx.Writef("database: %s\n", dbConf.Name)
		dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d", dbConf.Host, dbConf.Username, dbConf.Name, dbConf.Password, dbConf.Port)
		db, err := gorm.Open("postgres", dbURI)
		if err != nil {
			ctx.Writef("fail to connect to dbConf: %s", err.Error())
			continue
		}
		ctx.Writef("connection to '%s' successful", dbConf.Name)
		db.Close()
		ctx.Writef("\n")
	}
}
