package mapred

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inText = []string{
	"This is an   ²¹²æ 023 example.. text for mapreduce exc9",
	"exc9 does not need !! to be `97 #*  put in a Docker container",
}

func Test_Run(t *testing.T) {
	var mr MapReduce
	results := mr.Run(inText)
	expected := map[string]int{
		"not":       1,
		"a":         1,
		"container": 1,
		"docker":    1,
		"exc":       2,
		"to":        1,
		"put":       1,
		"in":        1,
		"example":   1,
		"text":      1,
		"does":      1,
		"need":      1,
		"be":        1,
		"this":      1,
		"is":        1,
		"an":        1,
		"for":       1,
		"mapreduce": 1,
	}
	assert.Equal(t, expected, results)
}

func Test_Run_Fail(t *testing.T) {
	var mr MapReduce
	results := mr.Run(inText)
	expected := map[string]int{
		"not":       1,
		"a":         1,
		"container": 1,
		"docker":    1,
		"exc":       2,
		"is":        1,
		"an":        1,
		"for":       1,
		"mapreduce": 1,
	}
	assert.NotEqual(t, expected, results)
}

func Test_wordCountMapper(t *testing.T) {
	var mr MapReduce
	results := mr.wordCountMapper("This this small BIG sentence sEnTenCE")
	expected := []KeyValue{
		{
			Key:   "this",
			Value: 1,
		},
		{
			Key:   "this",
			Value: 1,
		},
		{
			Key:   "small",
			Value: 1,
		},
		{
			Key:   "big",
			Value: 1,
		},
		{
			Key:   "sentence",
			Value: 1,
		},
		{
			Key:   "sentence",
			Value: 1,
		},
	}
	assert.ElementsMatch(t, expected, results)
}
func Test_wordCountMapper_Fail(t *testing.T) {
	var mr MapReduce
	results := mr.wordCountMapper("This this small BIG sentence sEnTenCE")
	expected := []KeyValue{
		{
			Key:   "this",
			Value: 1,
		},
		{
			Key:   "this",
			Value: 1,
		},
		{
			Key:   "sentence",
			Value: 1,
		},
		{
			Key:   "sentence",
			Value: 1,
		},
	}
	assert.NotElementsMatch(t, expected, results)
}

func Test_wordCountReducer(t *testing.T) {
	var mr MapReduce
	results := mr.wordCountReducer("test", []int{1, 1})
	expected := KeyValue{
		Key:   "test",
		Value: 2,
	}
	assert.Equal(t, expected.Key, results.Key)
	assert.Equal(t, expected.Value, results.Value)
}

func Test_wordCountReducer_Fail(t *testing.T) {
	var mr MapReduce
	results := mr.wordCountReducer("test", []int{1, 1})
	expected := KeyValue{
		Key:   "test",
		Value: 1,
	}
	assert.Equal(t, expected.Key, results.Key)
	assert.NotEqual(t, expected.Value, results.Value)
}
