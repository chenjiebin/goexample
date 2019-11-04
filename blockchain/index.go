package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

// 块实例
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// 块设置hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// 快实例序列化对象
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(b); err != nil {
		log.Fatalln("encode error:", err)
	}
	return result.Bytes()
}

// 对块进行反序列化
func DeserializeBlock(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	if err := decoder.Decode(&block); err != nil {
		log.Fatalln("decode error:", err)
	}
	return &block
}

// 创建新块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// 区块链数据结构
// 此处开始有区块链的数据结构和相应方法
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

// 增加新块
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlocksBucket"))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	// prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlocksBucket"))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash
		if err != nil {
			log.Fatal(err)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	// bc.blocks = append(bc.blocks, newBlock)
}

// 初始化创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// 实例化一个区块链
func NewBlockchain() *Blockchain {
	var tip []byte
	dbFile := "blockchain.db"
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlocksBucket"))
		// 如果bucket不存在，则进行初始化创世区块，并创建新bucket
		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte("BlocksBucket"))
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Fatal(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

// 主函数
func main() {
	// bc := NewBlockchain()

	// bc.AddBlock("Send 1 BTC to Ivan")
	// bc.AddBlock("Send 2 more BTC to Ivan")

	// bolt遍历输出
	dbFile := "blockchain.db"
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		fmt.Println("111")
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlocksBucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key = %s, value = %s\n", k, v)
			fmt.Println()

			block := DeserializeBlock(v)
			fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)

			pow := NewProofOfWork(block)
			fmt.Printf("Pow:%s\n", strconv.FormatBool(pow.Validate()))
			fmt.Println()
		}
		return nil
	})

	// 没有引入bolt前的输出方式
	// for _, block := range bc.blocks {
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)

	// 	pow := NewProofOfWork(block)
	// 	fmt.Printf("Pow:%s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }
}

// 此处开始挖矿部分
const targetBits = 12

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	maxNonce := math.MaxInt64

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		}
		nonce++
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

// 整形转换成16进制输出
func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}
