package resource

import "embed"

// The following following directive is necessary for compile.
//
//go:embed *
var FS embed.FS
