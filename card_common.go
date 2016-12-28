package main

import "fmt"

type Card interface {
	Name() string
	Type() string
	Set() string
	Id() string
	Rarity() string
	String() string
}

type card_common struct {
	c_name   string
	c_type   string
	c_set    string
	c_id     string
	c_rarity string
}

func (c card_common) Name() string {
	return c.c_name
}

func (c card_common) Type() string {
	return c.c_type
}

func (c card_common) Id() string {
	return c.c_id
}

func (c card_common) Set() string {
	return c.c_set
}

func (c card_common) Rarity() string {
	return c.c_rarity
}

func (c card_common) String() string {
	return fmt.Sprintf(
		"Name: %s\nType: %s\nSet: %s\nID#: %s\nRarity: %s",
		c.c_name,
		c.c_type,
		c.c_set,
		c.c_id,
		c.c_rarity,
	)
}

func NewCard(
	c_name string, c_type string, c_set string, c_id string,
	c_rarity string) *card_common {
	card := card_common{
		c_name, c_type, c_set, c_id, c_rarity,
	}

	return &card
}
