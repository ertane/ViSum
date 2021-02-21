package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var mainWindow fyne.Window

// SelectFile ... UI struct
type SelectFile struct {
	out *widget.Label
	add *widget.Button
}

func (c *SelectFile) openFile(win fyne.Window) {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err == nil && reader == nil {
			return
		}
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		c.out.SetText(reader.URI().Path())
	}, win)
}

func selectNewFile() *SelectFile {
	c := &SelectFile{}
	c.out = widget.NewLabel("No choosen file")
	c.add = widget.NewButton(" Select file ", func() {
		c.openFile(mainWindow)
	})

	return c
}

func showMD5Sum(lbl *widget.Label, cf string) {
	lbl.SetText("Please wait...")
	file, _ := os.Open(cf)
	defer file.Close()

	m := md5.New()
	io.Copy(m, file)
	lbl.SetText(fmt.Sprintf("MD5 : %x", m.Sum(nil)))
}

func showSHA1Sum(lbl *widget.Label, cf string) {
	lbl.SetText("Please wait...")
	file, _ := os.Open(cf)
	defer file.Close()

	s1 := sha1.New()
	io.Copy(s1, file)
	lbl.SetText(fmt.Sprintf("SHA1: %x", s1.Sum(nil)))
}

func showSHA2Sum(lbl *widget.Label, cf string) {
	lbl.SetText("Please wait...")
	file, _ := os.Open(cf)
	defer file.Close()

	s2 := sha256.New()
	io.Copy(s2, file)
	lbl.SetText(fmt.Sprintf("SHA2: %x", s2.Sum(nil)))
}

func main() {
	appUI := app.New()
	win := appUI.NewWindow(" ViSum ")
	mainWindow = win

	c := selectNewFile()

	labelMD5 := widget.NewLabel("")
	labelSHA1 := widget.NewLabel("")
	labelSHA2 := widget.NewLabel("")

	btnSum := widget.NewButton("Calculate", func() {
		if c.out.Text != "No choosen file" {
			go showMD5Sum(labelMD5, c.out.Text)
			go showSHA1Sum(labelSHA1, c.out.Text)
			go showSHA2Sum(labelSHA2, c.out.Text)
		}
	})

	topui := container.NewVBox(c.out, container.NewHBox(c.add, btnSum))
	bottomui := container.NewVBox(labelMD5, labelSHA1, labelSHA2)
	ui := container.NewVBox(topui, bottomui)
	win.SetContent(ui)

	win.Resize(fyne.NewSize(700, 400))

	win.ShowAndRun()
}
