package quote

import (
	"fmt"
	"testing"
)

func TestQuote(t *testing.T) {

	q := Quote{Text: "Brian is awesome", Author: "Brian"}
	s := q.Render()
	fmt.Println(s)
}
