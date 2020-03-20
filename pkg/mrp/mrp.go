package mrp

import (
	"errors"
	"strings"
)

type MRP struct {
	Request []byte
	Body    []byte
	Footers [][]byte
}

func NewMRP(reqType []byte, message []byte, footerlist ...[]byte) MRP {

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

	if splitMessage[len(splitMessage)-1] != "EOF" || len(splitMessage) < 3 {
		return retMRP, errors.New("error: MRP message not complete or missing lines")
	}

	retMRP.Request = []byte(splitMessage[0])
	retMRP.Body = []byte(splitMessage[1])
	for _, v := range splitMessage[2 : len(splitMessage)-1] {
		retMRP.Footers = append(retMRP.Footers, []byte(v))
	}

	return retMRP, nil
}

func MRPToByte(mrp MRP) []byte {

	var fullString = mrp.Request
	fullString = append(fullString, byte('\u000a'))
	fullString = append(fullString, mrp.Body...)

	for _, v := range mrp.Footers {
		fullString = append(fullString, byte('\u000a'))
		fullString = append(fullString, v...)
	}

	fullString = append(fullString, byte('\u000a'))
	fullString = append(fullString, []byte("EOF")...)

	var packet = []byte(fullString)

	return packet
}
