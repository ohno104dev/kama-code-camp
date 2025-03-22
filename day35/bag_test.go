package day35

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test ./day35 -v

func TestBagProblem2D(t *testing.T) {
	fmt.Println("bagProblem2D res:", bagProblem2D())
	assert.Equal(t, 35, bagProblem2D())
}

func TestBagProblem1D(t *testing.T) {
	fmt.Println("bagProblem1D res:", bagProblem1D())
	assert.Equal(t, 35, bagProblem1D())
}
