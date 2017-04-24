//
// schema.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"encoding/json"
	"io"
)

type RootSchema struct {
	Version string `json:"$schema,omitempty"`
	Schema
}

func NewRootSchema(r io.Reader) (*RootSchema, error) {
	ret := new(RootSchema)
	err := json.NewDecoder(r).Decode(ret)
	if err != nil {
		return nil, err
	}
	if ret.Version == "" {
		ret.Version = "http://json-schema.org/draft-04/schema#"
	}
	return ret, nil
}

type Schema struct {
	ID  string `json:"id,omitempty"`
	Ref string `json:"$ref,omitempty"`
	Metadata
	Constraints
	rest map[string]json.RawMessage
}

func (s *Schema) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, s)
	if err != nil {

	}
}
