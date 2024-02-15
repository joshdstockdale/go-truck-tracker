package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/joshdstockdale/go-truck-tracker/types"
)

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msgchan chan types.OBUData
	conn    *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgchan: make(chan types.OBUData, 128),
	}
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("client connected")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {

			log.Println("read error:", err)
			continue
		}
		fmt.Println("Received", data)
		dr.msgchan <- data

	}
}
