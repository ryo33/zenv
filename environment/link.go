package environment

import (
	"github.com/ryo33/zenv/util"
	"path"
	"strings"
)

type Link struct {
	name   string
	source string
}

const (
	LINKS = "links" // Both
)

func (link *Link) Name() string {
	return link.name
}

func (link *Link) Source() string {
	return link.source
}

func (env *Env) Links() []*Link {
	return env.links
}

func (env *Env) GetLinksPath() string {
	return env.GetPath(LINKS)
}

func NewLink(name, source string) *Link {
	return &Link{
		name:   name,
		source: source,
	}
}

func (env *Env) readLinks() {
	tmp := util.ReadFile(path.Join(env.dir, LINKS, LINKS))
	util.PrintDebug(strings.Join(tmp, "\n"))
	env.links = []*Link{}
	for _, li := range tmp {
		re := strings.Split(li, SEPARATOR)
		if len(re) == 2 {
			env.links = append(env.links, NewLink(re[0], re[1]))
		}
	}
}

func (env *Env) writeLinks() {
	linksPath := path.Join(env.dir, LINKS)
	util.PrepareDir(linksPath)
	for _, li := range env.links {
		//TODO is there smart way?
		util.ExecCommand("ln", "-s", li.source, path.Join(linksPath, li.name))
	}
	data := []string{}
	for _, li := range env.links {
		data = append(data, li.name+SEPARATOR+li.source)
	}
	util.WriteFile(path.Join(linksPath, LINKS), data)
}

func (env *Env) AddLink(link *Link, force bool) {
	exists := false
	for i, li := range env.links {
		if li.name == link.name {
			exists = true
			if force {
				env.links[i].source = link.source
			} else {
				util.PrintErrorMessage(`
				link already exists
				--force flag to overwrite
				`)
			}
			break
		}
	}
	if !exists {
		env.links = append(env.links, link)
	}
}

func (env *Env) RemoveLinks(links []string) {
	result := []*Link{}
	for _, li := range env.links {
		if !util.Contains(links, li.name) {
			result = append(result, li)
		}
	}
	env.links = result
}
