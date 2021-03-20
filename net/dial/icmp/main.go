package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() { // 用ICMP协议向在线的主机发送一个问候，并等待主机返回
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host")
		os.Exit(1)
	}
	service := os.Args[1]                      // address IP地址或域名
	conn, err := net.Dial("ip4:icmp", service) // 网络协议的名字
	checkError(err)
	var msg [512]byte
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum
	msg[3] = 0  // checksum
	msg[4] = 0  // identifier[0]
	msg[5] = 13 // identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[1]
	len := 8
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)
	_, err = conn.Write(msg[0:len])
	checkError(err)
	_, err = conn.Read(msg[0:])
	checkError(err)
	fmt.Println("Got response")
	if msg[5] == 13 {
		fmt.Println("Identifier = matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}
	os.Exit(0)
}

func checkError(err error) { // 对err的封装
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0
	n := 0
	for n+1 < len(msg) {
		sum += (int(msg[n]) << 8) | int(msg[n+1])
		n++
	}
	if n < len(msg) {
		sum += (int(msg[n]) << 8)
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum)
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
