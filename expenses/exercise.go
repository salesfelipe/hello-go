package expenses

import (
	"errors"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(inp []Record, predicate func(Record) bool) []Record {
	var result []Record = []Record{}

	for _, rec := range inp {
		if predicate(rec) {
			result = append(result, rec)
		}
	}

	return result
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(rec Record) bool {
		return rec.Day >= p.From && rec.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(rec Record) bool {
		return rec.Category == c
	}
}

func sumTotal(in []Record) float64 {
	var total float64 = 0

	for _, rec := range in {
		total += rec.Amount
	}

	return total
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	filteredRecords := Filter(in, ByDaysPeriod(p))

	return sumTotal(filteredRecords)
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	recordsInCategory := Filter(in, ByCategory(c))

	if len(recordsInCategory) == 0 {
		return 0, errors.New("no items found")
	}

	recordsInRange := Filter(recordsInCategory, ByDaysPeriod(p))

	return sumTotal(recordsInRange), nil
}
