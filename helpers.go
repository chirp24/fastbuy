package main

import (
	"xabbo.b7c.io/goearth/shockwave/in"
	"xabbo.b7c.io/goearth/shockwave/profile"
	"xabbo.b7c.io/goearth/shockwave/room"
)

func showMsg(msg string) {
	self := roomMgr.EntityByName(profileMgr.Name)
	if self == nil {
		// fmt.Println("self not found.")
		return
	}
	ext.Send(in.CHAT, self.Index, msg)
}


var roomMgr = room.NewManager(ext)
var profileMgr = profile.NewManager(ext)