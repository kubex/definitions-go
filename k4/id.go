package k4

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
)

func RandomString(n int) string {
	//return "aaaaalikfewhflkwheklfhewlkfhewlkfhwelkhfklwehflkew"[:n]
	if n < 1 {
		return ""
	}
	b := make([]byte, n+5)
	rand.Read(b)
	return strings.ReplaceAll(
		strings.ReplaceAll(base64.RawStdEncoding.EncodeToString(b), "+", ""),
		"/", "",
	)[:n]
}

func NewID() ID { return globalIDHost.New() }

func IDFromString(input string) ID {
	i := ID{}
	i.verification = input[:checksumSize]
	i.uniqueKey = input[checksumSize:]
	return i
}

type ID struct {
	uniqueKey    string
	verification string
}

func (i ID) String() string {
	return strings.TrimSpace(i.verification + i.uniqueKey)
}

func (i ID) UID() string {
	return i.uniqueKey
}

func (i ID) IsValid() bool {
	if i.verification == "" || len(i.verification) != checksumSize {
		return false
	}

	return i.verification == i.checkSum(i.uniqueKey)
}

func (i ID) checkSum(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	return fmt.Sprintf("%x", hasher.Sum(nil))[:checksumSize]
}
