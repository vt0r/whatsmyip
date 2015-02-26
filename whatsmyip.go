/*
 * Sal's What's My IP Daemon
 * --------------------------
 * It listens for requests, then it grabs
 * the IP from X-Real-IP and spits it back.
 * --------------------------
 * Copyright (c) 2015, Salvatore LaMendola <salvatore@lamendola.me>
 * All rights reserved.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func spitIP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, req.Header.Get("X-Real-IP"))
}

func main() {
	http.HandleFunc("/", spitIP)
	err := http.ListenAndServe("127.0.0.1:8999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
