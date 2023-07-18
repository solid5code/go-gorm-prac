package main

import (
	"log"

	"github.com/caarlos0/env/v6"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type (
	databaseEnv struct {
		DataCenter string `env:"DATA_CENTER_DATABASE_DSN" required:"true" envExpand:"true"`
	}

	genEnv struct {
		DatabaseDSN  string
		DatabaseName string
		TableNames   []string
	}
)

var _genEnv []genEnv

func init() {
	var databaseEnv databaseEnv
	if err := env.Parse(&databaseEnv); err != nil {
		log.Fatal(err)
	}

	_genEnv = []genEnv{
		{
			DatabaseName: "datacenter",
			DatabaseDSN:  databaseEnv.DataCenter,
			TableNames: []string{
				"user",
			},
		},
	}
}

func main() {
	for _, env := range _genEnv {
		db, err := gorm.Open(mysql.Open(env.DatabaseDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		g := gen.NewGenerator(gen.Config{
			OutPath:           "../model/dao/mysql/" + env.DatabaseName,
			FieldNullable:     true,
			FieldWithIndexTag: true,
			FieldWithTypeTag:  true,
			Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		})

		g.UseDB(db)

		for _, tn := range env.TableNames {
			g.ApplyBasic(g.GenerateModel(tn))
		}

		g.Execute()
	}
}
