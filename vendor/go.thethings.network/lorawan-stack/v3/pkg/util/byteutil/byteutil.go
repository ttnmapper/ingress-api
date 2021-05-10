// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package byteutil provides utilities for working with bytes using little endian binary encoding.
package byteutil

// MaxUint24 is the maximum value representable 3-octet uint.
const MaxUint24 = 1<<24 - 1

// ParseUint32 parses uint32 from b assuming little endian binary encoding.
func ParseUint32(b []byte) uint32 {
	switch len(b) {
	case 0:
		return 0
	case 1:
		_ = b[0]
		return uint32(b[0])
	case 2:
		_ = b[1]
		return uint32(b[0]) | uint32(b[1])<<8
	case 3:
		_ = b[2]
		return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16
	default:
		_ = b[3]
		return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
	}
}

// ParseUint64 parses uint64 from b assuming little endian binary encoding.
func ParseUint64(b []byte) uint64 {
	switch len(b) {
	case 0:
		return 0
	case 1:
		_ = b[0]
		return uint64(b[0])
	case 2:
		_ = b[1]
		return uint64(b[0]) | uint64(b[1])<<8
	case 3:
		_ = b[2]
		return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16
	case 4:
		_ = b[3]
		return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24
	case 5:
		_ = b[4]
		return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32
	case 6:
		_ = b[5]
		return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40
	case 7:
		_ = b[6]
		return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48
	default:
		_ = b[7]
		return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
	}
}

// AppendUint16 appends v to dst using little endian binary encoding using at most byteCount bytes.
func AppendUint16(dst []byte, v uint16, byteCount uint8) []byte {
	switch byteCount {
	case 0:
		return dst
	case 1:
		return append(dst, byte(v))
	default:
		dst = append(dst, byte(v), byte(v>>8))
		for i := uint8(2); i < byteCount; i++ {
			dst = append(dst, 0)
		}
		return dst
	}
}

// AppendUint32 appends v to dst using little endian binary encoding using at most byteCount bytes.
func AppendUint32(dst []byte, v uint32, byteCount uint8) []byte {
	switch byteCount {
	case 0, 1, 2:
		return AppendUint16(dst, uint16(v), byteCount)
	case 3:
		return append(dst, byte(v), byte(v>>8), byte(v>>16))
	default:
		dst = append(dst, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
		for i := uint8(4); i < byteCount; i++ {
			dst = append(dst, 0)
		}
		return dst
	}
}

// AppendUint64 appends v to dst using little endian binary encoding using at most byteCount bytes.
func AppendUint64(dst []byte, v uint64, byteCount uint8) []byte {
	switch byteCount {
	case 0, 1, 2, 3, 4:
		return AppendUint32(dst, uint32(v), byteCount)
	case 5:
		return append(dst, byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32))
	case 6:
		return append(dst, byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32), byte(v>>40))
	case 7:
		return append(dst, byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32), byte(v>>40), byte(v>>48))
	default:
		dst = append(dst, byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32), byte(v>>40), byte(v>>48), byte(v>>56))
		for i := uint8(8); i < byteCount; i++ {
			dst = append(dst, 0)
		}
		return dst
	}
}
