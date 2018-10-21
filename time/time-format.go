package main

import (
	"fmt"
	tm "time"
)

func main() {
	p := fmt.Println

	t := tm.Now()
	p(t)
	p(t.Format(tm.RFC3339))

	d, err := tm.ParseDuration("2h45m")
	p(d.Minutes(), err)

	t, err = tm.Parse(tm.RFC3339, "1991-11-06T00:00:00Z")
	p(t, err)

	d = tm.Since(t)
	p(d.Hours() / 24 / 365)

	p(tm.RFC3339)
}
