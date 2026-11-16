// Package typemotion renders a small animated text SVG for web embeds.
package typemotion

import (
	"fmt"
	"html"
	"math"
	"strings"
)

// Options controls the SVG canvas and text appearance. Zero values use defaults.
type Options struct {
	Width, Height, FontSize int
	TextColor, Background   string
	Duration                float64
}

// Render creates a standalone SVG with a simple fade and rise animation.
func Render(text string, options Options) (string, error) {
	if strings.TrimSpace(text) == "" {
		return "", fmt.Errorf("text must not be empty")
	}
	if options.Width == 0 {
		options.Width = 800
	}
	if options.Height == 0 {
		options.Height = 450
	}
	if options.FontSize == 0 {
		options.FontSize = 64
	}
	if options.Duration == 0 {
		options.Duration = 1.2
	}
	if options.TextColor == "" {
		options.TextColor = "#ffffff"
	}
	if options.Background == "" {
		options.Background = "#151624"
	}
	if options.Width < 1 || options.Height < 1 || options.FontSize < 1 || options.Duration <= 0 || math.IsNaN(options.Duration) || math.IsInf(options.Duration, 0) {
		return "", fmt.Errorf("invalid size or duration")
	}
	if !hexColor(options.TextColor) || !hexColor(options.Background) {
		return "", fmt.Errorf("colors must be #RRGGBB")
	}
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	if len(lines) > 12 {
		return "", fmt.Errorf("maximum 12 lines")
	}
	var out strings.Builder
	fmt.Fprintf(&out, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d" role="img" aria-label="%s">`, options.Width, options.Height, options.Width, options.Height, html.EscapeString(text))
	fmt.Fprintf(&out, `<rect width="100%%" height="100%%" fill="%s"/>`, options.Background)
	lineHeight := float64(options.FontSize) * 1.2
	startY := float64(options.Height)/2 - float64(len(lines)-1)*lineHeight/2
	for i, line := range lines {
		y := startY + float64(i)*lineHeight
		fmt.Fprintf(&out, `<text x="50%%" y="%.1f" text-anchor="middle" dominant-baseline="middle" font-family="Arial, sans-serif" font-size="%d" fill="%s" opacity="0">%s<animate attributeName="opacity" from="0" to="1" dur="%.2fs" begin="%.2fs" fill="freeze"/><animateTransform attributeName="transform" type="translate" from="0 18" to="0 0" dur="%.2fs" begin="%.2fs" fill="freeze"/></text>`, y, options.FontSize, options.TextColor, html.EscapeString(line), options.Duration, float64(i)*0.18, options.Duration, float64(i)*0.18)
	}
	out.WriteString(`</svg>`)
	return out.String(), nil
}

func hexColor(value string) bool {
	if len(value) != 7 || value[0] != '#' {
		return false
	}
	for _, c := range value[1:] {
		if !strings.ContainsRune("0123456789abcdefABCDEF", c) {
			return false
		}
	}
	return true
}
