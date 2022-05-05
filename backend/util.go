package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func getDices(dNum, dMax int) (dices []int) {
	dices = make([]int, dNum)

	for i := 0; i < dNum; i++ {
		dices[i] = rand.Intn(dMax) + 1
	}

	return
}
