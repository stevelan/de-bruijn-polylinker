package main

import (
	"fmt"
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
	doOverlapTest(t, motifs, "CCTTAAGGCCATGGATATAGGGGTTCCCCGCGAAAATTTTGGTGC")
}

func doOverlapTest(t *testing.T, motifs []string, output string) {
	t.Logf("%s : %s - %s\n", t.Name(), motifs, output)

	actual := PolyLinker(motifs)
	if output != actual {
		t.Error(fmt.Printf("Expected %s but got %s", output, actual))
	}
}
