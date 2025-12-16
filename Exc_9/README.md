# SBD Exercise 9
In this exercise we are going to implement [MapReduce](https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf) in Golang.
First the contents of Marcus Aurelius's Meditations (provided by [Project Gutenberg](https://www.gutenberg.org/ebooks/2680)) 
will be read into memory, then a MapReduce implementation will count the frequency of all words.

- [ ] Read contents of Meditations `res/meditations.txt` into a slice containing strings
- [ ] Implement the MapReduce algorithm and all functions required by the `MapReduceInterface`
  - The algorithm should run concurrently --> use go routines!
  - Filter out all special chars and numericals (RegEx)
- [ ] Your implementation works if it pases all supplied tests in `map_reduce_test.go`
  - Tests can be run with `go test ./mapred` 
- [ ] Once you've finished the implementation, compute the word frequency of the input text

## Resources
How the inner workings of a word frequency counter using MapReduce
could look like, is explained in great detail by [this article](https://sherbold.github.io/intro-to-data-science/12_Big-Data-and-Map-Reduce.html#Word-Count-with-MapReduce)
by Steffen Herbold of Uni Passau.

- https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf
- https://sherbold.github.io/intro-to-data-science/slides_pdf/12-BigDataMapReduce.pdf
- https://datascienceguide.github.io/map-reduce
- https://pdos.csail.mit.edu/6.824/labs/lab-mr.html