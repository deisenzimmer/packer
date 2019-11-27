package hcl2template

import "github.com/hashicorp/packer/packer"

// PackerConfig represents a loaded packer config
type PackerConfig struct {
	Sources map[SourceRef]*Source

	Variables PackerV1Variables

	Builds Builds

	Communicators map[CommunicatorRef]*Communicator
}

func (b Build) CoreBuildProvisioners() []packer.CoreBuildProvisioner {
	res := []packer.CoreBuildProvisioner{}
	for _, pg := range b.ProvisionerGroups {
		for _, p := range pg.Provisioners {
			res = append(res, packer.CoreBuildProvisioner{
				PType:       p.PType,
				Provisioner: p.Provisioner,
			})
		}
	}
	return res
}

func (b Build) CoreBuildPostProcessors() [][]packer.CoreBuildPostProcessor {
	res := [][]packer.CoreBuildPostProcessor{}
	for _, ppg := range b.PostProvisionerGroups {
		grp := []packer.CoreBuildPostProcessor{}
		for _, pp := range ppg.PostProcessors {
			grp = append(grp, packer.CoreBuildPostProcessor{
				PostProcessor: pp.PostProcessor,
				PType:         pp.PType,
			})
		}
		res = append(res, grp)
	}
	return res
}

func (p *PackerConfig) GetBuilds() []packer.Build {
	res := []packer.Build{}

	for _, build := range p.Builds {
		for _, from := range build.Froms {
			src := p.Sources[from]
			pcb := &packer.CoreBuild{
				Type:           src.Type,
				Builder:        src.Builder,
				BuilderConfig:  nil, // TODO(azr): do we really need that ?
				Hooks:          nil, // TODO(azr): what are hooks and do we really need those ?
				Provisioners:   build.CoreBuildProvisioners(),
				PostProcessors: build.CoreBuildPostProcessors(),
				TemplatePath:   "", // TODO(azr): do we really need that ?
				Variables:      p.Variables,
			}
			res = append(res, pcb)
		}
	}
	return res
}
