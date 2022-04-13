package main

import "github.com/tadvi/winc"

func NewEdit(window *winc.Form) (t *winc.Edit) {
	edt := winc.NewEdit(window)
	edt.SetMaxSize(1440, 900)
	edt.SetSize(800, 400)
	edt.SetPos(0, 28)
	winc.RGB(123, 44, 23)
	return edt
}
