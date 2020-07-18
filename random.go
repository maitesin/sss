package main

import (
	"crypto/rand"
	"math/big"
)

var max *big.Int

func init()  {
	//Max random value, a 130-bits integer, i.e 2^130 - 1
	max = new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
}

func randomBigInt() (*big.Int, error) {
	return rand.Int(rand.Reader, max)
}
