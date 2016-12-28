package main

import "fmt"

type Trainer struct {
	card_common
	c_text string
}

func (c Trainer) String() string {
	return fmt.Sprintf(
		"Name: %s\nType: %s\nSet: %s\nID#: %s\nRarity: %s\n\nText: %s",
		c.c_name,
		c.c_type,
		c.c_set,
		c.c_id,
		c.c_rarity,
		c.c_text,
	)
}

func NewTrainer(
	c_name string, c_type string, c_set string, c_id string,
	c_rarity string, c_text string) *Trainer {
	common := card_common{
		c_name,
		c_type,
		c_set,
		c_id,
		c_rarity,
	}
	trainer := Trainer{
		common,
		c_text,
	}
	return &trainer
}
