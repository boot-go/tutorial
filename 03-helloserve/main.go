/*
 * Copyright (c) 2021 boot-go
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package main

import (
	"github.com/boot-go/boot"
	"github.com/boot-go/stack/server/httpcmp"
	"io"
	"net/http"
)

// init() registers a factory method, which creates a hello component
func init() {
	boot.Register(func() boot.Component {
		return &hello{}
	})
}

// hello is a very simple http server example.
// It requires the Eventbus and the httpServer component. Both components
// are injected by the boot framework automatically
type hello struct {
	Eventbus boot.EventBus  `boot:"wire"`
	Server   httpcmp.Server `boot:"wire"`
}

// Init is the constructor of the component. The handler registration takes place here.
func (h *hello) Init() {
	// Subscribe to the registration event
	h.Eventbus.Subscribe(func(event httpcmp.InitializedEvent) {
		h.Server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			io.WriteString(writer, "boot-go says: 'Hello World'\n")
		})
	})
}

// Start the example and test with 'curl localhost:8080'
func main() {
	boot.Go()
}
