package sbf

import (
	"bytes"
	"reflect"
	"testing"
)

var (
	// Block ID 4020, GeoRAWL1 message from SVID 122 SouthPAN - Note, CRC might be incorrect
	bin4020   = []byte{'$', '@', 86, 148, 180, 15, 52, 0, 216, 116, 117, 11, 228, 8, 122, 1, 0, 24, 0, 26, 0, 0, 18, 83, 254, 239, 63, 0, 63, 254, 239, 255, 0, 0, 0, 224, 0, 0, 0, 0, 187, 185, 3, 0, 0, 0, 128, 163, 192, 47, 5, 20}
	block4020 = Block{
		CRC:    37974,
		ID:     4020,
		Length: 52,
		Data:   []byte{216, 116, 117, 11, 228, 8, 122, 1, 0, 24, 0, 26, 0, 0, 18, 83, 254, 239, 63, 0, 63, 254, 239, 255, 0, 0, 0, 224, 0, 0, 0, 0, 187, 185, 3, 0, 0, 0, 128, 163, 192, 47, 5, 20},
	}
	deserialized4020 = GEORawL1{
		TOW:        192247000,
		WNc:        2276,
		SVID:       122,
		CRCPassed:  1,
		ViterbiCnt: 0,
		Source:     24,
		FreqNr:     0,
		RxChannel:  26,
		NAVBits:    [8]uint32{1393688576, 4190206, 4293918271, 3758096384, 0, 244155, 2743074816, 335884224},
	}
)

func TestReadBlock(t *testing.T) {
	block, err := ReadBlock(bytes.NewReader(bin4020))
	if err != nil {
		t.Fatal("error in ReadBlock function:", err)
	}

	if !reflect.DeepEqual(block, block4020) {
		t.Fatalf("parsed block does not match expected result: %+v", block)
	}
}

func TestDeserializeBlock4020(t *testing.T) {
	deserializedBlock, err := DeserializeBlock4020(block4020)
	if err != nil {
		t.Fatal("error in DeserializeBlock4020 function:", err)
	}

	if !reflect.DeepEqual(deserializedBlock, deserialized4020) {
		t.Fatalf("deserialized block does not match expected result: %+v", deserializedBlock)
	}
}