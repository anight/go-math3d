package math3d

import "math"

func radians(degrees float64) float64 {
	return math.Pi * degrees / 180.0
}

func degrees(radians float64) float64 {
	return 180.0 * radians / math.Pi
}

func hypot(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func Angle(v1, v2 [3]float64) float64 {
	cosa := Mul_v_v_s(v1, v2)
	return math.Acos(cosa)
}

func Distance(v1, v2 [3]float64) float64 {
	dx, dy, dz := v1[0]-v2[0], v1[1]-v2[1], v1[2]-v2[2]
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func Normalize(v [3]float64) [3]float64 {
	d := Distance(v, [3]float64{0, 0, 0})
	return [3]float64{v[0] / d, v[1] / d, v[2] / d}
}

func Mul_m_v_v(m [3][3]float64, v [3]float64) [3]float64 {
	x := m[0][0]*v[0] + m[1][0]*v[1] + m[2][0]*v[2]
	y := m[0][1]*v[0] + m[1][1]*v[1] + m[2][1]*v[2]
	z := m[0][2]*v[0] + m[1][2]*v[1] + m[2][2]*v[2]
	return [3]float64{x, y, z}
}

func Mul_v_v_v(v1, v2 [3]float64) [3]float64 {
	return [3]float64{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}
}

func Mul_v_v_s(v1, v2 [3]float64) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func Ll2xyz(ll [2]float64) [3]float64 {
	lon, lat := radians(ll[0]), radians(ll[1])
	z := math.Sin(lat)
	x := math.Cos(lat) * math.Cos(lon)
	y := math.Cos(lat) * math.Sin(lon)
	return [3]float64{x, y, z}
}

func Xyz2ll(v [3]float64) [2]float64 {
	var lon, lat float64
	lat = degrees(math.Asin(v[2]))
	l := hypot(v[0], v[1])
	if l > 1e-10 {
		lon = math.Copysign(degrees(math.Acos(v[0])), v[1])
	} else {
		lon = 0
	}
	return [2]float64{lon, lat}
}

func Rotate(axis, v [3]float64, angle float64) [3]float64 {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)
	x, y, z := axis[0], axis[1], axis[2]
	xx, yy, zz, xy, xz, yz := x*x, y*y, z*z, x*y, x*z, y*z
	// http://en.wikipedia.org/wiki/Rotation_matrix#Axis_of_a_rotation
	m := [3][3]float64{
		{xx + (1-xx)*cosA, xy*(1-cosA) + z*sinA, xz*(1-cosA) - y*sinA},
		{xy*(1-cosA) - z*sinA, yy + (1-yy)*cosA, yz*(1-cosA) + x*sinA},
		{xz*(1-cosA) + y*sinA, yz*(1-cosA) - x*sinA, zz + (1-zz)*cosA},
	}
	return Mul_m_v_v(m, v)
}

func Det_m(m [3][3]float64) float64 {
	return m[0][0]*m[1][1]*m[2][2] + m[1][0]*m[2][1]*m[0][2] + m[2][0]*m[0][1]*m[1][2] -
		m[2][0]*m[1][1]*m[0][2] - m[0][0]*m[2][1]*m[1][2] - m[1][0]*m[0][1]*m[2][2]
}

func Inv_m(m [3][3]float64) [3][3]float64 {
	// http://en.wikipedia.org/wiki/Inverse_matrix#Inversion_of_3.C3.973_matrices
	dr := 1 / Det_m(m)
	r0 := Mul_v_v_v(m[1], m[2])
	r1 := Mul_v_v_v(m[2], m[0])
	r2 := Mul_v_v_v(m[0], m[1])
	return [3][3]float64{
		{dr * r0[0], dr * r1[0], dr * r2[0]},
		{dr * r0[1], dr * r1[1], dr * r2[1]},
		{dr * r0[2], dr * r1[2], dr * r2[2]},
	}
}

func Add_v(v1 [3]float64, v2 [3]float64) [3]float64 {
	return [3]float64{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func Neg_v(v [3]float64) [3]float64 {
	return [3]float64{-v[0], -v[1], -v[2]}
}

func Mul_v(v [3]float64, k float64) [3]float64 {
	return [3]float64{k * v[0], k * v[1], k * v[2]}
}

