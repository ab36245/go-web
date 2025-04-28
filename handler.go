package web

import (
	"net/http"

	"github.com/aivoicesystems/aivoice/common/websocket"
)

type requestFunc = func(http.ResponseWriter, *http.Request)

type socketFunc = func(websocket.Socket)

type Handler struct {
	get    requestFunc
	post   requestFunc
	socket socketFunc
}

func (h *Handler) Get(handler requestFunc) *Handler {
	h.get = handler
	return h
}

func (h *Handler) Post(handler requestFunc) *Handler {
	h.post = handler
	return h
}

func (h *Handler) Socket(handler socketFunc) *Handler {
	h.socket = handler
	return h
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if _, ok := r.Header["Upgrade"]; ok {
			h.doSocket(w, r)
		} else {
			h.doGet(w, r)
		}

	case "POST":
		h.doPost(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) doGet(w http.ResponseWriter, r *http.Request) {
	if h.get == nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	h.get(w, r)
}

func (h Handler) doPost(w http.ResponseWriter, r *http.Request) {
	if h.post == nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	h.post(w, r)
}

func (h Handler) doSocket(w http.ResponseWriter, r *http.Request) {
	if h.socket == nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	socket, err := websocket.Upgrade(w, r)
	if err != nil {
		// The websocket.Upgrade call should have sent an appropriate
		// error to the client.
		return
	}
	h.socket(socket)
}
