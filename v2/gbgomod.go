// bggomod is a homework for lesson3

package gbgomod

import (
	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
)

type Three struct {
	Field1, Field2 string
}

func NewThree() *Three {

	return &Three{
		Field1: "11111gfd",
		Field2: "22222tre",
	}
}

type Two struct {
	Field1 string
}

func NewTwo() *Two {
	return &Two{
		Field1: "val1",
	}
}
