package pal

import (
	"encoding/json"
	"flag"
	"iter"
	"os"
	"path/filepath"
)

var dataDir = flag.String("data_dir", "./data", "path to data dir")

type PalBaseId string
type PalBase struct {
	Id    PalBaseId
	Name  string
	Types []Type
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
