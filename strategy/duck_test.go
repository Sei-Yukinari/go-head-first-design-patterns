package strategy

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAllDucks(t *testing.T) {
	ducks := []Ducker{
		NewDecoyDuck(),
		NewMallardDuck(),
		NewModelDuck(),
		NewRubberDuck(),
		NewRedHeadDuck(),
		NewSimUDuck(),
	}
	for _, duck := range ducks {
		switch duck.(type) {
		case *DecoyDuck:
			assert.Empty(t, duck.PerformFly())
			assert.Empty(t, duck.PerformQuack())
		case *MallardDuck:
			assert.NotEmpty(t, duck.PerformFly())
			assert.Equal(t, duck.PerformQuack(), "Quack.")
		case *ModelDuck:
			assert.Empty(t, duck.PerformFly())
			assert.Equal(t, duck.PerformQuack(), "Quack.")
		case *RubberDuck:
			assert.Empty(t, duck.PerformFly())
			assert.Equal(t, duck.PerformQuack(), "Squeak.")
		case *RedHeadDuck:
			assert.NotEmpty(t, duck.PerformFly())
			assert.Equal(t, duck.PerformQuack(), "Quack.")
		case *SimUDuck:
			assert.NotEmpty(t, duck.PerformFly())
			assert.Equal(t, strings.Contains(duck.PerformFly(), "rocket"), true)
			assert.Empty(t, duck.PerformQuack())
		default:
			t.Error("invalid error duck type")
		}
	}
}
