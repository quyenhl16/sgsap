package messages

import (
        "fmt"
        "sgsap/ies"
)

type AlertAck struct {
        MsgType uint8
        IMSI    *ies.IE
        Payload []byte
}

// NewAlertAck creates a new AlertAck.
func NewAlertAck(Imsi string) *AlertAck {
        m := &AlertAck{
                MsgType: MsgTypeSGSAPAlertAck,
                IMSI:    ies.NewImsi(Imsi),
        }
        return m
}

// Marshal returns the byte sequence generated from a HeartbeatRequest.
func (m *AlertAck) Marshal() ([]byte, error) {
        b := make([]byte, 11)
        if err := m.MarshalTo(b); err != nil {
                return nil, err
        }
        fmt.Println("b: ", b)
        return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (m *AlertAck) MarshalTo(b []byte) error {
        fmt.Println("b []bypte", b)

        b[0] = MsgTypeSGSAPAlertAck
        offset := 1
        if i := m.IMSI; i != nil {
                if err := i.MarshalTo(b[offset:]); err != nil {
                        return err
                }
                offset += i.MarshalLen()
        }

        return nil
}

// MarshalLen returns the serial length of Data.
func (m *AlertAck) MarshalLen() int {
        l := 1 + 10 // 1: MsgType; 10: len(Imsi)
        return l
}
