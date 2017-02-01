package packet


type Packet struct{
	packerHeader PackerHeader
	data []byte
}
