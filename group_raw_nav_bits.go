package sbf

import (
	"encoding/binary"
	"fmt"
)

// GEORawL1/L5 messages are exactly the same - Some other Raw messages are also the same except with different NavBits lenghts, like BDSRaw (BDSRawB1C is not) and GALRawCNAV (probably others too)
type GEORaw struct {
	TOW        uint32
	WNc        uint16 // TimeStamp type for this and above (all messages use it?)
	SVID       uint8
	CRCPassed  uint8 // bool
	ViterbiCnt uint8
	Source     uint8 // Signal Type enum - 4.1.10 in spec
	FreqNr     uint8 // N/A in spec
	RxChannel  uint8
	NAVBits    [8]uint32
	// Padding ?
}

func deserializeBlockGEORaw(block Block) (GEORaw, error) {
	if len(block.Data) < 44 {
		return GEORaw{}, fmt.Errorf("block length less than minimum for block type")
	}
	return GEORaw{
		TOW:        binary.LittleEndian.Uint32(block.Data[:4]),
		WNc:        binary.LittleEndian.Uint16(block.Data[4:6]),
		SVID:       uint8(block.Data[6]),
		CRCPassed:  uint8(block.Data[7]),
		ViterbiCnt: uint8(block.Data[8]),
		Source:     uint8(block.Data[9]),
		FreqNr:     uint8(block.Data[10]),
		RxChannel:  uint8(block.Data[11]),
		NAVBits: [8]uint32{
			binary.LittleEndian.Uint32(block.Data[12:16]),
			binary.LittleEndian.Uint32(block.Data[16:20]),
			binary.LittleEndian.Uint32(block.Data[20:24]),
			binary.LittleEndian.Uint32(block.Data[24:28]),
			binary.LittleEndian.Uint32(block.Data[28:32]),
			binary.LittleEndian.Uint32(block.Data[32:36]),
			binary.LittleEndian.Uint32(block.Data[36:40]),
			binary.LittleEndian.Uint32(block.Data[40:44]),
		},
	}, nil
}

type GEORawL1 GEORaw // Block ID 4020

// TODO: Block name or number? DeserializeBlockGEORawL1. I think Block ID is good so caller doesn't need to refer to a table of Block ID to Block Name
func DeserializeBlock4020(block Block) (GEORawL1, error) {
	// TODO: alternatively just take []byte (block.Data), and assume the correct deserialize function was called
	//  the caller will presumably have to make this exact check anyway
	if block.ID != 4020 {
		return GEORawL1{}, fmt.Errorf("incorrect block ID: %d", block.ID)
	}
	geo, err := deserializeBlockGEORaw(block)
	return GEORawL1(geo), err
}

type GEORawL5 GEORaw // Block ID 4021

func DeserializeBlock4021(block Block) (GEORawL5, error) {
	if block.ID != 4021 {
		return GEORawL5{}, fmt.Errorf("incorrect block ID: %d", block.ID)
	}
	geo, err := deserializeBlockGEORaw(block)
	return GEORawL5(geo), err
}
