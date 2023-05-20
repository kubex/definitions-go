package k4

import (
	"math/big"
	"strings"
	"time"
)

const checksumSize = 1
const defaultIDLength = 15
const defaultTimeGenerator = TimeGeneratorMicro

var globalIDHost IDGenerator

type TimeGenerator int

const (
	//TimeGeneratorNano 11 Char string
	TimeGeneratorNano TimeGenerator = iota
	// TimeGeneratorMicro 9 Char string
	TimeGeneratorMicro
	// TimeGeneratorMilli 7-8 Char string
	TimeGeneratorMilli
	// TimeGeneratorSecond 6 Char string
	TimeGeneratorSecond
	// TimeGeneratorMinute 5 Char string
	TimeGeneratorMinute
	// TimeGeneratorHour 4 Char string
	TimeGeneratorHour
	// TimeGeneratorDay 3 Char string
	TimeGeneratorDay
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
	case TimeGeneratorMinute:
		i.SetInt64(now.Unix() / 60)
	case TimeGeneratorHour:
		i.SetInt64(now.Unix() / 3600)
	case TimeGeneratorDay:
		i.SetInt64(now.Unix() / 86400)
	}
	return i.Text(62)
}

func init() {
	globalIDHost = DefaultIDGenerator()
}

type IDGenerator struct {
	hostID       string
	hostIDLength int
	idLength     int
	timeSize     TimeGenerator
	generation   chan bool
}

func DefaultIDGenerator() IDGenerator {
	h := IDGenerator{
		idLength: defaultIDLength,
		timeSize: defaultTimeGenerator,
	}
	h.randomHostID()
	h.generation = make(chan bool, 1)
	return h
}

func (h *IDGenerator) SetHostID(id string) {
	h.hostID = id
	h.hostIDLength = len(h.hostID)
}

func (h *IDGenerator) GetHostID() string              { return h.hostID }
func (h *IDGenerator) randomHostID()                  { h.SetHostID(RandomString(2)) }
func (h *IDGenerator) SetBaseLength(size int)         { h.idLength = size }
func (h *IDGenerator) SetTimeSize(size TimeGenerator) { h.timeSize = size }

func (h *IDGenerator) New() ID {
	if h.hostID == "" {
		h.randomHostID()
	}

	h.generation <- true
	i := ID{}
	i.uniqueKey = h.randomID()
	i.verification = i.checkSum(i.uniqueKey)
	if h.timeSize == TimeGeneratorNano && h.idLength < 15 {
		//Sleep for a nanosecond to ensure uniqueness when precision is needed
		time.Sleep(time.Nanosecond)
	}
	<-h.generation
	return i
}

func (h *IDGenerator) randomID() string {
	tId := h.reverse(h.timeSize.timeID())
	useLen := h.idLength - len(tId) - len(h.hostID)
	return h.fixLen(tId+h.hostID+RandomString(useLen), h.idLength)
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
