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

// ResponseAlert alert text
const ResponseAlert = "x-kubex-alert"

// ResponseAlertInfo alert text
const ResponseAlertInfo = "x-kubex-alert-info"

// ResponseAlertSuccess alert text
const ResponseAlertSuccess = "x-kubex-alert-success"

// ResponseAlertWarning alert text
const ResponseAlertWarning = "x-kubex-alert-warning"

// ResponseAlertDanger alert text
const ResponseAlertDanger = "x-kubex-alert-danger"

// ResponseAppendElement append element to the DOM
const ResponseAppendElement = "x-kubex-append-element"

// ResponseVanishElement vanish element from the DOM
const ResponseVanishElement = "x-kubex-vanish-element"
