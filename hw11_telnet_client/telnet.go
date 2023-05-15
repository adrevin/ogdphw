package main

import (
	"bufio"
	"errors"
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
	scanner *bufio.Scanner
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
	client.scanner = bufio.NewScanner(client.conn)
	client.scanner.Split(bufio.ScanLines)
	return nil
}

func (client *telnetClient) Send() error {
	_, err := io.Copy(client.conn, client.in)
	if err != nil {
		return err
	}
	return nil
}

func (client *telnetClient) Receive() error {
	client.scanner.Scan()
	scannerErr := client.scanner.Err()
	if errors.Is(scannerErr, io.EOF) {
		return scannerErr
	}
	// bufio.ScanLines drops CR, return it
	bytes := client.scanner.Bytes()
	bytes = append(bytes, 10)
	_, err := client.out.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func (client *telnetClient) Close() error {
	return client.conn.Close()
}
