package main_test

import (
	sut "day05"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	r := sut.CreateMyRange(50, 98, 2)

	assert.Equal(t, r.IsInRange(98), true)
	assert.Equal(t, r.IsInRange(99), true)

	assert.Equal(t, r.IsInRange(100), false)
	assert.Equal(t, r.IsInRange(50), false)

	assert.Equal(t, r.GetDestination(98), 50)
	assert.Equal(t, r.GetDestination(99), 51)
}

func TestSecond(t *testing.T) {
	r := sut.CreateMyRange(52, 50, 48)

	assert.Equal(t, r.IsInRange(50), true)
	assert.Equal(t, r.IsInRange(79), true)
	assert.Equal(t, r.IsInRange(97), true)

	assert.Equal(t, r.IsInRange(100), false)
	assert.Equal(t, r.IsInRange(49), false)

	assert.Equal(t, r.GetDestination(50), 52)
	assert.Equal(t, r.GetDestination(79), 81)
	assert.Equal(t, r.GetDestination(96), 98)
}

func TestThird(t *testing.T) {
	c := sut.RangeCollection{
		MyRanges: []sut.MyRange{
			sut.CreateMyRange(52, 50, 48),
			sut.CreateMyRange(50, 98, 2),
		},
	}

	assert.Equal(t, c.GetDestination(79), 81)
	assert.Equal(t, c.GetDestination(14), 14)
	assert.Equal(t, c.GetDestination(55), 57)
	assert.Equal(t, c.GetDestination(13), 13)

	assert.Equal(t, c.GetDestination(99), 51)
}

func TestFourth(t *testing.T) {
	seed_to_soil := sut.RangeCollection{
		MyRanges: []sut.MyRange{
			sut.CreateMyRange(52, 50, 48),
			sut.CreateMyRange(50, 98, 2),
		},
	}

	soil_to_fertilizer := sut.RangeCollection{
		MyRanges: []sut.MyRange{
			sut.CreateMyRange(0, 15, 37),
			sut.CreateMyRange(37, 52, 2),
			sut.CreateMyRange(39, 0, 15),
		},
	}

	fertilizer_to_water := sut.RangeCollection{
		MyRanges: []sut.MyRange{
			sut.CreateMyRange(49, 53, 8),
			sut.CreateMyRange(0, 11, 42),
			sut.CreateMyRange(42, 0, 7),
			sut.CreateMyRange(57, 7, 4),
		},
	}

	water_to_light := sut.RangeCollection{
		MyRanges: []sut.MyRange{
			sut.CreateMyRange(88, 18, 7),
			sut.CreateMyRange(18, 25, 70),
		},
	}

	light_to_temperature := sut.RangeCollection{
		MyRanges: []sut.MyRange{
			sut.CreateMyRange(45, 77, 23),
			sut.CreateMyRange(81, 45, 19),
			sut.CreateMyRange(68, 64, 13),
		},
	}

	temperature_to_humidity := sut.RangeCollection{
		MyRanges: []sut.MyRange{
			sut.CreateMyRange(0, 69, 1),
			sut.CreateMyRange(1, 0, 69),
		},
	}

	humidity_to_location := sut.RangeCollection{
		MyRanges: []sut.MyRange{
			sut.CreateMyRange(60, 56, 37),
			sut.CreateMyRange(56, 93, 4),
		},
	}

	soil := seed_to_soil.GetDestination(79)
	assert.Equal(t, soil, 81)

	fertilizer := soil_to_fertilizer.GetDestination(soil)
	assert.Equal(t, fertilizer, 81)

	water := fertilizer_to_water.GetDestination(fertilizer)
	assert.Equal(t, water, 81)

	light := water_to_light.GetDestination(water)
	assert.Equal(t, light, 74)

	temperature := light_to_temperature.GetDestination(light)
	assert.Equal(t, temperature, 78)

	humidity := temperature_to_humidity.GetDestination(temperature)
	assert.Equal(t, humidity, 78)

	location := humidity_to_location.GetDestination(humidity)
	assert.Equal(t, location, 82)
}
