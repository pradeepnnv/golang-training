package car

import (
	"strings"
)

//go:generate stringer -type=Brand

type Brand int

const (
	UnknownBrand Brand = iota
	BMW
	Mercedes
	Audi
	Honda
	Toyota
	Volkswagon
	Ferrari
	Porsche
	Tesla
)

func (b Brand) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *Brand) UnmarshalText(text []byte) error {
	*b = BrandFromText(string(text))
	return nil
}

func BrandFromText(b string) Brand {
	switch strings.ToLower(b) {
	default:
		return UnknownBrand
	case "bmw":
		return BMW
	case "ferrari":
		return Ferrari
	case "mercedes":
		return Mercedes
	case "audi":
		return Audi
	case "honda":
		return Honda
	case "toyota":
		return Toyota
	case "volkswagon":
		return Volkswagon
	case "porsche":
		return Porsche
	case "tesla":
		return Tesla
	}
}
