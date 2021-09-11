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
	"log"
)

// hello is a simple component.
type hello struct{}

// say is the interface which can be used by other components.
type say interface {
	hello()
}

// hello is the provided interface implementation.
func (c *hello) hello() {
	log.Printf("boot-go says > 'Hello World'\n")
}

// init() registers both factory methods, which create a hello and example component.
func init() {
	boot.Register(func() boot.Component {
		return &hello{}
	})
	boot.Register(func() boot.Component {
		return &example{}
	})
}

// Init is the constructor of the component. It is called to initialize the component itself.
func (w *hello) Init() {
}

// example is component which uses the hello component.
type example struct{
	Say say `boot:"wire"`
}

// Init is the constructor of the component. It is called to initialize the component itself.
func (c *example) Init() {
	c.Say.hello()
}

// Start the example and exit after the component was completed
func main() {
	boot.Go()
}
