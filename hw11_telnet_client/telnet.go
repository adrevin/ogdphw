package main

import (
	"bufio"
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
	reader  *bufio.Reader
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{address: address, timeout: timeout, in: in, out: out}
}

func (client *telnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", client.address, timeout)
	if err != nil {
		return err
	}
	client.conn = conn
	client.reader = bufio.NewReader(client.conn)
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
	buf, err := client.reader.ReadBytes('\n')
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
	err := client.in.Close()
	if err != nil {
		return err
	}
	return client.conn.Close()
}
