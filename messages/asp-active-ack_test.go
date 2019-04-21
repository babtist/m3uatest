// Copyright 2018-2019 go-m3ua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import (
	"testing"

	"github.com/wmnsk/go-m3ua/messages/params"
)

func TestAspActiveAck(t *testing.T) {
	cases := []testCase{
		{
			"has-all-params",
			NewAspActiveAck(
				params.NewTrafficModeType(2),
				params.NewRoutingContext(1),
				params.NewInfoString("deadbeef"),
			),
			[]byte{
				// Header
				0x01, 0x00, 0x04, 0x03, 0x00, 0x00, 0x00, 0x2c,
				// TrafficModeType
				0x00, 0x0b, 0x00, 0x08, 0x00, 0x00, 0x00, 0x02,
				// RoutingContext
				0x00, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01,
				// InfoString
				0x00, 0x04, 0x00, 0x0c, 0x64, 0x65, 0x61, 0x64,
				0x62, 0x65, 0x65, 0x66,
			},
		}, {
			"has-tmt-rc",
			NewAspActiveAck(
				params.NewTrafficModeType(2),
				params.NewRoutingContext(1),
				nil,
			),
			[]byte{
				// Header
				0x01, 0x00, 0x04, 0x03, 0x00, 0x00, 0x00, 0x20,
				// TrafficModeType
				0x00, 0x0b, 0x00, 0x08, 0x00, 0x00, 0x00, 0x02,
				// RoutingContext
				0x00, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"has-tmt-info",
			NewAspActiveAck(
				params.NewTrafficModeType(2),
				nil,
				params.NewInfoString("deadbeef"),
			),
			[]byte{
				// Header
				0x01, 0x00, 0x04, 0x03, 0x00, 0x00, 0x00, 0x24,
				// TrafficModeType
				0x00, 0x0b, 0x00, 0x08, 0x00, 0x00, 0x00, 0x02,
				// InfoString
				0x00, 0x04, 0x00, 0x0c, 0x64, 0x65, 0x61, 0x64,
				0x62, 0x65, 0x65, 0x66,
			},
		}, {
			"has-rc-info",
			NewAspActiveAck(
				nil,
				params.NewRoutingContext(1),
				params.NewInfoString("deadbeef"),
			),
			[]byte{
				// Header
				0x01, 0x00, 0x04, 0x03, 0x00, 0x00, 0x00, 0x24,
				// RoutingContext
				0x00, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01,
				// InfoString
				0x00, 0x04, 0x00, 0x0c, 0x64, 0x65, 0x61, 0x64,
				0x62, 0x65, 0x65, 0x66,
			},
		}, {
			"has-tmt",
			NewAspActiveAck(
				params.NewTrafficModeType(2),
				nil,
				nil,
			),
			[]byte{
				// Header
				0x01, 0x00, 0x04, 0x03, 0x00, 0x00, 0x00, 0x18,
				// TrafficModeType
				0x00, 0x0b, 0x00, 0x08, 0x00, 0x00, 0x00, 0x02,
			},
		}, {
			"has-rc",
			NewAspActiveAck(
				nil,
				params.NewRoutingContext(1),
				nil,
			),
			[]byte{
				// Header
				0x01, 0x00, 0x04, 0x03, 0x00, 0x00, 0x00, 0x18,
				// RoutingContext
				0x00, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"has-info",
			NewAspActiveAck(
				nil,
				nil,
				params.NewInfoString("deadbeef"),
			),
			[]byte{
				// Header
				0x01, 0x00, 0x04, 0x03, 0x00, 0x00, 0x00, 0x1c,
				// InfoString
				0x00, 0x04, 0x00, 0x0c, 0x64, 0x65, 0x61, 0x64,
				0x62, 0x65, 0x65, 0x66,
			},
		}, {
			"has-none",
			NewAspActiveAck(
				nil, nil, nil,
			),
			[]byte{
				// Header
				0x01, 0x00, 0x04, 0x03, 0x00, 0x00, 0x00, 0x10,
			},
		},
	}

	runTests(t, cases, func(b []byte) (serializeable, error) {
		v, err := DecodeAspActiveAck(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
