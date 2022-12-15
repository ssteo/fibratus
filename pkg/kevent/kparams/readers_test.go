/*
 * Copyright 2021-2022 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kparams

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestReadBuffer(t *testing.T) {
	b := []byte{
		0x70, 0x3B, 0xD1, 0x24, 0x8E, 0xD6, 0xFF, 0xFF, 0x9C, 0x02, 0x00, 0x00, 0x2C, 0x00, 0x5C, 0x00,
		0x52, 0x00, 0x45, 0x00, 0x47, 0x00, 0x49, 0x00, 0x53, 0x00, 0x54, 0x00, 0x52, 0x00, 0x59, 0x00,
		0x5C, 0x00, 0x55, 0x00, 0x53, 0x00, 0x45, 0x00, 0x52, 0x00, 0x5C, 0x00, 0x53, 0x00, 0x2D, 0x00,
		0x31, 0x00, 0x2D, 0x00, 0x35, 0x00, 0x2D, 0x00, 0x32, 0x00, 0x31, 0x00, 0x2D, 0x00, 0x32, 0x00,
		0x32, 0x00, 0x37, 0x00, 0x31, 0x00, 0x30, 0x00, 0x33, 0x00, 0x34, 0x00, 0x34, 0x00, 0x35, 0x00,
		0x32, 0x00, 0x2D, 0x00, 0x32, 0x00, 0x36, 0x00, 0x30, 0x00, 0x36, 0x00, 0x32, 0x00, 0x37, 0x00,
		0x30, 0x00, 0x30, 0x00, 0x39, 0x00, 0x39, 0x00, 0x2D, 0x00, 0x39, 0x00, 0x38, 0x00, 0x34, 0x00,
		0x38, 0x00, 0x37, 0x00, 0x31, 0x00, 0x35, 0x00, 0x36, 0x00, 0x39, 0x00, 0x2D, 0x00, 0x31, 0x00,
		0x30, 0x00, 0x30, 0x00, 0x31, 0x00, 0x00, 0x00}
	p := uintptr(unsafe.Pointer(&b[0]))

	assert.Equal(t, uint64(18446698504724233072), ReadUint64(p, 0))
	assert.Equal(t, uint32(668), ReadUint32(p, 8))
	assert.Equal(t, uint16(44), ReadUint16(p, 12))
	assert.Equal(t, "\\REGISTRY\\USER\\S-1-5-21-2271034452-2606270099-984871569-1001", ConsumeUTF16String(p, 14, uint16(len(b))))
}