package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"sort"
	"strconv"
)

type superstring struct {
	motifIndices  []int
	value         string
	sizeReduction int
}

func main() {
	motifs := []string{"AC", "CA", "GA", "TA"}
	fmt.Println("Smallest " + PolyLinker(motifs))
}

func PolyLinker(motifs []string) string {

	workingMotifs := motifs
	count := 0
	for len(workingMotifs) >= 2 {

		n := len(workingMotifs)
		k := 3
		if n < k {
			k = n
		}
		fmt.Printf("\n***** Generating permutations with k = %d, n = %d \n", k, n)
		gen := combin.NewPermutationGenerator(n, k)

		smallest := superstring{
			motifIndices:  nil,
			value:         "",
			sizeReduction: -1,
		}

		for gen.Next() {
			picks := gen.Permutation(nil)

			candidate := ""
			for i := 0; i < len(picks); i++ {
				candidate = join(candidate, workingMotifs[picks[i]])
			}

			reduction := calcSizeReduction(candidate, workingMotifs, picks)
			if reduction > smallest.sizeReduction {
				smallest = superstring{motifIndices: picks, value: candidate, sizeReduction: reduction}
			} // else haven't found a smaller join
			count++
			// this is just for development when playing with different values of k,
			if count > 10_000_000 {
				panic(strconv.Itoa(count) + " too large")
			}
			if count%1000 == 0 {
				fmt.Printf("%d loops iterated - Current motifs: %d - %v\n", count, len(workingMotifs), workingMotifs)
			}

		}
		workingMotifs = removeElements(workingMotifs, smallest.motifIndices)
		workingMotifs = append(workingMotifs, smallest.value)
		fmt.Printf("Added %s to %s\n", smallest.value, workingMotifs)
	}

	return workingMotifs[0]
}

func removeElements(slice []string, indices []int) []string {
	fmt.Printf("Removing %v from %v\n", indices, slice)
	sorted := make([]int, len(indices))
	copy(sorted, indices)
	sort.Ints(sorted)
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
			prefix := left[i:]
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

func calcSizeReduction(candidate string, motifs []string, picks []int) int {
	candidateLen := len(candidate)
	originalLen := 0
	for _, pick := range picks {
		originalLen += len(motifs[pick])
	}
	return originalLen - candidateLen
}
