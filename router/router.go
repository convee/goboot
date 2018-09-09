package router

import (
	"net/http"
	"net/url"
	"regexp"
)

type Router struct {
	NotFoundHandler http.Handler
	handlers        []*Handler
}

type Handler struct {
	pattern        *regexp.Regexp
	allowedMethods map[string]struct{}
	http.Handler
}

func (r *Router) Handle(regex string, handler http.Handler) *Handler {
	pattern := regexp.MustCompile(regex)
	h := &Handler{pattern, nil, handler}
	r.handlers = append(r.handlers, h)
	return h
}

func (r *Router) HandleFunc(regex string, f func(http.ResponseWriter, *http.Request)) *Handler {
	return r.Handle(regex, http.HandlerFunc(f))
}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, h := range r.handlers {
		submatches := h.pattern.FindStringSubmatch(req.RequestURI)
		if len(submatches) == 0 {
			continue
		}
		if h.allowedMethods != nil {
			if _, ok := h.allowedMethods[req.Method]; !ok {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
		}
		names := h.pattern.SubexpNames()
		params := make(url.Values)
		for i := range names {
			params.Add(":"+names[i], submatches[i])
		}
		req.URL.RawQuery = params.Encode() + "&" + req.URL.RawQuery
		h.ServeHTTP(w, req)
		return
	}
	if r.NotFoundHandler != nil {
		r.NotFoundHandler.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func (h *Handler) method(m string) *Handler {
	if h.allowedMethods == nil {
		h.allowedMethods = make(map[string]struct{})
	}
	h.allowedMethods[m] = struct{}{}
	return h
}

func (h *Handler) Head() *Handler    { return h.method("HEAD") }
func (h *Handler) Get() *Handler     { return h.method("GET") }
func (h *Handler) Post() *Handler    { return h.method("POST") }
func (h *Handler) Put() *Handler     { return h.method("PUT") }
func (h *Handler) Delete() *Handler  { return h.method("DELETE") }
func (h *Handler) Options() *Handler { return h.method("OPTIONS") }
