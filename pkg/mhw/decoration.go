package mhw

type decoration struct {
	id int

	name    string
	size    int
	rarity  int
	skillId int
}

func newDecoration() *decoration {
	d := decoration{}
	return &d
}

type decorationInsertion struct {
	decorationId  int
	insertedCount int
}
