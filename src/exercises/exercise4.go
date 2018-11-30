package exercises

// Person interface
type Person interface {
	Pay(prize float64) float64
}

// Jubilado person struct
type Jubilado struct {
	name string
}

// Normal person struct
type Normal struct {
	Jubilado
}

// Invitado person struct
type Invitado struct {
	Jubilado
}

// Pay Jubilado func
func (p Jubilado) Pay(prize float64) float64 {
	return 0.75 * prize
}

// Pay Normal func
func (p Normal) Pay(prize float64) float64 {
	return prize
}

// Pay Invitado func
func (p Invitado) Pay(prize float64) float64 {
	return 0
}
