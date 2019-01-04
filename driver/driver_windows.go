package driver

import (
	"github.com/jholder85638/webapp"
	"github.com/jholder85638/webapp/internal/windows"
)

// ForPlatform returns the driver for your platform.
func ForPlatform() webapp.Driver {
	return windows.Driver()
}
