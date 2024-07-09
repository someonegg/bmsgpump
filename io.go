// Copyright 2023 someonegg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmsgpump

type Message []byte

type MessageReader interface {
	ReadMessage() (m Message, err error)
}

type MessageWriter interface {
	WriteMessage(m Message) error
}

type MessageReadWriter interface {
	MessageReader
	MessageWriter
}

type StopNotifier interface {
	OnStop()
}
