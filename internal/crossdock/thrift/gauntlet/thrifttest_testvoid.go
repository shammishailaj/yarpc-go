// Code generated by thriftrw v1.17.0. DO NOT EDIT.
// @generated

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gauntlet

import (
	fmt "fmt"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// ThriftTest_TestVoid_Args represents the arguments for the ThriftTest.testVoid function.
//
// The arguments for testVoid are sent and received over the wire as this struct.
type ThriftTest_TestVoid_Args struct {
}

// ToWire translates a ThriftTest_TestVoid_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestVoid_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ThriftTest_TestVoid_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestVoid_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestVoid_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestVoid_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestVoid_Args
// struct.
func (v *ThriftTest_TestVoid_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("ThriftTest_TestVoid_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestVoid_Args match the
// provided ThriftTest_TestVoid_Args.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestVoid_Args) Equals(rhs *ThriftTest_TestVoid_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ThriftTest_TestVoid_Args.
func (v *ThriftTest_TestVoid_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "testVoid" for this struct.
func (v *ThriftTest_TestVoid_Args) MethodName() string {
	return "testVoid"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *ThriftTest_TestVoid_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// ThriftTest_TestVoid_Helper provides functions that aid in handling the
// parameters and return values of the ThriftTest.testVoid
// function.
var ThriftTest_TestVoid_Helper = struct {
	// Args accepts the parameters of testVoid in-order and returns
	// the arguments struct for the function.
	Args func() *ThriftTest_TestVoid_Args

	// IsException returns true if the given error can be thrown
	// by testVoid.
	//
	// An error can be thrown by testVoid only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for testVoid
	// given the error returned by it. The provided error may
	// be nil if testVoid did not fail.
	//
	// This allows mapping errors returned by testVoid into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// testVoid
	//
	//   err := testVoid(args)
	//   result, err := ThriftTest_TestVoid_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from testVoid: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*ThriftTest_TestVoid_Result, error)

	// UnwrapResponse takes the result struct for testVoid
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if testVoid threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := ThriftTest_TestVoid_Helper.UnwrapResponse(result)
	UnwrapResponse func(*ThriftTest_TestVoid_Result) error
}{}

func init() {
	ThriftTest_TestVoid_Helper.Args = func() *ThriftTest_TestVoid_Args {
		return &ThriftTest_TestVoid_Args{}
	}

	ThriftTest_TestVoid_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	ThriftTest_TestVoid_Helper.WrapResponse = func(err error) (*ThriftTest_TestVoid_Result, error) {
		if err == nil {
			return &ThriftTest_TestVoid_Result{}, nil
		}

		return nil, err
	}
	ThriftTest_TestVoid_Helper.UnwrapResponse = func(result *ThriftTest_TestVoid_Result) (err error) {
		return
	}

}

// ThriftTest_TestVoid_Result represents the result of a ThriftTest.testVoid function call.
//
// The result of a testVoid execution is sent and received over the wire as this struct.
type ThriftTest_TestVoid_Result struct {
}

// ToWire translates a ThriftTest_TestVoid_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestVoid_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ThriftTest_TestVoid_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestVoid_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestVoid_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestVoid_Result) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestVoid_Result
// struct.
func (v *ThriftTest_TestVoid_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("ThriftTest_TestVoid_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestVoid_Result match the
// provided ThriftTest_TestVoid_Result.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestVoid_Result) Equals(rhs *ThriftTest_TestVoid_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ThriftTest_TestVoid_Result.
func (v *ThriftTest_TestVoid_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "testVoid" for this struct.
func (v *ThriftTest_TestVoid_Result) MethodName() string {
	return "testVoid"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *ThriftTest_TestVoid_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
