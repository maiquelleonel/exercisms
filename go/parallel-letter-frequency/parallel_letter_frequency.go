package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int
type prdFunc func(string) FreqMap

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup
	c := make(chan FreqMap)
	wg.Add(len(texts))

	for _, text := range texts {
		go concCall(text, Frequency, &wg, c)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	var F = make(FreqMap)
	for freq := range c {
		for r, total := range freq {
			F[r] += total
		}
	}
	return F

}

func concCall(t string, pred prdFunc, wg *sync.WaitGroup, c chan<- FreqMap) {
	defer wg.Done()
	c <- pred(t)
}
