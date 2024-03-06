package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	targetStr := "0001000000000000000000000000000000000000000000000000000000000000"
	tmpBigInt := new(big.Int)
	tmpBigInt.SetString(targetStr, 16)

	return &ProofOfWork{
		block:  block,
		target: tmpBigInt,
	}
}

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	var nonce uint64
	var hash [32]byte
	fmt.Println("start...")

	for {
		fmt.Printf("%x\r", hash[:])
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)

		tmpInt := new(big.Int)
		tmpInt.SetBytes(hash[:])

		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("success, hash :%x, nonce :%d\n", hash[:], nonce)
			break
		} else {
			nonce++
		}
	}

	return hash[:], nonce
}

func (pow *ProofOfWork) PrepareData(nonce uint64) []byte {
	b := pow.block

	tmp := [][]byte{
		Uint64ToBytes(b.Version), //将uint64转换为[]byte
		b.PrevHash,
		b.MerkleRoot,
		Uint64ToBytes(b.TimeStamp),
		Uint64ToBytes(b.Bits),
		Uint64ToBytes(nonce),
		// b.Hash,
		// b.Data,
	}
	data := bytes.Join(tmp, []byte{})
	return data
}
