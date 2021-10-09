package Car

var (
	Public_speed = 100
	private_name = "blablacar"
)

type My_Car struct {
	pr_number   int
	pr_weight   int
	Public_name string
}

func (car *My_Car) UpdateWeightCar(w int) {
	car.pr_weight = w
}
