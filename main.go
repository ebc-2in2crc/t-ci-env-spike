package main

import (
	"fmt"
	"os"
)

func main() {
	// コメント
	fmt.Println("t-ci-env-spike start")
	fmt.Printf("HOGE=%s\n", os.Getenv("HOGE"))
	fmt.Printf("FUGA=%s\n", os.Getenv("FUGA"))
	fmt.Printf("HOGERA (encrypted)=%s\n", os.Getenv("HOGERA"))
	fmt.Println("t-ci-env-spike end")
}
