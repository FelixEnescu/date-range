package daterange

import "sort"

// DataRanges is a collection of DateRange elements
// TODO(felix) explain the what si merged and sorted
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

// ToSlice returns the members of the collection as a slice
func (drs *DateRanges) ToSlice() []DateRange {
	copySlice := make([]DateRange, len(drs.dr))
	copy(copySlice, drs.dr)
	return copySlice
}

// String returns a string representation of the collection
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

// Append adds the given elements to the collection
func (drs *DateRanges) Append(dataRange ...DateRange) {
	drs.dr = append(drs.dr, dataRange...)
	drs.normalize()
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
