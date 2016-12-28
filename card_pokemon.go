package main

import "fmt"

type Pokemon struct {
	card_common
	c_hp int
}

func (c Pokemon) Hp() int {
	return c.c_hp
}

func (c Pokemon) String() string {
	return fmt.Sprintf(
		"Name: %s\nType: %s\nHP: %d\nSet: %s\nID#: %s\nRarity: %s",
		c.c_name,
		c.c_type,
		c.c_hp,
		c.c_set,
		c.c_id,
		c.c_rarity,
	)
}

func NewPokemon(
	c_name string, c_type string, c_set string, c_id string,
	c_rarity string, c_hp int) *Pokemon {
	common := card_common{
		c_name,
		c_type,
		c_set,
		c_id,
		c_rarity,
	}
	pokemon := Pokemon{
		common,
		c_hp,
	}
	return &pokemon
}
