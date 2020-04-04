package mrp

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type MRP struct {
	request []byte
	body    []byte
	footers [][]byte
}

func NewMRP(reqType []byte, message []byte, footerlist ...[]byte) *MRP {

	var packet = MRP{
		request: reqType,
		body:    message,
		footers: footerlist}

	return &packet
}

func ReadMRPFromBytes(packet []byte) (*MRP, error) {
	var retMRP = MRP{}

	var message = string(packet)

	var splitMessage = strings.Split(message, "\n")

	if splitMessage[len(splitMessage)-1] != "EOF" || len(splitMessage) < 3 {
		return &retMRP, errors.New("error: MRP message not complete or missing lines")
	}

	retMRP.request = []byte(splitMessage[0])
	retMRP.body = []byte(splitMessage[1])
	for _, v := range splitMessage[2 : len(splitMessage)-1] {
		retMRP.footers = append(retMRP.footers, []byte(v))
	}

	return &retMRP, nil
}

func ReadMRPFromConn(conn net.Conn, handleMRP func([]*MRP, net.Conn)) {
	var err error

	var carryOver []byte

	for {
		var message = make([]byte, 0)
		var newMRP *MRP
		var newMRPList []*MRP

		//if any message was not complete during the last pull
		//from buffer, carryOver stores it. After nil assignment
		//from above, carryOver fills in the beggining couple lines
		message = carryOver

		for {
			var buf = make([]byte, 1024)
			//conn.SetReadDeadline(time.Now().Add(10 * time.Second))
			_, err = conn.Read(buf)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			//This is a simple loop used to remove all null chars
			//from the buffer so that they aren't accidentally read
			for k, v := range buf {
				if v == 0 {
					buf = buf[0:k]
					break
				}
			}

			//we append the buffer to a message because it allows for
			//us to pull a larger message if one packet of 1024 bytes
			//was not enough. In that case buffer get overwrittin but
			//message would not change.
			message = append(message, buf...)
			//Here we are checking for multiple MRP's in the buffer,
			//This is useful incase we pull many small MRP's in one,
			//1024 byte buffer.
			messageString := strings.SplitAfter(string(message), "EOF")

			for _, v := range messageString {

				//We just check the segment to see if can be recongnised
				//as a MRP packet. If so we add it to the growing list of
				//MRP's and move on, else we assume the message is incomplete
				//and add it to the carryOver message.
				newMRP, err = ReadMRPFromBytes([]byte(v))
				if err == nil {
					newMRPList = append(newMRPList, newMRP)
				} else {
					carryOver = message[len(message)-len(v):]
				}

			}

			//If we have found any MRP's during the above loop
			//we are going to break out and begin processing them
			//before moving back into the list and continue pulling
			if len(newMRPList) != 0 {
				break
			}

		}

		handleMRP(newMRPList, conn)
	}
}

func (s *MRP) MRPToByte() []byte {

	var fullString = s.request
	fullString = append(fullString, byte('\u000a'))
	fullString = append(fullString, s.body...)

	for _, v := range s.footers {
		fullString = append(fullString, byte('\u000a'))
		fullString = append(fullString, v...)
	}

	fullString = append(fullString, byte('\u000a'))
	fullString = append(fullString, []byte("EOF")...)

	var packet = []byte(fullString)

	return packet
}

func (s *MRP) GetRequest() string {
	return string(s.request)
}

func (s *MRP) GetBody() string {
	return string(s.body)
}

func (s *MRP) GetFooters() []string {
	var footerString []string
	for _, item := range s.footers {
		for _, bytes := range item {
			footerString = append(footerString, string(bytes))
		}
	}
	return footerString
}
