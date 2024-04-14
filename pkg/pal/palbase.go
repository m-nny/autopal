package pal

import (
	"encoding/json"
	"flag"
	"fmt"
	"iter"
	"os"
	"path/filepath"
)

var dataDir = flag.String("data_dir", "./data", "path to data dir")

type PalBaseId string

var _ fmt.Stringer = (*PalBase)(nil)

type PalBase struct {
	Id          PalBaseId
	Name        string
	BaseHp      int `json:"base_hp"`
	BaseAttack  int `json:"base_attack"`
	BaseDefence int `json:"base_defence"`
	Types       []Type
}

func (pb *PalBase) String() string {
	return fmt.Sprintf("%s (%s)", pb.Name, pb.Id)
}

type PalBases []PalBase

func (b PalBases) All() iter.Seq[PalBase] {
	return func(yield func(PalBase) bool) {
		for _, item := range b {
			if !yield(item) {
				return
			}
		}
	}
}

var _palBases PalBases

func AllPalBases() iter.Seq[PalBase] {
	return _palBases.All()
}

func LoadPalBases() error {
	palBasesPath := filepath.Join(*dataDir, "pal_basic.json")
	fullFilePath, err := filepath.Abs(palBasesPath)
	if err != nil {
		return err
	}
	buf, err := os.ReadFile(fullFilePath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(buf, &_palBases); err != nil {
		return err
	}
	return nil
}
