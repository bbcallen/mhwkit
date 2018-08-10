package mhw

type skill struct {
	id int

	name     string
	maxLevel int
}

func newSkill() *skill {
	s := skill{}
	return &s
}

type skillEnhancement struct {
	skillId       int
	enhancedLevel int
}
