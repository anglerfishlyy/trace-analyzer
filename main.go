package main

import "fmt"

// Span represents a unit of work in a trace
type Span struct {
	ID       string
	Name     string
	Duration int // in ms
	Children []*Span
}

func getMockTrace() *Span {
	return &Span{
		ID:       "1",
		Name:     "gateway",
		Duration: 100,
		Children: []*Span{
			{
				ID:       "2",
				Name:     "auth",
				Duration: 120,
				Children: []*Span{
					{
						ID:       "3",
						Name:     "db",
						Duration: 200,
					},
				},
			},
			{
				ID:       "4",
				Name:     "payment",
				Duration: 300,
			},
		},
	}
}