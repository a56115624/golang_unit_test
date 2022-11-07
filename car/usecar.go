package car

type UseData struct {
	D dataInterface
}

func (d *UseData) UseDivision(a, b float64) (float64, error) {
	// return d.D.Division(a, b)
	c, err := d.D.Division(a+2, b)
	if err != nil {
		return 0, err
	}
	return c + 3, nil
}
