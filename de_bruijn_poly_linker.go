package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"sort"
	"strconv"
)

type superstring struct {
	motifIndices []int
	value        string
}

func main() {
	motifs := []string{"AC", "CA", "GA", "TA"}
	fmt.Println("Smallest " + PolyLinker(motifs))
}

func PolyLinker(motifs []string) string {

	workingMotifs := motifs
	count := 0
	for len(workingMotifs) >= 3 {

		n := len(workingMotifs)
		k := 3
		if n < 3 {
			k = n
		}
		fmt.Printf("Generating permutations with k = %d, n = %d ", k, n)
		gen := combin.NewPermutationGenerator(n, k)

		smallest := superstring{
			motifIndices: nil,
			value:        "",
		}

		for gen.Next() {
			picks := gen.Permutation(nil)

			candidate := ""
			for i := 0; i < len(picks); i++ {
				candidate = join(candidate, motifs[picks[i]])
			}
			fmt.Println(candidate)
			if smallest.value == "" || len(candidate) < len(smallest.value) {
				smallest = superstring{motifIndices: picks, value: candidate}
			} // else haven't found a smaller join
			count++
			if count > 1_000_000 {
				panic(strconv.Itoa(count) + " too large")
			}
			fmt.Println(count)

		}
		workingMotifs = removeElements(workingMotifs, smallest.motifIndices)
		workingMotifs = append(workingMotifs, smallest.value)
		fmt.Println(workingMotifs)
	}
	fmt.Println(len(workingMotifs))
	if len(workingMotifs) == 2 { //TODO handle joining in the right order if has prefix of the other
		return join(workingMotifs[0], workingMotifs[1])
	} else {
		return workingMotifs[0]
	}
}

func removeElements(slice []string, indices []int) []string {
	fmt.Printf("Removing %v from %v", indices, slice)
	sorted := make([]int, len(indices))
	copy(sorted, indices)
	sort.Ints(sorted)
	fmt.Println(sorted)
	for i := len(sorted) - 1; i >= 0; i-- {
		index := sorted[i]
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}

func join(left string, right string) string {
	suffix := findSuffix(left, right)
	if len(suffix) > 0 {
		if len(suffix) < len(right) {
			return left[0:len(left)-len(suffix)] + right
		} else { // right is fully contained within left
			return left
		}
	} else {
		return left + right
	}
}

func findSuffix(left string, right string) string {
	for i := 0; i < len(left); i++ {
		if right[0] == left[i] {
			prefix := left[i:len(left)]
			if isPrefix(prefix, right) {
				return prefix
			}
		}
	}
	return ""
}

func isPrefix(prefix string, target string) bool {
	for i := 0; i < len(prefix) && i < len(target); i++ {
		// TODO handle extended genetic alphabet
		if prefix[i] != target[i] {
			return false
		}
	}
	return true
}
