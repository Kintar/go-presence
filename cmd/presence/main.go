package main

import (
	"fmt"
	"github.com/kintar/go-presence/icons"
	"math/rand"
	"os"
	"time"

	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
)

func main() {
	systray.Run(onReady, onExit)
}

func onExit() {
	// GNDN - Goes Nowhere, Does Nothing
}

var paused bool

func onReady() {
	paused = true

	systray.SetTemplateIcon(icons.Waiting, icons.Waiting)
	//systray.SetTitle("Presence Faker")
	systray.SetTooltip("Fakes mouse activity")
	mPause := systray.AddMenuItemCheckbox("Pause", "Stop moving the mouse", paused)
	mPause.SetIcon(icons.Pause)
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit")
	mQuit.SetIcon(icons.Quit)

	go presenceFunc()

	fmt.Printf("Running : PID %d\n", os.Getpid())

	for {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
		case <-mPause.ClickedCh:
			paused = !paused
			if paused {
				mPause.Check()
				systray.SetTemplateIcon(icons.Waiting, icons.Waiting)
			} else {
				mPause.Uncheck()
				systray.SetTemplateIcon(icons.Working, icons.Working)
			}
		}
	}
}

func presenceFunc() {
	sx, sy := robotgo.GetScreenSize()
	sx -= 50
	sy -= 50

	for {
		select {
		case <-time.After(time.Second * 10):
			if !paused {
				targetX := rand.Intn(sx) + 25
				targetY := rand.Intn(sy) + 25
				robotgo.MoveSmooth(targetX, targetY, 0.25, 1.0)
			}
		}
	}
}
