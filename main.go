package main

import (
	"fmt"
	"github.com/afdemirbas/gongo/gongo"
	"github.com/afdemirbas/gongo/gongo/examples"
)

func main() {

	userModel := gongo.Model(examples.User{})
	articleModel := gongo.Model(examples.Article{})

	r := userModel.New("")
	r2 := articleModel.New("")

	fmt.Println(r, r2)
}