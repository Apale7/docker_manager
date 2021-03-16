package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	b, _ := os.ReadFile("./tmp/dockerfile")
	fmt.Println(string(b))
	bs, err := json.Marshal(string(b))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}
