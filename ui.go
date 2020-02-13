package main

import (
	"github.com/tadvi/winc"
)

func main() {
	gui()
	winc.RunMainLoop()
}

func gui() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300)
	mainWindow.SetText("getEmailList")

	// Main window menu. Context menus on controls also available.
	menu := mainWindow.NewMenu()
	fileMn := menu.AddSubMenu("File")
	newFileMn := fileMn.AddItem("Open file", winc.Shortcut{winc.ModControl, winc.KeyN})
	menu.Show()
	edt := winc.NewEdit(mainWindow)
	edt.SetPos(10, 20)
	edt.SetText("Choose file")
	var path string
	// Menu items can be disabled and checked.
	newFileMn.OnClick().Bind(func(e *winc.Event) {
		path, _ = winc.ShowOpenFileDlg(mainWindow, "Select File", ".csv", 1, "home")
		edt.SetText(path)
	})

	// Most Controls have default size unless SetSize is called.

	btn := winc.NewPushButton(mainWindow)
	btn.SetText("Generate List")
	btn.SetPos(40, 50)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(func(e *winc.Event) {
		err := process(path)
		if err != nil {
			winc.MsgBoxOk(mainWindow, "Failed", "Something went wrong "+err.Error())
		} else {
			winc.MsgBoxOk(mainWindow, "Generated List", "EmailList.txt got generated!")
		}
	})

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)

}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}
