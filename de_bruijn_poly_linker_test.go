package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestNoOverlap(t *testing.T) {
	motifs := []string{"AAAA", "GGGG", "CCCC"}
	doOverlapTest(t, motifs, "AAAAGGGGCCCC")
}

func TestOverlap(t *testing.T) {
	motifs := []string{"AAAA", "AAGG", "GGCC", "CCTT", "TTCC"}
	doOverlapTest(t, motifs, "AAAAGGCCTTCC")
}

func TestLargeSample(t *testing.T) {
	motifs := []string{"AAAA", "AAGG", "GGCC", "CCTT", "TTCC", "GGGG", "CCCC", "TTTT", "TTAA", "AATT", "TTGG", "GGTT", "AAAT", "ATAT", "TATA", "CGCG", "CATG", "GTGC", "GGAT", "CCAT"}
	doOverlapTest(t, motifs, "AAAATTTTGGGGTTCCCCTTAAGGCCATGTGCGCGGATATA")

}

func TestAllContained(t *testing.T) {
	contains := "AAAGGGATATCA"
	motifs := []string{contains, "TATC", "AGGGA", "GGGAT", "ATAT", "AAGG"}
	doOverlapTest(t, motifs, contains)
}

func doOverlapTest(t *testing.T, motifs []string, output string) {
	t.Logf("%s : %s - %s\n", t.Name(), motifs, output)

	actual := PolyLinker(append([]string{}, motifs...))
	if output != actual {
		t.Error(fmt.Printf("Expected %s - %d but got %s - %d", output, len(output), actual, len(actual)))
	}
	originalLen := 0
	for _, motif := range motifs {
		originalLen += len(motif)
		if !strings.Contains(output, motif) {
			t.Error(fmt.Printf("Expected %s to contain %s", output, motif))
		}
	}
	reduction := originalLen - len(output)
	fmt.Printf("Achieved reduction of %d - %.2f\n", originalLen-len(output), float64(reduction)*100/float64(originalLen))
}
