package day04

import (
	"aoc/utils/aocinput"
	"aoc/utils/aocmath"
	"aoc/utils/aocparse"
	"fmt"
	"log"
	"strings"
)

func Part1() string {
	content := aocinput.ReadInputAsString(5)
	content = strings.TrimSuffix(content, "\n")

	groups := strings.Split(content, "\n\n")

	seeds := parseSeeds(groups[0])

	almanac := almanac{
		groupToMap(groups[1]),
		groupToMap(groups[2]),
		groupToMap(groups[3]),
		groupToMap(groups[4]),
		groupToMap(groups[5]),
		groupToMap(groups[6]),
		groupToMap(groups[7]),
	}

	locations := make([]int, 0, len(seeds))
	for _, seed := range seeds {
		location := almanac.traverse(seed)
		locations = append(locations, location)
	}

	return fmt.Sprint(aocmath.MinInt(locations))
}

type almanac struct {
	seedToSoil            farmMap
	soilToFertilizer      farmMap
	fertilizerToWater     farmMap
	waterToLight          farmMap
	lightToTemperature    farmMap
	temperatureToHumidity farmMap
	humidityToLocation    farmMap
}

func (a almanac) traverse(seed int) (location int) {
	soil := a.seedToSoil.get(seed)
	fertilizer := a.soilToFertilizer.get(soil)
	water := a.fertilizerToWater.get(fertilizer)
	light := a.waterToLight.get(water)
	temperature := a.lightToTemperature.get(light)
	humidity := a.temperatureToHumidity.get(temperature)
	location = a.humidityToLocation.get(humidity)
	log.Printf("%d => %d => %d => %d => %d => %d => %d => %d",
		seed, soil, fertilizer, water, light, temperature, humidity, location)
	return
}

func parseSeeds(seedLine string) []int {
	parts := strings.Split(seedLine, ": ")
	spec := parts[1]
	strings := strings.Split(spec, " ")

	ints := make([]int, 0, len(strings))

	for _, str := range strings {
		ints = append(ints, aocparse.MustAtoi(str))
	}
	return ints
}

func groupToMap(group string) farmMap {
	lines := strings.Split(group, "\n")
	mappingStrings := lines[1:]

	mappings := make(farmMap, 0, len(mappingStrings))
	for _, mappingString := range mappingStrings {
		mappings = append(mappings, lineToMapping(mappingString))
	}
	return mappings
}

func lineToMapping(line string) farmMapping {
	parts := strings.Split(line, " ")
	destinationRangeStart := aocparse.MustAtoi(parts[0])
	sourceRangeStart := aocparse.MustAtoi(parts[1])
	length := aocparse.MustAtoi(parts[2])
	return farmMapping{destinationRangeStart, sourceRangeStart, length}
}

type farmMap []farmMapping

func (m farmMap) get(key int) int {
	for _, mapping := range m {
		if val, ok := mapping.getWithin(key); ok {
			return val
		}
	}
	return key
}

type farmMapping struct {
	destinationRangeStart int
	sourceRangeStart      int
	length                int
}

func (m farmMapping) contains(key int) bool {
	return key >= m.sourceRangeStart && key < m.sourceRangeStart+m.length
}

func (m farmMapping) getWithin(key int) (int, bool) {
	if m.contains(key) {
		dist := key - m.sourceRangeStart
		return m.destinationRangeStart + dist, true
	}
	return 0, false
}

func Part2() string {
	content := aocinput.ReadInputAsString(5)
	content = strings.TrimSuffix(content, "\n")

	groups := strings.Split(content, "\n\n")

	ranges := parseSeedRanges(groups[0])

	almanac := almanac{
		groupToMap(groups[1]),
		groupToMap(groups[2]),
		groupToMap(groups[3]),
		groupToMap(groups[4]),
		groupToMap(groups[5]),
		groupToMap(groups[6]),
		groupToMap(groups[7]),
	}

	locations := make([]int, 0)
	for _, r := range ranges {
		for i := r.start; i < r.start+r.length; i++ {
			location := almanac.traverse(i)
			locations = append(locations, location)
		}
	}

	return fmt.Sprint(aocmath.MinInt(locations))
}

func parseSeedRanges(line string) []seedRange {
	parts := strings.Split(line, ": ")
	rangeParts := strings.Split(parts[1], " ")

	ranges := make([]seedRange, 0)
	for i := 0; i < len(rangeParts); i += 2 {
		start := aocparse.MustAtoi(rangeParts[i])
		length := aocparse.MustAtoi(rangeParts[i+1])
		ranges = append(ranges, seedRange{start, length})
	}
	return ranges
}

type seedRange struct {
	start  int
	length int
}
