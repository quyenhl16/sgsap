package messages

import (
        "fmt"
        "sgsap/ies"
)

type LocationUpdateAccept struct {
        MsgType        uint8
        IMSI           *ies.IE
        LocationAreaId *ies.IE
        Payload        []byte
}

// NewAlertAck creates a new AlertAck.
func NewLocationUpdateAccept(Imsi, Lai string) *LocationUpdateAccept {
        m := &LocationUpdateAccept{
                MsgType:        MsgTypeSGSAPLocationUpdateAccept,
                IMSI:           ies.NewImsi(Imsi),
                LocationAreaId: ies.NewLAI(Lai),
        }
        return m
}

// Marshal returns the byte sequence generated from a HeartbeatRequest.
func (m *LocationUpdateAccept) Marshal() ([]byte, error) {
        b := make([]byte, 11)
        if err := m.MarshalTo(b); err != nil {
                return nil, err
        }
        fmt.Println("b: ", b)
        return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (m *LocationUpdateAccept) MarshalTo(b []byte) error {
        fmt.Println("b []bypte", b)
        if m.Payload != nil {
                m.Payload = nil
        }
        m.Payload = make([]byte, 20)
        b[0] = MsgTypeSGSAPLocationUpdateAccept
        offset := 1
        if i := m.IMSI; i != nil {
                if err := i.MarshalTo(m.Payload[offset:]); err != nil {
                        return err
                }
                offset += i.MarshalLen()
        }
        fmt.Println("OFFSET || b", offset, b)
        if i := m.LocationAreaId; i != nil {
                if err := i.MarshalTo(m.Payload[offset:]); err != nil {
                        return err
                }
                offset += i.MarshalLen()
        }
        copy(b[1:], m.Payload)
        return nil
}

// MarshalLen returns the serial length of Data.
func (m *LocationUpdateAccept) MarshalLen() int {
        l := 1 + 10 + 6 // 1: MsgType; 10: len(Imsi); 6: len(LAI)
        return l
}
