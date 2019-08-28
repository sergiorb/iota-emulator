package generator

import "math/rand"
import "fmt" 

type RandomFloatGenerator struct {

}

func (g *RandomFloatGenerator) Generate() string {

	return fmt.Sprintf("%f", rand.Float32());
}
