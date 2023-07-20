package uuid

import (
	"github.com/go-fires/generator"
	uuid "github.com/satori/go.uuid"
)

type Generator struct{}

var _ generator.IDGenerator = (*Generator)(nil)

func (g *Generator) Generate() string {
	return uuid.NewV4().String()
}

var global = &Generator{}

func Generate() string {
	return global.Generate()
}
