package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

//func main() {
//	conn, err := tls.Dial("tcp", "218.76.62.118:3389", nil)
//	if err != nil {
//		log.Fatalln(err.Error())
//	}
//	log.Println("Client Connect To ", conn.RemoteAddr())
//	//status := conn.ConnectionState()
//	//fmt.Printf("%#v\n", status)
//	buf := make([]byte, 1024)
//	ticker := time.NewTicker(1 * time.Millisecond * 500)
//	for {
//		select {
//		case <-ticker.C:
//			{
//				log.Println("开始发送数据")
//				_, err = io.WriteString(conn, "GET / HTTP/1.1\nHost: www.baidu.com\r\n\r\n")
//				if err != nil {
//					log.Fatalln("错误节点一",err.Error())
//				}
//				log.Println("发送数据成功")
//				log.Println("开始接受数据")
//				len, err := conn.Read(buf)
//				if err != nil {
//					fmt.Println("错误节点二",err.Error())
//				} else {
//					fmt.Println("Receive From Server:", string(buf[:len]))
//				}
//				conn.Close()
//			}
//		}
//	}
//}

func main() {

	conn, err := net.DialTimeout("tcp", "218.76.62.118:3389", time.Second*2)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Client Connect To ", conn.RemoteAddr())
	buf := make([]byte, 1024)
	log.Println("开始发送数据")
	s := `\x16\x03\0\0\x69\x01\0\0\x65\x03\x03U\x1c\xa7\xe4random1random2random3random4\0\0\x0c\0/\0\x0a\0\x13\x009\0\x04\0\xff\x01\0\0\x30\0\x0d\0,\0*\0\x01\0\x03\0\x02\x06\x01\x06\x03\x06\x02\x02\x01\x02\x03\x02\x02\x03\x01\x03\x03\x03\x02\x04\x01\x04\x03\x04\x02\x01\x01\x01\x03\x01\x02\x05\x01\x05\x03\x05\x02`
	s = strings.Replace(s, `\0`, `\x00`, -1)
	s = Unicode2Str(s)
	_, err = io.WriteString(conn, s)
	if err != nil {
		log.Fatalln("错误节点一", err.Error())
	}
	log.Println("发送数据成功")
	log.Println("开始接受数据")
	length, err := conn.Read(buf)
	if err != nil {
		fmt.Println("错误节点二", err.Error())
	} else {
		fmt.Println("Receive From Server:", string(buf[:length]))
	}
	conn.Close()
}

func Unicode2Str(s string) string {
	length := len(s)
	res := ""
	for i := 0; i < length; i++ {
		if s[i:i+1] == "\\" {
			cHexStr := s[i+2 : i+4]
			c, _ := hex.DecodeString(cHexStr)
			cStr := string(c)
			i += 3
			res += cStr
			continue
		}
		res += s[i : i+1]
	}
	fmt.Println(res)
	return res
}
