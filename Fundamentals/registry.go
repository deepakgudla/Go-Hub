package main

import "sort"

// Example represents a runnable code
type Example interface {
	Name() string
	Run()
}

// registry holds all registered examples
var registry []Example

// Register adds an example to the registry
func Register(ex Example) {
	registry = append(registry, ex)
}

func GetAllExamples() []Example {
	examples := make([]Example, len(registry))
	copy(examples, registry)

	sort.Slice(examples, func(i, j int) bool {
		return examples[i].Name() < examples[j].Name()
	})

	return examples
}
