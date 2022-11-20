package main

import (
	"ginVueBlog/model"
	"ginVueBlog/routers"
)

func main() {
	model.InitDb()
	routers.InitRouter()
}
