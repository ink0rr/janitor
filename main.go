package main

import (
	"fmt"
	"janitor/janitor"
	"time"
)

func main() {
	start := time.Now()

	janitor.CheckBP()
	janitor.CheckRP()

	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}
