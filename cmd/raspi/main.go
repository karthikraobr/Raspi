package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "11")
	button := gpio.NewButtonDriver(r, "7")

	work := func() {
		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button pressed")
			if err := led.On(); err != nil {
				fmt.Println("led error")
			}
		})

		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button released")
			if err := led.Off(); err != nil {
				fmt.Println("led error")
			}
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button, led},
		work,
	)

	robot.Start()
	// r := raspi.NewAdaptor()
	// led := gpio.NewLedDriver(r, "7")

	// work := func() {
	// 	gobot.Every(1*time.Second, func() {
	// 		led.Toggle()
	// 	})
	// }

	// robot := gobot.NewRobot("blinkBot",
	// 	[]gobot.Connection{r},
	// 	[]gobot.Device{led},
	// 	work,
	// )

	// robot.Start()
}
