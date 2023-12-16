package day5

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func d5p2(input string) int {
	lines := strings.Split(input, "\n")

	toNumbersMapper := aocutils.Mapper(func(str string) int {
		res, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal("failed to map to numbers")
		}
		return res
	})

	seedRanges := toNumbersMapper(strings.Split(strings.Split(lines[0], ": ")[1], " "))

	lines = lines[1:]

	var activeMapping *mapping
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			if activeMapping != nil {
				mappings[activeMapping.srcCode] = *activeMapping
				activeMapping = nil
			}
			continue
		}
		if activeMapping == nil {
			// reading a new mapping details
			parts := strings.Split(strings.Split(line, " ")[0], "-")
			activeMapping = &mapping{
				srcCode: parts[0],
				dstCode: parts[2],
			}
		} else {
			// parsing new mapping rule here
			numbers := toNumbersMapper(strings.Split(line, " "))
			activeMapping.rules = append(activeMapping.rules, mappingRule{
				src:   numbers[1],
				dst:   numbers[0],
				count: numbers[2],
			})
		}
	}
	if activeMapping != nil {
		mappings[activeMapping.srcCode] = *activeMapping
	}

	rangeMinimums := make(chan int, len(seedRanges)/2)

	for i := 0; i < len(seedRanges); i += 2 {
		rangeStart := seedRanges[i]
		rangeEnd := rangeStart + seedRanges[i+1]
		go func() {
			fmt.Println(rangeStart, rangeEnd)
			var minLocation int = math.MaxInt
			for seed := rangeStart; seed < rangeEnd; seed++ {
				location := doMappings("seed", "location", seed)
				if location < minLocation {
					minLocation = location
				}
			}
			rangeMinimums <- minLocation
		}()
	}

	totalMin := math.MaxInt
	resultsRead := 0
	for rangeMin := range rangeMinimums {
		if rangeMin < totalMin {
			totalMin = rangeMin
		}
		resultsRead += 1
		if resultsRead == len(seedRanges)/2 {
			break
		}
	}

	fmt.Println(totalMin)
	return totalMin
}
