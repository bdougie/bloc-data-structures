# Data Structures

These are exercises from the Bloc [Software Engineering Principles](https://www.bloc.io/software-engineering-track)

## Why GO?

The course itself is pretty much language agnostic, but provides recommendations to use Ruby, however being a Rubyist for the past 2 years, I am looking for a bit more of a challenge and have opted to use [Golang](https://golang.org/) for the exercises. 

### Tests

While going through these exercises in DataStructures, I will be writing Unit Tests for everything to ensure I understand and can prove these my solutions work. I am going through this course unguided and without a mentor and hope to leverage the GO community if I come across any limitations(there most likely will not) in my choice for this language.

## Completed Exercises
- [Intro to Data Structures](https://github.com/bdougie/bloc-data-structures/tree/master/intro_to_data_structures)
- [Stacks and Queues](https://github.com/bdougie/bloc-data-structures/tree/master/stacks_and_queues)
- [Linked List](https://github.com/bdougie/bloc-data-structures/tree/master/linked_lists)
- [Hashes](https://github.com/bdougie/bloc-data-structures/tree/master/hashes_1)
- [Binary Trees](https://github.com/bdougie/bloc-data-structures/tree/master/binary_trees)

## Setup

- install GO 1.6 from [source](https://golang.org/dl/)
- set up your [$GOPATH](https://www.kajabinext.com/marketplace/courses/1222)
- clone repo to your $GOPATH
- run `$ go run bloc_datastructures/intro_to_data_structures/assignment.go`

## Test

- `$ go get "github.com/stretchr/testify/assert"`
- `$ go test ./...`

If you want to run the test separately:

- `cd intro_to_data_structures` You can cd into each directory to run
tests separately
- `$ go test`

If you want to run a single test:

- `$ go test -v -run=TestPixelIsAddedToTheScreen`
