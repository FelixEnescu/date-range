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
		name     string
		from     time.Time
		to       time.Time
		wantFrom time.Time
		wantTo   time.Time
	}{
		{
			name:     "zero zero",
			from:     time.Time{},
			to:       time.Time{},
			wantFrom: time.Time{},
			wantTo:   time.Time{},
		},
		{
			name:     "zero non zero utc",
			from:     time.Time{},
			to:       time.Date(2019, 1, 1, 2, 3, 4, 5, time.UTC),
			wantFrom: time.Time{},
			wantTo:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "zero non zero est",
			from:     time.Time{},
			to:       time.Date(2019, 1, 1, 21, 5, 6, 7, time.FixedZone("EST", -5*60*60)),
			wantFrom: time.Time{},
			wantTo:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non zero utc zero",
			from:     time.Date(2019, 1, 1, 2, 3, 4, 5, time.UTC),
			to:       time.Time{},
			wantFrom: time.Time{},
			wantTo:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non zero est zero",
			from:     time.Date(2019, 1, 1, 21, 5, 6, 7, time.FixedZone("EST", -5*60*60)),
			to:       time.Time{},
			wantFrom: time.Time{},
			wantTo:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non zero utc non zero utc equal date no time",
			from:     time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			wantFrom: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			wantTo:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non zero utc non zero utc equal date time",
			from:     time.Date(2019, 1, 1, 2, 3, 4, 5, time.UTC),
			to:       time.Date(2019, 1, 1, 2, 3, 4, 5, time.UTC),
			wantFrom: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			wantTo:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non zero utc non zero utc equal date time different location",
			from:     time.Date(2019, 1, 1, 2, 3, 4, 5, time.FixedZone("EST", -5*60*60)),
			to:       time.Date(2019, 1, 1, 2, 3, 4, 5, time.FixedZone("EEST", +2*60*60)),
			wantFrom: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			wantTo:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non zero utc non zero utc equal date time different location",
			from:     time.Date(2019, 1, 1, 2, 3, 4, 5, time.FixedZone("EST", -5*60*60)),
			to:       time.Date(2019, 1, 1, 5, 4, 3, 2, time.FixedZone("EEST", +2*60*60)),
			wantFrom: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			wantTo:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non zero non zero from before to",
			from:     time.Date(2019, 1, 1, 0, 7, 0, 0, time.UTC),
			to:       time.Date(2019, 1, 2, 0, 8, 0, 0, time.FixedZone("EST", -5*60*60)),
			wantFrom: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			wantTo:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non zero non zero from after to",
			from:     time.Date(2019, 1, 2, 5, 0, 0, 0, time.UTC),
			to:       time.Date(2019, 1, 1, 6, 0, 0, 0, time.UTC),
			wantFrom: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			wantTo:   time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			got := dr.NewDateRange(c.from, c.to)
			if got.From() != c.wantFrom || got.To() != c.wantTo {
				t.Errorf("NewDateRange(%v, %v) = %v, want {%v - %v}", c.from, c.to, got, c.wantFrom, c.wantTo)
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
		{
			name: "non zero non zero from same day time before to",
			from: time.Date(2019, 1, 1, 6, 7, 8, 9, time.UTC),
			to:   time.Date(2019, 1, 1, 2, 3, 8, 9, time.UTC),
			want: dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
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
			if r == "from date (2019-01-02 00:00:00 +0000 UTC) is after to date (2019-01-01 00:00:00 +0000 UTC)" {
				t.Logf("MustNewDateRange() panicked with correct message: %v", r)
			} else {
				t.Errorf("MustNewDateRange() panicked with wrong message: %v", r)
			}
		}
	}()

	// test panic case
	dr.MustNewDateRange(
		time.Date(2019, 1, 2, 3, 4, 5, 6, time.UTC),
		time.Date(2019, 1, 1, 6, 5, 4, 3, time.UTC),
	)
}

// test dr.DateRange.From
func TestDateRangeFrom(t *testing.T) {
	cases := []struct {
		name string
		d    dr.DateRange
		want time.Time
	}{
		{
			name: "zero",
			d:    dr.DateRange{},
			want: time.Time{},
		},
		{
			name: "non zero",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
			want: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			if got := c.d.From(); !reflect.DeepEqual(got, c.want) {
				t.Errorf("DateRange.From() = %v, want %v", got, c.want)
			}
		})
	}
}

// test dr.DateRange.To
func TestDateRangeTo(t *testing.T) {
	cases := []struct {
		name string
		d    dr.DateRange
		want time.Time
	}{
		{
			name: "zero",
			d:    dr.DateRange{},
			want: time.Time{},
		},
		{
			name: "non zero",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC)),
			want: time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, c := range cases {
		t.Logf("Running test %s", c.name)
		t.Run(c.name, func(t *testing.T) {
			if got := c.d.To(); !reflect.DeepEqual(got, c.want) {
				t.Errorf("DateRange.To() = %v, want %v", got, c.want)
			}
		})
	}
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
			want: dr.NewDateRanges(),
		},
		{
			name: "zero range, non zero other",
			d:    dr.DateRange{},
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, zero other",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.DateRange{},
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, outside before",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, outside before, adjacent",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, outside after",
			d:    dr.NewDateRange(time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC))),
		}, {
			name: "non zero range, non zero other, outside after, adjacent",
			d:    dr.NewDateRange(time.Date(2019, 1, 15, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, inside",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, overlap from",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, equal from",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, overlap to",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, equal to",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, equal range",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, overlap both",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, overlap both, single day",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, overlap both, single day both",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC))),
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
			want: dr.NewDateRanges(),
		},
		{
			name: "zero range, non zero other",
			d:    dr.DateRange{},
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(),
		},
		{
			name: "non zero range, zero other",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.DateRange{},
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, outside before",
			d:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, outside after",
			d:    dr.NewDateRange(time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 14, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 16, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, inside",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(),
		},
		{
			name: "non zero range, non zero other, overlap from",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC))),
		},

		{
			name: "non zero range, non zero other, equal from",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, overlap to",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, equal to",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, equal range",
			d:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(),
		},
		{
			name: "non zero range, non zero other, overlap both",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 8, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 9, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC))),
		},
		{
			name: "non zero range, non zero other, overlap both, single day",
			d:    dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC)),
			o:    dr.NewDateRange(time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 5, 0, 0, 0, 0, time.UTC)),
			want: dr.NewDateRanges(
				dr.NewDateRange(time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC)),
				dr.NewDateRange(time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC), time.Date(2019, 1, 18, 0, 0, 0, 0, time.UTC))),
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
