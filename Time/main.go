package main

import (
	"fmt"
	"time"
)

func main() {
	timing := time.Date(2022, time.January, 10, 10, 23, 56, 56, time.UTC)
	fmt.Println(timing.Format("01-02-2006 Monday"))

}
