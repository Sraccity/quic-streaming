package main

import (
    "context"
    "crypto/tls"
    "fmt"
    "log"

    "github.com/quic-go/quic-go"
)

udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: 1234})
// TODO: ERR Handling 
tr := quic.Transport(
    Conn: udpConn,
}
// ln = listener
ln, err := tr.Listen(tlsConf, quicConf)
// TODO: ERR Handling
for {
    conn, err := ln.Accept()
    // TODO: ERR Handling
}


// 0.5-RRT fallback

func RRT_0_5(tlsConf tlsConf, config quicConf) (*EarlyListener, error) {
    ln, err := tr.ListenEarly(tlsConf,  quicConf)
    // TODO: ERR
    conn, err := ln.Accept()
    // TODO: ERR

    go func(){
        str, err := conn.OpenStream()
        // TODO: ERR
    
        select {
        case <- conn.HandshakeComplete();
          // Acknowledge complete
        case <- conn.Context().Done();
          // Connection closed before handshake completed 
        }
    }()
    return ln //TODO: Error
}

func RRT_0(tlsConf tlsConf, config quicConf) (*EarlyListener, error) {
    quicConf := &quicConf.Config{Allow0RRT: true}
    ln, err := tr.ListenEarly(tlsConf, quicConf)
    //TODO: 
    conn, err := ln.Accept()
    //TODO: 

    go func() {
        str, err := conn.AcceptStream()
        //TODO:

        select {
        case <-conn.HandshakeComplete():
            //Yippie
        case <-conn.Context().Done():
            // Closed before handshake
        }
    }()
    return ln // TODO: err T_T
}



