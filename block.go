package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	Version    uint64
	PrevHash   []byte
	MerkleRoot []byte
	TimeStamp  uint64
	Bits       uint64
	Nonce      uint64
	Hash       []byte
	Data       []byte
}

func Uint64ToBytes(num uint64) []byte {
	var buff bytes.Buffer
	if err := binary.Write(&buff, binary.BigEndian, num); err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func NewBlock(data string, PrevHash []byte) *Block {
	block := Block{
		Version:    0,
		PrevHash:   PrevHash,
		MerkleRoot: nil,
		TimeStamp:  uint64(time.Now().Unix()),
		Bits:       0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}

	// block.SetHash()
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return &block
}

func (b *Block) SetHash() {
	blockInfo := [][]byte{
		Uint64ToBytes(b.Version),
		b.PrevHash,
		b.MerkleRoot,
		Uint64ToBytes(b.TimeStamp),
		Uint64ToBytes(b.Bits),
		Uint64ToBytes(b.Nonce),
		b.Data,
	}

	blockInfoBytes := bytes.Join(blockInfo, []byte{})

	hash := sha256.Sum256(blockInfoBytes)
	b.Hash = hash[:]
}
