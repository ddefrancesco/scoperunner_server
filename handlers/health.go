package handlers

import (
	"net/http"
	"strings"

	"github.com/go-http-utils/headers"
	"github.com/go-http-utils/negotiator"

	"github.com/dotse/go-health"
)

const (
	ContentType = "application/health+json"
)

// Handle serves a health response over HTTP. It supports the GET, HEAD and
// OPTIONS methods as well as content negotiation.
func HealthCommandHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodOptions:
		w.Header().Set(headers.Allow, strings.Join([]string{
			http.MethodGet,
			http.MethodHead,
			http.MethodOptions,
		}, ", "))
		w.WriteHeader(http.StatusNoContent)

		return

	case http.MethodGet, http.MethodHead:

	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	negotiator := negotiator.New(req.Header)

	if negotiator.Charset("UTF-8") == "" {
		http.Error(w, "", http.StatusNotAcceptable)
		return
	}

	ct := negotiator.Type(ContentType, "application/json")
	if ct == "" {
		http.Error(w, "", http.StatusNotAcceptable)
		return
	}

	w.Header().Set(headers.ContentEncoding, "UTF-8")
	w.Header().Set(headers.ContentType, ct)

	if req.Method == http.MethodHead {
		w.Header().Set(headers.ContentLength, "0")
		return
	}

	resp, err := health.CheckHealthContext(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if _, err := resp.Write(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
