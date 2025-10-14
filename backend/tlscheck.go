package main

import (
    "crypto/tls"
    "fmt"
    "net"
    "time"
)

func tryHost(host string) {
    addr := net.JoinHostPort(host, "27017")
    fmt.Println("Testing:", addr)
    dialer := &net.Dialer{Timeout: 10 * time.Second}
    conn, err := tls.DialWithDialer(dialer, "tcp", addr, &tls.Config{ServerName: host})
    if err != nil {
        fmt.Println("TLS dial error:", err)
        return
    }
    defer conn.Close()
    fmt.Println("TLS handshake OK. Peer certificates:")
    for i, cert := range conn.ConnectionState().PeerCertificates {
        fmt.Printf(" cert[%d]: Subject=%+v Issuer=%+v\n", i, cert.Subject, cert.Issuer)
    }
}

func main() {
    hosts := []string{
        "ac-dcvbn90-shard-00-00.mzeqzcd.mongodb.net",
        "ac-dcvbn90-shard-00-01.mzeqzcd.mongodb.net",
        "ac-dcvbn90-shard-00-02.mzeqzcd.mongodb.net",
    }
    for _, h := range hosts {
        tryHost(h)
        fmt.Println("---")
        time.Sleep(500 * time.Millisecond)
    }
}
