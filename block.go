package sbf

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Block struct {
	// Sync int16 // "$@" or 0x24,0x40
	CRC    uint16 // Computed for ID+Length+Data as []byte
	ID     uint16
	Length uint16 // multiple of 4
	Data   []byte // Length-8 bytes
}

// Reads from reader until a valid SBF Block is found (based on Block sync bits - not calculating CRC)
func ReadBlock(reader io.Reader) (block Block, err error) {
	r := bufio.NewReader(reader)

	// TODO: Should we skip bytes like this, or just return err if first two bytes are not '$@'?
	// 	Could return an error and leave it up to the invoker to track skipped bytes
	//  Would we peek or read?
	if _, err = r.ReadBytes('$'); err != nil {
		return block, err
	}

	b, err := r.ReadByte()
	if err != nil {
		return block, err
	}

	if b != byte('@') {
		return ReadBlock(reader) // This does seem strange - see above TODO
	}

	if err := binary.Read(r, binary.LittleEndian, &block.CRC); err != nil {
		return block, err
	}

	if err := binary.Read(r, binary.LittleEndian, &block.ID); err != nil {
		return block, err
	}

	if err := binary.Read(r, binary.LittleEndian, &block.Length); err != nil {
		return block, err
	}

	block.Data = make([]byte, block.Length-8)
	_, err = r.Read(block.Data)

	return block, err
}
