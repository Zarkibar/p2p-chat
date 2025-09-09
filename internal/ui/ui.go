package ui

import (
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	PageMenu    string = "menu"
	PageConnect string = "connect"
	PageAddress string = "address"
	PageChat    string = "chat"
)

var (
	App               *tview.Application
	Messages          *tview.TextView
	InputField        *tview.InputField
	AddrList          *tview.List
	ConnectInputField *tview.InputField
	MenuList          *tview.List
	ChatLayout        *tview.Flex
	MenuLayout        *tview.Flex
	Pages             *tview.Pages
)

func InitializeUI() {
	App = tview.NewApplication()

	App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlC:
			App.Stop()
		case tcell.KeyEsc:
			page, _ := Pages.GetFrontPage()
			if page != PageMenu {
				Pages.SwitchToPage(PageMenu)
			}
		}

		// Or check for rune keys
		if event.Rune() == 'h' || event.Rune() == 'H' {
			// println("You pressed H!")
		}

		return event // return nil to swallow the key
	})

	// Menu
	MenuList = tview.NewList().
		AddItem("Connect", "Type other peer's address to connect", 's', func() {
			Pages.SwitchToPage(PageConnect)
		}).
		AddItem("Addresses", "View your addresses", 'a', func() {
			Pages.SwitchToPage(PageAddress)
		}).
		AddItem("Quit", "Exit the program", 'q', func() {
			App.Stop()
		})

	MenuLayout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(MenuList, 0, 1, true)

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
			Pages.SwitchToPage(PageMenu)
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

	// Page
	Pages = tview.NewPages()
	Pages.AddPage(PageMenu, MenuLayout, true, true)
	Pages.AddPage(PageConnect, ConnectInputField, true, false)
	Pages.AddPage(PageAddress, AddrList, true, false)
	Pages.AddPage(PageChat, ChatLayout, true, false)
}

func StartUI() {
	if err := App.SetRoot(Pages, true).Run(); err != nil {
		panic(err)
	}
}

func SwitchPage(name string) {
	Pages.SwitchToPage(name)
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

func AddMessage(msg string) {
	Messages.Write([]byte(msg + "\n"))
}
