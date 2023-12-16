package day04

import (
	"aoc/utils/aocasync"
	"aoc/utils/aocinput"
	"aoc/utils/aocmath"
	"aoc/utils/aocparse"
	"fmt"
	"log"
	"math"
	"runtime"
	"strings"
	"time"
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
	// log.Printf("%d => %d => %d => %d => %d => %d => %d => %d",
	// 	seed, soil, fertilizer, water, light, temperature, humidity, location)
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

// It may be surprising at first, but this implementation is actually about 10x SLOWER
// than the synchronous implementation above. I believe this is due to the overhead of
// channel communication and goroutine spawing.
//
// I'm leaving it here because it was still a GREAT way for me to learn more about using
// channels, but it's also a good reminder that multithreading isn't a panacea and
// premature optimization is the root of all evil.
func Part2Async() string {
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

	seeds := genSeeds(almanac, ranges)

	concurrency := runtime.NumCPU()
	log.Printf("Processing seeds with %d mappers", concurrency)
	outs := make([]<-chan int, 0, concurrency)
	for i := 0; i < concurrency; i++ {
		out := seedMapper(i, seeds, almanac)
		outs = append(outs, out)
	}
	results := aocasync.Merge(outs...)

	totalSeeds := 0
	for _, r := range ranges {
		totalSeeds += r.length
	}
	progress := aocasync.ProgressTracker(results, totalSeeds)
	return fmt.Sprint(aocmath.MinIntChan(progress))
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

	totalSeeds := 0
	for _, r := range ranges {
		totalSeeds += r.length
	}
	hundredth := int(math.Max(math.Floor(float64(totalSeeds)/100), 1))
	log.Printf("1%% is about %d items", hundredth)
	done := 0

	min := math.MaxInt
	start := time.Now()
	for _, r := range ranges {
		for i := r.start; i < r.start+r.length; i++ {
			seed := i
			result := almanac.traverse(seed)
			if result < min {
				min = result
			}
			done += 1
			if done%hundredth == 0 {
				percent := (float32(done) / float32(totalSeeds)) * 100
				duration := time.Since(start)
				log.Printf("Completed %d of %d (%f%%) in %v", done, totalSeeds, percent, duration)
				start = time.Now()
			}
		}
	}

	return fmt.Sprint(min)
}

func genSeeds(a almanac, ranges []seedRange) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, r := range ranges {
			for i := r.start; i < r.start+r.length; i++ {
				//log.Printf("INPUT: %d", i)
				out <- i
			}
		}
	}()

	return out
}

func seedMapper(id int, in <-chan int, a almanac) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for seed := range in {
			result := a.traverse(seed)
			//log.Printf("MAP(%d): %d => %d", id, seed, result)
			out <- result
		}
	}()
	return out
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
