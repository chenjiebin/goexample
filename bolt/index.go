package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	// 在当前目录打开mydb.db文件
	// 如果不存在，则文件会被创建
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	defer db.Close()

	// 写数据
	err = db.Update(func(tx *bolt.Tx) error {
		// 创建一个bucket
		b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		err = b.Put([]byte("answer"), []byte("fun"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	// 读数据
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
}
