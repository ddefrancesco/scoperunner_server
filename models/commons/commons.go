package commons

type RequestAddress struct {
	Address string `json:"address"`
}
type ScopeRequest struct {
	Body string `json:"body"`
}

type ScopeSetRequest struct {
	Body map[string]string `json:"body"`
}

type ScopeInitRequest struct {
	Body map[string]map[string]string `json:"body"`
}
type ScopeResponse struct {
	Code     int    `json:"code"`
	Response string `json:"response"`
	Cmd      string `json:"cmd"`
}

type ScopeErr struct {
	Err            int    `json:"error_code"`
	ErrDescription string `json:"error_description"`
	ScopeFunction  string `json:"scope_function"`
	Cmd            any    `json:"cmd"`
}
