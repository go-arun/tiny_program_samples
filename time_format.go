package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	fmt.Println(now) // Local time

	now = time.Now().UTC() // Here changed to UTC time
	fmt.Println(now)

	now = time.Now()
//	now = now.Format("2006-01-02T15:04:05Z")
	fmt.Println(now.Format("2006-01-02T15:04:05Z")) // by modifying this order we can alter format
	
	fmt.Println(now.Format("02-01-2006"))
	fmt.Println(now.Format("02-01-2006T15:04"))
	fmt.Println(now.Format("02-01-2006 15:04"))
}

/*
GMT is a time zone officially used in some European and African countries. The time can be displayed using both the 24-hour format (0 - 24) or the 12-hour format (1 - 12 am/pm).

UTC is not a time zone, but a time standard that is the basis for civil time and time zones worldwide. This means that no country or territory officially uses UTC as a local time.
*/
