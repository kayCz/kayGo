package protocol

import (
	"io"
	"encoding/json"
	log "github.com/kayCz/kaylog"
)

type JsonProtocol struct {
	rw io.ReadWriter
	decoder *json.Decoder
	encoder *json.Encoder
}

func NewJsonProtocol() *JsonProtocol {
	return &JsonProtocol{}
}

func (self *JsonProtocol) Fill(rw io.ReadWriter) {
	self.rw = rw
	self.decoder = json.NewDecoder(rw)
	self.encoder = json.NewEncoder(rw)
}

func (self *JsonProtocol) Send(v interface{}) error {
	log.Info("Json Protocol | Send self:" , self)
	return self.encoder.Encode(v)
}

func (self *JsonProtocol) Receive() (v interface{}, err error) {
	self.decoder.Decode(&v)
	log.Info("Json Protocol | Receive :" ,self,  v)
	return v , nil
}
