package lettermark_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/botforge-pro/lettermark"
)

type corpus struct {
	Initials []struct {
		Case       string `yaml:"case"`
		Name       string `yaml:"name"`
		Expect     string `yaml:"expect"`
		Codepoints string `yaml:"expect_codepoints"`
	} `yaml:"initials"`
	Slot []struct {
		Case   string `yaml:"case"`
		ID     int64  `yaml:"id"`
		Slots  int    `yaml:"slots"`
		Expect int    `yaml:"expect"`
	} `yaml:"slot"`
}

func read(t *testing.T) corpus {
	t.Helper()
	body, err := os.ReadFile("cases.yaml")
	require.NoError(t, err)
	var cases corpus
	require.NoError(t, yaml.Unmarshal(body, &cases))
	return cases
}

func TestCorpus_AnswersEveryCase(t *testing.T) {
	cases := read(t)
	require.NotEmpty(t, cases.Initials, "the corpus is the contract and it came back empty")

	for _, one := range cases.Initials {
		t.Run(one.Case, func(t *testing.T) {
			got := lettermark.Initials(one.Name)
			assert.Equal(t, one.Expect, got, "%q", one.Name)
			if one.Codepoints != "" {
				assert.Equal(t, one.Codepoints, codepointsOf(got),
					"the letters compare equal but are written differently, which a port that skips normalising would also pass")
			}
		})
	}
}

func TestCorpus_TheSlotsItNames(t *testing.T) {
	cases := read(t)
	require.NotEmpty(t, cases.Slot)

	for _, one := range cases.Slot {
		t.Run(one.Case, func(t *testing.T) {
			palette := lettermark.NewPalette(one.Slots)
			assert.Equal(t, one.Slots, palette.Slots())
			assert.Equal(t, one.Expect, palette.Slot(one.ID))
		})
	}
}

func TestNewPalette_RefusesACountThatPaintsNothing(t *testing.T) {
	for _, slots := range []int{0, -1} {
		assert.Panics(t, func() { lettermark.NewPalette(slots) },
			"a palette of %d colours means the code and the theme disagree, and a slot handed out now would be a colour nobody painted", slots)
	}
}

func TestZeroPalette_SaysWhatIsMissing(t *testing.T) {
	var unbuilt lettermark.Palette
	assert.PanicsWithValue(t,
		"lettermark: the zero Palette holds no colours; build it with NewPalette",
		func() { unbuilt.Slot(1) },
		"the bare remainder would divide by zero and name neither the type nor the way to build one")
}

func codepointsOf(text string) string {
	var out []string
	for _, r := range text {
		out = append(out, fmt.Sprintf("%04X", r))
	}
	return strings.Join(out, " ")
}

func ExampleInitials() {
	fmt.Println(lettermark.Initials("Wiki Layer"))
	// Output: WL
}
