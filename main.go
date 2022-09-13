package main

import (
	"github.com/afdemirbas/gongo/src/entity"
	"github.com/afdemirbas/gongo/src/examples"
)

func main() {
	entity.NewModel(examples.User{})
	entity.NewModel(examples.Article{})
}