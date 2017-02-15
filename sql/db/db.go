package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
)

var (
	mu              sync.RWMutex
	instances       = map[string]*Conn{}
	nullEvtReceiver = &dbr.NullEventReceiver{}

	_ Executor = &Conn{}
	_ Executor = &Tx{}
)

type Executor interface {
	ExecStmt(stmt dbr.Builder, args ...interface{}) (sql.Result, error)
	QueryStmt(stmt dbr.Builder, args ...interface{}) (*sql.Rows, error)
	QueryRowStmt(stms dbr.Builder, args ...interface{}) *sql.Row
	SQL() *Builder
}

func New(engine, source string, opts ...Option) (*Conn, error) {
	var d dbr.Dialect
	switch engine {
	case MYSQL:
		d = dialect.MySQL

	case POSTGRESQL:
		d = dialect.PostgreSQL

	case SQLITE3:
		d = dialect.SQLite3

	default:
		return nil, fmt.Errorf("unsupport engine %q", engine)
	}

	db, err := sql.Open(engine, source)
	if err != nil {
		return nil, err
	}

	conn := &Conn{
		DB: db,
		builder: &Builder{
			d: d,
		},
	}

	for _, one := range opts {
		if one != nil {
			one(conn)
		}
	}

	return conn, nil
}

func Open(engine, source string, opts ...Option) error {
	if engine == "" {
		engine = defaultEngine
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := instances[engine]; ok {
		return fmt.Errorf("duplicate db instance %q", engine)
	}

	conn, err := New(engine, source, opts...)
	if err != nil {
		return err
	}

	instances[engine] = conn

	return nil
}

func Get(engine string) (*Conn, error) {
	if engine == "" {
		engine = defaultEngine
	}

	mu.RLock()
	instace := instances[engine]
	mu.RUnlock()

	if instace == nil {
		return nil, fmt.Errorf("db instance for %q not found", engine)
	}

	return instace, nil
}

func MustGet(engine string) *Conn {
	instance, err := Get(engine)
	if err != nil {
		panic(err)
	}

	return instance
}

type Conn struct {
	*sql.DB
	builder *Builder
}

func (this *Conn) ExecStmt(stmt dbr.Builder, args ...interface{}) (sql.Result, error) {
	return this.DB.Exec(this.builder.MustBuild(stmt), args...)
}

func (this *Conn) QueryStmt(stmt dbr.Builder, args ...interface{}) (*sql.Rows, error) {
	return this.DB.Query(this.builder.MustBuild(stmt), args...)
}

func (this *Conn) QueryRowStmt(stmt dbr.Builder, args ...interface{}) *sql.Row {
	return this.DB.QueryRow(this.builder.MustBuild(stmt), args...)
}

func (this *Conn) SQL() *Builder {
	return this.builder
}

func (this *Conn) BeginTx() (*Tx, error) {
	tx, err := this.DB.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{
		Tx:      tx,
		builder: this.builder,
	}, nil
}

type Tx struct {
	*sql.Tx
	builder *Builder
}

func (this *Tx) ExecStmt(stmt dbr.Builder, args ...interface{}) (sql.Result, error) {
	return this.Tx.Exec(this.builder.MustBuild(stmt), args...)
}

func (this *Tx) QueryStmt(stmt dbr.Builder, args ...interface{}) (*sql.Rows, error) {
	return this.Tx.Query(this.builder.MustBuild(stmt), args...)
}

func (this *Tx) QueryRowStmt(stmt dbr.Builder, args ...interface{}) *sql.Row {
	return this.Tx.QueryRow(this.builder.MustBuild(stmt), args...)
}

func (this *Tx) SQL() *Builder {
	return this.builder
}
