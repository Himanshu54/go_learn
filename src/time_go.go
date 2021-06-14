package main

import (
	"fmt"
	"time"
)
// time.Now(), time.Date() class methods as static methods 
// timeObj.Month(), timeObj.Format() object methods
// format date : Monday Jan 2 3:04:5 2006 -7
func main() {

	today := time.Date(2021, time.June, 10, 10, 03, 0, 0, time.UTC)
	fmt.Println("time IST:", today)

  now := time.Now()
  fmt.Println("now: ",now)

  fmt.Println("Months :",now.Month())

  longFormat := "Monday, January 3,2006"
  fmt.Println("Tommorrow is:", today.AddDate(0,0,1).Format(longFormat))
}
