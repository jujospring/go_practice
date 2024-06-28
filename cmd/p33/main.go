package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		for costs[i].day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[costs[i].day] += costs[i].value
	}
	return costsByDay
}
