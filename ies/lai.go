package ies

import "sgsap/internal/utils"

func NewLAI(LAI string) *IE {
        v, err := utils.StrToSwappedBytes(LAI, "f")
        if err != nil {
                return nil
        }
        return &IE{
                IEI:     LocationAreaIdentifier,
                Length:  uint8(len(v)),
                Payload: v,
        }
}
