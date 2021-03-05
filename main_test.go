package main

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	ti, _ := time.Parse(time.RFC3339Nano, "2020-12-18T12:21:06.637698261Z")
	fmt.Println(ti.Unix())
}
