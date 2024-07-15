package messages

const (
        MsgTypeSGSAPAlertAck             uint8 = 14
        MsgTypeSGSAPLocationUpdateAccept uint8 = 10
)

type Message interface {
        MarshalTo([]byte) error
        UnmarshalBinary(b []byte) error
        MarshalLen() int
        MessageType() uint8
        MessageTypeName() string
        Version() int
        SEID() uint64
        Sequence() uint32
        SetSequenceNumber(seq uint32)
}
