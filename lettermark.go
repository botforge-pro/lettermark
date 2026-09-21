package lettermark

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/rivo/uniseg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

var fullUppercase = cases.Upper(language.Und)

type Palette struct {
	slots int
}

func NewPalette(slots int) Palette {
	if slots < 1 {
		panic(fmt.Sprintf("lettermark: a palette of %d colours has no slot to hand out", slots))
	}
	return Palette{slots: slots}
}

func (p Palette) Slots() int {
	return p.slots
}

func (p Palette) Slot(id int64) int {
	if p.slots < 1 {
		panic("lettermark: the zero Palette holds no colours; build it with NewPalette")
	}
	slots := int64(p.slots)
	return int(((id % slots) + slots) % slots)
}

func Initials(name string) string {
	words := wordsOf(norm.NFC.String(name))

	var units []string
	for _, word := range words {
		unit := firstSyllableOf(word)
		if unit == "" {
			continue
		}
		units = append(units, upperKeepingLength(unit))
		if len(units) == 2 || writesAsOneSyllableWhole(unit) {
			break
		}
	}
	if len(units) == 0 {
		return norm.NFC.String(firstDrawableCluster(norm.NFC.String(name)))
	}
	return norm.NFC.String(strings.Join(units, ""))
}

func wordsOf(name string) []string {
	return strings.FieldsFunc(name, isWordBreak)
}

func isWordBreak(r rune) bool {
	switch r {
	case 0x0085, 0x00A0, 0x1680, 0x2028, 0x2029, 0x202F, 0x205F, 0x3000:
		return true
	}
	return (r >= 0x0009 && r <= 0x000D) || r == 0x0020 || (r >= 0x2000 && r <= 0x200A)
}

func firstSyllableOf(word string) string {
	clusters := clustersOf(word)
	for i, cluster := range clusters {
		if !opensAUnit(firstScalar(cluster)) {
			continue
		}
		return grownIntoASyllable(cluster, clusters[i+1:])
	}
	return ""
}

func grownIntoASyllable(unit string, rest []string) string {
	for len(rest) > 0 {
		switch {
		case isVirama(lastScalar(unit)), isPrefixVowel(unit):
			unit += rest[0]
			rest = rest[1:]
		default:
			return unit
		}
	}
	return unit
}

func isPrefixVowel(unit string) bool {
	scalars := []rune(unit)
	if len(scalars) != 1 {
		return false
	}
	r := scalars[0]
	return (r >= 0x0E40 && r <= 0x0E44) || (r >= 0x0EC0 && r <= 0x0EC4)
}

func isVirama(r rune) bool {
	switch r {
	case 0x094D, 0x09CD, 0x0A4D, 0x0ACD, 0x0B4D, 0x0BCD, 0x0C4D, 0x0CCD,
		0x0D4D, 0x0DCA, 0x0F84, 0x1039, 0x17D2:
		return true
	}
	return false
}

func opensAUnit(r rune) bool {
	return unicode.In(r, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lo, unicode.Nd)
}

func writesAsOneSyllableWhole(unit string) bool {
	r := firstScalar(unit)
	for _, span := range [][2]rune{
		{0x0600, 0x06FF}, {0x0700, 0x074F}, {0x0750, 0x077F}, {0x07C0, 0x07FF},
		{0x0860, 0x086F}, {0x08A0, 0x08FF}, {0xFB50, 0xFDFF}, {0xFE70, 0xFEFF},
		{0x1E900, 0x1E95F},
		{0x1100, 0x11FF}, {0x3130, 0x318F}, {0xA960, 0xA97F}, {0xAC00, 0xD7FF},
		{0xFFA0, 0xFFDC},
		{0x2E80, 0x2EFF}, {0x3400, 0x4DBF}, {0x4E00, 0x9FFF}, {0xF900, 0xFAFF},
		{0x20000, 0x3134F},
		{0x3040, 0x30FF}, {0x31F0, 0x31FF}, {0xFF66, 0xFF9F},
		{0x1800, 0x18AF},
	} {
		if r >= span[0] && r <= span[1] {
			return true
		}
	}
	return false
}

func upperKeepingLength(unit string) string {
	if isGeorgianMkhedruli(firstScalar(unit)) {
		return unit
	}
	uppercased := fullUppercase.String(unit)
	if len([]rune(uppercased)) > len([]rune(unit)) {
		return unit
	}
	return uppercased
}

func isGeorgianMkhedruli(r rune) bool {
	return r >= 0x10D0 && r <= 0x10FF
}

func firstDrawableCluster(name string) string {
	for _, cluster := range clustersOf(name) {
		r := firstScalar(cluster)
		if isWordBreak(r) || unicode.In(r, unicode.M, unicode.C) {
			continue
		}
		return cluster
	}
	return ""
}

func clustersOf(text string) []string {
	var out []string
	state := -1
	for len(text) > 0 {
		var cluster string
		cluster, text, _, state = uniseg.FirstGraphemeClusterInString(text, state)
		out = append(out, cluster)
	}
	return out
}

func firstScalar(cluster string) rune {
	for _, r := range cluster {
		return r
	}
	return 0
}

func lastScalar(unit string) rune {
	scalars := []rune(unit)
	if len(scalars) == 0 {
		return 0
	}
	return scalars[len(scalars)-1]
}
