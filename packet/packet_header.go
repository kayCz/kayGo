package packet

import (
	"encoding/binary"
	"io"
)

const (
	MAGIC int16 = 0x0221
)


type PackerHeader struct {
	magic int16
	length int32
	version int8
	bodyLen int32
}

func (self *PackerHeader) Marshal(w io.Writer) error	{
	binary.Write(w, binary.BigEndian, MAGIC)
	binary.Write(w, binary.BigEndian, self.length)
	binary.Write(w, binary.BigEndian, self.version)
	binary.Write(w, binary.BigEndian, self.bodyLen)
	return nil
}


func (self *PackerHeader) UnMarshal(b io.Reader) (err error){
	err = binary.Read(b, binary.BigEndian, &self.magic)
	if err != nil {
		return err
	}
	err = binary.Read(b, binary.BigEndian, &self.length)
	if err != nil {
		return err
	}
	err = binary.Read(b, binary.BigEndian, &self.length)
	if err != nil {
		return err
	}
	err = binary.Read(b, binary.BigEndian, &self.bodyLen)
	if err != nil {
		return err
	}
	return nil
}


