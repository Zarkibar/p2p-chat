package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	App        *tview.Application
	Messages   *tview.TextView
	InputField *tview.InputField
	Layout     *tview.Flex
)

func InitializeUI() {
	App = tview.NewApplication()

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

	Layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(Messages, 0, 1, false).
		AddItem(InputField, 1, 1, false)
}

func InputDoneCallback(callback func(key tcell.Key)) {
	InputField.SetDoneFunc(callback)
}

func StartUI() {
	if err := App.SetRoot(Layout, true).SetFocus(InputField).Run(); err != nil {
		panic(err)
	}
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
