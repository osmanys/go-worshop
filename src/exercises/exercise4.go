package exercises

type Person interface {
	Pay(prize float64) float64
}

type Jubilado struct {
	name string
}

type Normal struct {
	Jubilado
}

type Invitado struct {
	Jubilado
}

func (p Jubilado) Pay(prize float64) float64 {
	return 0.75 * prize
}

func (p Normal) Pay(prize float64) float64 {
	return prize
}

func (p Invitado) Pay(prize float64) float64 {
	return 0
}
