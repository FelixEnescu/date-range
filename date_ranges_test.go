package daterange_test

import (
	"reflect"
	"testing"
	"time"

	dr "github.com/felixenescu/date-range"
)

// test dr.DateRanges.NewDateRanges
func TestNewDateRanges(t *testing.T) {
	cases := []struct {
		name string
		drs  []dr.DateRange
		want []dr.DateRange
	}{
		{
			name: "empty",
			drs:  []dr.DateRange{},
			want: []dr.DateRange{},
		},
		{
			name: "zero",
			drs:  []dr.DateRange{dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{})},
			want: []dr.DateRange{},
		},
		{
			name: "one day",
			drs:  []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "two days",
			drs:  []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "multiple days",
			drs:  []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC))},
		},
		// we get sorted ranges
		{
			name: "two multiple days unsorted",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			},
			want: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "three multiple days unsorted",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 22, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			},
			want: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 22, 0, 0, 0, 0, time.UTC)),
			},
		},
		// zero ranges are removed
		{
			name: "zero zero",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{}),
				dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{})},
			want: []dr.DateRange{},
		},
		{
			name: "zero zero zero",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{}),
				dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{}),
				dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{})},
			want: []dr.DateRange{},
		},
		{
			name: "zero multiple days zero",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{}),
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{})},
			want: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "zero multiple days zero multiple days",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{}),
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 13, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{}),
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 13, 0, 0, 0, 0, time.UTC))},
		},
		// ranges are merged
		{
			name: "two overlapping multiple days",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "three overlapping multiple days",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "two overlapping by one day multiple days",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "two adjacent multiple days",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "two adjacent multiple days single day",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "two inclusive multiple days",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "three inclusive multiple days",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 2, 9, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 2, 6, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 2, 3, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 2, 9, 0, 0, 0, 0, time.UTC))},
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := dr.NewDateRanges(c.drs...)
			if !reflect.DeepEqual(got.ToSlice(), c.want) {
				t.Errorf("NewDateRanges(%v) = %v, want %v", c.drs, got, c.want)
			}
		})
	}
}

// test dr.DateRanges.String
func TestDateRangesString(t *testing.T) {
	cases := []struct {
		name string
		drs  []dr.DateRange
		want string
	}{
		{
			name: "empty",
			drs:  []dr.DateRange{},
			want: "[]",
		},
		{
			name: "zero",
			drs:  []dr.DateRange{dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{})},
			want: "[]",
		},
		{
			name: "one day",
			drs:  []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))},
			want: "[{2019-01-01 - 2019-01-01}]",
		},
		{
			name: "two days",
			drs:  []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC))},
			want: "[{2019-01-01 - 2019-01-02}]",
		},
		{
			name: "multiple days",
			drs:  []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC))},
			want: "[{2019-01-01 - 2019-01-03}]",
		},
		{
			name: "two multiple days unsorted",
			drs: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			},
			want: "[{2019-01-02 - 2019-01-04} {2019-01-09 - 2019-01-12}]",
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := dr.NewDateRanges(c.drs...)
			if got.String() != c.want {
				t.Errorf("NewDateRanges(%v).String() = %v, want %v", c.drs, got, c.want)
			}
		})
	}
}

// test dr.DateRanges.IsZero
func TestDateRangesIsZero(t *testing.T) {
	cases := []struct {
		name string
		drs  []dr.DateRange
		want bool
	}{
		{
			name: "empty",
			drs:  []dr.DateRange{},
			want: true,
		},
		{
			name: "zero",
			drs:  []dr.DateRange{dr.NewDateRange(time.Time{}, time.Time{}), dr.NewDateRange(time.Time{}, time.Time{})},
			want: true,
		},
		{
			name: "one day",
			drs:  []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))},
			want: false,
		},
		{
			name: "two days",
			drs:  []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC))},
			want: false,
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := dr.NewDateRanges(c.drs...)
			if got.IsZero() != c.want {
				t.Errorf("NewDateRanges(%v).IsZero() = %v, want %v", c.drs, got, c.want)
			}
		})
	}
}

// test dr.DateRanges.Append
func TestDateRangesAppend(t *testing.T) {
	cases := []struct {
		name          string
		drs           []dr.DateRange
		newDataRanges []dr.DateRange
		want          []dr.DateRange
	}{
		// append to a empty collection
		{
			name:          "empty zero",
			drs:           []dr.DateRange{},
			newDataRanges: []dr.DateRange{dr.NewDateRange(time.Time{}, time.Time{})},
			want:          []dr.DateRange{},
		},
		{
			name:          "empty one day",
			drs:           []dr.DateRange{},
			newDataRanges: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))},
			want:          []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "empty multiple days multiple days overlapping",
			drs:  []dr.DateRange{},
			newDataRanges: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 12, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			},
			want: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			},
		},
		// append to a non empty collection
		{
			name:          "one day zero",
			drs:           []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))},
			newDataRanges: []dr.DateRange{dr.NewDateRange(time.Time{}, time.Time{})},
			want:          []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC))},
		},
		{
			name:          "one day one day",
			drs:           []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC))},
			newDataRanges: []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC))},
			want: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "one day multiple days",
			drs:  []dr.DateRange{dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC))},
			newDataRanges: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 25, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 29, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC)),
			},
			want: []dr.DateRange{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 25, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 29, 0, 0, 0, 0, time.UTC)),
			},
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			drs := dr.NewDateRanges(c.drs...)
			drs.Append(c.newDataRanges...)
			if !reflect.DeepEqual(drs.ToSlice(), c.want) {
				t.Errorf("NewDateRanges(%v).Append(%v) = %v, want %v", c.drs, c.newDataRanges, drs, c.want)
			}
		})
	}
}