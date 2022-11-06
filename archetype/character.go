package archetype

type Character struct {
	Body       Body
	Hair       Hair
	FacialHair Hair
	Equipment  Equipment
}

type Body struct {
	Type  int
	Color int
}

type Hair struct {
	Type  int
	Color int
}

type Equipment struct {
	Head     Armor
	Chest    Armor
	Legs     Armor
	Feet     Armor
	MainHand Weapon
	OffHand  Weapon
}

type Armor struct {
	ID      int
	Defense int
}

type Weapon struct {
	ID int
}
