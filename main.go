package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jekasell/api"
	db "github.com/jekasell/db/sqlc"
	"github.com/jekasell/util"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := util.LoadConfig(".")

	if err != nil {
		fmt.Printf("cannot load config: %v\n", err)
		os.Exit(1)
	}
	conn, err := sql.Open(cfg.DBDriver, cfg.DBUrl)

	if err != nil {
		fmt.Printf("cannot connect to db: %v\n", err)
		os.Exit(1)
	}
	str := db.NewStore(conn)
	svr := api.NewServer(str)

	log.Fatal(svr.Start(cfg.ServerAddr))
}
