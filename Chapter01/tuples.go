// main package has examples shown
// in Go Data Structures and algorithms book
// this one uses goroutines and channels to compare concurrency and sequential processing
package main

// importing fmt package
import (
	"fmt"
	"sync"
	"time"
)

// A tuple is a finite sorted list of elements. It is a data structure that
// groups data. Tuples are  typically immutable sequential collections.
// The element has related fields of different  datatypes.
type tuple struct {
	squr int
	cube int
}

// gets the powerseries of integer a and returns tuple of square of a
// and cube of a
func powerSeriesConcurrent(a int, wg *sync.WaitGroup, c chan tuple) {
	data := new(tuple)
	data.cube, data.squr = a*a, a*a*a
	c <- *data
	time.Sleep(20 * time.Microsecond)

	wg.Done()
}

func powerSeries(a int) (int, int) {
	time.Sleep(20 * time.Microsecond)
	return a * a, a * a * a
}

func main() {
	wg := new(sync.WaitGroup)

	// depending on the number of loops we need to increase the buffer
	ch := make(chan tuple, 1024*1024)
	loops := 20 // how many elements in the series to calculate

	start_time := time.Now()
	fmt.Println("Calculating concurrently ...")
	for i := 3; i < loops; i++ {
		go powerSeriesConcurrent(i, wg, ch)
		wg.Add(1)
	}
	wg.Wait()
	close(ch)

	fmt.Println(time.Since(start_time), " passed")

	start_time = time.Now()
	fmt.Println("Calculating sequentially ...")
	for i := 3; i < loops; i++ {
		powerSeries(i)
	}
	fmt.Println(time.Since(start_time), " passed")

}
