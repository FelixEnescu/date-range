[![Build Status](https://github.com/felixenescu/date-range/actions/workflows/ci.yml/badge.svg)](https://github.com/felixenescu/date-range/actions/workflows/ci.yml/badge.svg)
[![Code coverage](https://codecov.io/gh/FelixEnescu/date-range/graph/badge.svg?token=Q4A0CDVYHB)](https://codecov.io/gh/FelixEnescu/date-range)
[![Go Report Card](https://goreportcard.com/badge/github.com/felixenescu/date-range)](https://goreportcard.com/report/github.com/felixenescu/date-range)
[![GoDoc](https://godoc.org/github.com/felixenescu/date-range?status.svg)](http://godoc.org/github.com/felixenescu/date-range)


# date-range

## Introduction

**daterange** is a powerful and intuitive Go package designed for handling date intervals efficiently and effectively. This library simplifies operations such as comparing dates, checking overlaps, and processing date ranges in Go applications.

## Features
Easy comparison of date ranges
Checking for overlaps and inclusions between date intervals
Automatic normalization of date ranges
Processing and manipulation of date intervals

## Potential Use Cases

 - **Event Planning and Scheduling:** The Date-Range library can be used in applications that deal with event scheduling, helping to check for conflicts in dates and times.

 - **Reservation Systems:** For systems that handle reservations, such as hotel booking or venue reservation systems, this library can assist in managing availability and booking dates.

 - **Data Analysis and Reporting:** In data analysis, particularly when dealing with time series data, the Date-Range library can help in segmenting data into specific time intervals for better insights.

 - **Financial Applications:** In financial applications, the library can be used to calculate durations for interest calculations, investment maturity periods, and more.

 - **Educational Tools:** For educational software that schedules courses or exams, the library can help in managing and visualizing course timelines.

 - **Healthcare Applications:** In healthcare applications, the library can be used to manage patient appointments, medication schedules, and other time-sensitive data.

## Usage


### DateRange

#### Overview

**DateRange** represents an **inclusive** range of dates. It is defined by two time.Time values, `from` and `to`.

**Note:** Only the date portion of the time.Time values is compared. The time portion is ignored.

#### Constructors

 - **NewDateRange(from, to time.Time):** Creates a new `DateRange` instance. The input dates are automatically ordered.
 - **MustNewDateRange(from, to time.Time):** Similar to `NewDateRange` but panics if the `from` date is after the `to` date.

#### Methods

 - **String() string:** Returns a string representation of the `DateRange`.
 - **IsZero() bool:** Checks if both dates in the range are zero values.
 - **Contains(date time.Time) bool:** Returns true if the given date is within the range.
 - **Overlaps(other DateRange) bool:** Checks if the given `DateRange` overlaps with this range.
 - **Includes(other DateRange) bool:** Checks if the given `DateRange` is included within this range.
 - **Intersection(other DateRange) DateRange:** Returns the intersection of two `DateRanges`.
 - **Union(other DateRange) DateRanges:** Returns a `DateRanges` collection that is the union of two `DateRanges`.
 - **Difference(other DateRange) DateRanges:** Returns a `DateRanges` collection that is the difference between two `DateRanges`.

#### Use Cases and Examples

 - **Checking if a date is within a range**

```go
package main

import (
	"fmt"
	"time"

	dr "github.com/felixenescu/date-range"
)

func main() {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	dateRange := dr.NewDateRange(start, end)

	dateToCheck := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
	fmt.Println("Date within range:", dateRange.Contains(dateToCheck))
	// Date within range: true
}
```

 - **Finding overlap between two date ranges**
```go
package main

import (
	"fmt"
	"time"

	dr "github.com/felixenescu/date-range"
)

func main() {
	firstRange := dr.NewDateRange(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC))
	secondRange := dr.NewDateRange(time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC))

	fmt.Println("Ranges overlap:", firstRange.Overlaps(secondRange))
	// Ranges overlap: true
}
```

 - **Calculating the union of two date ranges**

```go
package main

import (
	"fmt"
	"time"

	dr "github.com/felixenescu/date-range"
)

func main() {
	firstRange := dr.NewDateRange(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC))
	secondRange := dr.NewDateRange(time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC))

	unionRanges := firstRange.Union(secondRange)
	fmt.Println(unionRanges)
	// [{2024-01-01 - 2024-06-30}]
}
```

 - **Calculating the difference of two date ranges**

```go
package main

import (
	"fmt"
	"time"

	dr "github.com/felixenescu/date-range"
)

func main() {
	firstRange := dr.NewDateRange(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC))
	secondRange := dr.NewDateRange(time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC))

	diffRanges := firstRange.Difference(secondRange)
	fmt.Println(diffRanges)
	// [{2024-01-01 - 2024-01-31} {2024-03-11 - 2024-03-31}]
}
```

### DateRanges

#### Overview

**DateRanges** represents a collection of `DateRange` instances. It provides methods for managing and processing multiple date ranges.

#### Constructor

 - **NewDateRanges(dataRanges ...DateRange):** Creates a new `DataRanges` collection with given elements.

#### Methods

 - **String() string:** Returns a string representation of the collection. Items are guaranteed to be sorted and non-overlapping. Any adjacent periods are merged.
 - **ToSlice() []DateRange:** Returns the members of the collection as a slice. Items are guaranteed to be sorted, non-overlapping and non-zero. Any adjacent periods are merged..
 - **IsZero() bool:** Returns true if the collection is empty.
 - **Len() int:** Returns the number of elements in the collection.
 - **FirstDate() time.Time:** Returns the first date of the collection.
 - **LastDate() time.Time:** Returns the last date of the collection.
 - **Equal(other DateRanges) bool:** Returns true if the collection is equal to the given collection.
 - **Append(dataRange ...DateRange):** Adds the given elements to the collection.
 - **Contains(date time.Time) bool:** Returns true if the given date is in the collection.
 - **IsAnyDateIn(date time.Time) bool:** Returns true if any date in the given DateRange is in the collection. Zero DateRange is always considered to be in the collection.
 - **IsAllDatesIn(date time.Time) bool:** Returns true if all dates in the given DateRange are in the collection. Zero DateRange is always considered to be in the collection.
 - **SplitInclusive(date time.Time) (DateRanges, DateRanges):** Splits the collection into two collections at the given date. The given date is included in both collections.

#### Use Cases and Examples

 - **Normalize a list of intervals**

```go
package main

import (
	"fmt"
	"time"

	dr "github.com/felixenescu/date-range"
)

func main() {

	drs := []dr.DateRange{
		dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 11, 0, 0, 0, 0, time.UTC)),
	}
	cleanRanges := dr.NewDateRanges(drs...)
	fmt.Println(cleanRanges)
	// [{2019-01-02 - 2019-01-04} {2019-01-09 - 2019-01-14}]

	// extract normalized intervals
	for idx, dataRange := range cleanRanges.ToSlice() {
		fmt.Println(idx, dataRange.From(), dataRange.To())
	}
	// 0 2019-01-02 00:00:00 +0000 UTC 2019-01-04 00:00:00 +0000 UTC
	// 1 2019-01-09 00:00:00 +0000 UTC 2019-01-14 00:00:00 +0000 UTC
}
```

 - **Check if a date is in a collection**

```go
package main

import (
	"fmt"
	"time"

	dr "github.com/felixenescu/date-range"
)

func main() {

	drs := []dr.DateRange{
		dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 13, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 15, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 24, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 11, 0, 0, 0, 0, time.UTC)),
	}
	reservations := dr.NewDateRanges(drs...)
	fmt.Println(reservations)
	// [{2019-01-09 - 2019-01-15} {2019-01-20 - 2019-01-24}]

	newReservation := dr.NewDateRange(time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 17, 0, 0, 0, 0, time.UTC))
	if reservations.IsAnyDateIn(newReservation) {
		fmt.Println(newReservation, "conflicts with existing reservations")
	} else {
		fmt.Println(newReservation, "does not conflicts with existing reservations")
	}
	// {2019-01-10 - 2019-01-17} conflicts with existing reservations

}
```

 - **Check if all dates are in a collection**

```go
package main

import (
	"fmt"
	"time"

	dr "github.com/felixenescu/date-range"
)

func main() {

	drs := []dr.DateRange{
		dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 13, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 15, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 24, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 11, 0, 0, 0, 0, time.UTC)),
	}
	availableDates := dr.NewDateRanges(drs...)
	fmt.Println(availableDates)
	// [{2019-01-09 - 2019-01-15} {2019-01-20 - 2019-01-24}]

	newReservation := dr.NewDateRange(time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 17, 0, 0, 0, 0, time.UTC))
	if availableDates.IsAllDatesIn(newReservation) {
		fmt.Println(newReservation, "there is availability for new reservations")
	} else {
		fmt.Println(newReservation, "there is no availability for new reservations")
	}
	// {2019-01-10 - 2019-01-17} there is no availability for new reservations
}
```

 - **Split a collection at a given date**

```go
package main

import (
	"fmt"
	"time"

	dr "github.com/felixenescu/date-range"
)

func main() {

	drs := []dr.DateRange{
		dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 13, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 15, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 24, 0, 0, 0, 0, time.UTC)),
		dr.NewDateRange(time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 11, 0, 0, 0, 0, time.UTC)),
	}
	reservations := dr.NewDateRanges(drs...)
	fmt.Println(reservations)
	// [{2019-01-09 - 2019-01-15} {2019-01-20 - 2019-01-24}]

	splitDate := time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)
	before, after := reservations.SplitInclusive(splitDate)
	fmt.Println("Before:", before)
	// [{2019-01-09 - 2019-01-14}]
	fmt.Println("After:", after)
	// [{2019-01-14 - 2019-01-15} {2019-01-20 - 2019-01-24}]
}
```
