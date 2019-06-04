package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

func main() {
	time, _ := ntp.Time("ntp1.stratum2.ru")
	fmt.Println("Time Now: ", time)
}
