package main

import (
	"log"
	"time"

	"github.com/UnnoTed/go-appindicator"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	menu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal(err)
	}

	item, err := gtk.MenuItemNewWithLabel("item-label")
	if err != nil {
		log.Fatal(err)
	}

	indicator := appindicator.New("indicator-xyz", "network-transmit-receive", appindicator.CategoryApplicationStatus)
	indicator.SetTitle("indi-title")
	indicator.SetLabel("indi-label", "")
	indicator.SetStatus(appindicator.StatusActive)
	indicator.SetMenu(menu)

	item.Connect("activate", func() {
		indicator.SetLabel("activated", "")
	})

	menu.Add(item)
	menu.ShowAll()

	go func() {
		for {
			<-time.After(time.Second)
			label := time.Now().Format(time.RFC1123)
			indicator.SetLabel(label, "")
		}
	}()

	gtk.Main()
}
