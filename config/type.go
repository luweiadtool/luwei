package config

import "image"

type ForDealCon struct {
	Puzzle    puzzle  `toml:"puzzle"`
	CatId     int64   `toml:"catId"`
	Limit     []int64 `toml:"limit"`
	DetailIds []int64 `toml:"detailIds"`
	RootPath  string
}

type puzzle struct {
	Style    int    `toml:"style"`
	Color    string `toml:"color"`
	Fontsize int    `toml:"fontsize"`
	Count    int    `toml:"count"`
	Logo     string  `toml:"logo"`
	LogoImage image.Image
}
