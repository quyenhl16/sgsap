package ies

import (
        "errors"
        "fmt"
)

// IE definitions - 3GPP TS 29.118 Ver 14.1.0
const (
        IMSI                            uint8 = 1
        VLRName                         uint8 = 2
        TMSI                            uint8 = 3
        LocationAreaIdentifier          uint8 = 4
        ChannelNeeded                   uint8 = 5
        eMLPPPriority                   uint8 = 6
        TMSIStatus                      uint8 = 7
        SGsCause                        uint8 = 8
        MMEName                         uint8 = 9
        EPSLocationUpdateType           uint8 = 10
        GlobalCNId                      uint8 = 11
        MobileIdentity                  uint8 = 14
        RejectCause                     uint8 = 15
        IMSIDetachFromEPSServiceType    uint8 = 16
        IMSIDetachFromNonEPSServiceType uint8 = 17
        IMEISV                          uint8 = 21
        NASMessageContainer             uint8 = 22
        MMInformation                   uint8 = 23
        ErroneousMessage                uint8 = 27
        CLI                             uint8 = 28
        LCSClientIdentity               uint8 = 29
        LCSIndicator                    uint8 = 30
        SSCode                          uint8 = 31
        ServiceIndicator                uint8 = 32
        UETimeZone                      uint8 = 33
        MobileStationClassmark2         uint8 = 34
        TrackingAreaIdentity            uint8 = 35
        EUTRANCellGlobalIdentity        uint8 = 36
        UEEMMMode                       uint8 = 37
        AdditionalPagingIndicators      uint8 = 38
        TMSIBasedNRIContainer           uint8 = 39
)

type IE struct {
        IEI     uint8
        Length  uint8
        Payload []byte
}

// Marshal returns the byte sequence generated from an IE instance.
func (i *IE) Marshal() ([]byte, error) {
        b := make([]byte, i.MarshalLen())
        if err := i.MarshalTo(b); err != nil {
                return nil, err
        }
        return b, nil
}

// MarshalTo puts the byte sequence in the byte array given as b.
func (i *IE) MarshalTo(b []byte) error {
        l := len(b)
        if l < 2 {
                return errors.New("length value is invalid")
        }
        fmt.Println("IEI | ", i.IEI, "i.Payload |", i.Payload)
        b[0] = i.IEI
        b[1] = i.Length
        copy(b[2:i.MarshalLen()], i.Payload)
        return nil
}

// MarshalLen returns field length in integer.
func (i *IE) MarshalLen() int {
        l := 1 + 1 // len(IEI) + len(Length)
        return l + len(i.Payload)
}

// SetLength sets the length in Length field.
func (i *IE) SetLength() {
        l := 0
        i.Length = uint8(l + len(i.Payload))
}
