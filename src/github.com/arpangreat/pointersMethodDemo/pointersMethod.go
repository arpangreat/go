package main

import "fmt"

const uSixteenMax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / uSixteenMax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / uSixteenMax / kmh_multiple)
}

func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{gas_pedal: 2234, brake_pedal: 0, steering_wheel: 12561, top_speed_kmh: 225.0}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}
