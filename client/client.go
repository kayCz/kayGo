package client

import (
	"net"
	"github.com/kayCz/kayGo/protocol"
	log "github.com/kayCz/kaylog"
)


type Client struct {
	conn net.Conn
	protocol protocol.Protocol
	address string
}

func Dail(network, address string, protocol protocol.Protocol) (*Client, error){
	conn , err := net.Dial(network, address)
	if err != nil {
		log.Error("Client | Dail | Error : " , err)
		return nil, err
	}
	return NewClient(conn, protocol), nil
}

func NewClient(conn net.Conn, protocol protocol.Protocol) *Client {
	protocol.Fill(conn)
	return &Client{
		conn : conn,
		address : conn.RemoteAddr().String(),
		protocol : protocol,
	}
}

func (self *Client) Send(v interface{}) error{
	log.Info("Client | Send  : ", v, self)
	return self.protocol.Send(v)
}

func (self *Client) Receive() (v interface{}, err error){
	return self.protocol.Receive()
}

func (self *Client) Close() error {
	return self.conn.Close()
}