package main

import (
	"github.com/getlantern/systray"
	"github.com/gobuffalo/packr"
	"github.com/pkg/browser"
	"log"
	"net/http"
	"os"
	"photon_explorer_serv/icon"
	"time"
)

func main() {
	// TODO: check if other instance of photon_explorer_serv is already running and if so just open browser
	go systray.Run(onReady, onExit)

	box := packr.NewBox("./dist")
	http.Handle("/", http.FileServer(box))
	go func() {
		time.Sleep(100 * time.Millisecond)
		// TODO: search for free port
		err := browser.OpenURL("http://localhost:3000")
		if err != nil {
			log.Fatal(err)
		}
	}()
	err := http.ListenAndServe("127.0.0.1:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func onReady() {
	systray.SetIcon(icon.Photonicon)
	systray.SetTooltip("Photon Explorer")
	open := systray.AddMenuItem("Open in browser", "Open Photon Explorer in your browser")
	quit := systray.AddMenuItem("Quit", "Quit Photon Explorer")
	go func() {
		<-quit.ClickedCh
		systray.Quit()
		os.Exit(0)
	}()
	go func() {
		<-open.ClickedCh
		time.Sleep(100 * time.Millisecond)
		// TODO: check on which port this is running
		err := browser.OpenURL("http://localhost:3000")
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func onExit() {
}
