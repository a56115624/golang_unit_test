package car

type UseData struct {
	D dataInterface
}

// 實際執行的地方UseDivision
func (d *UseData) UseDivision(a, b float64) (float64, error) {
	// return d.D.Division(a, b)
	ans, err := d.D.Division(a+2, b)
	if err != nil {
		return 0, err
	}
	return ans + 3, nil
}

// 實際執行的地方UseDot
func (d *UseData) UseDot(a, b float64) (float64, error) {
	anser, err := d.D.DOT(a, b)
	if err != nil {
		return 1, err
	}
	return anser + 1, nil
}
func (d *UseData) UseAdd(a, b float64) (float64, error) {
	anser, err := d.D.Add(a+1, b)
	if err != nil {
		return 1, err
	}
	return anser + 2, nil
}
