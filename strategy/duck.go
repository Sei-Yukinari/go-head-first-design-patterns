package strategy

type Ducker interface {
	Display() string
	Swim() string
	PerformFly() string
	PerformQuack() string
}

type Duck struct {
	flyBehavior   Flyer
	quackBehavior Quacker
}

func (duck *Duck) SetFlyBehavior(flyBehavior Flyer) {
	duck.flyBehavior = flyBehavior
}

func (duck *Duck) SetQuackBehavior(quackBehavior Quacker) {
	duck.quackBehavior = quackBehavior
}

func (duck Duck) PerformFly() string {
	return duck.flyBehavior.Fly()
}

func (duck Duck) PerformQuack() string {
	return duck.quackBehavior.Quack()
}

func (Duck) Swim() string {
	return "All ducks float, even decoys."
}

type MallardDuck struct {
	Duck
}

func (MallardDuck) Display() string {
	return "I'm a real Mallarduck duck."
}

func NewMallardDuck() *MallardDuck {
	mallardDuck := MallardDuck{}
	mallardDuck.SetFlyBehavior(FlyWithWings{})
	mallardDuck.SetQuackBehavior(RealQuack{})
	return &mallardDuck
}

type ModelDuck struct {
	Duck
}

func (ModelDuck) Display() string {
	return "I'm a model duck."
}

func NewModelDuck() *ModelDuck {
	modelDuck := ModelDuck{}
	modelDuck.SetFlyBehavior(FlyNoWay{})
	modelDuck.SetQuackBehavior(RealQuack{})
	return &modelDuck
}

type RedHeadDuck struct {
	Duck
}

func (RedHeadDuck) Display() string {
	return "I'm a real Reduck Headeduck duck."
}

func NewRedHeadDuck() *RedHeadDuck {
	redHeadDuck := RedHeadDuck{}
	redHeadDuck.SetFlyBehavior(FlyWithWings{})
	redHeadDuck.SetQuackBehavior(RealQuack{})
	return &redHeadDuck
}

type RubberDuck struct {
	Duck
}

func (RubberDuck) Display() string {
	return "I'm a rubber duck."
}

func NewRubberDuck() *RubberDuck {
	rubberDuck := RubberDuck{}
	rubberDuck.SetFlyBehavior(FlyNoWay{})
	rubberDuck.SetQuackBehavior(Squeak{})
	return &rubberDuck
}

type DecoyDuck struct {
	Duck
}

func (DecoyDuck) Display() string {
	return "I'm a duck Decoy."
}

func NewDecoyDuck() *DecoyDuck {
	decoyDuck := DecoyDuck{}
	decoyDuck.SetFlyBehavior(FlyNoWay{})
	decoyDuck.SetQuackBehavior(MuteQuack{})
	return &decoyDuck
}

type SimUDuck struct {
	Duck
}

func (SimUDuck) Display() string {
	return "I'm a simU Duck."
}

func NewSimUDuck() *SimUDuck {
	s := SimUDuck{}
	s.SetFlyBehavior(FlyRocketPowered{})
	s.SetQuackBehavior(MuteQuack{})
	return &s
}
