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

	if err := outputGoFile(dir, sqlConnFile, "sqlconn", struct {
		Package string
	}{
		Package: pkgName,
	}); err != nil {
		return err
	}

	tables := map[string][]Table{}

	for _, data := range objects {
		if data.Engine == "" {
			data.Engine = opt.DefaultEngine
		}

		structFile := fmt.Sprintf("gen_%s_struct.go", data.Table)
		if err := outputGoFile(dir, structFile, "struct", data); err != nil {
			return err
		}

		sqlMgrFile := fmt.Sprintf("gen_%s_sql_mgr.go", data.Table)
		if err := outputGoFile(dir, sqlMgrFile, "sqlmgr", data); err != nil {
			return err
		}

		tables[data.Engine] = append(tables[data.Engine], Table{Object: data, opt: opt})
	}

	for engine, ts := range tables {
		sqlScriptTableFile := fmt.Sprintf("gen_script_table_%s.sql", engine)

		if err := outputNormalFile(dir, sqlScriptTableFile, "sql_script_table", ts); err != nil {
			return err
		}
	}

	return nil
}
