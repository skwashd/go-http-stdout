// Copyright 2018 Dave Hall <skwashd@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of
// the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
// WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var listenFlag = flag.String("listen", ":8888", "Port to listen on (supports address:port)")

// Pretty prints JSON to stdout
func HandleRequest(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		log.Println("Invalid request method (" + request.Method + ") received at " + request.URL.Path + " from IP: " + request.RemoteAddr)
		fmt.Fprintln(writer, "\"POST only\"")
		return
	}

	log.Println("Request received at " + request.URL.Path + " from IP: " + request.RemoteAddr)

	// We always return JSON.
	writer.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Panic(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, body, "", "  ")
	if err != nil {
		log.Panic(err)
	}

	log.Println(out.String())

	// Return a null, so the client gets something back.
	fmt.Fprintln(writer, "null")
}

func main() {
	log.SetOutput(os.Stdout)
	flag.Parse()
	http.HandleFunc("/", HandleRequest)
	log.Println("Server listening on " + *listenFlag)
	http.ListenAndServe(*listenFlag, nil)
}
