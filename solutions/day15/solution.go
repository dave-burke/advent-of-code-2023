package day15

import (
	"aoc/utils/aocinput"
	"aoc/utils/aocparse"
	"container/list"
	"fmt"
	"log"
	"strings"
)

func Part1() string {
	// input := aocinput.ReadSampleAsString(15)
	input := aocinput.ReadInputAsString(15)

	input = strings.TrimSpace(input) // remove newline at the end

	steps := strings.Split(input, ",")

	sum := 0
	for _, step := range steps {
		result := Hash(step)
		log.Printf("%s => %d", step, result)
		sum += result
	}

	return fmt.Sprint(sum)
}

func Hash(s string) int {
	currentValue := 0

	for _, char := range s {
		code := int(char)
		currentValue += code
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

type Lens struct {
	label       string
	focalLength int
}

func Part2() string {
	input := aocinput.ReadSampleAsString(15)
	// input := aocinput.ReadInputAsString(15)

	input = strings.TrimSpace(input) // remove newline at the end

	steps := strings.Split(input, ",")

	boxes := make(map[int]list.List)
	for _, step := range steps {
		if step[len(step)-1] == '-' {
			label := step[:len(step)-1]
			hash := Hash(label)
			if existingList, ok := boxes[hash]; ok {
				log.Printf("Remove %s from box %d with %d items", label, hash, existingList.Len())
				for elem := existingList.Front(); elem != nil; elem = elem.Next() {
					if elem.Value.(Lens).label == label {
						existingList.Remove(elem)
						break
					}
				}
			} else {
				log.Printf("%s is not in box %d", label, hash)
			}
		} else {
			parts := strings.Split(step, "=")
			lens := Lens{parts[0], aocparse.MustAtoi(parts[1])}
			hash := Hash(lens.label)
			log.Printf("Add %v to box %d", lens, hash)
			if existingList, ok := boxes[hash]; ok {
				existingList.PushBack(lens)
			} else {
				newList := *list.New()
				newList.PushBack(lens)
				boxes[hash] = newList
			}
		}
	}

	return fmt.Sprint(len(input))
}
