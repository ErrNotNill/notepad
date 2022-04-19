package main

import (
	"fmt"
	"github.com/tadvi/winc"
)

func FileButton(window *winc.Form) {
	NewPushButton(window,
		0, 0,
		60, 30,
		"File").OnClick().Bind(func(e *winc.Event) {
		tabs := winc.NewMultiPanel(window)

		tabs.SetPos(30, 30)
		tabs.SetSize(100, 92)
		fmt.Println("Create,NewWindow,Open,Save,SaveAs,ParametersPage,Print,Exit")
	})
}
func EditButton(window *winc.Form) {
	NewPushButton(window,
		60, 0,
		60, 30,
		"Edit").OnClick().Bind(func(e *winc.Event) {
		fmt.Println("Cancel,Cut,Copy,Paste,Delete,Search,FindNext,FindEarlier,ReplaceText")
	})
}
func FormatButton(window *winc.Form) {
	NewPushButton(window,
		120, 0,
		60, 30,
		"Format").OnClick().Bind(func(e *winc.Event) {
		fmt.Println("Create,NewWindow,Open,Save,SaveAs,ParametersPage,Print,Exit")
	})
}
func ViewButton(window *winc.Form) {
	NewPushButton(window,
		180, 0,
		60, 30,
		"View").OnClick().Bind(func(e *winc.Event) {
		fmt.Println("Create,NewWindow,Open,Save,SaveAs,ParametersPage,Print,Exit")
	})
}
func HelpButton(window *winc.Form) {
	NewPushButton(window,
		240, 0,
		60, 30,
		"Help").OnClick().Bind(func(e *winc.Event) {
		fmt.Println("Create,NewWindow,Open,Save,SaveAs,ParametersPage,Print,Exit")
	})
}
func ApiButton(window *winc.Form) {
	NewPushButton(window,
		300, 0,
		60, 30,
		"Api").OnClick().Bind(func(e *winc.Event) {
		fmt.Println("Create,NewWindow,Open,Save,SaveAs,ParametersPage,Print,Exit")
	})
}

func NewPushButton(wf *winc.Form, xPos int, yPos int, xSize int, ySize int, btnName string) *winc.PushButton {
	btn := winc.NewPushButton(wf)
	btn.SetText(btnName)
	btn.SetPos(xPos, yPos)
	btn.SetSize(xSize, ySize)
	return btn
}

func NewEditForm(wf *winc.Form, xPos int, yPos int, xSize int, ySize int) *winc.PushButton {
	edt := winc.NewEdit(wf)
	btn := winc.NewPushButton(wf)
	btn.SetText("save")
	btn.SetPos(xPos, yPos)
	btn.SetSize(xSize, ySize)

	btn.OnClick().Bind(func(e *winc.Event) {
		if edt.Visible() {
			edt.Hide()
		} else {
			edt.Show()
		}
	})
	return btn
}
