package generator

import (
	"path/filepath"
)

type Option struct {
	Input       string
	Package     string
	OutputDir   string
	ExportTable bool
	ExportIndex bool
}

func (this *Option) PackageName() string {
	if this.Package != "" {
		return this.Package
	}

	return filepath.Base(this.Input)
}
