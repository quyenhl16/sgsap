package ies

import "sgsap/internal/utils"

func NewImsi(imsi string) *IE {
        v, err := utils.StrToSwappedBytes(imsi, "f")
        if err != nil {
                return nil
        }
        return &IE{
                IEI:     IMSI,
                Length:  uint8(len(v)),
                Payload: v,
        }
}
