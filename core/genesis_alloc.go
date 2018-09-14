// Copyright 2017 The go-wabei Authors
// This file is part of the go-wabei library.
//
// The go-wabei library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-wabei library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-wabei library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData ="\xf8\xaa\xe1\x94\x11\xf1\xd8\U0009b5fd\xf1\x87\xdd\xe5;\x13\u0744\xb1\f{\x9c\x89\x8b\x03O\bo;3\xb6\x84\x00\x00\x00\xe1\x94\x1fTP\xf5\x81\xc5^f?\xb1\x85MW-\xfb.K\x12D\xe2\x8b\x03O\bo;3\xb6\x84\x00\x00\x00\xe1\x94)\xb1\x94\x946\u0790\x81\xb8\xaf\xc9\xff\t7\x17P\xf6h\xc6\x03\x8b\x03O\bo;3\xb6\x84\x00\x00\x00\xe1\x94d>\x84\u0217\x16\x0f\xbd\xb6a\x13\x81\xe5\xfdVi\x92\x86\u054f\x8b\x03O\bo;3\xb6\x84\x00\x00\x00\xe1\x94\xf0^#4:\x9b\x9b\x10N \x88\x18\x9f\x83\u0354v\xbd\x05\xa3\x8b\x03O\bo;3\xb6\x84\x00\x00\x00"
const testnetAllocData = ""
const rinkebyAllocData = ""
