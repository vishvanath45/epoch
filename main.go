package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	args := os.Args

	if len(args) > 1 {
		epochStr := args[1]

		epoch, err := strconv.ParseInt(epochStr, 10, 64)

		if err != nil {
			os.Exit(-1)
		} else {
			dateInUTC := time.Unix(epoch, 0).UTC()

			loc := time.FixedZone("WIB", 7*60*60)
			dateInJKT := time.Unix(epoch, 0).In(loc)

			loc = time.FixedZone("IST", 5*60*60 + 30*60)
			dateInIST := time.Unix(epoch, 0).In(loc)

			fmt.Println("UTC\t", dateInUTC)
			fmt.Println("IST\t", dateInIST)
			fmt.Println("JKT\t", dateInJKT)
		}
	} else {
		epochTime := time.Now().Unix()
		fmt.Println(epochTime)
	}
}
