package devinput

import (
	"golang.org/x/sys/unix"
)

type Event struct {
	// Timestamp is the time at which the event happend.
	Timestamp unix.Timeval
	// Type is the type of event, e.g., EV_KEY
	Type uint16
	// Code is the event code, e.g., REL_X or KEY_BACKSPACE
	Code uint16
	// Value is the event's value, e.g.,
	// 0 for EV_KEY for release,
	// 1 for keypress,
	// 2 for autorepeat.
	Value uint32
}

const (
	TypeKey = 0x01

	KeyBrightnessDown           = 224
	KeyBrightnessUp             = 225
	KeyScale                    = 120
	KeyExpose                   = KeyScale
	KeyDashboard                = 204
	KeyKeyboardIlluminationDown = 229
	KeyKeyboardIlluminationUp   = 230
	KeyPreviousSong             = 165
	KeyPlayPause                = 164
	KeyNextSong                 = 163
	KeyMute                     = 113
	KeyVolumnDown               = 114
	KeyVolumnUp                 = 115

	ValueKeyRelease = 0
	ValueKeyPress   = 1
	ValueKeyRepeat  = 2
)
