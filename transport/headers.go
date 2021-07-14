package transport

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"
)

const HeaderWorkspaceID = "x-kx-workspace-id" // Workspace UUID
const HeaderUserID = "x-kx-user-id"           // User UID
const HeaderSignature = "x-kx-signature"      // Kubex Signature
const HeaderTraceID = "x-kx-trace-id"         // Kubex request Trace ID

const HeaderUserIP = "x-kx-user-ip"       // User IP
const HeaderUserAgent = "x-kx-user-agent" // User Agent

const HeaderAuthorization = "x-kx-authorization"   // Json Authorizations, [{"k":"permission-key","e":"A","r":"*"},"k":"permission-key-2","e":"D","r":"uid1"]
const HeaderAuthentication = "x-kx-authentication" // JSON access credentials, provided by the app e.g. {"accessToken":"xx"}

func Verify(headers map[string]string, signatureKey string, maxTimeDiff int64) error {

	if sig, ok := headers[HeaderSignature]; ok && strings.Contains(sig, "/") {
		splits := strings.SplitN(sig, "/", 2)
		timestamp, _ := strconv.ParseInt(splits[1], 10, 64)

		now := time.Now().Unix()
		if timestamp > (now+maxTimeDiff) || timestamp < (now-maxTimeDiff) {
			return errors.New("signature outside of available time window")
		}

		verifyString := ""
		if headerValue, ok := headers[HeaderWorkspaceID]; ok {
			verifyString += headerValue
		}
		if headerValue, ok := headers[HeaderUserID]; ok {
			verifyString += headerValue
		}
		verifyString += signatureKey
		if headerValue, ok := headers[HeaderTraceID]; ok {
			verifyString += headerValue
		}
		if headerValue, ok := headers[HeaderUserIP]; ok {
			verifyString += headerValue
		}
		if headerValue, ok := headers[HeaderUserAgent]; ok {
			verifyString += headerValue
		}

		verifyString += splits[1]

		signature := sha256.New()
		signature.Write([]byte(verifyString))

		if sig == hex.EncodeToString(signature.Sum(nil)) {
			return nil
		}
		return errors.New("unable to verify signature")

	}
	return errors.New("invalid or missing signature")

}
