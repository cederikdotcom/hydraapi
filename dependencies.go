package hydraapi

import (
	"path"
	"runtime/debug"
	"strings"
)

// HydraDependencies returns internal hydra module deps from the running binary.
func HydraDependencies() map[string]string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil
	}
	deps := make(map[string]string)
	for _, dep := range info.Deps {
		if strings.HasPrefix(dep.Path, "github.com/cederikdotcom/") {
			deps[path.Base(dep.Path)] = dep.Version
		}
	}
	if len(deps) == 0 {
		return nil
	}
	return deps
}
