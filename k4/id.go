package k4

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

func RandomString(n int) string {
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
	if len(input) < checksumSize {
		return i
	}
	i.verification = input[:checksumSize]
	i.uniqueKey = input[checksumSize:]
	return i
}

func IDFromUUID(input string) ID {
	rawID, _ := hex.DecodeString(strings.ReplaceAll(input, "-", ""))
	return IDFromString(strings.ReplaceAll(encodeB63(rawID), "?", ""))
}

type ID struct {
	uniqueKey    string
	verification string
}

func (i ID) String() string {
	return strings.TrimSpace(i.verification + i.uniqueKey)
}

func (i ID) UUID() string {
	if len(i.String()) > 21 {
		// IDs generated that are over 21 characters long, have the potential to be invalid uuids
		return ""
	}

	var rawID []byte
	var uuid string
	for xi := 20; xi < 24; xi++ {
		if len(uuid) < 32 {
			rawID, _ = decodeB63((i.String() + "??????????????????????")[:xi])
			uuid = hex.EncodeToString(rawID)
		}
	}

	return fmt.Sprintf("%s-%s-%s-%s-%s", uuid[:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:32])
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
