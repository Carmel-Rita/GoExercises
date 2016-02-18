// redis service will act as clone of redis, with the following services like PUT, GET, DELETE and COUNT.
package main

import (
	"fmt"
	"net/http"
	"strings"
)

var redismap map[string]string

func handler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	method := r.Method

	if method == "GET" {
		fmt.Fprintf(w, "%s", redismap[key])
	}

	if method == "PUT" {
		keyvalue := strings.Split(key, "/")
		redismap[keyvalue[0]] = keyvalue[1]
		result := "key: " + keyvalue[0] + " value: " + keyvalue[1] + " saved."
		fmt.Fprint(w, result)
	}

	if method == "DELETE" {
		delete(redismap, key)
		result := "key: " + key + " deleted."
		fmt.Fprint(w, result)
	}

	if method == "COUNT" {
		count := 0
		if key == "" {
			count=len(redismap)
		}else{
		for mapkey, _ := range redismap {
			if strings.Contains(mapkey, key) {
				count++
			}
		}
	}
		fmt.Fprint(w, count)
	}
}

func main() {
	redismap = make(map[string]string)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
