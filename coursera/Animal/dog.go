package Animal

type Dog struct {
	Name string
	id   int
}

func GetName(dog *Dog) int {
	return dog.id
}
