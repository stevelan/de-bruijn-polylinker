# De Bruijn Poly Linker

A polylinker is a fragment of DNA that contains multiple target sites for different restriction enzymes, which can then be injected into a genome to offer target motifs to restriction enzymes which cut DNA into fragments. An optimal poly linker for a set of target motifs is the shortest sequence that contains each motif once.

This project attempts to construct an optimal poly linker from a set of target motifs by constructing a de bruijn graph from the motifs and finding the shortest Euclidean path through the graph that does not introduce each motif into the sequence more than once.
 

## Getting started

Install required tools with \`make wintools`

Build the project with `make build` will produce a binary in the root of the project directory

Clean the project with `make clean` 

## Context from emails:

Plan is loosely this: 
* When designing plasmid vectors researchers will sometimes want to include a Multiple Cloning Site (MCS) that contains motifs that can be cut with their favourite restriction enzymes (REs).
* Short sequences are cheaper to synthesize.
* DeBruijn graphs are good at finding the shortest path/s through a set of sequences
* Gets slightly complicated when you account for REs that can have ambiguous motifs i.e. ATGANNNNATTG when N is any base.
* Write a module that takes a list of motifs and finds the shortest overlapping arrangement.
