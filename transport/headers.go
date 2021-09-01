package transport

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"
)

// RequestWorkspaceID Workspace UUID
const RequestWorkspaceID = "x-kx-workspace-id"

// RequestUserID User UID
const RequestUserID = "x-kx-user-id"

// RequestSignature Kubex Signature
const RequestSignature = "x-kx-signature"

// RequestTraceID Kubex request Trace ID
const RequestTraceID = "x-kx-trace-id"

// RequestUserIP User IP
const RequestUserIP = "x-kx-user-ip"

// RequestUserAgent User Agent
const RequestUserAgent = "x-kx-user-agent"

// RequestAuthorization Json Authorizations, [{"k":"permission-key","e":"A","r":"*"},"k":"permission-key-2","e":"D","r":"uid1"]
const RequestAuthorization = "x-kx-authorization"

// RequestAuthentication JSON access credentials, provided by the app e.g. {"accessToken":"xx"}
const RequestAuthentication = "x-kx-authentication"

// ResponseUri Uri to set in the address bar for the current request
const ResponseUri = "x-kubex-uri"

// ResponseDebug Debug object for the browser
const ResponseDebug = "x-kubex-debug"

// ResponseZeroPad When set to true, padding will be removed for the container
const ResponseZeroPad = "x-kubex-zeropad"

func Verify(headers map[string]string, signatureKey string, maxTimeDiff int64) error {

	if sig, ok := headers[RequestSignature]; ok && strings.Contains(sig, "/") {
		splits := strings.SplitN(sig, "/", 2)
		timestamp, _ := strconv.ParseInt(splits[1], 10, 64)

		now := time.Now().Unix()
		if timestamp > (now+maxTimeDiff) || timestamp < (now-maxTimeDiff) {
			return errors.New("signature outside of available time window")
		}

		verifyString := ""
		if headerValue, ok := headers[RequestWorkspaceID]; ok {
			verifyString += headerValue
		}
		if headerValue, ok := headers[RequestUserID]; ok {
			verifyString += headerValue
		}
		verifyString += signatureKey
		if headerValue, ok := headers[RequestTraceID]; ok {
			verifyString += headerValue
		}
		if headerValue, ok := headers[RequestUserIP]; ok {
			verifyString += headerValue
		}
		if headerValue, ok := headers[RequestUserAgent]; ok {
			verifyString += headerValue
		}

		verifyString += splits[1]

		signature := sha256.New()
		signature.Write([]byte(verifyString))

		if splits[0] == hex.EncodeToString(signature.Sum(nil)) {
			return nil
		}
		return errors.New("unable to verify signature")

	}
	return errors.New("invalid or missing signature")

}
