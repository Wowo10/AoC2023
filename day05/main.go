package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	seeds, seed_to_soil, soil_to_fertilizer, fertilizer_to_water, water_to_light,
		light_to_temperature, temperature_to_humidity, humidity_to_location := parseFile()

	min := 9999999999999
	min2 := 9999999999999

	for _, seed := range seeds {
		soil := seed_to_soil.GetDestination(seed)
		fertilizer := soil_to_fertilizer.GetDestination(soil)
		water := fertilizer_to_water.GetDestination(fertilizer)
		light := water_to_light.GetDestination(water)
		temperature := light_to_temperature.GetDestination(light)
		humidity := temperature_to_humidity.GetDestination(temperature)
		location := humidity_to_location.GetDestination(humidity)

		if location < min {
			min = location
		}
	}

	count := 0

	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedRange := seeds[i+1]

		for j := seedStart; j < seedStart+seedRange; j++ {
			count++
			soil := seed_to_soil.GetDestination(j)
			fertilizer := soil_to_fertilizer.GetDestination(soil)
			water := fertilizer_to_water.GetDestination(fertilizer)
			light := water_to_light.GetDestination(water)
			temperature := light_to_temperature.GetDestination(light)
			humidity := temperature_to_humidity.GetDestination(temperature)
			location := humidity_to_location.GetDestination(humidity)

			if location < min2 {
				min2 = location
			}
		}
	}

	fmt.Println(min)
	fmt.Println(min2)
	fmt.Println(count)
}

func parseFile() (seeds []int, seed_to_soil, soil_to_fertilizer,
	fertilizer_to_water, water_to_light, light_to_temperature,
	temperature_to_humidity, humidity_to_location RangeCollection) {

	inputContent, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	splitContent := strings.Split(string(inputContent), "\n\n")

	seedContent := splitContent[0]

	seedsStr := strings.Trim(strings.Split(seedContent, ":")[1], " ")

	for _, s := range strings.Split(seedsStr, " ") {
		i, _ := strconv.Atoi(s)

		seeds = append(seeds, i)
	}

	ranges := []RangeCollection{}
	for _, rangeContent := range splitContent[1:] {
		parsed := strings.Trim(strings.Split(rangeContent, ":")[1], "\n")

		rangeContents := strings.Split(parsed, "\n")

		rc := RangeCollection{}

		for _, rangeContent := range rangeContents {
			data := strings.Split(rangeContent, " ")

			destination, _ := strconv.Atoi(data[0])
			source, _ := strconv.Atoi(data[1])
			length, _ := strconv.Atoi(data[2])

			rc.MyRanges = append(rc.MyRanges, CreateMyRange(destination, source, length))
		}

		ranges = append(ranges, rc)
	}

	return seeds, ranges[0], ranges[1], ranges[2], ranges[3], ranges[4], ranges[5], ranges[6]
}
