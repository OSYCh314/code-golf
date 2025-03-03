package config

import (
	"cmp"
	"slices"
	"strings"
)

var (
	// Standard languages.
	LangByID = map[string]*Lang{}
	LangList []*Lang

	// Experimental languages.
	ExpLangByID = map[string]*Lang{}
	ExpLangList []*Lang

	// All languages.
	AllLangByID = map[string]*Lang{}
	AllLangList []*Lang

	// Redirects.
	LangRedirects = map[string]string{}
)

type Lang struct {
	Args, Redirects, Env []string `json:"-"`
	ArgsQuine            []string `json:"-" toml:"args-quine"`
	Example              string   `json:"example"`
	Experiment           int      `json:"experiment,omitempty"`
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Size                 string   `json:"size"`
	Version              string   `json:"version"`
	Website              string   `json:"website"`
}

func init() {
	var langs map[string]*Lang
	unmarshal("data/langs.toml", &langs)

	for name, lang := range langs {
		lang.Example = strings.TrimSuffix(lang.Example, "\n")
		lang.ID = ID(name)
		lang.Name = name

		// Redirects.
		for _, redirect := range lang.Redirects {
			LangRedirects[redirect] = lang.ID
		}

		AllLangByID[lang.ID] = lang
		AllLangList = append(AllLangList, lang)

		if lang.Experiment == 0 {
			LangByID[lang.ID] = lang
			LangList = append(LangList, lang)
		} else {
			ExpLangByID[lang.ID] = lang
			ExpLangList = append(ExpLangList, lang)
		}
	}

	// Case-insensitive sort.
	for _, langs := range [][]*Lang{LangList, ExpLangList, AllLangList} {
		slices.SortFunc(langs, func(a, b *Lang) int {
			return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
		})
	}
}
