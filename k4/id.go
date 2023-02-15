package k4

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

const spacer = "-"
const checksumSize = 2
const verifyLength = 4

var baseTime = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

// Support dash removal if required
const p1Len = 6
const p2Len = 11

var minVerifyLen int
var hostId string

func init() {
	hostId = RandomString(2)
	minVerifyLen = p1Len + p2Len + len(spacer)
}

type ID struct {
	UniqueKey    string
	Verification string
}

func (i ID) Full() string {
	return strings.TrimRight(i.UniqueKey+spacer+i.Verification, spacer)
}

func (i ID) UID() string {
	return i.UniqueKey
}

func IDFromString(input string) ID {
	i := ID{}
	//random-time-verify
	exploded := strings.Split(input, spacer)
	expLen := len(exploded)
	lastPart := expLen
	if expLen > 2 && len(input) > minVerifyLen {
		lastPart = expLen - 1
		i.Verification = exploded[expLen-1]
		if len(i.Verification) != verifyLength+checksumSize {
			i.Verification = ""
		}
	}
	i.UniqueKey = strings.Join(exploded[:lastPart], spacer)

	return i
}

func (i ID) HasValidVerification() bool {
	if i.Verification == "" || len(i.Verification) != (verifyLength+checksumSize) {
		return false
	}

	preVerify := i.Verification[:verifyLength]
	return i.Verification == preVerify+checkSum(preVerify+i.UniqueKey)
}

func NewID() ID {
	i := ID{}
	i.UniqueKey = randomID()
	i.Verification = RandomString(verifyLength)
	i.Verification = i.Verification + checkSum(i.Verification+i.UniqueKey)
	return i
}

func checkSum(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	return fmt.Sprintf("%x", hasher.Sum(nil))[:checksumSize]
}

func RandomString(n int) string {
	b := make([]byte, n+5)
	rand.Seed(time.Now().UnixNano())
	rand.Read(b)
	return strings.ReplaceAll(
		strings.ReplaceAll(base64.RawStdEncoding.EncodeToString(b), "+", ""),
		"/", "",
	)[:n]
}

func randomID() string {
	return fixLen(RandomString(p1Len), p1Len) + spacer +
		fixLen(hostId+timeID(), p2Len)
}

func timeID() string {
	var i big.Int
	now := time.Now()
	now.Sub(baseTime)
	i.SetInt64(int64(math.Floor(float64(now.UnixNano() / 1000))))
	return i.Text(62)
}

func fixLen(input string, reqLen int) string {
	srcLen := len(input)
	if srcLen == reqLen {
		return input
	}
	if srcLen > reqLen {
		return input[:reqLen]
	}
	return input + strings.Repeat("X", reqLen-srcLen)
}
