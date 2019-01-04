package driver

import (
	"github.com/jholder85638/webapp"
	"github.com/jholder85638/webapp/internal/macos"
)

// ForPlatform returns the driver for your platform.
func ForPlatform() webapp.Driver {
	return macos.Driver()
}
