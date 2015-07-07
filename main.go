package main

import (
	"log"

	"github.com/bernerdschaefer/applemediakeyd/devinput"
	"github.com/bernerdschaefer/applemediakeyd/oss"
)

const (
	eventInputDevice     = "/dev/input/event0"
	backlightDir         = "/sys/class/backlight/acpi_video0/"
	keyboardBacklightDir = "/sys/devices/platform/applesmc.768/leds/smc::kbd_backlight/"
	lightSensorFile      = "/sys/devices/platform/applesmc.768/light"
	masterMixerDevice    = "/dev/mixer1"
)

func main() {
	mixer, err := oss.OpenMixer(masterMixerDevice)
	if err != nil {
		log.Fatal(err)
	}

	source, err := devinput.Open(eventInputDevice)
	if err != nil {
		log.Fatal(err)
	}

	var ev devinput.Event

	for {
		if err := source.Read(&ev); err != nil {
			log.Fatal(err)
		}

		if ev.Type != devinput.TypeKey || ev.Value != devinput.ValueKeyRelease {
			continue
		}

		switch ev.Code {
		case devinput.KeyBrightnessUp:
			adjustBrightness(backlightDir, 1)
		case devinput.KeyBrightnessDown:
			adjustBrightness(backlightDir, -1)
		case devinput.KeyKeyboardIlluminationUp:
			adjustBrightness(keyboardBacklightDir, 1)
		case devinput.KeyKeyboardIlluminationDown:
			adjustBrightness(keyboardBacklightDir, -1)
		case devinput.KeyVolumnUp:
			mixer.AdjustVolume(+10)
		case devinput.KeyVolumnDown:
			mixer.AdjustVolume(-10)
		case devinput.KeyMute:
			mixer.ToggleMute()
		}
	}
}
