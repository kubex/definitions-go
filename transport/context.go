package transport

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/kubex/definitions-go/app"
	"io"
	"net/textproto"
	"strconv"
	"strings"
	"time"
)

type RequestContext struct {
	WorkspaceID    string
	UserID         string
	signature      string
	TraceID        string
	UserIP         string
	UserAgent      string
	Authorization  []app.PermissionStatement
	Authentication json.RawMessage
}

func NewContext(headers map[string][]string) *RequestContext {
	c := &RequestContext{}
	c.ApplyHeaders(headers)
	return c
}

func NewContextFromRaw(rawHeaders io.Reader) (*RequestContext, error) {
	reader := textproto.NewReader(bufio.NewReader(rawHeaders))
	headers, err := reader.ReadMIMEHeader()
	if headers != nil && io.EOF != err && err != nil {
		return nil, err
	}
	return NewContext(headers), nil
}

func (r *RequestContext) ApplyHeaders(headers map[string][]string) {
	for k, vs := range headers {
		switch strings.ToLower(k) {
		case RequestWorkspaceID:
			r.WorkspaceID = vs[0]
		case RequestUserID:
			r.UserID = vs[0]
		case RequestSignature:
			r.signature = vs[0]
		case RequestTraceID:
			r.TraceID = vs[0]
		case RequestUserIP:
			r.UserIP = vs[0]
		case RequestUserAgent:
			r.UserAgent = vs[0]
		case RequestAuthentication:
			r.Authentication = json.RawMessage(vs[0])
		case RequestAuthorization:
			_ = json.Unmarshal([]byte(vs[0]), &r.Authorization)
		}
	}
}

func (r RequestContext) Verify(signatureKey string, maxTimeDiff int64) error {
	if r.signature != "" && strings.Contains(r.signature, "/") {
		splits := strings.SplitN(r.signature, "/", 2)
		timestamp, _ := strconv.ParseInt(splits[1], 10, 64)

		now := time.Now().Unix()
		if timestamp > (now+maxTimeDiff) || timestamp < (now-maxTimeDiff) {
			return errors.New("signature outside of available time window")
		}

		verifyString := ""
		verifyString += r.WorkspaceID
		verifyString += r.UserID
		verifyString += signatureKey
		verifyString += r.TraceID
		verifyString += r.UserIP
		verifyString += r.UserAgent

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

func (r RequestContext) HasPermission(perm app.ScopedKey) bool {
	for _, statement := range r.Authorization {
		if perm.Key == statement.Permission.Key && statement.Effect == app.PermissionEffectAllow {
			return true
		}
	}
	return false
}

func (r RequestContext) HasResourcePermission(perm app.ScopedKey, resource string) bool {
	for _, statement := range r.Authorization {
		if perm.Key == statement.Permission.Key && statement.Effect == app.PermissionEffectAllow {
			if statement.Resource == app.PermissionResourceAll || resource == statement.Resource {
				return true
			}
			if strings.HasSuffix(statement.Resource, app.PermissionResourceAll) &&
				strings.HasPrefix(statement.Resource, statement.Resource[:len(statement.Resource)-1]) {
				return true
			}
		}
	}
	return false
}

func (r RequestContext) PermittedResources(perm app.ScopedKey) []string {
	var resources []string
	for _, statement := range r.Authorization {
		if perm.Key == statement.Permission.Key && statement.Effect == app.PermissionEffectAllow {
			resources = append(resources, statement.Resource)
			break
		}
	}
	return resources
}
func (r RequestContext) DeniedResources(perm app.ScopedKey) []string {
	var resources []string
	for _, statement := range r.Authorization {
		if perm.Key == statement.Permission.Key && statement.Effect == app.PermissionEffectDeny {
			resources = append(resources, statement.Resource)
			break
		}
	}
	return resources
}
