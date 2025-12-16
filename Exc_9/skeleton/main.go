package main

import (
	"exc9/mapred"

	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// Main function
func main() {
	//STEP 1: Open the input file
	//open meditations.txt from the res/ folder
	file, err := os.Open("res/meditations.txt")
	if err != nil {
		//if the file cannot be opened, log an error and exit.
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	//STEP 2: Read the file line by line
	//each line of text will become an element in a slice of strings
	//this slice will be the input for MapReduce
	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	//STEP 3: Run the MapReduce word count
	//create a MapReduce instance and call Run on the input text
	var mr mapred.MapReduce
	results := mr.Run(text)
	//'results' is a map[string]int where the key is the word and the value is the frequency

	//STEP 4: Sort results for readability
	//sort words by frequency in descending order
	//helps to see the most common words first
	type wordCount struct {
		Word  string
		Count int
	}

	var sorted []wordCount
	for word, count := range results {
		sorted = append(sorted, wordCount{word, count})
	}

	//Sort slice by Count descending
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Count > sorted[j].Count
	})

	//STEP 5: Print the results
	//print the top 20 most frequent words
	fmt.Println("Top 50 most frequent words in Meditations:")
	for i := 0; i < 50 && i < len(sorted); i++ {
		fmt.Printf("%2d. %-15s %d\n", i+1, sorted[i].Word, sorted[i].Count)
	}
}
