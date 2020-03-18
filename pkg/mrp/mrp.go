package mrp

import (
	"errors"
	"strings"
)

type MRP struct {
	request string
	body    string
	footers []string
}

func NewMRP(reqType string, message string, footerlist ...string) MRP {

	var packet = MRP{
		request: reqType,
		body:    message,
		footers: footerlist}

	return packet
}

func ReadMRP(packet []byte) (MRP, error) {
	var retMRP = MRP{}

	var message = string(packet)

	var splitMessage = strings.Split(message, "/n")

	var lastValue = len(splitMessage) - 1

	if len(splitMessage) < 3 {
		return retMRP, errors.New("error: MRP message not complete or missing lines")
	}

	retMRP.request = splitMessage[0]
	retMRP.body = splitMessage[1]
	retMRP.footers = splitMessage[1:lastValue]

	return retMRP, nil
}

func MRPToByte(mrp MRP) []byte {

	var fullString = mrp.request + "\n" + mrp.request

	for _, v := range mrp.footers {
		fullString = "\n" + fullString + v
	}

	var packet = []byte(fullString)

	return packet
}
