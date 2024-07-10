package main

import (
	"server/pool"
	"testing"
)

func TestInit(t *testing.T) {
	err := pool.Init()
	if err != nil {
		t.Fatal(err)
	}
}
