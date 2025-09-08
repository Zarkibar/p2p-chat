package ui

import (
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	App               *tview.Application
	Messages          *tview.TextView
	InputField        *tview.InputField
	AddrList          *tview.List
	ConnectInputField *tview.InputField
	ChatLayout        *tview.Flex
	MenuLayout        *tview.Flex
)

func InitializeUI() {
	App = tview.NewApplication()

	// Menu
	menuList := tview.NewList().
		AddItem("Connect", "Type other peer's address to connect", 's', func() {
			App.SetRoot(ConnectInputField, true)
		}).
		AddItem("Addresses", "View your addresses", 'a', func() {
			App.SetRoot(AddrList, true)
		}).
		AddItem("Quit", "Exit the program", 'q', func() {
			App.Stop()
		})

	MenuLayout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(menuList, 0, 1, true)

	// Connect Screen
	ConnectInputField = tview.NewInputField().
		SetLabel("> ").
		SetPlaceholder("Enter Address...").
		SetPlaceholderTextColor(tcell.ColorWhite).
		SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.ColorBlack) //.SetDoneFunc(func(key tcell.Key) {})

	// Addresses List
	AddrList = tview.NewList().
		AddItem("Back", "", 0, func() {
			App.SetRoot(MenuLayout, true)
		})

	// Chat
	Messages = tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			Messages.ScrollToEnd()
			App.Draw()
		})

	InputField = tview.NewInputField().
		SetLabel("> ").
		SetPlaceholder("Enter your message...").
		SetPlaceholderTextColor(tcell.ColorWhite).
		SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.ColorBlack) //.SetDoneFunc(func(key tcell.Key) {})

	ChatLayout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(Messages, 0, 1, false).
		AddItem(InputField, 1, 1, true)
}

func StartUI() {
	if err := App.SetRoot(MenuLayout, true).Run(); err != nil {
		panic(err)
	}
}

func SwitchPage(page tview.Primitive) {
	App.SetRoot(page, true)
}

func NewMultiAddress(text string) {
	AddrList.AddItem(text, "", 0, func() {
		if err := clipboard.WriteAll(text); err != nil {
			panic(err)
		}
		println("\n Copied Address")
	})
}

func AddSystemMessage(msg string) {
	Messages.Write([]byte("[yellow:b]" + msg + "\n"))
}

func AddMessage(msg string, color ...string) {
	c := "green"
	if len(color) > 0 && color[0] != "" {
		c = color[0]
	}

	Messages.Write([]byte("[" + c + "::b]" + msg + "\n"))
}
