// Package config provides a struct to store the applications config
package config

import (
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"
)

// AppConfig contains the application configuration structure.
var AppConfig struct {
	CRDB        crdbx.Config
	Logging     loggingx.Config
	Events      events.Config
	OIDC        echojwtx.AuthConfig
	Permissions permissions.Config
	Server      echox.Config
	Tracing     otelx.Config
}
