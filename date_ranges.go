package daterange

import (
	"sort"
	"time"
)

// DataRanges is a collection of DateRange elements
type DateRanges struct {
	dr []DateRange
}

// NewDateRanges returns a new collection with given elements
func NewDateRanges(dataRanges ...DateRange) DateRanges {
	drs := DateRanges{
		dr: make([]DateRange, len(dataRanges)),
	}
	copy(drs.dr, dataRanges)
	drs.normalize()
	return drs
}

// ToSlice returns the members of the collection as a slice.
// Items are guaranteed to be sorted, non-overlapping and non-zero.
// Any adjacent periods are merged.
func (drs *DateRanges) ToSlice() []DateRange {
	copySlice := make([]DateRange, len(drs.dr))
	copy(copySlice, drs.dr)
	return copySlice
}

// String returns a string representation of the collection.
// Items are guaranteed to be sorted, non-overlapping and non-zero.
// Any adjacent periods are merged.
func (drs DateRanges) String() string {
	if drs.IsZero() {
		return "[]"
	}
	str := "["
	for i, dr := range drs.dr {
		if i > 0 {
			str += " "
		}
		str += dr.String()
	}
	str += "]"
	return str
}

// IsZero returns true if the collection is empty
func (drs *DateRanges) IsZero() bool {
	return len(drs.dr) == 0
}

// Len returns the number of elements in the collection
func (drs *DateRanges) Len() int {
	return len(drs.dr)
}

// FirstDate returns the first date of the collection
func (drs *DateRanges) FirstDate() time.Time {
	if drs.IsZero() {
		return time.Time{}
	}
	return drs.dr[0].from
}

// LastDate returns the last date of the collection
func (drs *DateRanges) LastDate() time.Time {
	if drs.IsZero() {
		return time.Time{}
	}
	return drs.dr[len(drs.dr)-1].to
}

// Equal returns true if the collection is equal to the given collection
func (drs *DateRanges) Equal(other DateRanges) bool {
	if len(drs.dr) != len(other.dr) {
		return false
	}
	for i, dr := range drs.dr {
		if dr != other.dr[i] {
			return false
		}
	}
	return true
}

// Append adds the given elements to the collection
func (drs *DateRanges) Append(dataRange ...DateRange) {
	drs.dr = append(drs.dr, dataRange...)
	drs.normalize()
}

// Contains returns true if the given date is in the collection
func (drs *DateRanges) Contains(date time.Time) bool {
	for _, dr := range drs.dr {
		if dr.Contains(date) {
			return true
		}
	}
	return false
}

// IsAnyDateIn returns true if any date in the given DateRange is in the collection
// Zero DateRange is always considered to be in the collection
func (drs *DateRanges) IsAnyDateIn(other DateRange) bool {
	if other.IsZero() {
		return true
	}
	for _, dr := range drs.dr {
		if dr.Overlaps(other) {
			return true
		}
	}
	return false
}

// IsAllDatesIn returns true if all dates in the given DateRange are in the collection
// Zero DateRange is always considered to be in the collection
func (drs *DateRanges) IsAllDatesIn(other DateRange) bool {
	if other.IsZero() {
		return true
	}
	for _, dr := range drs.dr {
		if dr.Includes(other) {
			return true
		}
	}
	return false
}

// SplitInclusive splits the collection into two collections based on the given date
// The given date is included in both collections
func (drs *DateRanges) SplitInclusive(date time.Time) (DateRanges, DateRanges) {
	before := NewDateRanges()
	after := NewDateRanges()
	for _, dr := range drs.dr {
		if dr.to.Before(date) {
			before.Append(dr)
		} else if dr.from.After(date) {
			after.Append(dr)
		} else {
			before.Append(DateRange{from: dr.from, to: date})
			after.Append(DateRange{from: date, to: dr.to})
		}
	}
	return before, after
}

// normalize sorts the collection and merges overlapping periods
func (drs *DateRanges) normalize() *DateRanges {
	if len(drs.dr) == 0 {
		return drs
	}
	return drs.sort().removeZero().merge()
}

// sort sorts the collection
func (drs *DateRanges) sort() *DateRanges {
	sort.Slice(drs.dr, func(i, j int) bool {
		return drs.dr[i].from.Before(drs.dr[j].from)
	})
	return drs
}

// removeZero removes zero periods from a sorted collection
func (drs *DateRanges) removeZero() *DateRanges {
	if len(drs.dr) == 0 {
		return drs
	}
	firstNonZero := 0
	for _, dr := range drs.dr {
		if !dr.IsZero() {
			break
		}
		firstNonZero++
	}
	drs.dr = drs.dr[firstNonZero:]
	return drs
}

// merge merges overlapping periods from a sorted collection
func (drs *DateRanges) merge() *DateRanges {
	if len(drs.dr) == 0 {
		return drs
	}
	// merge overlapping periods
	merged := []DateRange{}
	var current = drs.dr[0]
	for _, period := range drs.dr[1:] {
		// we add 2 because we want to merge periods that are 1 day apart
		// e.g. 2019-01-01 - 2019-01-03 and 2019-01-04 - 2019-01-05
		if period.from.Before(current.to.AddDate(0, 0, 2)) {
			if period.to.After(current.to) {
				current.to = period.to
			}
		} else {
			merged = append(merged, current)
			current = period
		}
	}
	merged = append(merged, current)
	drs.dr = merged
	return drs
}
