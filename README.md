# TypeMotion.Art SVG for Go

A dependency-free Go library and CLI that makes a basic animated text SVG. The SVG fades and lifts up to 12 lines of text. It runs locally and does not use the TypeMotion.Art service. For the full visual editor and video export, use [TypeMotion.Art](https://www.typemotion.art/).

## Install

```sh
go get github.com/YananLee/typemotion-art-svg-go@v0.1.0
```

## Go API

```go
svg, err := typemotion.Render("Hello world", typemotion.Options{})
if err != nil { panic(err) }
_ = os.WriteFile("title.svg", []byte(svg), 0644)
```

Import `typemotion "github.com/YananLee/typemotion-art-svg-go"` and `os`. Options include canvas size, font size, colors, and duration.

## CLI

```sh
go run github.com/YananLee/typemotion-art-svg-go/cmd/typemotion-art-svg@v0.1.0 -text "Hello world" -out title.svg
```

Omit `-text` to read stdin or `-out` to print SVG. Embed the generated SVG inline or as an image on a web page. SVG animation support varies outside browsers.

## License

MIT.
