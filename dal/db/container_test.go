package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

func TestGetContainer(t *testing.T) {
	c, err := GetContainer(ctx, 1, "")
	if err != nil {
		logrus.Error(err)
		t.FailNow()
	}
	for _, i := range c {
		fmt.Println(*i)
	}
}

func TestDeleteContainer(t *testing.T) {
	err := DeleteContainer(ctx, 2, "3")
	if err != nil {
		t.FailNow()
	}
}
