package main

import (
	"github.com/ipid/chatroom-backend/def"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func getDices(requests []def.DiceRequest) (dices []def.Dice) {
	for _, request := range requests {
		for i := 0; i < request.Num; i++ {
			dices = append(dices, def.Dice{
				Max:   request.Max,
				Value: rand.Intn(request.Max) + 1,
			})
		}
	}

	return
}
