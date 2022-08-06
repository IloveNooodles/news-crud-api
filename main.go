package main

import (
	"github.com/ilovenooodles/news-crud-api/cmd"
	"github.com/ilovenooodles/news-crud-api/config"
)

func main() {
	config.Init()
	cmd.Init()
}
