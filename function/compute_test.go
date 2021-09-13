package p

import (
	"github.com/stretchr/testify/assert"
	"time"

	"math/rand"
	"testing"
)

func getRand() (n1, n2 int){
	rand.Seed(time.Now().UnixNano())
	n1 = 100 + rand.Intn(-100-100+1)
	n2 = 100 + rand.Intn(-100-100+1)
	return n1, n2
}

func testAdd(t *testing.T) {
	n1, n2 := getRand()
	res := n1 + n2
	assert.Equal(t, ComputationFunctions["add"], res)
}

func testSub(t *testing.T) {
	n1, n2 := getRand()
	res := n1 + n2
	assert.Equal(t, ComputationFunctions["sub"], res)
}