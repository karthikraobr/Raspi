package main

import (
	"github.com/karthikraobr/raspi/led"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	robot := led.Blink(r, "11", "7")
	robot.Start()
}
