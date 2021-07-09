package main

import (
	"fmt"
	"log"
	"net"
)

const (
	Ping           = "Ping"
	Pong           = "Pong"
	Connect        = "Connect"
	FindNode       = "FindNode"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 25026})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("本地地址: <%s> \n", listener.LocalAddr().String())
	peers := make([]net.UDPAddr, 0, 2)
	//回复ping
	Pongs := make(chan *net.UDPAddr)

	go sendPong(listener, Pongs)

	for {
		data := make([]byte, 1024)
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])
		switch string(data[:n]){
		case Ping :
			Pongs <- remoteAddr
		case Connect :
			peers = append(peers, *remoteAddr)
		case FindNode:
			peerstr := ""
			for _,item := range peers{
				peerstr = fmt.Sprintf("%s,%s",peerstr,item.String())
			}
			listener.WriteToUDP([]byte(peerstr), remoteAddr)
		}
		//if len(peers) == 2 {
		//	log.Printf("进行UDP打洞,建立 %s <--> %s 的连接\n", peers[0].String(), peers[1].String())
		//	listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
		//	listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
		//	time.Sleep(time.Second * 8)
		//	log.Println("中转服务器退出,仍不影响peers间通信")
		//	return
		//}
	}
}



//
func sendPong(conn *net.UDPConn, Pongs <-chan *net.UDPAddr) {
	for addr := range Pongs {
		_, _ = conn.WriteToUDP([]byte(Pong), addr)
	}
}
