package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("t-ci-env-spike start")
	fmt.Printf("HOGE=%s\n", os.Getenv("HOGE"))
	fmt.Printf("FUGA (encrpted)=%s\n", os.Getenv("FUGA"))
	fmt.Println("t-ci-env-spike end")
}
