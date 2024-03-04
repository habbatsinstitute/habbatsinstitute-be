package websocket

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Server struct {
	clients     map[string]*Client
	client_refs map[string]string
	Rooms       map[int]*Room
}

func NewServer() *Server {
	return &Server{
		clients:     make(map[string]*Client),
		client_refs: make(map[string]string),
		Rooms:       make(map[int]*Room),
	}
}

func (s *Server) FindClient(ref string) *Client {
	sign := s.client_refs[ref]
	return s.clients[sign]
}

func (s *Server) FindRoom(sign int) *Room {
	return s.Rooms[sign]
}

func (s *Server) CreateClient(ctx echo.Context, user int, role int) string {
	ref, client := NewClient(ctx, s, user, role)
	s.clients[client.sign] = client
	s.client_refs[ref] = client.sign
	go s.clients[client.sign].Send()
	go s.clients[client.sign].Recv()
	return client.sign
}

func (s *Server) CreateRoom(id int, sign int, refs ...string) *Room {
	clients := func() (buffer []*Client) {
		for _, ref := range refs {
			buffer = append(buffer, s.FindClient(ref))
		}
		return buffer
	}()
	s.Rooms[sign] = NewRoom(id, sign, clients...)
	go s.Rooms[sign].Listen()
	return s.Rooms[sign]
}

func (s *Server) JoinRoom(sign int, ref string) {
	s.Rooms[sign].join <- s.FindClient(ref)
}

func (s *Server) DeleteClient(sign string) {
	for i, room := range s.clients[sign].rooms {
		logrus.Infof("[ws.server]: client@%s keluar dari room %d", sign, i)
		room.leave <- s.clients[sign]
	}
	delete(s.clients, sign)
}

func (s *Server) DeleteRoom(sign int) {
	for client := range s.Rooms[sign].clients {
		logrus.Infof("[ws.server]: client@%s keluar dari room %d", client.sign, sign)
		s.Rooms[sign].leave <- client
	}
	// close(s.rooms[sign].join)
	// close(s.rooms[sign].leave)
	close(s.Rooms[sign].message)
	delete(s.Rooms, sign)
}