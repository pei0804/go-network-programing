package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)

const BOM = '\ufffe'

func main() {
	service := "0.0.0.0:1210"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		str := "j'ai arrêté"
		shorts := utf16.Encode([]rune(str))
		writeShorts(conn, shorts)

		conn.Close() // we're finished
	}
}

// &    ビット演算 and          整数
// |    ビット演算 or           整数
// ^    ビット演算 xor          整数
// &^   ビットクリア(and not)   整数
//
// <<   左シフト                整数 << 符号なし整数
// >>   右シフト                整数 >> 符号なし整数
func writeShorts(conn net.Conn, shorts []uint16) {
	var bytes [2]byte
	// send the BOM as first two bytes
	// '\ufffe' なら255になる
	bytes[0] = BOM >> 8
	// '\ufffe' なら254になる
	bytes[1] = BOM & 255
	_, err := conn.Write(bytes[0:])
	if err != nil {
		return
	}
	for _, v := range shorts {
		bytes[0] = byte(v >> 8)
		bytes[1] = byte(v & 255)
		_, err = conn.Write(bytes[0:])
		if err != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
