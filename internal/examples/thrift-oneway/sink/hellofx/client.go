// Code generated by thriftrw-plugin-yarpc
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

package hellofx

import (
	fx "go.uber.org/fx"
	yarpc "go.uber.org/yarpc"
	thrift "go.uber.org/yarpc/encoding/thrift"
	helloclient "go.uber.org/yarpc/internal/examples/thrift-oneway/sink/helloclient"
)

// Params defines the dependencies for the Hello client.
type Params struct {
	fx.In

	Provider yarpc.ClientConfig
}

// Result defines the output of the Hello client module. It provides a
// Hello client to an Fx application.
type Result struct {
	fx.Out

	Client helloclient.Interface

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// Client provides a Hello client to an Fx application using the given name
// for routing.
//
// 	fx.Provide(
// 		hellofx.Client("..."),
// 		newHandler,
// 	)
func Client(name string, opts ...thrift.ClientOption) interface{} {
	return func(p Params) Result {
		client := helloclient.New(p.Provider.ClientConfig(name), opts...)
		return Result{Client: client}
	}
}
