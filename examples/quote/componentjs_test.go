// build +js
package quote

import (
	"fmt"
	"testing"
)

func TestJSQuote(t *testing.T) {

	q := Quote{Text: "Client is awesome", Author: "Brian"}
	s := q.Render()
	fmt.Println(s)
}
