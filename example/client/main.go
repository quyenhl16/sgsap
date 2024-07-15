package main

import (
        "fmt"
        "log"
        "sgsap/messages"
)

func main() {
        //var (
        //      server = flag.String("s", "127.0.0.2:8805", "server's addr/port")
        //)
        //flag.Parse()
        //
        //raddr, err := net.ResolveUDPAddr("udp", *server)
        //if err != nil {
        //      log.Fatal(err)
        //}
        //
        //conn, err := net.DialUDP("udp", nil, raddr)
        //if err != nil {
        //      log.Fatal(err)
        //}
        //_ = conn
        uchbreq := messages.NewAlertAck("452040916843227")
        //uchbreq := messages.NewLocationUpdateAccept("452040916843227", "4520441527")
        hbreq, err := uchbreq.Marshal()
        fmt.Println("UnencodeMsg: ", uchbreq, "Encoded: ", "hbreq", hbreq)
        if err != nil {
                log.Fatal(err)
        }

        //if _, err := conn.Write(hbreq); err != nil {
        //      log.Fatal(err)
        //}
        //log.Printf("sent Heartbeat Request to: %s", raddr)
        //
        //if err := conn.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
        //      log.Fatal(err)
        //}
}
