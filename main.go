package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

type Notification struct {
	Title, Content string
}

const title = "Goal Planing App"

var todoFileName = ".todo.json"

func timer(clock *widget.Label) {
	formatted := time.Now().Format("03:04")
	clock.SetText(formatted)
}

func FileFunc() {

}

func main() {

	a := app.New()
	w := a.NewWindow(title)
	newItem := fyne.NewMenuItem("Save", nil)

	file := fyne.NewMenu("File", newItem)
	//if !fyne.CurrentDevice().IsMobile() {
	//		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), newItem)
	//	}
	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		file,
	)
	w.SetMainMenu(mainMenu)
	w.SetMaster()

	a.SetIcon(theme.FyneLogo())
	makeTray(a)

	label1 := widget.NewLabel("Goal:")
	textbox := widget.NewMultiLineEntry()
	textbox.SetPlaceHolder("Text here")
	label2 := widget.NewLabel("Time:")
	clocker := widget.NewLabel("")
	timer(clocker)
	//	saveme := widget.NewButton("SAVE", func() {
	//		log.Println("Saved")
	//	})
	grid := container.New(layout.NewFormLayout(), label1, textbox, label2, clocker)
	///box1 := container.New(layout.NewVBoxLayout(), saveme)
	w.SetContent(grid)

	// w.SetContent(box1)
	go func() {
		for range time.Tick(time.Second) {
			timer(clocker)

		}
	}()
	alert := NewNotifaction("bra", "biz")
	a.SendNotification((*fyne.Notification)(alert))
	w.ShowAndRun()
}

func NewNotifaction(title, content string) *Notification {
	return &Notification{Title: "Time Is:", Content: "Foo"}
}

func makeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		clocker := widget.NewLabel("")
		timer(clocker)
		h := fyne.NewMenuItem("Hello", func() {})

		h.Icon = theme.HomeIcon()
		menu := fyne.NewMenu(title, h)
		h.Action = func() {
			log.Println("System tray menu tapped")
			h.Label = "Welcome"
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}

func savefunc() {

}
