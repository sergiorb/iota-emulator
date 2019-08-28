package generator

type GeneratorType string

const (
    RandomInt GeneratorType = "random_int"
    RandomFloat GeneratorType = "random_float"
)

type Generator interface {

    Generate() string
}

func BuildGenerator(gtype GeneratorType) Generator {

    if (gtype == RandomFloat) {

        return &RandomFloatGenerator{}

    } else {

        return &RandomIntGenerator{}
    }

}
