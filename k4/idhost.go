package k4

import (
	"math/big"
	"strings"
	"time"
)

const checksumSize = 2
const defaultIDLength = 14
const defaultTimeGenerator = TimeGeneratorMicro

var globalIDHost IDGenerator

type TimeGenerator int

const (
	TimeGeneratorNano TimeGenerator = iota
	TimeGeneratorMicro
	TimeGeneratorMilli
	TimeGeneratorSecond
)

func (t TimeGenerator) timeID() string {
	var i big.Int
	now := time.Now()
	switch t {
	case TimeGeneratorNano:
		i.SetInt64(now.UnixNano())
	case TimeGeneratorMicro:
		i.SetInt64(now.UnixMicro())
	case TimeGeneratorMilli:
		i.SetInt64(now.UnixMilli())
	case TimeGeneratorSecond:
		i.SetInt64(now.Unix())
	}
	return i.Text(62)
}

func init() {
	globalIDHost = IDGen()
}

type IDGenerator struct {
	HostID       string
	hostIDLength int
	idLength     int
	timeSize     TimeGenerator
	generation   chan bool
}

func IDGen() IDGenerator {
	h := IDGenerator{
		idLength: defaultIDLength,
		timeSize: defaultTimeGenerator,
	}
	h.genHostID()
	h.generation = make(chan bool, 1)
	return h
}

func (h *IDGenerator) genHostID() {
	h.HostID = RandomString(2)
	h.hostIDLength = len(h.HostID)
}

func (h *IDGenerator) SetBaseLength(size int)         { h.idLength = size }
func (h *IDGenerator) SetTimeSize(size TimeGenerator) { h.timeSize = size }

func (h *IDGenerator) New() ID {
	if h.HostID == "" {
		h.genHostID()
	}

	h.generation <- true
	i := ID{}
	i.uniqueKey = h.randomID()
	i.verification = i.checkSum(i.uniqueKey)
	time.Sleep(time.Nanosecond)
	<-h.generation
	return i
}

func (h *IDGenerator) randomID() string {
	//tId := h.reverse(h.timeSize.timeID())
	tId := h.timeSize.timeID()
	useLen := h.idLength - len(tId) - len(h.HostID)
	return h.fixLen(tId+h.HostID+RandomString(useLen), h.idLength)
}

func (h *IDGenerator) reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (h *IDGenerator) fixLen(input string, reqLen int) string {
	srcLen := len(input)
	if srcLen == reqLen {
		return input
	}
	if srcLen > reqLen {
		return input[:reqLen]
	}
	return input + strings.Repeat("X", reqLen-srcLen)
}
