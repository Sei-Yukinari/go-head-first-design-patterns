package strategy

type Quacker interface {
	Quack() string
}

type RealQuack struct{}

func (RealQuack) Quack() string {
	return "Quack."
}

type Squeak struct{}

func (Squeak) Quack() string {
	return "Squeak."
}

type MuteQuack struct{}

func (MuteQuack) Quack() string {
	return ""
}
