package samples

import (
	"fmt"
	"strconv"
	"time"
)

func Switches(d string, t string) {

	if t != "" {
		SwitchType(t)
	} else {
		switch d {
		case "day":
			SwitchDay()
		case "weekday":
			SwitchWeekday()
		case "time":
			SwitchTime()
		default:
			fmt.Println("Unknown time type")
		}
	}
}

func SwitchDay() {
	d := time.Now().Weekday()
	switch d {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thursday")
	case time.Friday:
		fmt.Println("Friday")
	case time.Saturday:
		fmt.Println("Saturday")
	case time.Sunday:
		fmt.Println("Sunday")
	default:
		fmt.Println("Unknown day", d)
	}
}

func SwitchWeekday() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is a weekend", time.Now().Weekday())
	default:
		fmt.Println("Today is a weekday", time.Now().Weekday())
	}
}

func SwitchTime() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning", time.Now().Format("15:04:05"))
	case t.Hour() < 18:
		fmt.Println("Good Afternoon", time.Now().Format("15:04:05"))
	default:
		fmt.Println("Good Evening", time.Now().Format("15:04:05"))
	}
}

func SwitchType(s string) {
	switch s {
	case "true", "false":
		b, _ := strconv.ParseBool(s)
		fmt.Printf("Type: %T Value: %v\n", b, b)
	case "Nil":
		fmt.Printf("Type: <nil> Value: %v\n", s)
	default:
		if i, err := strconv.Atoi(s); err == nil {
			fmt.Printf("Type: %T Value: %v\n", i, i)
		} else {
			fmt.Printf("Unknown type %T of %v\n", s, s)
		}
	}
}
