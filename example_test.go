package daterange_test

import (
	"fmt"
	"time"

	daterange "github.com/felixenescu/date-range"
)

func ExampleNewDateRange() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.String())
	// Output: {2024-01-26 - 2024-01-28}
}

func ExampleMustNewDateRange() {
	// Create a new DateRange
	dr := daterange.MustNewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.String())
	// Output: {2024-01-26 - 2024-01-28}
}

func ExampleDateRange_From() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.From().Format("2006-01-02"))
	// Output: 2024-01-26
}

func ExampleDateRange_To() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.To().Format("2006-01-02"))
	// Output: 2024-01-28
}

func ExampleDateRange_String() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.String())
	// Output: {2024-01-26 - 2024-01-28}
}

func ExampleDateRange_IsZero() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Time{}, time.Time{})
	fmt.Println(dr.IsZero())
	// Output: true
}

func ExampleDateRange_Contains() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.Contains(time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC)))
	// Output: true
}

func ExampleDateRange_Overlaps() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.Overlaps(daterange.NewDateRange(time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC))))
	// Output: true
}

func ExampleDateRange_Includes() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.Includes(daterange.NewDateRange(time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC))))
	// Output: true
}

func ExampleDateRange_Intersection() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.Intersection(daterange.NewDateRange(time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC))).String())
	// Output: {2024-01-27 - 2024-01-28}
}

func ExampleDateRange_Union() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.Union(daterange.NewDateRange(time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC))).String())
	// Output: [{2024-01-26 - 2024-01-29}]
}

func ExampleDateRange_Difference() {
	// Create a new DateRange
	dr := daterange.NewDateRange(time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC))
	fmt.Println(dr.Difference(daterange.NewDateRange(time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC))).String())
	// Output: [{2024-01-16 - 2024-01-19} {2024-01-25 - 2024-01-28}]
}

func ExampleNewDateRanges() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
		daterange.NewDateRange(time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println(drs.String())
	// Output: [{2024-01-26 - 2024-01-31}]
}

func ExampleDateRanges_ToSlice() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
		daterange.NewDateRange(time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println(drs.ToSlice())
	// Output: [{2024-01-26 - 2024-01-31}]
}

func ExampleDateRanges_String() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
		daterange.NewDateRange(time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println(drs.String())
	// Output: [{2024-01-26 - 2024-01-31}]
}

func ExampleDateRanges_IsZero() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges()
	fmt.Println(drs.IsZero())
	// Output: true
}

func ExampleDateRanges_Len() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
		daterange.NewDateRange(time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println(drs.Len())
	// Output: 1
}

func ExampleDateRanges_FirstDate() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println(drs.FirstDate().Format("2006-01-02"))
	// Output: 2024-01-26
}

func ExampleDateRanges_LastDate() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println(drs.LastDate().Format("2006-01-02"))
	// Output: 2024-01-31
}

func ExampleDateRanges_Equal() {
	// Create a new DateRanges
	drs1 := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
		daterange.NewDateRange(time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
	)
	drs2 := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)),
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println(drs1.Equal(drs2))
	// Output: true
}

func ExampleDateRanges_Append() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
	)
	drs.Append(daterange.NewDateRange(time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)))
	fmt.Println(drs.String())
	// Output: [{2024-01-26 - 2024-01-31}]
}

func ExampleDateRanges_Contains() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println(drs.Contains(time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC)))
	// Output: true
}

func ExampleDateRanges_IsAnyDateIn() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println(drs.IsAnyDateIn(
		daterange.NewDateRange(time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 29, 0, 0, 0, 0, time.UTC)),
	))
	// Output: true
}

func ExampleDateRanges_IsAllDatesIn() {
	// Create a new DateRanges
	drs := daterange.NewDateRanges(
		daterange.NewDateRange(time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println(drs.IsAllDatesIn(
		daterange.NewDateRange(time.Date(2024, 1, 27, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC)),
	))
	// Output: true
}
