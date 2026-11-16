package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	typemotion "github.com/YananLee/typemotion-art-svg-go"
)

func main() {
	text := flag.String("text", "", "text to animate (reads stdin if omitted)")
	out := flag.String("out", "", "output SVG file (stdout if omitted)")
	flag.Parse()
	input := *text
	if input == "" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fail(err)
		}
		input = strings.TrimSpace(string(data))
	}
	svg, err := typemotion.Render(input, typemotion.Options{})
	if err != nil {
		fail(err)
	}
	if *out == "" {
		fmt.Println(svg)
		return
	}
	if err := os.WriteFile(*out, []byte(svg), 0644); err != nil {
		fail(err)
	}
}

func fail(err error) { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
