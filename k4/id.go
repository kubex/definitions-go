package k4

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"
)

const spacer = "-"
const checksumSize = 4
const verifyLength = 10

// Support dash removal if required
const p1Len = 6
const p2Len = 5
const p3Len = 5
const p4Len = 9

var minVerifyLen int
var hostId string

func init() {
	hostId = RandomString(3)
	minVerifyLen = p4Len + p2Len + p1Len + p3Len + (len(spacer) * 4)
}

type ID struct {
	Prefix       string
	UniqueKey    string
	Verification string
}

func (i ID) Full() string {
	return strings.TrimRight(i.Prefix+spacer+i.UniqueKey+spacer+i.Verification, spacer)
}

func (i ID) Compact() string {
	return i.Prefix + spacer + i.UniqueKey
}

func (i ID) UID() string {
	return i.UniqueKey
}

func IDFromString(input string) ID {
	i := ID{}

	exploded := strings.Split(input, spacer)
	i.Prefix = exploded[0]
	expLen := len(exploded)
	lastPart := expLen
	hasVerify := expLen > 2 && len(input) > minVerifyLen+len(i.Prefix)
	if hasVerify {
		lastPart = expLen - 1
		i.Verification = exploded[expLen-1]
		if len(i.Verification) != verifyLength+checksumSize {
			i.Verification = ""
		}
	}
	i.UniqueKey = strings.Join(exploded[1:lastPart], spacer)

	return i
}

func (i ID) HasValidVerification() bool {
	if i.Verification == "" || len(i.Verification) != (verifyLength+checksumSize) {
		return false
	}

	preVerify := i.Verification[:verifyLength]
	return i.Verification == strings.ToLower(preVerify+checkSum(preVerify+i.UniqueKey))
}

func NewID(prefix string) ID {
	i := ID{}
	if prefix == "" {
		prefix = "K4"
	}
	i.Prefix = strings.ToLower(prefix)
	i.UniqueKey = strings.ToLower(base36Time())
	i.Verification = strings.ToLower(RandomString(verifyLength))
	i.Verification = i.Verification + strings.ToLower(checkSum(i.Verification+i.UniqueKey))
	return i
}

func checkSum(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	return fmt.Sprintf("%x", hasher.Sum(nil))[:checksumSize]
}

func RandomString(n int) string {
	b := make([]byte, n+3)
	_, err := rand.Read(b)

	if err != nil {
		now := time.Now()
		mrand.Seed(now.UnixNano())
		shuffled := []rune(strings.ReplaceAll(strconv.FormatInt(now.UnixNano(), 36)+now.Format("20060102150405.999999999"), ".", ""))
		mrand.Shuffle(len(shuffled), func(i, j int) {
			shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
		})
		return string(shuffled)[:n]
	}

	return base32.StdEncoding.EncodeToString(b)[:n]
}
func base36Time() string {
	now := time.Now()
	timeKey := ""
	intString := strconv.FormatInt(int64(now.Nanosecond()+now.Second()), 36)
	for x := len(intString); x > 0; x-- {
		timeKey += string(intString[x-1])
	}

	return fixLen(timeKey, p1Len) + spacer +
		RandomString(p2Len) + spacer +
		RandomString(p3Len) + spacer +
		fixLen(hostId+strconv.FormatInt(now.Unix(), 36), p4Len)
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
