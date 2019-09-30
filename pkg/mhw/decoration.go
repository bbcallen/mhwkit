package mhw

type decoration struct {
	id int

	name           string
	size           int
	rarity         int
	skillId        int
    enhancedLevel  int
}

func newDecoration() *decoration {
	d := decoration{}
    d.enhancedLevel = 1
	return &d
}

type decorationInsertion struct {
	decorationId  int
	insertedCount int
}
