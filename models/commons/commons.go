package commons

type ScopeRequest struct {
	Body string `json:"body"`
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
	Cmd            string `json:"cmd"`
}
