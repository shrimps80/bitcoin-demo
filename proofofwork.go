package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const targetBits = 4 //4=1个0

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	// targetStr := "0001000000000000000000000000000000000000000000000000000000000000"
	// tmpBigInt := new(big.Int)
	// tmpBigInt.SetString(targetStr, 16)
	tmpBigInt := big.NewInt(1)
	tmpBigInt.Lsh(tmpBigInt, 256-targetBits)
	return &ProofOfWork{
		block:  block,
		target: tmpBigInt,
	}
}

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	var hashInt big.Int
	var nonce uint64
	var hash [32]byte

	for {
		fmt.Printf("%x\r", hash[:])
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])

		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if hashInt.Cmp(pow.target) == -1 {
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

func (pow *ProofOfWork) IsValid() bool {
	var hashInt big.Int

	data := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
