package transport

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
