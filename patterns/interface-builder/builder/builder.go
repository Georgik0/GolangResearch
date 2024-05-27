package builder

type Data struct {
	name string
}

func (d *Data) modifiedName() string {
	return d.name
}

type dataProvider interface {
	modifiedName() string
}
