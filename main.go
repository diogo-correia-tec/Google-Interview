//Challenge availabe here: https://www.youtube.com/watch?v=rw4s4M3hFfs&t=332s

package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

const MaxUint = ^uint(0)

type param map[string]bool

type distance map[string]int

type Blocks struct {
	establishments []param
}

type Result struct {
	Index       int      `json:"index"`
	Distances   distance `json:"distances"`
	MaxDistance int      `json:"maxDistance"`
}

func main() {

	var blocks Blocks
	blocks.establishments = []param{
		map[string]bool{"gym": false, "school": true, "store": false},
		map[string]bool{"gym": true, "school": false, "store": false},
		map[string]bool{"gym": true, "school": true, "store": false},
		map[string]bool{"gym": false, "school": true, "store": false},
		map[string]bool{"gym": false, "school": true, "store": true},
	}

	request := []string{"gym", "school", "store"}

	result := make([]Result, len(blocks.establishments))

	for index, establishment := range blocks.establishments {
		result[index].Index = index
		result[index].Distances = make(map[string]int, len(request))
		maxDistance := 0
		for _, req := range request {
			if establishment[req] {
				result[index].Distances[req] = 0
				continue
			}

			nextDistance := getNextDistance(blocks.establishments, index, req)
			previousDistance := getPreviousDistance(blocks.establishments, index, req)

			if nextDistance < previousDistance {
				result[index].Distances[req] = nextDistance
			} else {
				result[index].Distances[req] = previousDistance
			}

			if result[index].Distances[req] > maxDistance {
				maxDistance = result[index].Distances[req]
			}
		}
		result[index].MaxDistance = maxDistance
	}

	finalResult, _ := json.Marshal(SortResult(result))

	fmt.Printf("%+v\n", string(finalResult))
}

func SortResult(result []Result) []Result {
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].MaxDistance < result[j].MaxDistance
	})

	return result
}

func getNextDistance(establishments []param, index int, req string) int {
	if index+1 > len(establishments) {
		return int(MaxUint >> 1)
	}

	for i := index + 1; i < len(establishments); i++ {
		if establishments[i][req] {
			return i - index
		}
	}

	return int(MaxUint >> 1)
}
func getPreviousDistance(establishments []param, index int, req string) int {
	if index-1 < 0 {
		return int(MaxUint >> 1)
	}

	for i := index - 1; i > 0; i-- {
		if establishments[i][req] {
			return index - i
		}
	}

	return int(MaxUint >> 1)
}
