/*
 * Copyright (c) 2022 VMware, Inc. or its affiliates
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func _GetServerPort() int {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		intPort, err := strconv.Atoi(port)
		if err != nil {
			log.Fatal("Cannot parse environment variable: SERVER_PORT")
		}
		return intPort
	}
	return 8080
}

func _HandleGreetings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Hello world!")
}

func main() {
	log.Printf("Process PID: %d", os.Getpid())

	http.HandleFunc("/greetings", _HandleGreetings)

	static := http.FileServer(http.Dir("static"))
	http.Handle("/", static)

	port := _GetServerPort()
	log.Printf("Listening on port: %d", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("Cannot start web server", err)
	}
}
