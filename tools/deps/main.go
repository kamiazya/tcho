package main

import (
	"fmt"
	"go/build"
	"strings"
)

var (
	buildContext = build.Default

	packageCtx = map[string]*Package{}

	idCounter int

	groups = []*Group{
		{
			name:     "tcho",
			path:     "bitbucket.org/kamiazya",
			packages: []string{},
			subGroups: []*Group{
				{
					name:     "cli",
					path:     "bitbucket.org/kamiazya/tcho/cli",
					packages: []string{},
					subGroups: []*Group{
						{
							name:     "adapter",
							path:     "bitbucket.org/kamiazya/tcho/cli/adapter",
							packages: []string{},
						},
						// {
						// 	name:     "extarnal",
						// 	path:     "bitbucket.org/kamiazya/tcho/cli/extarnal",
						// 	packages: []string{},
						// },
						{
							name:     "module",
							path:     "bitbucket.org/kamiazya/tcho/cli/module",
							packages: []string{},
							subGroups: []*Group{
								{
									name:     "local",
									path:     "bitbucket.org/kamiazya/tcho/cli/module/local",
									packages: []string{},
									subGroups: []*Group{
										{
											name:     "component",
											path:     "bitbucket.org/kamiazya/tcho/cli/module/local/component",
											packages: []string{},
										},
									},
								},
							},
						},
					},
				},
				{
					name:     "infrastructure",
					path:     "bitbucket.org/kamiazya/tcho/core/infrastructure",
					packages: []string{},
					subGroups: []*Group{
						{
							name:      "database",
							path:      "bitbucket.org/kamiazya/tcho/core/infrastructure/database",
							packages:  []string{},
							subGroups: []*Group{},
							line:      2,
						},
					},
				},
				{
					name:     "core",
					path:     "bitbucket.org/kamiazya/tcho/core",
					packages: []string{},
					line:     1,
					subGroups: []*Group{
						{
							name:     "domain",
							path:     "bitbucket.org/kamiazya/tcho/core/domain",
							packages: []string{},
							subGroups: []*Group{
								{
									name:     "model",
									path:     "bitbucket.org/kamiazya/tcho/core/domain/model",
									packages: []string{},
									line:     1,
								},
								{
									name:     "repository",
									path:     "bitbucket.org/kamiazya/tcho/core/domain/repository",
									packages: []string{},
									line:     2,
								},
								{
									name:     "value",
									path:     "bitbucket.org/kamiazya/tcho/core/domain/value",
									packages: []string{},
									line:     1,
								},
							},
						},
						{
							name:     "application",
							path:     "bitbucket.org/kamiazya/tcho/core/application",
							packages: []string{},
							subGroups: []*Group{
								{
									name:     "usecase",
									path:     "bitbucket.org/kamiazya/tcho/core/application/usecase",
									packages: []string{},
								},
							},
						},
					},
				},
			},
		},
		{
			name:     "",
			path:     "",
			packages: []string{},
		},
	}
)

func main() {

	pkgRoot := "bitbucket.org/kamiazya/tcho/cli"
	pkg := scanPackage(pkgRoot)
	pkg.group()

	fmt.Println("digraph godep {")

	for _, g := range groups {
		g.print()
	}
	pkg.print()
	fmt.Println("}")

}

type Group struct {
	name      string
	path      string
	packages  []string
	subGroups []*Group
	line      int
}

func (g *Group) match(p *Package) bool {
	return strings.Contains(p.path, g.path)
}

func (g *Group) add(p *Package) {
	if g.subGroups != nil {
		for i := range g.subGroups {
			if g.subGroups[i].match(p) {
				g.subGroups[i].add(p)
				return
			}
		}
	}
	for _, pkg := range g.packages {
		if pkg == p.path {
			return
		}
	}
	g.packages = append(g.packages, p.path)
}

func (g *Group) print() {
	fmt.Printf("subgraph cluster_%s {\nlabel = \"%s\";\n", g.name, g.name)

	for _, path := range g.packages {
		pkg := packageCtx[path]
		if !pkg.goroot {
			// if g.name == "" {
			// 	fmt.Printf("_%d [label=\"%s\" style=\"filled\"];\n", pkg.id, pkg.path)
			// } else

			if pkg.name != "main" {
				fmt.Printf("_%d [label=\"%s\" style=\"filled\"];\n", pkg.id, pkg.name)
			}
		}
	}

	for _, sub := range g.subGroups {
		sub.print()
	}

	// if g.path != "" {
	fmt.Println("}")
	// }
}

type Package struct {
	id          int
	name        string
	path        string
	subPackages []*Package
	printed     bool
	goroot      bool
}

func (p *Package) print() {
	if !p.printed {
		for _, sub := range p.subPackages {
			if !p.goroot && !sub.goroot {
				if p.name != "main" {
					fmt.Printf("_%d -> _%d;\n", p.id, sub.id)
				}
				sub.print()
			}
		}
		p.printed = true
	}
}

func (p *Package) group() {
	for _, g := range groups {
		if g.match(p) {
			g.add(p)
		}
	}
	for _, sub := range p.subPackages {
		sub.group()
	}
}

func scanPackage(pkgRoot string) (p *Package) {
	p = new(Package)
	p.path = pkgRoot
	if pkg, exist := packageCtx[pkgRoot]; exist {
		return pkg
	}
	idCounter++
	p.id = idCounter
	packageCtx[pkgRoot] = p
	pkg, err := buildContext.Import(pkgRoot, "", 0)
	if err != nil || pkg.Goroot {
		p.goroot = pkg.Goroot
		return nil
	}

	p.name = pkg.Name
	p.subPackages = make([]*Package, 0, len(pkg.Imports))
	for _, subPkgPath := range pkg.Imports {
		subPkg := scanPackage(subPkgPath)
		if subPkg != nil {
			p.subPackages = append(p.subPackages, subPkg)
		}

	}
	return
}
