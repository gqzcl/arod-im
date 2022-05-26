// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package json

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Marshaler struct{}

func (j Marshaler) Marshal(v interface{}) ([]byte, error) {
	if pb, ok := v.(proto.Message); ok {
		return protojson.Marshal(pb)
	}
	return json.Marshal(v)
}

func (j Marshaler) Unmarshal(d []byte, v interface{}) error {
	if pb, ok := v.(proto.Message); ok {
		return protojson.Unmarshal(d, pb)
	}
	return json.Unmarshal(d, v)
}

func (j Marshaler) Name() string {
	return "json"
}
