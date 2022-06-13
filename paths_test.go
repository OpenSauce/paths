package paths

import (
	"testing"
)

var p *Path

func BenchmarkGetPathFromCells(b *testing.B) {
	firstMap := NewGrid(200, 200, 16, 16)
	for i := 0; i < b.N; i++ {
		p = firstMap.GetPathFromCells(firstMap.Get(0, 0), firstMap.Get(10, 5), true, true)
	}
}

func BenchmarkGetPathFromCellsA(b *testing.B) {
	firstMap := NewGrid(200, 200, 16, 16)
	for i := 0; i < b.N; i++ {
		p = firstMap.GetPathFromCellsAStar(firstMap.Get(0, 0), firstMap.Get(10, 5), true, true)
	}
}

func TestGetPathFromCellsA(t *testing.T) {
	firstMap := NewGrid(200, 200, 16, 16)
	p1 := firstMap.GetPathFromCells(firstMap.Get(0, 0), firstMap.Get(10, 5), true, true)
	p2 := firstMap.GetPathFromCellsAStar(firstMap.Get(0, 0), firstMap.Get(10, 5), true, true)
	if !p1.Same(p2) {
		t.Fatalf("Path 1 %f:\n%s\nPath 2 %f:\n%s", p1.TotalCost(), p1, p2.TotalCost(), p2)
	}
}
