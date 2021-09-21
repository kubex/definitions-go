package transport

import (
	"bytes"
	"testing"
)

func TestRequestContextFromRaw(t *testing.T) {
	raw := bytes.NewReader([]byte("Host: connectors.chargehive.cubex-local.com:8873\nUser-Agent: Kubex Rubix/1.0\nContent-Type: \nX-Kx-Authentication: {\"chive-access-token\":\"dGVzdC1wcm9qZWN0OmQzNDdmZDI3LTFiMTMtNGE0Mi1hNDZiLTBlYTA1MzUzZjFiOTpDN0RCQ1lVVzdLMFdeNzkxYWVmYjc0YmJhNGU2MjllNzg=\",\"chive-project-id\":\"test-project\"}\nX-Kx-Authorization: [{\"effect\":\"Allow\",\"permission\":{\"Key\":\"view-configuration\"},\"resource\":\"*\"}]\nX-Kx-Signature: bb557ad1806c359a6ac9d52046c7fd6dcb2e216d60312f516182e8482d03cf75/1632220779\nX-Kx-Trace-Id: 32eb2a12-9ab4-4dd6-8143-4388bb982ee3\nX-Kx-User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36\nX-Kx-User-Id: EFIDFIID-ZFA8TK5L6-MISCR-JY6QKDP\nX-Kx-User-Ip: 127.0.0.1\nX-Kx-Workspace-Id: 6daeb9c0-898a-4ee8-bc1e-96ac82e9fe6b\nX-Requested-With: XMLHttpRequest\nAccept-Encoding: gzip"))
	r, err := NewContextFromRaw(raw)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	if r.WorkspaceID == "" {
		t.Fatal("Unable to decode headers")
	}
}
