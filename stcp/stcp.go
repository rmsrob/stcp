package stcp

import (
	"fmt"
	"net"
)

type Message struct {
	From    string
	Payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitChan   chan struct{}
	MsgChan    chan Message
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitChan:   make(chan struct{}),
		MsgChan:    make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()
	<-s.quitChan
	close(s.MsgChan)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}

		fmt.Println("New conn to the server: ", conn.RemoteAddr().String())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error", err)
			continue
		}

		s.MsgChan <- Message{
			From:    conn.RemoteAddr().String(),
			Payload: buf[:n],
		}

		conn.Write([]byte("I have spoken!\n\n"))
	}
}
