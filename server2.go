package main

import (
    "context"
    "crypto/tls"
    "fmt"
    "log"

    "github.com/quic-go/quic-go"
)

func main() {
  // Set Up TLS config
  log.SetFlags(log.Lshortfile)

  cer, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
  if err != nil {
      log.Fatal("Error loading certificate and key:", err)
    return
  } 
  
  tlsConf := &tls.Config{Certificates: []tls.Certificate{cer}}
  
  quicConf := &quic.Config{}

  listener, err := quic.ListenAddr("localhost:4242",tlsConf,quicConf)
  if err != nil {
    log.Fatal("Error starting QUIC listner:",err)
    return
  }
  fmt.Println("Listening on localhost:4242")

  for {
    sess, err := Listener.Accept(context.Background())
    if err != nil {
        log.Fatal("Error accepting connection:",err)
        return
    }

    go func(sess quic.Session) {
      for {
        stream, err := sess.AcceptStream(context.Background())
        if err != nil {
            log.Fatal("Error accepting stream:", err)
            return
        }

        go handleStream(stream)
      }
    }(sess)
  }
}

func handleStream(stream quic.Stream) {
  defer stream.Close()

  buf := make([]byte, 1024)
  for {
    n, err := stream.Read(buf)
    if err != nil {
      log.Println("Stream closed:", err)
      return
    }
    stream.Write(buf[:n])
  }
}
