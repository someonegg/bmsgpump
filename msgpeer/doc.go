// Copyright 2024 someonegg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package msgpeer implements a synchronous request response model over message-pump.
// It considers the client and server peers, allowing each to send requests to the
// other concurrently.
package msgpeer
