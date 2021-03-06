# About

./bin/task -file ./data/187 -> 210-240ms

![pprof.svg](./github/cpu.svg)

- [cpu.pprof](./github/cpu.pprof)
- [mem.pprof](./github/mem.pprof)
- [trace.out](./github/trace.out)

## Description
Your program will be given a list of accepted words (`/data/vocabulary.txt`).
For each word **W** in the input file, you must find word **W'** from the list of accepted words such that
the number of changes from **W** to **W'** is minimized.
It is possible that **W** is already **W'** and thus the number of changes necessary is zero.
A change is defined as replacing a single letter with another letter, adding a letter in any position,
or removing a letter from any position.
The total score that you need to output is the minimum number of changes necessary to make all words acceptable.

## Input specification
Your program must take a single string argument,
representing the file name of the content to analyze.
The input file consists entirely of lower case letters and space characters.
You are guaranteed that the input file will start with a lower case letter,
and that all words are separated by at least one space character.
The file may or may not end with a new line character.

Example input file:
`tihs sententcnes iss nout varrry goud`

You are guaranteed that your program will run against well formed input files and that the accepted word
list is identical to the one provided for testing.

**Example input file:**
Example input file located in `/data/example_input`

```
tihs sententcnes iss nout varrry goud
```

You are guaranteed that your program will run against well formed input
files and that the accepted word list is identical to the one provided
for testing.

## Output Specification
Your program must print out the minimum number of changes necessary
to turn all words in the input file into accepted words as defined by
the vocabulary.
Words may not be joined together, or separated into multiple words.
A change in a word is defined as one of the following:
1. Replacing any single letter with another letter.
2. Adding a single letter in any position.
3. Removing any single letter.

This score must be printed out as an integer and followed by a single new line.

**Example Command:**

`./bin/task -file data/example_input`
OR
`./bin/task -file data/187`

**Example Output (newline after number):**

```
8
```
OR
```
187
```

## Programming Language and Efficiency
Your solution has to be fast in order for us to consider you for this role.
You can test your solution with file 187 to make sure that:
 * your solution is correct (it should output 187/22717)
 * it’s fast enough (it should run under 2 seconds (2 minutes for 22717 file) in Go on your regular machine).
