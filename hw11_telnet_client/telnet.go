package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}
type telnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{address: address, timeout: timeout, in: in, out: out}
}

func (client *telnetClient) Connect() error {
	conn, err := net.Dial("tcp", client.address)
	if err != nil {
		return err
	}
	client.conn = conn
	return nil
}

func (client *telnetClient) Send() error {
	buf, err := io.ReadAll(client.in)
	if err != nil {
		return err
	}
	_, err = client.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
func (client *telnetClient) Receive() error {
	buf, err := io.ReadAll(client.conn)
	if err != nil {
		return err
	}
	_, err = client.out.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
func (client *telnetClient) Close() error {
	return client.conn.Close()
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
