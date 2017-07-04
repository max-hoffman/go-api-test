package main

import (
	"github.com/jmoiron/sqlx"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p *user) getUser(db *sqlx.DB) error {
	return db.Get(&p, "SELECT name FROM users WHERE id=$1", p.ID)
}

func (p *user) updateUser(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE users SET name=$1 WHERE id=$2", p.Name, p.ID)
	return err
}

func (p *user) deleteUser(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", p.ID)
	return err
}

func (p *user) createUser(db *sqlx.DB) error {
	return db.QueryRowx("INSERT INTO users(name) VALUES($1) RETURNING id", p.Name).StructScan(&p)
}

func getUsers(db *sqlx.DB, start, count int) ([]user, error) {
	pp := []user{}
	err := db.Select(&pp, "SELECT id, name, price FROM products LIMIT $1 OFFSET $2", count, start)
	return pp, err
}
