package main

import (
	"context"

	"github.com/gdamore/tcell/v2"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	peerstore "github.com/libp2p/go-libp2p/core/peer"

	"p2p-net/internal/p2p"
	"p2p-net/internal/ui"
	"p2p-net/pkg/utility"
)

var (
	username  string = "User"
	userColor string = "white"
	node      host.Host
	ctx       context.Context
	peerinfo  *peerstore.AddrInfo
)

func main() {
	ctx = context.Background()
	username = utility.GenerateName()
	userColor = utility.GenerateColor()

	node = p2p.CreateNode()
	defer node.Close()

	ui.InitializeUI()
	ui.InputField.SetDoneFunc(setInputCommands)

	for i := 0; i < len(node.Addrs()); i++ {
		addr := node.Addrs()[i].String() + "/p2p/" + node.ID().String()
		ui.NewMultiAddress(addr)
	}

	ui.ConnectInputField.SetDoneFunc(func(key tcell.Key) {
		peerinfo = p2p.ConnectToPeer(node, ctx, ui.ConnectInputField.GetText())

		s, err := node.NewStream(ctx, peerinfo.ID, "/connection/1.0.0")
		if err != nil {
			panic(err)
		}
		s.Write([]byte(node.Addrs()[0].String() + "/p2p/" + node.ID().String()))

		ui.SwitchPage(ui.ChatLayout)
	})

	node.SetStreamHandler("/connection/1.0.0", func(s network.Stream) {
		peerinfo = p2p.HandleConnectionStream(s, node, ctx)
		ui.SwitchPage(ui.ChatLayout)
	})

	node.SetStreamHandler("/chat/1.0.0", func(s network.Stream) {
		p2p.HandleChatStream(s, func(msg string) { ui.AddMessage(msg, userColor) })
	})

	ui.StartUI()
}

func setInputCommands(key tcell.Key) {
	if ui.InputField.GetText()[0] == '/' {
		if ui.InputField.GetText()[1:10] == "username " {
			username = ui.InputField.GetText()[10:]
		}
	} else {
		msg := username + ": " + ui.InputField.GetText()
		ack := p2p.SendMsg(node, ctx, peerinfo, msg)
		if ack == "ack" {
			ui.AddMessage(msg)
		} else {
			ui.AddSystemMessage("Message not delivered")
		}
	}
	ui.InputField.SetText("")
}
