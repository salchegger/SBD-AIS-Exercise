package mapred

type MapReduceInterface interface {
	Run(input []string) map[string]int
	wordCountMapper(text string) []KeyValue
	wordCountReducer(key string, values []int) KeyValue
}

// KeyValue represents a key-value pair
type KeyValue struct {
	Key   string
	Value int
}
