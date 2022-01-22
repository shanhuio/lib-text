// Copyright (C) 2022  Shanhu Tech Inc.
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
// for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package lexing

import (
	"fmt"
)

// Error is a parsing error
type Error struct {
	Pos  *Pos   // Pos can be null for error not related to any position
	Err  error  // Err is the error message, human friendly.
	Code string // Code is the error code, machine friendly.
}

// Error returns the error string.
func (e *Error) Error() string {
	if e.Pos == nil {
		return e.Err.Error()
	}

	return fmt.Sprintf("%s:%d: %s",
		e.Pos.File, e.Pos.Line,
		e.Err.Error(),
	)
}

// JSON returns a JSON marshable object of the error.
func (e *Error) JSON() interface{} {
	var ret struct {
		File string `json:"file"`
		Line int    `json:"line"`
		Col  int    `json:"col"`
		Code string `json:"code"`
		Err  string `json:"err"`
	}

	pos := e.Pos
	if pos != nil {
		ret.File = pos.File
		ret.Line = pos.Line
		ret.Col = pos.Col
	}
	ret.Code = e.Code
	ret.Err = e.Err.Error()
	return ret
}

// CodeErrorf creates a lex8.Error with ErrCode
func CodeErrorf(c string, f string, args ...interface{}) *Error {
	e := fmt.Errorf(f, args...)
	return &Error{Err: e, Code: c}
}

// Errorf creates a lex8.Error similar to fmt.Errorf
func Errorf(f, c string, args ...interface{}) *Error {
	return CodeErrorf("", f, args...)
}
