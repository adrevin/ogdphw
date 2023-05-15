package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "connect timeout, such as 300ms, 3s, ...")
}

func main() {
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	address := net.JoinHostPort(flag.Args()[0], flag.Args()[1])
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	client := NewTelnetClient(address, timeout, io.NopCloser(in), out)

	err := client.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanBytes)
		for {
			select {
			case <-ctx.Done():
				client.Close()
				return
			default:
				scanner.Scan()
				if errors.Is(scanner.Err(), io.EOF) {
					stop()
					break
				}
				in.Write(scanner.Bytes())
				err = client.Send()
				if err != nil {
					stop()
					return
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := client.Receive()
				if errors.Is(err, io.EOF) {
					stop()
					fmt.Fprintln(os.Stderr, "...Connection was closed by peer")
					break
				}
				fmt.Fprint(os.Stdout, out.String())
				out.Reset()
			}
		}
	}()

	wg.Wait()
}
