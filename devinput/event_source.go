package devinput

import (
	"encoding/binary"
	"io"
	"os"
)

type EventSource struct {
	device io.ReadCloser
}

func Open(device string) (*EventSource, error) {
	f, err := os.Open(device)
	if err != nil {
		return nil, err
	}

	return &EventSource{device: f}, nil
}

func (es *EventSource) Read(ev *Event) error {
	return binary.Read(es.device, binary.LittleEndian, ev)

}

func (es *EventSource) Close() error {
	return es.device.Close()
}
