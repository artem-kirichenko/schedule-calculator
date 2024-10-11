package main

import (
	"fmt"
	"time"
)

type Task struct {
	Name     string
	Duration time.Duration
}

func main() {
	fmt.Println("Building a plan for the day")
	var startHour, startMinute int
	fmt.Print("Enter hours: ")
	fmt.Scanln(&startHour)
	fmt.Print("Enter minutes: ")
	fmt.Scanln(&startMinute)
	fmt.Println("--- Schedule: ---")

	tasks := []Task{
		{"wash and exercise", 30 * time.Minute},
		{"English Vanessa", 15 * time.Minute},
		{"Developing 1", 55 * time.Minute},
		{"Breakfast (20 min break)", 0 * time.Minute},
		{"Developing 2", 55 * time.Minute},
		{"Developing 3", 55 * time.Minute},
		{"Lunch and nap time", 70 * time.Minute},
		{"English Vanessa", 25 * time.Minute},
		{"English Rachel", 30 * time.Minute},
		{"Pick up a child from school", 30 * time.Minute},
		{"Gym", 30 * time.Minute},
		{"Dinner", 30 * time.Minute},
		{"English Vanessa", 30 * time.Minute},
	}

	currentTime := time.Date(0, 1, 1, startHour, startMinute, 0, 0, time.UTC)

	for i, task := range tasks {
		startHour12 := currentTime.Hour()
		if startHour12 > 12 {
			startHour12 -= 12
		}
		if startHour12 == 0 {
			startHour12 = 12
		}
		period := "AM"
		if currentTime.Hour() >= 12 {
			period = "PM"
		}

		endHour := currentTime.Add(task.Duration).Hour()
		endHour12 := endHour
		if endHour12 > 12 {
			endHour12 -= 12
		}
		if endHour12 == 0 {
			endHour12 = 12
		}
		endPeriod := "AM"
		if endHour >= 12 {
			endPeriod = "PM"
		}
		fmt.Printf("   %02d:%02d %s - %02d:%02d %s (%d) %s\n",
			startHour12, currentTime.Minute(), period,
			endHour12, currentTime.Add(task.Duration).Minute(), endPeriod,
			int(task.Duration.Minutes()),
			task.Name,
		)
		currentTime = currentTime.Add(task.Duration)

		if i < len(tasks)-1 {
			breakDuration := 10 * time.Minute
			currentTime = currentTime.Add(breakDuration)
		}
	}
}
