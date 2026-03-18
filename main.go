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
func printTrace(span *Span, level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("%s (%dms)\n", span.Name, span.Duration)

	for _, child := range span.Children {
		printTrace(child, level+1)
	}
}
func findSlowest(span *Span) *Span {
	slowest := span

	for _, child := range span.Children {
		candidate := findSlowest(child)
		if candidate.Duration > slowest.Duration {
			slowest = candidate
		}
	}

	return slowest
}
func findCriticalPath(span *Span) ([]string, int) {
	path := []string{span.Name}
	total := span.Duration

	if len(span.Children) == 0 {
		return path, total
	}

	var maxChild *Span
	for _, child := range span.Children {
		if maxChild == nil || child.Duration > maxChild.Duration {
			maxChild = child
		}
	}

	childPath, childTotal := findCriticalPath(maxChild)

	return append(path, childPath...), total + childTotal
}
func main() {
	traceID := "abc123"
	trace := getMockTrace()

	fmt.Println("TraceID:", traceID)
	fmt.Println("\nTrace Tree:")
	printTrace(trace, 0)

	path, total := findCriticalPath(trace)

	fmt.Println("\nCritical Path:")
	for i, p := range path {
		if i > 0 {
			fmt.Print(" → ")
		}
		fmt.Print(p)
	}
	fmt.Println()

	fmt.Printf("\nTotal Latency: %dms\n", total)

	slowest := findSlowest(trace)
	fmt.Printf("\nSlowest Span:\n%s (%dms)\n", slowest.Name, slowest.Duration)
}