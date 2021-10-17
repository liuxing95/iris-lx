package controllers

import (
	"iris-lx/web/wrapper"
)

type UserController struct {
}
type Book struct {
	Title string `json:"title"`
}

func (c UserController) Login(ctx *wrapper.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	ctx.JSON(books)
}
