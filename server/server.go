package server

import (
	"net"
	log "github.com/kayCz/kaylog"
	"../client"
	"../protocol"
)

type Handler interface {
	serveHandler(v interface{}) (r interface{}, err error)
}

type ServerHandler func(v interface{})  (r interface{}, err error)

func (self ServerHandler) serveHandler(v interface{}) (r interface{}, err error) {
	return self(v)
}

type Server struct {
	host string
	isShutDown bool
	Address string
	protocol protocol.Protocol
	handler Handler
}

func (self *Server) Start() error{

	add , err := net.ResolveTCPAddr("tcp", self.host)
	if err != nil {
		return err
	}

	self.Address = add.String()

	l , er := net.ListenTCP("tcp", add)
	if er != nil {
		return er
	}
	log.Info("Server | Start  : ", self)
	go func() {
		for !self.isShutDown {
			conn, e := l.Accept()
			if e != nil {
				log.Error("Server | Accept | Error : ", e)
				continue
			}
			client := client.NewClient(conn, self.protocol)
			log.Info("Client | ", client)
			go self.serve(client)
		}
	}()

	return nil

}

func (self *Server) serve(client *client.Client) {
	v ,err := client.Receive()
	if err != nil {
		log.Error("Server | serve Error : " , err)
	}
	ans, _ := self.handler.serveHandler(v)
	client.Send(ans)
}


func NewServer(host string, protocol protocol.Protocol, handler Handler) *Server {
	return &Server{host:host, isShutDown:false, protocol:protocol, handler:handler}
}

func (self *Server) ShutDown() {
	self.isShutDown = true
	log.Info("Server ShutDown")
}


