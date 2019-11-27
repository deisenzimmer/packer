package main

import (
	"fmt"

	"github.com/hashicorp/packer/packer"
)

type mapOfProvisioner map[string]func() (packer.Provisioner, error)

func (mop mapOfProvisioner) Get(provisioner string) (packer.Provisioner, error) {
	p, found := mop[provisioner]
	if !found {
		return nil, fmt.Errorf("Unknown provisioner %s", provisioner)
	}
	return p()
}

func (mod mapOfProvisioner) List() []string {
	res := []string{}
	for k := range mod {
		res = append(res, k)
	}
	return res
}

type mapOfPostProcessor map[string]func() (packer.PostProcessor, error)

func (mop mapOfPostProcessor) Get(provisioner string) (packer.PostProcessor, error) {
	p, found := mop[provisioner]
	if !found {
		return nil, fmt.Errorf("Unknown post-processor %s", provisioner)
	}
	return p()
}

func (mod mapOfPostProcessor) List() []string {
	res := []string{}
	for k := range mod {
		res = append(res, k)
	}
	return res
}

type mapOfBuilder map[string]func() (packer.Builder, error)

func (mob mapOfBuilder) Get(builder string) (packer.Builder, error) {
	d, found := mob[builder]
	if !found {
		return nil, fmt.Errorf("Unknown builder %s", builder)
	}
	return d()
}

type mapOfCommunicator map[string]func() packer.ConfigurableCommunicator

func (mob mapOfCommunicator) Get(communicator string) (packer.ConfigurableCommunicator, error) {
	c, found := mob[communicator]
	if !found {
		return nil, fmt.Errorf("Unknown communicator %s", communicator)
	}
	return c(), nil
}
