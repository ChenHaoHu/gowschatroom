package server

import (
	"errors"
	"strings"

	"github.com/gorilla/websocket"
)

type Member struct {
	Uid  string
	Name string
	Conn *websocket.Conn
}

func ParseToken(token string) (*Member, error) {

	strs := strings.Split(token, ",")

	if len(strs) == 2 {

		return &Member{
			Uid:  strs[0],
			Name: strs[1],
		}, nil
	}

	return nil, errors.New("token error")
}
