package binary

type bigEndian struct{}

var BigEndian bigEndian

// Int8 turn []byte to bigEndian int8
func (bigEndian) Int8(b []byte) int8 {
	return int8(b[0])
}

// PutInt8 turn int8 to bigEndian []byte
func (bigEndian) PutInt8(b []byte, v int8) {
	b[0] = byte(v)
}

// Int16 turn []byte to bigEndian int16
func (bigEndian) Int16(b []byte) int16 { return int16(b[1]) | int16(b[0])<<8 }

// PutInt16 turn int16 to bigEndian []byte
func (bigEndian) PutInt16(b []byte, v int16) {
	_ = b[1]
	b[0] = byte(v >> 8)
	b[1] = byte(v)
}

// Int32 turn []byte to bigEndian int32
func (bigEndian) Int32(b []byte) int32 {
	return int32(b[3]) | int32(b[2])<<8 | int32(b[1])<<16 | int32(b[0])<<24
}

// PutInt32 turn int32 to bigEndian []byte
func (bigEndian) PutInt32(b []byte, v int32) {
	_ = b[3]
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
}
