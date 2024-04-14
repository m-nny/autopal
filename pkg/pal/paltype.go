package pal

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Type string

var (
	TypeNeutral  Type = "neutral"
	TypeFire     Type = "fire"
	TypeWater    Type = "water"
	TypeGrass    Type = "grass"
	TypeElectric Type = "electric"
	TypeIce      Type = "ice"
	TypeGround   Type = "ground"
	TypeDark     Type = "dark"
	TypeDragon   Type = "dragon"
)

var TypeAll = []Type{
	TypeNeutral,
	TypeFire,
	TypeWater,
	TypeGrass,
	TypeElectric,
	TypeIce,
	TypeGround,
	TypeDark,
	TypeDragon,
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	s = strings.ToLower(s)
	for _, item := range TypeAll {
		if string(item) == s {
			*t = item
			return nil
		}
	}
	return fmt.Errorf("Type %s not found", s)
}
