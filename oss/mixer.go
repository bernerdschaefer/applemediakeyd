package oss

import (
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	SOUND_MIXER_READ_VOLUME  = 0x80044D00
	SOUND_MIXER_WRITE_VOLUME = 0xC0044D00
)

type Mixer struct {
	device *os.File

	volume        int
	restoreVolume int
}

func OpenMixer(device string) (*Mixer, error) {
	f, err := os.Open(device)

	if err != nil {
		return nil, err
	}

	mixer := &Mixer{device: f}
	volume, err := mixer.readVolume()
	if err != nil {
		return nil, err
	}
	mixer.volume = volume

	return mixer, nil
}

func (m *Mixer) ToggleMute() error {
	if m.volume == 0 && m.restoreVolume == 0 {
		// volume is muted with no previous known state,
		// guess at 30.
		return m.SetVolume(30)
	}

	volume := m.volume

	if err := m.SetVolume(m.restoreVolume); err != nil {
		return err
	}

	m.restoreVolume = volume
	return nil
}

func (m *Mixer) AdjustVolume(amount int) error {
	volume := m.volume + amount
	if volume < 0 {
		volume = 0
	} else if volume > 100 {
		volume = 100
	}

	return m.SetVolume(volume)
}

func (m *Mixer) Volume() int {
	return m.volume
}

func (m *Mixer) SetVolume(volume int) error {
	if err := m.writeVolume(volume | (volume << 8)); err != nil {
		return err
	}

	m.volume = volume
	return nil
}

func (m *Mixer) readVolume() (int, error) {
	var currentVolume int

	_, _, err := unix.Syscall(
		unix.SYS_IOCTL,
		m.device.Fd(),
		SOUND_MIXER_READ_VOLUME,
		uintptr(unsafe.Pointer(&currentVolume)),
	)

	if err != 0 {
		return 0, err
	}

	return currentVolume & 0xff, nil
}

func (m *Mixer) writeVolume(channelVolume int) error {
	_, _, err := unix.Syscall(
		unix.SYS_IOCTL,
		m.device.Fd(),
		SOUND_MIXER_WRITE_VOLUME,
		uintptr(unsafe.Pointer(&channelVolume)),
	)

	if err != 0 {
		return err
	}

	return nil
}

func (m *Mixer) Close() error {
	return m.device.Close()
}
