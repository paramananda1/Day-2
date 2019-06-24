package main

import (
	"fmt"
	"time"
)

func countDays() {

	var secondtry string
	var m1, m2, y1, y2, d1, d2 int
	fmt.Println("Please First date (YYYY MM  DD): ")
	fmt.Scanf("%d %d %d", &y1, &m1, &d1)
 	fmt.Println("Please Second  date (YYYY MM  DD): ")
	fmt.Scanf("%d %d %d", &y2, &m2, &d2)

	// use of time.Month(int)
	pastDate := time.Date(y1, time.Month(m1), d1, 0, 0, 0, 0, time.UTC)
	now := time.Date(y2, time.Month(m2), d2, 0, 0, 0, 0, time.UTC)
	diff := now.Sub(pastDate)
	days := int(diff.Hours() / 24)
	fmt.Printf("Diffrence in days : %d days\n", days)

	// use of t.Month()
	fmt.Println("Do you want to try once: y/n")
	fmt.Scanf("%s",&secondtry)
	if secondtry == "Y" || secondtry == "y" {
		var startdate, enddate string
		fmt.Println("Please First date (YYYY-MM-DD): ")
		fmt.Scanf("%s", &startdate)
		fmt.Println("Please Second date (YYYY-MM-DD): ")
		fmt.Scanf("%s", &enddate)

		fmt.Println(daysBetween(date(startdate), date(enddate)))
	}
}

func daysBetween(a, b time.Time) int {

	if a.After(b) {
		a, b = b, a
	}

	pastDate := time.Date(a.Year(), a.Month(), a.Day(), 0, 0, 0, 0, time.UTC)
	now := time.Date(b.Year(), b.Month(), b.Day(), 0, 0, 0, 0, time.UTC)
	diff := now.Sub(pastDate)
	days := int(diff.Hours() / 24)
	return days
}

func date(s string) time.Time {
	d, _ := time.Parse("2006-01-02", s)
	return d
}

