package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type blockController struct {
	beego.Controller
}
type Block struct {
	//区块高度
	Height int64
	//上一个区块的hash
	PrevBlockHash []byte
	//本区块的hash
	Hash []byte
	//交易数据
	Data []byte
	//时间戳
	Timestamp int64
}

func (b *blockController) NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	return &Block{Height: height, PrevBlockHash: prevBlockHash, Data: []byte(data), Timestamp: time.Now().Unix()}
}
