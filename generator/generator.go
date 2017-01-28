package generator

import (
	"fmt"
	"path/filepath"

	"github.com/tukdesk/tuk-gen/parser"
)

func Generate(opt Option) error {
	pkgName := opt.PackageName()
	objects, err := parser.Parse(opt.Input, pkgName)
	if err != nil {
		return err
	}

	dir := filepath.Join(opt.OutputDir, pkgName)
	sqlConnFile := "gen_sql_conn.go"

	if err := outputFile(dir, sqlConnFile, "sqlconn", struct {
		Package string
	}{
		Package: pkgName,
	}); err != nil {
		return err
	}

	for _, data := range objects {
		structFile := fmt.Sprintf("gen_%s_struct.go", data.Table)
		if err := outputFile(dir, structFile, "struct", data); err != nil {
			return err
		}

		sqlMgrFile := fmt.Sprintf("gen_%s_sql_mgr.go", data.Table)
		if err := outputFile(dir, sqlMgrFile, "sqlmgr", data); err != nil {
			return err
		}
	}

	return nil
}
