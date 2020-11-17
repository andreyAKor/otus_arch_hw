package actions

var _ Doer = (*Move)(nil)

// Структура движения
type Move struct {
	m Movable
}

func NewMove(m Movable) *Move {
	return &Move{m}
}

func (m *Move) Do() {
	m.m.SetPosition(m.m.GetPosition().Add(m.m.GetVelocity()))
}
