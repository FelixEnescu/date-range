package daterange

import (
	"reflect"
	"testing"
	"time"
)

// test DateRange.NewDateRange
func TestNewDateRange(t *testing.T) {
	cases := []struct {
		name string
		from time.Time
		to   time.Time
		want DateRange
	}{
		{
			name: "zero zero",
			from: time.Time{},
			to:   time.Time{},
			want: DateRange{},
		},
		{
			name: "zero non zero",
			from: time.Time{},
			to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: DateRange{
				from: time.Time{},
				to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "non zero zero",
			from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Time{},
			want: DateRange{
				from: time.Time{},
				to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "non zero non zero equal",
			from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "non zero non zero from before to",
			from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			want: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "non zero non zero from after to",
			from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := NewDateRange(c.from, c.to)
			if got != c.want {
				t.Errorf("NewDateRange(%v, %v) = %v, want %v", c.from, c.to, got, c.want)
			}
		})
	}
}

// test DateRange.MustNewDateRange
func TestMustNewDateRange(t *testing.T) {
	cases := []struct {
		name string
		from time.Time
		to   time.Time
		want DateRange
	}{
		{
			name: "zero zero",
			from: time.Time{},
			to:   time.Time{},
			want: DateRange{},
		},
		{
			name: "zero non zero",
			from: time.Time{},
			to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: DateRange{
				from: time.Time{},
				to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "non zero non zero equal",
			from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "non zero non zero from before to",
			from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			want: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := MustNewDateRange(c.from, c.to)
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

	MustNewDateRange(
		time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	)
}

// test DateRange.String
func TestDateRangeString(t *testing.T) {
	cases := []struct {
		name string
		d    DateRange
		want string
	}{
		{
			name: "zero",
			d:    DateRange{},
			want: "{0001-01-01 - 0001-01-01}",
		},
		{
			name: "non zero",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)},
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

// test DateRange.IsZero
func TestDateRangeIsZero(t *testing.T) {
	cases := []struct {
		name string
		d    DateRange
		want bool
	}{
		{
			name: "zero",
			d:    DateRange{},
			want: true,
		},
		{
			name: "zero-initialized",
			d: DateRange{
				from: time.Time{},
				to:   time.Time{},
			},
			want: true,
		},
		{
			name: "non-zero",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			},
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

// test DateRange.Contains
func TestDateRangeContains(t *testing.T) {
	cases := []struct {
		name string
		d    DateRange
		t    time.Time
		want bool
	}{
		{
			name: "zero range, zero time",
			d:    DateRange{},
			t:    time.Time{},
			want: false,
		},
		{
			name: "zero range, non zero time",
			d:    DateRange{},
			t:    time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "non zero range, non zero time, before from",
			d: DateRange{
				from: time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			t:    time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "non zero range, non zero time, equal from",
			d: DateRange{
				from: time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			t:    time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "non zero range, non zero time, inside",
			d: DateRange{
				from: time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			t:    time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "non zero range, non zero time, equal to",
			d: DateRange{
				from: time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			t:    time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "non zero range, non zero time, after to",
			d: DateRange{
				from: time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			t:    time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "all equal",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			},
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

// test DateRange.Overlaps
func TestDateRangeOverlaps(t *testing.T) {
	cases := []struct {
		name string
		d    DateRange
		o    DateRange
		want bool
	}{
		{
			name: "zero",
			d:    DateRange{},
			o:    DateRange{},
			want: false,
		},
		{
			name: "zero range, non zero other",
			d:    DateRange{},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)},
			want: false,
		},
		{
			name: "non zero range, zero other",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)},
			o:    DateRange{},
			want: false,
		},
		{
			name: "non zero range, non zero other, before",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: false,
		},
		{
			name: "non zero range, non zero other, after",
			d: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)},
			want: false,
		},
		{
			name: "non zero range, non zero other, same from",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
		{
			name: "non zero range, non zero other, same to",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
		{
			name: "non zero range, non zero other, same from and to",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
		{
			name: "non zero range, non zero other, inside",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
		{
			name: "non zero range, non zero other, outside before",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},

			o: DateRange{
				from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)},
			want: false,
		},
		{
			name: "non zero range, non zero other, outside after",
			d: DateRange{
				from: time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)},

			o: DateRange{
				from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)},
			want: false,
		},
		{
			name: "non zero range, non zero other, overlap from",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: true,
		},
		{
			name: "non zero range, non zero other, overlap to",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
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

// test DateRange.Overlaps
func TestDateRangeIncludes(t *testing.T) {
	cases := []struct {
		name string
		d    DateRange
		o    DateRange
		want bool
	}{
		{
			name: "zero zero",
			d:    DateRange{},
			o:    DateRange{},
			want: false,
		},
		{
			name: "non-zero zero",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)},
			o:    DateRange{},
			want: false,
		},
		{
			name: "zero range, non zero other",
			d:    DateRange{},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)},
			want: false,
		},
		{
			name: "non zero range, non zero other, before",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
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

// test DateRange.Intersection
func TestDateRangeIntersection(t *testing.T) {
	cases := []struct {
		name string
		d    DateRange
		o    DateRange
		want DateRange
	}{
		{
			name: "zero range, zero other",
			d:    DateRange{},
			o:    DateRange{},
			want: DateRange{},
		},
		{
			name: "zero range, non zero other",
			d:    DateRange{},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: DateRange{},
		},
		{
			name: "non zero range, zero other",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o:    DateRange{},
			want: DateRange{},
		},
		{
			name: "non zero range, non zero other, inside",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
		},
		{
			name: "non zero range, non zero other, outside before",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)},
			want: DateRange{},
		},
		{
			name: "non zero range, non zero other, outside after",
			d: DateRange{
				from: time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)},
			want: DateRange{},
		},
		{
			name: "non zero range, non zero other, overlap from",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
		},
		{
			name: "non zero range, non zero other, overlap to",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
			want: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
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
		d    DateRange
		o    DateRange
		want DateRanges
	}{
		{
			name: "zero range, zero other",
			d:    DateRange{},
			o:    DateRange{},
			want: DateRanges{},
		},
		{
			name: "zero range, non zero other",
			d:    DateRange{},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "non zero range, zero other",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			o: DateRange{},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "non zero range, non zero other, outside before",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			o: DateRange{
				from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC),
			},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
				},
				{
					from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "non zero range, non zero other, outside before, adjacent",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			o: DateRange{
				from: time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC),
			},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "non zero range, non zero other, outside after",
			d: DateRange{
				from: time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC),
			},
			o: DateRange{
				from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC),
			},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC),
				},
				{
					from: time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC),
				},
			},
		}, {
			name: "non zero range, non zero other, outside after, adjacent",
			d: DateRange{
				from: time.Date(2019, 1, 15, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC),
			},
			o: DateRange{
				from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC),
			},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "non zero range, non zero other, inside",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "non zero range, non zero other, overlap from",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, equal from",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, overlap to",
			d: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, equal to",
			d: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, equal range",
			d: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
			},
			o: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
			},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "non zero range, non zero other, overlap both",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC),
			},
			o: DateRange{
				from: time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
			},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "non zero range, non zero other, overlap both, single day",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC),
			},
			o: DateRange{
				from: time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "non zero range, non zero other, overlap both, single day both",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			o: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
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
		d    DateRange
		o    DateRange
		want DateRanges
	}{
		{
			name: "zero range, zero other",
			d:    DateRange{},
			o:    DateRange{},
			want: DateRanges{},
		},
		{
			name: "zero range, non zero other",
			d:    DateRange{},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{},
		},
		{
			name: "non zero range, zero other",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o: DateRange{},
			want: DateRanges{{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, outside before",
			d: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, outside after",
			d: DateRange{
				from: time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{{
				from: time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, inside",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{},
		},
		{
			name: "non zero range, non zero other, overlap from",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{{
				from: time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, equal from",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{{
				from: time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, overlap to",
			d: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, equal to",
			d: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, equal range",
			d: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{},
		},
		{
			name: "non zero range, non zero other, overlap both",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
				{
					from: time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "non zero range, non zero other, overlap both, single day",
			d: DateRange{
				from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)},
			o: DateRange{
				from: time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)},
			want: DateRanges{
				{
					from: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)},
				{
					from: time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC),
					to:   time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)},
			},
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
