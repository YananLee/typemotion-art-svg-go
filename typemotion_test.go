package typemotion

import (
	"strings"
	"testing"
)

func TestRenderEscapesText(t *testing.T) {
	svg, err := Render(`<script>&`, Options{})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(svg, `<script>`) || !strings.Contains(svg, `&lt;script&gt;&amp;`) {
		t.Fatal("text was not escaped")
	}
}

func TestRenderRejectsInvalidColor(t *testing.T) {
	if _, err := Render("hello", Options{Background: `red" onload="alert(1)`}); err == nil {
		t.Fatal("expected error")
	}
}
