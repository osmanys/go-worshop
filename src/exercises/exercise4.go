package exercises

type Person interface {
  Pay()
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

func (p Jubilado) Pay(prize int) int{
  return 0.75 * prize
}

func (p Normal) Pay(prize int) int{
  return prize
}

func (p Invitado) Pay(prize int) int{
  return 0
}

