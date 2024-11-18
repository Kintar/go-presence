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

type controlMessage int

const (
	togglePause controlMessage = iota + 1
	quit
)

var controlChannel = make(chan controlMessage, 1)

func main() {
	systray.Run(onReady, onExit)
}

func onExit() {
	// GNDN - Goes Nowhere, Does Nothing
}

var paused bool

func absi(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func onReady() {
	paused = true

	systray.SetTemplateIcon(icons.Waiting, icons.Waiting)
	systray.SetTitle("Presence")
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
			controlChannel <- quit
		case <-mPause.ClickedCh:
			if !paused {
				mPause.Check()
				systray.SetTemplateIcon(icons.Waiting, icons.Waiting)
			} else {
				mPause.Uncheck()
				systray.SetTemplateIcon(icons.Working, icons.Working)
			}
			controlChannel <- togglePause
		}
	}
}

const mouseJitter = 10

var moveDelay time.Duration = 10

func presenceFunc() {
	timer := time.NewTimer(time.Second * moveDelay)

	sx, sy := robotgo.GetScreenSize()
	sx -= 50
	sy -= 50
	lx, ly := robotgo.GetMousePos()

	for {
		select {
		case msg := <-controlChannel:
			switch msg {
			case togglePause:
				paused = !paused
				if paused {
					timer.Stop()
				} else {
					timer.Reset(time.Second * moveDelay)
				}
			case quit:
				systray.Quit()
			}
			return
		default:
		case <-timer.C:
			timer.Reset(time.Second * moveDelay)

			// if the mouse has been moved by more than a small amount, skip the auto-move
			cx, cy := robotgo.GetMousePos()
			if absi(cx-lx) > mouseJitter || absi(cy-ly) > mouseJitter {
				lx, ly = cx, cy
				continue
			}

			targetX := rand.Intn(sx) + 25
			targetY := rand.Intn(sy) + 25
			robotgo.MoveSmooth(targetX, targetY, 0.25, 1.0)
			lx, ly = targetX, targetY
		}
	}
}
