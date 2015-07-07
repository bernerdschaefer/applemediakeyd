package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func adjustBrightness(device string, sign int) {
	current, err := readIntFromFile(device + "brightness")
	if err != nil {
		log.Fatalf("unable to read brightness: %s", err)
	}
	max, err := readIntFromFile(device + "max_brightness")
	if err != nil {
		log.Fatalf("unable to read max brightness: %s", err)
	}

	var (
		step = max * sign / 10
		next = current + step
	)

	if next > max {
		next = max
	} else if next < 0 {
		next = 0
	}

	if err := writeIntToFile(device+"brightness", next); err != nil {
		log.Fatalf("unable to write new brightness: %s", err)
	}
}

func readIntFromFile(filename string) (int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	var value int
	if _, err := fmt.Sscanf(string(data), "%d", &value); err != nil {
		return 0, err
	}

	return value, nil
}

func writeIntToFile(filename string, value int) error {
	return ioutil.WriteFile(filename, []byte(strconv.Itoa(value)), os.ModePerm)
}
