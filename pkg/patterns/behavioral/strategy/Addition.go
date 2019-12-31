package strategy

type Addition struct{}

func (Addition) Apply(leftVal, rightVal int) int {
	return leftVal + rightVal
}
