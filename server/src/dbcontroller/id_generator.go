package dbcontroller

import (
	"math/rand"
)

const idLen = 16
const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type IDGenerator struct {
	idLen      int
	charset    string
	charsetLen int
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{
		idLen:      idLen,
		charset:    charset,
		charsetLen: len(charset),
	}
}

func (g *IDGenerator) generate() string {
	id := ""
	var charIdx int
	var randChar byte
	for i := 0; i < g.idLen; i++ {
		charIdx = rand.Intn(g.charsetLen)
		randChar = g.charset[charIdx]
		id += string(randChar)
	}

	return id
}