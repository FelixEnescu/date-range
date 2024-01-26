[![Build Status](https://github.com/felixenescu/date-range/actions/workflows/ci.yml/badge.svg)](https://github.com/felixenescu/date-range/actions/workflows/ci.yml/badge.svg)
[![Code coverage](https://codecov.io/gh/FelixEnescu/date-range/graph/badge.svg?token=Q4A0CDVYHB)](https://codecov.io/gh/FelixEnescu/date-range)
[![Go Report Card](https://goreportcard.com/badge/github.com/felixenescu/date-range)](https://goreportcard.com/report/github.com/felixenescu/date-range)
[![GoDoc](https://godoc.org/github.com/felixenescu/date-range?status.svg)](http://godoc.org/github.com/felixenescu/date-range)


# date-range

## Introduction

**daterange** is a powerful and intuitive Go package designed for handling date intervals efficiently and effectively. This library simplifies operations such as comparing dates, checking overlaps, and processing date ranges in Go applications.

## Features
Easy comparison of date ranges
Checking for overlaps between date intervals
Processing and manipulation of date intervals

## Installation

Use `go get` to install this package.

```shell
go get github.com/felixenescu/date-range
```

Import it with:

```go
import dr "github.com/felixenescu/date-range"
```

and use `dr` as the package name inside the code.

## Usage


### DateRange

#### Overview

**DateRange** represents an inclusive range of dates. It is defined by two time.Time values, from and to.


#### Constructors

 - **NewDateRange(from, to time.Time**):** Creates a new `DateRange` instance. The input dates are automatically ordered.
 - **MustNewDateRange(from, to time.Time**):** Similar to `NewDateRange` but panics if the from date is after the to date.

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


 - **Checking if a Date is Within a Range**

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

 - **Finding Overlap Between Two Date Ranges**
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

 - **Calculating the Union of Two Date Ranges**

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

 - **Calculating the Difference of Two Date Ranges**

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


## Potential Use Cases

 - **Event Planning and Scheduling:** The Date-Range library can be used in applications that deal with event scheduling, helping to check for conflicts in dates and times.

 - **Reservation Systems:** For systems that handle reservations, such as hotel booking or venue reservation systems, this library can assist in managing availability and booking dates.

 - **Data Analysis and Reporting:** In data analysis, particularly when dealing with time series data, the Date-Range library can help in segmenting data into specific time intervals for better insights.

 - **Financial Applications:** In financial applications, the library can be used to calculate durations for interest calculations, investment maturity periods, and more.

 - **Educational Tools:** For educational software that schedules courses or exams, the library can help in managing and visualizing course timelines.