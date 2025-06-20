package something

import (
	"math/rand"
	"time"
)

type Service struct {
}

func (this *Service) Process(value int) int {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Minute)
	return value * 2
}
