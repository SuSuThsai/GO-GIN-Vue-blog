package main

import (
	"GO-GIN-Vue-blog/model"
	"GO-GIN-Vue-blog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
