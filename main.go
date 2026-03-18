package main

import "fmt"

// Span represents a unit of work in a trace
type Span struct {
	ID       string
	Name     string
	Duration int // in ms
	Children []*Span
}