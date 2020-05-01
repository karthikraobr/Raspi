package led

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func Blink(raspi *raspi.Adaptor, ledPin string, buttonPin string) *gobot.Robot {
	led := gpio.NewLedDriver(raspi, ledPin)
	button := gpio.NewButtonDriver(raspi, buttonPin)
	work := func() {
		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.On()
		})

		button.On(gpio.ButtonPush, func(data interface{}) {
			led.Off()
		})
	}
	return gobot.NewRobot("buttonBot",
		[]gobot.Connection{raspi},
		[]gobot.Device{button, led},
		work,
	)
}
