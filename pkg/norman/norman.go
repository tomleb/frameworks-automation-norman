// Package norman is a stand-in module for the automation playground.
// Imports wrangler so go mod tidy keeps the require entry.
package norman

import "github.com/tomleb/frameworks-automation-wrangler/pkg/wrangler"

const Version = "0.8"

func Hello() string { return "hello from norman " + Version + " using " + wrangler.Hello() }
