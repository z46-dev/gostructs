package sstable

import "github.com/z46-dev/gostructs/misc"

// Table structure:
// 1: BloomFilter Size, ... BloomFilter
// 2: Index Size, ... Index
// 3: ... Data

func uint32ToBytes(input uint32) (output []byte) {
	output = make([]byte, 4)
	output[0] = byte(input >> 24)
	output[1] = byte(input >> 16)
	output[2] = byte(input >> 8)
	output[3] = byte(input)
	return
}

func bytesToUint32(input []byte) (output uint32) {
	output = 0
	output |= uint32(input[0]) << 24
	output |= uint32(input[1]) << 16
	output |= uint32(input[2]) << 8
	output |= uint32(input[3])
	return
}

type StructuredIndexNode struct {
	Key     string
	Pointer uint64
}

type SSData struct {
	Key   string
	Value []byte
}

func (s *SSData) ToBytes() (output []byte) {
	output = make([]byte, 0)
	output = append(output, uint32ToBytes(uint32(len(s.Key)))...)
	output = append(output, []byte(s.Key)...)
	output = append(output, uint32ToBytes(uint32(len(s.Value)))...)
	output = append(output, s.Value...)

	return
}

type SSTable struct {
	Filter *misc.BloomFilter[string]
	Index  []StructuredIndexNode
}
