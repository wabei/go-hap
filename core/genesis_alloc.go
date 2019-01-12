// Copyright 2017 The go-hap Authors
// This file is part of the go-hap library.
//
// The go-hap library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-hap library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-hap library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData ="\xf8f\xe1\x94NH\x1a\x89\x0f\u0443\xbd7\x13\x90\x0f\x81\x9d\xfb,\x1e\t\x9e\x88\x8b)[\xe9nd\x06ir\x00\x00\x00\xe1\x94`)\xd1\xfc\xa8\xb6xy\x99\u07ce4\"\u0689\xceIt\xa8L\x8b!\x16TXP\x05!(\x00\x00\x00\xe1\x94\xe7\xb7\x1c\xae\x8f>\xd5\u0141v\x8a\xe6_\xe8Y\x9bt\x15;p\x8b\bE\x95\x16\x14\x01HJ\x00\x00\x00"
const testnetAllocData = "\xf8f\xe1\x94NH\x1a\x89\x0f\u0443\xbd7\x13\x90\x0f\x81\x9d\xfb,\x1e\t\x9e\x88\x8b)[\xe9nd\x06ir\x00\x00\x00\xe1\x94`)\xd1\xfc\xa8\xb6xy\x99\u07ce4\"\u0689\xceIt\xa8L\x8b!\x16TXP\x05!(\x00\x00\x00\xe1\x94\xe7\xb7\x1c\xae\x8f>\xd5\u0141v\x8a\xe6_\xe8Y\x9bt\x15;p\x8b\bE\x95\x16\x14\x01HJ\x00\x00\x00"
const rinkebyAllocData = "\xf8f\xe1\x94NH\x1a\x89\x0f\u0443\xbd7\x13\x90\x0f\x81\x9d\xfb,\x1e\t\x9e\x88\x8b)[\xe9nd\x06ir\x00\x00\x00\xe1\x94`)\xd1\xfc\xa8\xb6xy\x99\u07ce4\"\u0689\xceIt\xa8L\x8b!\x16TXP\x05!(\x00\x00\x00\xe1\x94\xe7\xb7\x1c\xae\x8f>\xd5\u0141v\x8a\xe6_\xe8Y\x9bt\x15;p\x8b\bE\x95\x16\x14\x01HJ\x00\x00\x00"
