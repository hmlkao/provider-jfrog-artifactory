package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
)

const (
	resourcePrefix = "artifactory"
	modulePath     = "github.com/hmlkao/provider-jfrog-artifactory"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string
