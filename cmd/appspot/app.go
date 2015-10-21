package sapinapp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/buildkite/terminal"
	"github.com/gorilla/schema"
	"github.com/moul/go-sapin"
)

func init() {
	http.HandleFunc("/", handler)
}

type Options struct {
	Size  int  `schema:"size"`
	Balls int  `schema:"balls"`
	Star  bool `schema:"star"`
	Emoji bool `schema:"emoji"`
	Color bool `schema:"color"`
}

var htmlTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <title>Sapin</title>
    <meta charset="UTF-8" />
    <link type="text/css" rel="stylesheet" href="terminal.css" media="all" />
  </head>
  <body bgcolor="#171717">
    <div class="term-container">&nbsp;
CONTENT</div>
  </body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	// extract query from url
	u, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Fprintf(w, "URL error: %v:\n", err)
		return
	}

	// parse query
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		fmt.Fprintf(w, "URL query error: %v:\n", err)
		return
	}

	// check if no arguments
	if len(m) == 0 {
		http.Redirect(w, r, "?size=5&balls=4&star=true&emoji=false&color=false", http.StatusFound)
	}

	// unmarshal arguments
	opts := Options{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(&opts, m)
	if err != nil {
		fmt.Fprintf(w, "Parameters error: %v:\n", err)
		return
	}

	// check size argument
	if opts.Size > 10 {
		fmt.Fprintf(w, "Max size is: 10\n")
		return
	}
	if opts.Size < 0 {
		fmt.Fprintf(w, "Min size is: 0\n")
		return
	}

	if opts.Balls < 0 {
		fmt.Fprintf(w, "Min value for balls is 0\n")
		return
	}
	if opts.Balls > 100 {
		fmt.Fprintf(w, "Max value for balls is 100\n")
		return
	}

	sapin := sapin.NewSapin(opts.Size)
	if opts.Star {
		sapin.AddStar()
	}
	sapin.AddBalls(opts.Balls)
	if opts.Emoji {
		sapin.Emojize()
	}
	if opts.Color {
		w.Header().Set("Content-Type", "text/html")
		sapin.Colorize()
		coloredOutput := string(terminal.Render([]byte(sapin.String())))
		html := strings.Replace(htmlTemplate, "CONTENT", coloredOutput, 1)
		fmt.Fprint(w, html)
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "%s\n", sapin)
	}
}
