package strategy

type Multiplication struct{}

func (Multiplication) Apply(leftVal, rightVal int) int {
	return leftVal * rightVal
}
