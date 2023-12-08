package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type mappingRule struct {
	src   int
	dst   int
	count int
}

type mapping struct {
	srcCode string
	dstCode string
	rules   []mappingRule
}

var mappings = map[string]mapping{}

func (m *mapping) mapValue(src int) int {
	for _, rule := range m.rules {
		if src >= rule.src && src < rule.src+rule.count {
			return rule.dst + src - rule.src
		}
	}
	return src
}

func doMappings(srcCode string, dstCode string, srcValue int) int {
	mapping := mappings[srcCode]
	result := srcValue
	for mapping.srcCode != dstCode {
		result = mapping.mapValue(result)
		dstMapping, ok := mappings[mapping.dstCode]
		if !ok {
			break
		}
		mapping = dstMapping
	}
	return result
}

func Mapper[T any, R any](transform func(T) R) func([]T) []R {
	return func(arr []T) []R {
		transformed := []R{}
		for _, item := range arr {
			transformed = append(transformed, transform(item))
		}
		return transformed
	}
}

func main() {
	content, err := os.ReadFile("d5p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	lines := strings.Split(input, "\n")

	toNumbersMapper := Mapper(func(str string) int {
		res, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal("failed to map to numbers")
		}
		return res
	})

	seeds := toNumbersMapper(strings.Split(strings.Split(lines[0], ": ")[1], " "))

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

	fmt.Println(mappings)

	var minLocation int = math.MaxInt

	for _, seed := range seeds {
		location := doMappings("seed", "location", seed)
		if location < minLocation {
			minLocation = location
		}
	}

	fmt.Println(minLocation)
}
