package mrp

import (
	"errors"
	"strings"
)

type MRP struct {
	Request string
	Body    string
	Footers []string
}

func NewMRP(reqType string, message string, footerlist ...string) MRP {

	var packet = MRP{
		Request: reqType,
		Body:    message,
		Footers: footerlist}

	return packet
}

func ReadMRP(packet []byte) (MRP, error) {
	var retMRP = MRP{}

	var message = string(packet)

	var splitMessage = strings.Split(message, "\n")

	if splitMessage[len(splitMessage)-1] != "EOF" {
		return retMRP, errors.New("error: MRP message not complete or missing lines")
	}

	retMRP.Request = splitMessage[0]
	retMRP.Body = splitMessage[1]
	retMRP.Footers = splitMessage[2 : len(splitMessage)-1]

	return retMRP, nil
}

func MRPToByte(mrp MRP) []byte {

	var fullString = mrp.Request + "\n" + mrp.Body

	for _, v := range mrp.Footers {
		fullString = fullString + "\n" + v
	}

	fullString = fullString + "\n" + "EOF"

	var packet = []byte(fullString)

	return packet
}
