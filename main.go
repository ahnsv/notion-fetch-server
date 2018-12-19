package main

import (
	"github.com/ahnsv/notion-fetch-server/notion"
	"github.com/ahnsv/notion-fetch-server/server"
)

func main() {
	notion.Init()
	server.Serve()
}
