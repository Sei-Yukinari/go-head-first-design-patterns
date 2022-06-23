package strategy

type Flyer interface {
	Fly() string
}

type FlyWithWings struct {
}

func (FlyWithWings) Fly() string {
	return "I am Flying."
}

type FlyNoWay struct {
}

func (FlyNoWay) Fly() string {
	return ""
}

type FlyRocketPowered struct {
}

func (FlyRocketPowered) Fly() string {
	return "I'm flying with a rocket."
}
