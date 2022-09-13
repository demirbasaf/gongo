package main

import (
	"github.com/afdemirbas/gongo/gongo"
	"github.com/afdemirbas/gongo/gongo/configuration"
	"github.com/afdemirbas/gongo/gongo/examples"
)

func main() {
	gc := configuration.GongoConfig{URI: ""}

	err := gongo.Connect(gc)

	if err != nil {
		panic("")
	}

	userModel := gongo.Model(examples.User{})

	userModel.New()
	userModel.Find()

}