package sapinapp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/buildkite/terminal"
	"github.com/gorilla/schema"
	"github.com/moul/sapin"
)

func init() {
	http.HandleFunc("/", handler)
}

type Options struct {
	Size     int  `schema:"size"`
	Balls    int  `schema:"balls"`
	Garlands int  `schema:"garlands"`
	Star     bool `schema:"star"`
	Emoji    bool `schema:"emoji"`
	Color    bool `schema:"color"`
	Presents bool `schema:"presents"`
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
    <a href="https://github.com/moul/sapin"><img style="position: absolute; top: 0; right: 0; border: 0;" src="https://camo.githubusercontent.com/52760788cde945287fbb584134c4cbc2bc36f904/68747470733a2f2f73332e616d617a6f6e6177732e636f6d2f6769746875622f726962626f6e732f666f726b6d655f72696768745f77686974655f6666666666662e706e67" alt="Fork me on GitHub" data-canonical-src="https://s3.amazonaws.com/github/ribbons/forkme_right_white_ffffff.png"></a>
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
		http.Redirect(w, r, "?size=5&balls=4&star=true&emoji=false&color=false&presents=true&garlands=5", http.StatusFound)
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
	if opts.Garlands < 0 {
		fmt.Fprintf(w, "Min value for garlands is 0\n")
		return
	}
	if opts.Garlands > 50 {
		fmt.Fprintf(w, "Max value for garlands is 50\n")
		return
	}

	sapin := sapin.NewSapin(opts.Size)
	if opts.Star {
		sapin.AddStar()
	}
	sapin.AddBalls(opts.Balls)
	sapin.AddGarlands(opts.Garlands)
	if opts.Emoji {
		sapin.Emojize()
	}
	if opts.Color {
		w.Header().Set("Content-Type", "text/html")
		sapin.Colorize()
		if opts.Presents {
			sapin.AddPresents()
		}
		coloredOutput := string(terminal.Render([]byte(sapin.String())))
		html := strings.Replace(htmlTemplate, "CONTENT", coloredOutput, 1)
		fmt.Fprint(w, html)
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		if opts.Presents {
			sapin.AddPresents()
		}
		fmt.Fprintf(w, "%s\n", sapin)
	}
}
