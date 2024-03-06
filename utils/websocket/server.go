package websocket

import (
	realtimechat "institute/features/realtime_chat"

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

func (s *Server) FindRoom(roomId int) *Room {
	return s.Rooms[roomId]
}

func (s *Server) CreateClient(ctx echo.Context, user int, role string) string {
	ref, client := NewClient(ctx, s, user, role)
	s.clients[client.sign] = client
	s.client_refs[ref] = client.sign
	go s.clients[client.sign].Send()
	go s.clients[client.sign].Recv()
	return client.sign
}

func (s *Server) CreateRoom(roomId int, user realtimechat.User, refs ...string) *Room {
	clients := func() (buffer []*Client) {
		for _, ref := range refs {
			buffer = append(buffer, s.FindClient(ref))
		}
		return buffer
	}()
	s.Rooms[roomId] = NewRoom(roomId, user, clients...)
	go s.Rooms[roomId].Listen()
	return s.Rooms[roomId]
}

func (s *Server) JoinRoom(roomId int, ref string) {
	s.Rooms[roomId].join <- s.FindClient(ref)
}

func (s *Server) DeleteClient(sign string) {
	for i, room := range s.clients[sign].rooms {
		logrus.Infof("[ws.server]: client@%s keluar dari room %d", sign, i)
		room.leave <- s.clients[sign]
	}
	delete(s.clients, sign)
}

func (s *Server) DeleteRoom(roomId int) {
	for client := range s.Rooms[roomId].clients {
		logrus.Infof("[ws.server]: client@%s keluar dari room %d", client.sign, roomId)
		s.Rooms[roomId].leave <- client
	}
	// close(s.Rooms[roomid].join)
	// close(s.Rooms[roomid].leave)
	close(s.Rooms[roomId].message)
	delete(s.Rooms, roomId)
}