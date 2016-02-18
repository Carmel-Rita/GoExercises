// when a request is raised, md5http service returns the md5 hash code of the request body content.
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	message, _ := ioutil.ReadAll(r.Body)
	msgbyte := []byte(message)
	encrypted := md5.Sum(msgbyte)
	md5val := hex.EncodeToString(encrypted[:])
	fmt.Fprintf(w, md5val)
	}

func main() {
	http.HandleFunc("/md5", handler)
	http.ListenAndServe(":8888", nil)
}
