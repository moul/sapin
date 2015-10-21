package sapinapp

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/moul/go-sapin"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Fprintf(w, "Invalid URL: %v\n", err)
		return
	}
	m, _ := url.ParseQuery(u.RawQuery)

	size := 5
	if len(m["size"]) > 0 {
		size, err = strconv.Atoi(m["size"][0])
		if err != nil {
			fmt.Fprintf(w, "Invalid size %q: %v\n", m["size"][0], err)
			return
		}
		if size > 10 {
			fmt.Fprintf(w, "Max size is: 10\n")
			return
		}
		if size < 0 {
			fmt.Fprintf(w, "Min size is: 0\n")
			return
		}
	}
	sapin := sapin.NewSapin(size)
	fmt.Fprintf(w, "%s\n", sapin)
}
