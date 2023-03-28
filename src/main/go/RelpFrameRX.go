package main

import (
	"errors"
	"strconv"
)

type RelpFrameRX struct {
	RelpFrame
}

func (rxFrame *RelpFrameRX) ParseResponseCode() (int, error) {
	var code = make([]byte, 3, 3)
	for i, v := range rxFrame.data {
		if i == 3 && v == ' ' {
			num, err := strconv.ParseInt(string(code), 10, 64)
			if err != nil {
				return 0, errors.New("relpFrameRX: could not parse response code")
			} else {
				return int(num), nil
			}
		} else if i >= 3 {
			return 0, errors.New("relpFrameRX: unexpected error code length: longer than 3 numbers")
		}

		if v >= 48 && v <= 57 {
			//0-9 ascii
			code[i] = v
		} else {
			return 0, errors.New("relpFrameRX: response code had a non-number ascii char")
		}
	}

	return 0, errors.New("relpFrameRX: response code could not be found")
}
