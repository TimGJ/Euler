package main

import "fmt"

/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

type Month int
type DayOfWeek int

const (
	Sunday DayOfWeek = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	January Month = iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (d DayOfWeek) String() string {
	switch d {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	default:
		return "** UNKNOWN DAY **"
	}

}

func (m Month) String() string {
	switch m {
	case January:
		return "January"
	case February:
		return "February"
	case March:
		return "March"
	case April:
		return "April"
	case May:
		return "May"
	case June:
		return "June"
	case July:
		return "July"
	case August:
		return "August"
	case September:
		return "September"
	case October:
		return "October"
	case November:
		return "November"
	case December:
		return "December"
	default:
		return "** UNKNOWN MONTH **"
	}
}

func IsLeap(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}

func DaysInMonth(month Month, year int) int {
	// Return the number of days in a given month

	switch month {
	case September, April, June, November:
		return 30
	case February:
		if IsLeap(year) {
			return 29
		} else {
			return 28
		}
	default:
		return 31
	}
}

type FirstOfMonth struct {
	Year     int       // The year e.g. 1964
	Month    Month     // The Month e.g. May
	FirstDay DayOfWeek //What day of the week the 1st fell on
}

func (f FirstOfMonth) String() string {
	return fmt.Sprintf("%-10s %d starts on %-9s", f.Month, f.Year, f.FirstDay)
}

func GetDay(days int) DayOfWeek {
	// Returns the day of the week for a certain offset numnber of days from 01 January 1900
	const firstday = Monday

	return DayOfWeek((int(firstday) + days) % 7)

}
func main() {

	months := []FirstOfMonth{}
	var offset int = 365
	var sundays int
	for year := 1901; year <= 2000; year++ {
		for month := January; month <= December; month++ {
			thismonth := FirstOfMonth{year, month, GetDay(offset)}
			if thismonth.FirstDay == Sunday {
				sundays++
			}
			months = append(months, thismonth)
			offset += DaysInMonth(month, year)
			fmt.Printf("%s %6d\n", thismonth, offset)
		}
	}
	fmt.Println("The number of months beginning on a Sunday was", sundays)
}
