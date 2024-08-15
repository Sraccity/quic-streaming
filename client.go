package main

import (
    "crypto/tls"
    "fmt"
    "log"
    "time"

    "github.com/quic-go/quic-go"
)

func main() {
    // Set up TLS configuration
    tlsConf := &tls.Config{
        // Load your TLS certificate and key here if needed
        InsecureSkipVerify: true, // For testing purposes
    }

    // Create a QUIC connection
    sess, err := quic.DialAddr("localhost:4242", tlsConf, &quic.Config{})
    if err != nil {
        log.Fatal(err)
    }
    defer sess.Close()

    // Create a stream
    stream, err := sess.OpenStreamSync(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    defer stream.Close()

    // Send data
    _, err = stream.Write([]byte("Hello, QUIC server!"))
    if err != nil {
        log.Fatal(err)
    }

    // Wait for response
    buf := make([]byte, 1024)
    _, err = stream.Read(buf)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Received response:", string(buf))

    // Allow time for any cleanup
    time.Sleep(1 * time.Second)
}
