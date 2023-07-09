package matrix
type matrix struct {
	length int
	width int
}

func NewMatrix(x,y int) *matrix {
	m:=new(matrix)
	m.length=x
	m.width=y

	return m
}
