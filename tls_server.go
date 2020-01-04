package main

import (
    "log"
    "crypto/tls"
    "net"
)

/*
 * TLS Handshake:
 * 
 */
 const (
     service string = "0.0.0.0:8000"
 )

 func main() {
     // The certPem and the keyPem are the CA of the TLS server and need to be
     // modified accordingly
     cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
     if err != nil {
         log.Fatal(err)
     }

     cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

     listener, err := tls.Listen("tcp", service, cfg)
     if err != nil {
         log.Fatal(err)
     }
     defer listener.Close()

     log.Print("Server: Listening")
     for {
         conn, err := listener.Accept()
         if err != nil {
             log.Println(err)
             continue
         }

         log.Printf("Server: Connection accepted with %s", conn.RemoteAddr())
         go handleClient(conn)
     }
}

func handleClient(conn net.Conn) {
    buf := make([]byte, 512)
    for {
        log.Print("Server: connection waiting")
        numbytes, err := conn.Read(buf)

        if err != nil {
            log.Printf("Server: read error: %s", err)
            break
        }

        log.Printf("Server: connection echo %q\n", string(buf[:numbytes]))
        numbytes, err = conn.Write(buf[:numbytes])
        log.Printf("Server: connection wrote %d bytes", numbytes)

        if err != nil {
            log.Printf("Server: write error: %s", err)
            break
        }
    }

    log.Println("Server: connection closed")
}
