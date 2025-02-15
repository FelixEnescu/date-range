/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2024 Felix Enescu

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package daterange

import (
	"fmt"
	"time"
)

// DateRange is an **inclusive** range of dates. The range is defined by two dates.
type DateRange struct {
	from time.Time // inclusive dates
	to   time.Time
}

// NewDateRange returns a new DateRange from the given dates. This automatically
// order input dates and truncates the time portion of the dates, ignoring the
// time zone (for example 2024-01-26 9pm EST will still be the 26th of January 2024).
// Use MustNewDateRange if you want to panic if `from` date if after the `to`date .
// Note: Only the date portion of the time.Time values is compared. The time portion is ignored.
func NewDateRange(from, to time.Time) DateRange {
	from = toDateUTC(from)
	to = toDateUTC(to)
	return DateRange{
		from: minTime(from, to),
		to:   maxTime(from, to),
	}
}

// MustNewDateRange returns a new DateRange from the given dates. This automatically
// truncates the time portion of the dates, ignoring the time zone (for example
// 2024-01-26 9pm EST will still be the 26th of January 2024). This panics if the
// truncated `from` date is after the truncated `to` date.
// Use NewDateRange if you want to automatically order input dates.
// Note: Only the date portion of the time.Time values is compared. The time portion is ignored.
func MustNewDateRange(from, to time.Time) DateRange {
	from = toDateUTC(from)
	to = toDateUTC(to)
	if from.After(to) {
		panic(fmt.Sprintf("from date (%s) is after to date (%s)", from, to))
	}
	return NewDateRange(from, to)
}

// From returns the start date of the range, as midnight of that day, UTC time.
func (d DateRange) From() time.Time {
	return d.from
}

// To returns the end date of the range, as midnight of that day, UTC time.
func (d DateRange) To() time.Time {
	return d.to
}

// String returns a string representation of the DateRange
func (d DateRange) String() string {
	return "{" + d.from.Format("2006-01-02") + " - " + d.to.Format("2006-01-02") + "}"
}

// IsZero returns true if the both dates of range are zero
func (d DateRange) IsZero() bool {
	return d.from.IsZero() && d.to.IsZero()
}

// Contains returns true if the given date is in the range. The range is inclusive.
func (d DateRange) Contains(date time.Time) bool {
	if d.IsZero() {
		return false
	}
	return !date.Before(d.from) && !date.After(d.to)
}

// Overlaps returns true if the given range overlaps with the range. The range is inclusive.
func (d DateRange) Overlaps(other DateRange) bool {
	if d.IsZero() || other.IsZero() {
		return false
	}
	return d.Contains(other.from) || d.Contains(other.to) || other.Contains(d.from) || other.Contains(d.to)
}

// Includes returns true if the given range is included in the range. The range is inclusive.
func (d DateRange) Includes(other DateRange) bool {
	if d.IsZero() || other.IsZero() {
		return false
	}
	return d.Contains(other.from) && d.Contains(other.to)
}

// Intersection returns the intersection of the two DateRanges
func (d DateRange) Intersection(other DateRange) DateRange {
	if d.Overlaps(other) {
		return DateRange{
			from: maxTime(d.from, other.from),
			to:   minTime(d.to, other.to),
		}
	}
	return DateRange{}
}

// Union returns a DateRanges collection that is the union of the two DateRanges
func (d DateRange) Union(other DateRange) DateRanges {
	switch {
	case d.IsZero() && other.IsZero():
		return NewDateRanges()
	case d.IsZero():
		return NewDateRanges(other)
	case other.IsZero():
		return NewDateRanges(d)
	}

	// non zero, check for overlapping
	if d.Overlaps(other) {
		return NewDateRanges(
			DateRange{
				from: minTime(d.from, other.from),
				to:   maxTime(d.to, other.to),
			},
		)
	}

	// non zero, no overlapping, check for adjacent
	if d.to.AddDate(0, 0, 1).Equal(other.from) {
		return NewDateRanges(
			DateRange{
				from: d.from,
				to:   other.to,
			},
		)
	}
	if other.to.AddDate(0, 0, 1).Equal(d.from) {
		return NewDateRanges(
			DateRange{
				from: other.from,
				to:   d.to,
			},
		)
	}

	// non zero, no overlapping, disjoint ranges, return in ascending order
	if d.from.Before(other.from) {
		return NewDateRanges(
			d,
			other,
		)
	}

	return NewDateRanges(
		other,
		d,
	)
}

// Difference returns a DateRanges collection that is the difference of the two DateRanges
func (d DateRange) Difference(other DateRange) DateRanges {
	if d.IsZero() {
		return NewDateRanges()
	}

	if !d.Overlaps(other) {
		return NewDateRanges(d)
	}

	ranges := NewDateRanges()
	if other.from.After(d.from) {
		ranges.Append(DateRange{
			from: d.from,
			to:   other.from.AddDate(0, 0, -1),
		})
	}
	if other.to.Before(d.to) {
		ranges.Append(DateRange{
			from: other.to.AddDate(0, 0, 1),
			to:   d.to,
		})
	}
	return ranges
}
