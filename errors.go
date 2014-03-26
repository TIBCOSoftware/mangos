// Copyright 2014 Garrett D'Amore
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sp

import (
	"errors"
)

var (
	ErrBadAddr     = errors.New("invalid address")
	ErrBadHeader   = errors.New("invalid header received")
	ErrBadVersion  = errors.New("invalid protocol version")
	ErrTooShort    = errors.New("message is too short")
	ErrTooLong     = errors.New("message is too long")
	ErrClosed      = errors.New("connection closed")
	ErrConnRefused = errors.New("connection refused")
	ErrSendTimeout = errors.New("send time out")
	ErrRecvTimeout = errors.New("receive time out")
	ErrBadTran     = errors.New("invalid or unsupported transport")
	ErrBadProto    = errors.New("invalid or unsupported protocol")
	ErrPipeFull    = errors.New("pipe full")
	ErrPipeEmpty   = errors.New("pipe empty")
	ErrBadOption   = errors.New("invalid or unsupported option")
	ErrBadValue    = errors.New("invalid option value")
	ErrGarbled     = errors.New("message garbled")
	ErrAddrInUse   = errors.New("address in use")
)
