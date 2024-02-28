package websocket

type Server struct {
	clients     map[string]*Client
	client_refs map[string]string
	rooms       map[int]*Room
}

func NewServer() *Server {
	return &Server{
		clients:     make(map[string]*Client),
		client_refs: make(map[string]string),
		rooms:       make(map[int]*Room),
	}
}