package settings

import (
	"github.com/ryo33/zenv/util"
	"path"
	"strings"
)

type Setting struct {
	read       func(string) []string
	write      func([]string) string
	activate   func([]string, *Info) bool
	deactivate func([]string, *Info) bool
	initialize func(string)
	equal      func([]string, []string) bool
}

type Info struct {
	zenv   string
	envdir string
}

type Items map[string][][]string

var settings = map[string]Setting{
	"link": link,
}

const (
	RC = "rc" // File
)

func NewInfo(zenv, envdir string) *Info {
	return &Info{
		zenv:   zenv,
		envdir: envdir,
	}
}

func Read(dir string) Items {
	items := make(map[string][][]string)
	// each settings
	for key, se := range settings {
		if pa := path.Join(dir, key); util.Exists(pa) {
			its := [][]string{}
			// each lines of file
			for line, str := range util.ReadFile(pa) {
				// if succeeded reading
				if it := se.read(str); len(it) != 0 {
					its = append(its, it)
				} else {
					util.PrintErrorMessage(util.FormatErrorAtLine(pa, line, "Can't parse"))
				}
			}
			if len(its) > 0 {
				items[key] = its
			}
		}
	}
	pa := path.Join(dir, RC)
	// each lines of file
	for line, str := range util.ReadFile(pa) {
		read := false
		if fir := strings.Index(str, " "); fir != -1 {
			key := str[:fir]
			str := str[fir+1:]
			// if setting exists
			if se, ok := settings[key]; ok {
				// if succeeded reading
				if it := se.read(str); len(it) != 0 {
					is, ok := items[key]
					if !ok {
						is = [][]string{}
					}
					items[key] = append(is, it)
					read = true
				}
			}
		}
		if !read {
			util.PrintErrorMessage(util.FormatErrorAtLine(pa, line, "Can't parse"))
		}
	}
	return Items(items)
}

func (items Items) Write(dir string) {
	// each lables
	for key, its := range items.ToMap() {
		result := []string{}
		se, ok := settings[key]
		if !ok {
			util.PrintErrorMessageContinue(key + " lable is undefined")
		}
		for _, it := range its {
			if str := se.write(it); len(str) > 0 {
				result = append(result, str)
			} else {
				util.Print("黄金") // This should not be called.
			}
		}
		if len(result) > 0 {
			util.WriteFile(path.Join(dir, key), result)
		}
	}
}

func (items Items) Get(lable string) [][]string {
	its, ok := items.ToMap()[lable]
	if ok {
		result := [][]string{}
		for _, it := range its {
			result = append(result, []string(it))
		}
		return result
	} else {
		return [][]string{}
	}
}

func (items Items) Activate(info *Info) {
	for key, its := range items.ToMap() {
		se, ok := settings[key]
		if !ok {
			util.PrintErrorMessageContinue(key + " lable is undefined")
		} else {
			for _, it := range its {
				if !se.activate(it, info) {
					util.PrintErrorMessageContinue("aaaa") // TODO can't activate
				}
			}
		}
	}
}

func (items Items) Deactivate(info *Info) {
	for key, its := range items.ToMap() {
		se, ok := settings[key]
		if !ok {
			util.PrintErrorMessageContinue(key + " lable is undefined")
		} else {
			for _, it := range its {
				if !se.deactivate(it, info) {
					util.PrintErrorMessageContinue("aaaa") // TODO can't deactivate
				}
			}
		}
	}
}

func (items Items) AddItems(lable string, force bool, its2 [][]string) {
	its, ok := items.ToMap()[lable]
	se, ok2 := settings[lable]
	if !ok {
		its = [][]string{}
	}
	if !ok2 {
		util.PrintErrorMessageContinue(lable + " lable is undefined")
	} else {
		for idx, it := range its {
			tmp := its2
			for idx2, it2 := range tmp {
				if se.equal(it, it2) {
					if !force {
						util.PrintErrorMessage(strings.Join(it2, " ") + " already exists\n--force to overwrite")
					} else {
						items.ToMap()[lable][idx] = it2
					}
					its2 = append(its2[:idx2], its2[idx2+1:]...)
				}
			}
		}
		items.ToMap()[lable] = append(its, its2...)
	}
}

func (items Items) ToMap() map[string][][]string {
	return map[string][][]string(items)
}

func Initialize(dir string) {
	for _, se := range settings {
		se.initialize(dir)
	}
}
