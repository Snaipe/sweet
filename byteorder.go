/* sweet
 *
 * Copyright (C) 2018  Franklin "Snaipe" Mathieu <me@snai.pe>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/binary"
	"unsafe"
)

type native struct{}

func (native) Uint16(data []byte) uint16 {
	if len(data) < int(unsafe.Sizeof(uint16(0))) {
		panic("NativeByteOrder.Uint16: data too short")
	}
	return *(*uint16)(unsafe.Pointer(&data[0]))
}

func (native) Uint32(data []byte) uint32 {
	if len(data) < int(unsafe.Sizeof(uint32(0))) {
		panic("NativeByteOrder.Uint32: data too short")
	}
	return *(*uint32)(unsafe.Pointer(&data[0]))
}

func (native) Uint64(data []byte) uint64 {
	if len(data) < int(unsafe.Sizeof(uint64(0))) {
		panic("NativeByteOrder.Uint64: data too short")
	}
	return *(*uint64)(unsafe.Pointer(&data[0]))
}

func (native) PutUint16(dest []byte, data uint16) {
	if len(dest) < int(unsafe.Sizeof(uint16(0))) {
		panic("NativeByteOrder.PutUint16: destination too short")
	}
	*(*uint16)(unsafe.Pointer(&dest[0])) = data
}

func (native) PutUint32(dest []byte, data uint32) {
	if len(dest) < int(unsafe.Sizeof(uint32(0))) {
		panic("NativeByteOrder.PutUint32: destination too short")
	}
	*(*uint32)(unsafe.Pointer(&dest[0])) = data
}

func (native) PutUint64(dest []byte, data uint64) {
	if len(dest) < int(unsafe.Sizeof(uint64(0))) {
		panic("NativeByteOrder.PutUint64: destination too short")
	}
	*(*uint64)(unsafe.Pointer(&dest[0])) = data
}

func (native) String() string {
	return "NativeByteOrder"
}

var NativeByteOrder binary.ByteOrder = native{}
