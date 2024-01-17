package daterange_test

import (
	"reflect"
	"testing"
	"time"

	dr "github.com/felixenescu/date-range"
)

// test dr.DateRange.NewDateRange
func TestNewDateRange(t *testing.T) {
	cases := []struct {
		name string
		from time.Time
		to   time.Time
		want dr.DateRange
	}{
		{
			name: "zero zero",
			from: time.Time{},
			to:   time.Time{},
			want: dr.DateRange{},
		},
		{
			name: "zero non zero",
			from: time.Time{},
			to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: dr.NewDateRange(time.Time{}, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "non zero zero",
			from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Time{},
			want: dr.NewDateRange(time.Time{}, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "non zero non zero equal",
			from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "non zero non zero from before to",
			from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			want: dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "non zero non zero from after to",
			from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := dr.NewDateRange(c.from, c.to)
			if got != c.want {
				t.Errorf("NewDateRange(%v, %v) = %v, want %v", c.from, c.to, got, c.want)
			}
		})
	}
}

// test dr.DateRange.MustNewDateRange
func TestMustNewDateRange(t *testing.T) {
	cases := []struct {
		name string
		from time.Time
		to   time.Time
		want dr.DateRange
	}{
		{
			name: "zero zero",
			from: time.Time{},
			to:   time.Time{},
			want: dr.DateRange{},
		},
		{
			name: "zero non zero",
			from: time.Time{},
			to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: dr.NewDateRange(time.Time{}, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "non zero non zero equal",
			from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "non zero non zero from before to",
			from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			want: dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := dr.MustNewDateRange(c.from, c.to)
			if got != c.want {
				t.Errorf("MustNewDateRange(%v, %v) = %v, want %v", c.from, c.to, got, c.want)
			}
		})
	}

	// test panic case
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustNewDateRange() should have panicked")
		} else {
			if r == "from date is after to date" {
				t.Logf("MustNewDateRange() panicked with correct message: %v", r)
			} else {
				t.Errorf("MustNewDateRange() panicked with wrong message: %v", r)
			}
		}
	}()

	dr.MustNewDateRange(
		time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	)
}

// test dr.DateRange.String
func TestDateRangeString(t *testing.T) {
	cases := []struct {
		name string
		d    dr.DateRange
		want string
	}{
		{
			name: "zero",
			d:    dr.DateRange{},
			want: "{0001-01-01 - 0001-01-01}",
		},
		{
			name: "non zero",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
			want: "{2019-01-01 - 2019-01-02}",
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			if got := c.d.String(); got != c.want {
				t.Errorf("DateRange.String() = %v, want %v", got, c.want)
			}
		})
	}
}

// test dr.DateRange.IsZero
func TestDateRangeIsZero(t *testing.T) {
	cases := []struct {
		name string
		d    dr.DateRange
		want bool
	}{
		{
			name: "zero",
			d:    dr.DateRange{},
			want: true,
		},
		{
			name: "zero-initialized",
			d:    dr.NewDateRange(time.Time{}, time.Time{}),
			want: true,
		},
		{
			name: "non-zero",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
			want: false,
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			if got := c.d.IsZero(); got != c.want {
				t.Errorf("DateRange.IsZero() = %v, want %v", got, c.want)
			}
		})
	}
}

