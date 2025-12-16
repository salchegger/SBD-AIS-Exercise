package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce
//  Map -> Shuffle -> Reduce

// ////////////////////////////////
// RUN function
// ////////////////////////////////
func (mr MapReduce) Run(input []string) map[string]int {
	//Channel to collect mapper outputs
	mapChan := make(chan []KeyValue)

	var mapWg sync.WaitGroup

	//MAP PHASE
	for _, text := range input {
		mapWg.Add(1)

		//Each input line is processed concurrently
		go func(t string) {
			defer mapWg.Done()
			mapChan <- mr.wordCountMapper(t)
		}(text)
	}

	//Close channel once all mappers are done
	go func() {
		mapWg.Wait()
		close(mapChan)
	}()

	//SHUFFLE/GROUP PHASE
	grouped := make(map[string][]int)

	for kvSlice := range mapChan {
		for _, kv := range kvSlice {
			grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
		}
	}

	//REDUCE PHASE
	result := make(map[string]int)
	var reduceWg sync.WaitGroup
	var mu sync.Mutex

	for key, values := range grouped {
		reduceWg.Add(1)

		go func(k string, v []int) {
			defer reduceWg.Done()
			kv := mr.wordCountReducer(k, v)

			//Protect shared map with mutex
			mu.Lock()
			result[kv.Key] = kv.Value
			mu.Unlock()
		}(key, values)
	}

	reduceWg.Wait()
	return result
}

// ////////////////////////////////
// MAPPER function
// ////////////////////////////////
func (mr MapReduce) wordCountMapper(text string) []KeyValue {
	//Converting everything to lowercase to make the counting case-insensitive
	text = strings.ToLower(text)

	//Regex: keeping only letters and spaces
	re := regexp.MustCompile(`[^a-z\s]`) //remove everything that is NOT a letter or whitespace
	cleanText := re.ReplaceAllString(text, "")

	//Split text into words
	words := strings.Fields(cleanText)

	//Create key-value pairs
	var result []KeyValue
	for _, word := range words {
		result = append(result, KeyValue{Key: word, Value: 1})
	}
	return result
}

// ////////////////////////////////
// REDUCER function
// ////////////////////////////////
func (mr MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return KeyValue{Key: key, Value: sum}
}
