package interp

type Func func(t float64) float64

type MultiFunc func(t float64) []float64

func Linear(x0, x1 float64, n int) []float64 {
	f := LinearFunc(x0, x1, 0, float64(n))
	
}

func LinearFunc(x0, x1, t0, t1 float64) Func {
	if t0 == t1 {
		panic("interp: t0 == t1")
	}
	k := (x1 - x0) / (t1 - t0)
	return func(t float64) float64 {
		return k*t + t0
	}
}

func LinearMultiFunc(x0, x1 []float64, t0, t1 float64) MultiFunc {
	return nil
}
