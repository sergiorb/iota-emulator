package generator

import "math/rand"
import "strconv" 

type RandomIntGenerator struct {

}

func (g *RandomIntGenerator) Generate() string {
	
	return strconv.Itoa(rand.Intn(100));
}
