//package space converts the amount of time, in seconds, to the number of times any Sol. planet completes an orbit, planet specific year
package space

//type Planet is the only acceptable input type accepted by the tester
type Planet string

//constants used where sec is the amount of time in one Earth year
const sec float64 = 31557600

//constants table for all the planets using the Earth Year as a basis
func (P Planet) Reference() float64 {

	value := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61518726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return value[P]
}

//func Age converts the time, s, to the number of years of a specific planet, other_planet_year
func Age(s float64, p Planet) float64 {

	return (s / p.Reference() / sec)

}
