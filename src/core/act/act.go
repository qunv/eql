package act

type Act interface {
	Calculate(input [][]float64) float64
}
