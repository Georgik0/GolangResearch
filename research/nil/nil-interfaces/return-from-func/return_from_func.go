package return_from_func

import "fmt"

type Data struct {
	s string
}

func (d *Data) Info() {
	fmt.Println(d.s)
}

type Provider interface {
	Info()
}

func Research() {
	p := getNilProvider()

	p.Info()
}

func getNilProvider() Provider {
	return nil
}
