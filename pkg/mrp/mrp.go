// Package mrp defines custom communcation protocol and related functions.
package mrp

import (
	"errors"
	"strings"
)

// MRP (Maritime Royale Packet) is a custom message protocol used between the server and client
// applications to communicate game state information.
type MRP struct {
	Request []byte
	Body    []byte
	Footers [][]byte
}

// NewMRP creates a new MRP message that consists of a request verb, message, and an array of footers.
// There is currently no validation of the type of data that can be placed in an MRP message.
func NewMRP(reqType []byte, message []byte, footerlist ...[]byte) MRP {

	var packet = MRP{
		Request: reqType,
		Body:    message,
		Footers: footerlist}

	return packet
}

// ReadMRP parses a MPR byte packet into a proper Go struct.
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

// ToByte takes a Go MRP construct and translates to a byte slice
// that can be communcated through a tcp connection.
func ToByte(mrp MRP) []byte {

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
