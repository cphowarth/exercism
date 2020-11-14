// Package space - caclulate the when living on various planets
package space

// Planet type string definition
type Planet string

// Age - work out the age on a planet in years
func Age(ageSeconds float64, planetname Planet) float64 {
	orbitTime := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return ageSeconds / (orbitTime[planetname] * 365.25 * 24 * 60 * 60)
}
