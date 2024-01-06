package config

import "github.com/gorilla/websocket"

var WSUpgrader websocket.Upgrader

func initWebsocket() {
	WSUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}
