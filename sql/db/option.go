package db

import (
	"time"
)

type Option func(*Conn)

func ConnMaxLifeTime(d time.Duration) Option {
	return func(conn *Conn) {
		conn.SetConnMaxLifetime(d)
	}
}

func MaxIdleConns(n int) Option {
	return func(conn *Conn) {
		conn.SetMaxIdleConns(n)
	}
}

func MaxOpenConns(n int) Option {
	return func(conn *Conn) {
		conn.SetMaxOpenConns(n)
	}
}
