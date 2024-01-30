package sbf

import (
	"bytes"
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

func deserializeBlockGEORaw(block Block) (gr GEORaw, err error) {
	if len(block.Data) < 44 {
		return GEORaw{}, fmt.Errorf("block length less than minimum for block type")
	}

	err = binary.Read(bytes.NewReader(block.Data), binary.LittleEndian, &gr)
	return gr, err
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
