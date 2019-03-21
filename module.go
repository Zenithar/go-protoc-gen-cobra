package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type cobraGen struct {
	pgs.ModuleBase
	pgsgo.Context
}

func (*cobraGen) Name() string {
	return "cobra"
}

func (m *cobraGen) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.Context = pgsgo.InitContext(c.Parameters())
}

func (m *cobraGen) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, file := range targets {
		name := m.Context.OutputPath(file).SetExt(".cobra.go").String()

		fm := fileModel{
			Package: m.Context.PackageName(file).String(),
		}

		imbs := map[string]importModel{}

		for _, service := range file.Services() {
			sm := serviceModel{
				Name:    service.Name().String(),
				FullName: service.Name().UpperCamelCase().String(),
				UseName: service.Name().LowerCamelCase().String(),
			}

			for _, method := range service.Methods() {
				mb := methodModel{
					Name:   m.Context.Name(method).String(),
					UseName: m.Context.Name(method).LowerCamelCase().String(),
					Input:  "*" + m.Context.Name(method.Input()).String(),
					Output: "*" + m.Context.Name(method.Output()).String(),
				}

				if !method.Input().BuildTarget() {
					path := m.Context.ImportPath(method.Input()).String()
					imbs[path] = importModel{
						Value: path,
					}

					mb.Input = "*" + m.Context.PackageName(method.Input()).String() + "." + m.Context.Name(method.Input()).String()
				}

				if !method.Output().BuildTarget() {
					path := m.Context.ImportPath(method.Output()).String()
					imbs[path] = importModel{
						Value: path,
					}

					mb.Output = "*" + m.Context.PackageName(method.Output()).String() + "." + m.Context.Name(method.Output()).String()
				}

				sm.Methods = append(sm.Methods, mb)
			}

			fm.Services = append(fm.Services, sm)
		}

		if len(fm.Services) == 0 {
			continue
		}

		for _, imb := range imbs {
			fm.Imports = append(fm.Imports, imb)
		}

		m.OverwriteGeneratorTemplateFile(
			name,
			T.Lookup("File"),
			&fm,
		)
	}

	return m.Artifacts()
}

// -----------------------------------------------------------------------------

type fileModel struct {
	Services []serviceModel
	Package  string
	Imports  []importModel
}

type serviceModel struct {
	Name    string
	FullName string
	UseName string
	Methods []methodModel
}

type methodModel struct {
	Name   string
	UseName string
	Input  string
	Output string
}

type importModel struct {
	Key   string
	Value string
}
