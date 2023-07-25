package server

import (
	"fmt"
	"time"
)

type server struct {
	token string
}

func NewServer(token string) *server {
	return &server{
		token: token,
	}
}

func RunLoop() {
	for {
		fmt.Println("test")
		time.Sleep(1000)
	}
}
