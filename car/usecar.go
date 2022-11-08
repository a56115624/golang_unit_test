package car

type UseData struct {
	D dataInterface
}

func (d *UseData) UseDivision(a, b float64) (float64, error) {
	// return d.D.Division(a, b)
	ans, err := d.D.Division(a+2, b)
	if err != nil {
		return 0, err
	}
	return ans + 3, nil
}
func (d *UseData) UseDot(a, b float64) (float64, error) {
	anser, err := d.D.DOT(a, b)
	if err != nil {
		return 1, err
	}
	return anser + 1, nil
}
