package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"strconv"
)

var (
	in_addr = flag.String("in", "0.0.0.0:30002", "incoming socket")
)

func main() {
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp4", *in_addr)
	if err != nil {
		log.Fatalf("Can't resolve address: '%v'", err)
	}

	sock, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatalf("Can't open UDP socket: '%v'", err)
	}

	log.Printf("Start listening on udp://%v\n", *in_addr)

	defer sock.Close()

	for {
		var buf = make([]byte, 65536)
		rlen, _, err := sock.ReadFromUDP(buf)
		if err != nil {
			log.Fatalf("Error on sock.ReadFrom, %v", err)
		}
		if rlen == 0 {
			continue
		}

		request := &Request{}
		proto.Unmarshal(buf[0:rlen], request)
		fmt.Printf("%15s %30s: %3.2f %s\n",
			*request.ServerName,
			*request.ScriptName,
			*request.RequestTime,
			request.Tags(),
		)
		for _, timer := range GetTimers(request) {
			fmt.Printf("\t%s\n", timer)
		}
	}
}

func (request *Request) Tags() string {
	var tags bytes.Buffer
	if request.Status != nil {
		tags.WriteString(" status=")
		tags.WriteString(strconv.FormatInt(int64(*request.Status), 10))
	}
	for idx, val := range request.TagValue {
		tags.WriteString(" ")
		tags.WriteString(request.Dictionary[request.TagName[idx]])
		tags.WriteString("=")
		tags.WriteString(request.Dictionary[val])
	}
	return tags.String()
}

func GetTimers(request *Request) []string {
	offset := 0
	timers := make([]string, len(request.TimerValue))
	for idx, val := range request.TimerValue {
		var timer bytes.Buffer
		var cputime float64 = 0.0
		if len(request.TimerUtime) == len(request.TimerValue) {
			cputime = float64(request.TimerUtime[idx] + request.TimerStime[idx])
		}

		timer.WriteString("Val: ")
		timer.WriteString(strconv.FormatFloat(float64(val), 'f', 4, 64))
		timer.WriteString(" Hit: ")
		timer.WriteString(strconv.FormatInt(int64(request.TimerHitCount[idx]), 10))
		timer.WriteString(" CPU: ")
		timer.WriteString(strconv.FormatFloat(cputime, 'f', 4, 64))
		timer.WriteString(" Tags: ")

		for k, key_idx := range request.TimerTagName[offset : offset+int(request.TimerTagCount[idx])] {
			val_idx := request.TimerTagValue[int(offset)+k]
			if val_idx >= uint32(len(request.Dictionary)) || key_idx >= uint32(len(request.Dictionary)) {
				continue
			}
			timer.WriteString(" ")
			timer.WriteString(request.Dictionary[key_idx])
			timer.WriteString("=")
			timer.WriteString(request.Dictionary[val_idx])
		}

		timers[idx] = timer.String()
		offset += int(request.TimerTagCount[idx])
	}
	return timers
}
