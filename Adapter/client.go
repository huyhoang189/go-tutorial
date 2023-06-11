package Adapter

type Client struct {
}

func (client *Client) InsertPortToComputer(c Computer) {
	c.insertByUsbA()
}
