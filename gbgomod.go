package gbgomod

import (
	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
)

type One struct {
	Field1, Field2 string
}

func NewOne() *One {

	return &One{
		Field1: "gfd",
		Field2: "tre",
	}
}
