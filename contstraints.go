//
// type.go
// Copyright (C) 2017 weirdgiraffe <giraffe@cyberzoo.xyz>
//
// Distributed under terms of the MIT license.
//

package main

import "encoding/json"

type Constraints struct {
	// MultipleOf MUST be a number, strictly greater than 0
	MultipleOf *uint `json:"multipleOf,omitempty"`

	// Maximum MUST be a number
	Maximum *int `json:"maximum,omitempty"`

	// ExclusiveMaximum MUST be a boolean
	// an undefined value is the same as false
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`

	// Minimum MUST be a number
	Minimum *int `json:"minimum,omitempty"`

	// ExclusiveMinimum MUST be a boolean
	// an undefined value is the same as false.
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`

	// MaxLength MUST be an integer
	// this integer MUST be greater than, or equal to, 0
	MaxLength *uint `json:"maxLength,omitempty"`

	// MinLength MUST be an integer
	// this integer MUST be greater than, or equal to, 0
	// if absent, may be considered as being present with integer value 0
	MinLength uint `json:"minLength,omitempty"`

	// Pattern MUST be a string
	// this string SHOULD be a valid regular expression
	// according to the ECMA 262 regular expression dialect.
	Pattern string `json:"pattern,omitempty"`

	// Items MUST be either a schema or array of schemas
	Items json.RawMessage `json:"items,omitempty"`

	// AdditionalItems MUST be either a boolean or an object
	// if it is an object, this object MUST be a valid JSON Schema
	AdditionalItems json.RawMessage `json:"additionalItems,omitempty"`

	// MaxItems MUST be an integer
	// this integer MUST be greater than, or equal to, 0
	MaxItems *uint `json:"maxItems,omitempty"`

	// MinItems MUST be an integer
	// this integer MUST be greater than, or equal to, 0
	// if this keyword is not present, it may be considered present with a value of 0
	MinItems uint `json:"minItems,omitempty"`

	// UniqueItems MUST be a boolean
	// If not present, this keyword may be considered present with boolean value false
	UniqueItems bool `json:"uniqueItems,omitempty"`

	// MaxProperties MUST be an integer
	// this integer MUST be greater than, or equal to, 0
	MaxProperties *uint `json:"maxProperties,omitempty"`

	// MinProperties MUST be an integer
	// this integer MUST be greater than, or equal to, 0
	// if this keyword is not present, it may be considered present with a value of 0
	MinProperties uint `json:"minProperties,omitempty"`

	// Required MUST be an array
	// this array MUST have at least one element
	// elements of this array MUST be strings, and MUST be unique
	Required []string `json:"required,omitempty"`

	// Properties MUST be an object
	// each value of this object MUST be an object, and each object MUST be a valid JSON Schema
	// if absent, it can be considered the same as an empty object
	Properties json.RawMessage `json:"properties,omitempty"`

	// PatternProperties MUST be an object
	// each property name of this object SHOULD be a valid regular expression, according to the ECMA 262 regular expression dialect.
	// each property value of this object MUST be an object, and each object MUST be a valid JSON Schema
	// if absent, it can be considered the same as an empty object
	PatternProperties json.RawMessage `json:"patternProperties,omitempty"`

	// AdditionalItems MUST be a boolean or a schema
	// if "additionalProperties" is absent, it may be considered present with an empty schema as a value.
	AdditionalProperties json.RawMessage `json:"additionalProperties,omitempty"`

	// Dependencies MUST be an object
	// each property specifies a dependency
	// each dependency value MUST be an object or an array
	// if the dependency value is an object, it MUST be a valid JSON Schema
	// if the dependency value is an array, it MUST have at least one element, each element MUST be a string, and elements in the array MUST be unique
	Dependencies json.RawMessage `json:"dependencies,omitempty"`

	// Enum MUST be an array. This array SHOULD have at least one element. Elements in the array SHOULD be unique
	// elements in the array MAY be of any type, including null
	Enum []interface{} `json:"enum,omitempty"`

	// ValueType MUST be either a string or an array
	// if it is an array, elements of the array MUST be strings and MUST be unique.
	// string values MUST be one of the seven primitive types defined by the core specification.
	ValueType interface{} `json:"type,omitempty"`

	// AllOf MUST be an array
	// this array MUST have at least one element
	// elements of the array MUST be objects. Each object MUST be a valid JSON Schema.
	AllOf []interface{} `json:"allOf,omitempty"`

	// AnyOf MUST be an array. This array MUST have at least one element.
	// Elements of the array MUST be objects. Each object MUST be a valid JSON Schema.
	AnyOf []interface{} `json:"anyOf,omitempty"`

	// OneOf MUST be an array. This array MUST have at least one element.
	// Elements of the array MUST be objects. Each object MUST be a valid JSON Schema.
	OneOf []interface{} `json:"oneOf,omitempty"`

	// Not MUST be an object. This object MUST be a valid JSON Schema.
	Not json.RawMessage `json:"not,omitempty"`

	// Definitions MUST be an object.
	// Each member value of this object MUST be a valid JSON Schema.
	Definitions json.RawMessage `json:"definitions,omitempty"`
}
