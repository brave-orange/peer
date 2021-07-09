package main

type DataBody struct {
	Type string  //消息类型
	Code string  //消息code,同一个code是同一个消息
	Size int64
	data []byte
}
