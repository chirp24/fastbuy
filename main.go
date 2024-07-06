package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	g "xabbo.b7c.io/goearth"
	"xabbo.b7c.io/goearth/shockwave/in"
	"xabbo.b7c.io/goearth/shockwave/out"
)

var ext = g.NewExt(g.ExtInfo{
	Title:       "fastbuy",
	Description: "An extension to purchase items in bulk from the store.",
	Author:      "chirp",
	Version:     "1.0",
})

var packet0 *g.Packet
var buybool bool

func main() {

	var a int      // number of furni
	fmt.Println(a) //

	ext.Initialized(onInitialized)
	ext.Connected(onConnected)
	ext.Disconnected(onDisconnected)
	ext.Intercept(in.CHAT, in.CHAT_2, in.CHAT_3).With(handleChat)
	ext.Intercept(out.PURCHASE_FROM_CATALOG).With(func(e *g.Intercept) {
		packet0 = e.Packet.Copy()
		log.Println(packet0)
	})
	ext.Intercept(in.PURCHASE_OK).With(func(e *g.Intercept) {
		if buybool {
			e.Block()
		}
	})
	ext.Run()
}

func onInitialized(e g.InitArgs) {
	log.Println("Extension initialized")
}

func onConnected(e g.ConnectArgs) {
	log.Printf("Game connected (%s)\n", e.Host)
}

func onDisconnected() {
	log.Println("Game disconnected")
}

func handleChat(e *g.Intercept) {
	e.Packet.ReadInt() // skip entity index
	message1 := e.Packet.ReadString()
	if strings.Contains(message1, ":buy") { // :buy msg
		e.Block()
		log.Println(message1)

		parts := strings.Fields(message1) // split message into parts
		var a int
		for _, part := range parts {
			if num, err := strconv.Atoi(part); err == nil { // extracting int from str
				a = num
				break
			}
		}
		if a != 0 {
			fmt.Println("Captured Integer:", a)
			go buyitems(a)
		} else {
			fmt.Println("No Integer found in the string.") // no int found in :buy command
		}
	}
}

func buyitems(a int) {
	buybool = true
	defer func() {
		buybool = false
	}()
	for i := 1; i <= a; i++ { // stop looping when reach variable a value
		fmt.Println(i)
		if packet0 != nil {
			ext.SendPacket(packet0) // repeat buy packet
			time.Sleep(600 * time.Millisecond)
		} else {
			fmt.Println("No packet set to send.")
		}
	}
}