// test dr.DateRange.Contains
func TestDateRangeContains(t *testing.T) {
	cases := []struct {
		name string
		d    dr.DateRange
		t    time.Time
		want bool
	}{
		{
			name: "zero range, zero time",
			d:    dr.DateRange{},
			t:    time.Time{},
			want: false,
		},
		{
			name: "zero range, non zero time",
			d:    dr.DateRange{},
			t:    time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "non zero range, non zero time, before from",
			d:    dr.NewDateRange(time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			t:    time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "non zero range, non zero time, equal from",
			d:    dr.NewDateRange(time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			t:    time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "non zero range, non zero time, inside",
			d:    dr.NewDateRange(time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			t:    time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "non zero range, non zero time, equal to",
			d:    dr.NewDateRange(time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			t:    time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "non zero range, non zero time, after to",
			d:    dr.NewDateRange(time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			t:    time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "all equal",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
			t:    time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: true,
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			if got := c.d.Contains(c.t); got != c.want {
				t.Errorf("DateRange.Contains() = %v, want %v", got, c.want)
			}
		})
	}
}

// test dr.DateRange.Overlaps
func TestDateRangeOverlaps(t *testing.T) {
	cases := []struct {
		name string
		d    dr.DateRange
		o    dr.DateRange
		want bool
	}{
		{
			name: "zero",
			d:    dr.DateRange{},
			o:    dr.DateRange{},
			want: false,
		},
		{
			name: "zero range, non zero other",
			d:    dr.DateRange{},
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
			want: false,
		},
		{
			name: "non zero range, zero other",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
			o:    dr.DateRange{},
			want: false,
		},
		{
			name: "non zero range, non zero other, before",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: false,
		},
		{
			name: "non zero range, non zero other, after",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
			want: false,
		},
		{
			name: "non zero range, non zero other, same from",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: true,
		},
		{
			name: "non zero range, non zero other, same to",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: true,
		},
		{
			name: "non zero range, non zero other, same from and to",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: true,
		},
		{
			name: "non zero range, non zero other, inside",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: true,
		},
		{
			name: "non zero range, non zero other, outside before",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: false,
		},
		{
			name: "non zero range, non zero other, outside after",
			d:    dr.NewDateRange(time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: false,
		},
		{
			name: "non zero range, non zero other, overlap from",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: true,
		},
		{
			name: "non zero range, non zero other, overlap to",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			want: true,
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			if got := c.d.Overlaps(c.o); got != c.want {
				t.Errorf("%v.DateRange.Overlaps(%v) = %v, want %v", c.d, c.o, got, c.want)
			}
		})
	}
}

// test dr.DateRange.Overlaps
func TestDateRangeIncludes(t *testing.T) {
	cases := []struct {
		name string
		d    dr.DateRange
		o    dr.DateRange
		want bool
	}{
		{
			name: "zero zero",
			d:    dr.DateRange{},
			o:    dr.DateRange{},
			want: false,
		},
		{
			name: "non-zero zero",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
			o:    dr.DateRange{},
			want: false,
		},
		{
			name: "zero range, non zero other",
			d:    dr.DateRange{},
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
			want: false,
		},
		{
			name: "non zero range, non zero other, before",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: false,
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			if got := c.d.Includes(c.o); got != c.want {
				t.Errorf("%v.DateRange.Includes(%v) = %v, want %v", c.d, c.o, got, c.want)
			}
		})
	}
}

// test dr.DateRange.Intersection
func TestDateRangeIntersection(t *testing.T) {
	cases := []struct {
		name string
		d    dr.DateRange
		o    dr.DateRange
		want dr.DateRange
	}{
		{
			name: "zero range, zero other",
			d:    dr.DateRange{},
			o:    dr.DateRange{},
			want: dr.DateRange{},
		},
		{
			name: "zero range, non zero other",
			d:    dr.DateRange{},
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRange{},
		},
		{
			name: "non zero range, zero other",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.DateRange{},
			want: dr.DateRange{},
		},
		{
			name: "non zero range, non zero other, inside",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "non zero range, non zero other, outside before",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRange{},
		},
		{
			name: "non zero range, non zero other, outside after",
			d:    dr.NewDateRange(time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRange{},
		},
		{
			name: "non zero range, non zero other, overlap from",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "non zero range, non zero other, overlap to",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := c.d.Intersection(c.o)
			if got != c.want {
				t.Errorf("%v.DateRange.Intersection(%v) = %v, want %v", c.d, c.o, got, c.want)
			}
		})
	}
}

// test DataRange.Union
func TestDateRangeUnion(t *testing.T) {
	cases := []struct {
		name string
		d    dr.DateRange
		o    dr.DateRange
		want dr.DateRanges
	}{
		{
			name: "zero range, zero other",
			d:    dr.DateRange{},
			o:    dr.DateRange{},
			want: dr.DateRanges{},
		},
		{
			name: "zero range, non zero other",
			d:    dr.DateRange{},
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, zero other",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.DateRange{},
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, outside before",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, outside before, adjacent",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, outside after",
			d:    dr.NewDateRange(time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC))},
		}, {
			name: "non zero range, non zero other, outside after, adjacent",
			d:    dr.NewDateRange(time.Date(2019, 1, 15, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, inside",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, overlap from",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, equal from",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, overlap to",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, equal to",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, equal range",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, overlap both",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, overlap both, single day",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, overlap both, single day both",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC))},
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := c.d.Union(c.o)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("%v.DateRange.Union(%v) = %#v, want %#v", c.d, c.o, got, c.want)
			}
		})
	}
}

// test DataRange.Difference
func TestDateRangeDifference(t *testing.T) {
	cases := []struct {
		name string
		d    dr.DateRange
		o    dr.DateRange
		want dr.DateRanges
	}{
		{
			name: "zero range, zero other",
			d:    dr.DateRange{},
			o:    dr.DateRange{},
			want: dr.DateRanges{},
		},
		{
			name: "zero range, non zero other",
			d:    dr.DateRange{},
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{},
		},
		{
			name: "non zero range, zero other",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.DateRange{},
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, outside before",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, outside after",
			d:    dr.NewDateRange(time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, inside",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{},
		},
		{
			name: "non zero range, non zero other, overlap from",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC))},
		},

		{
			name: "non zero range, non zero other, equal from",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, overlap to",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, equal to",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, equal range",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{},
		},
		{
			name: "non zero range, non zero other, overlap both",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC))},
		},
		{
			name: "non zero range, non zero other, overlap both, single day",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.DateRanges{
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC))},
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := c.d.Difference(c.o)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("%v.DateRange.Difference(%v) = %#v, want %#v", c.d, c.o, got, c.want)
			}
		})
	}
}
