package main_test

import (
	"log"
	"os"
	"testing"

	"."
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}
	// a.Initialize(
	// 	os.Getenv("TEST_DB_USERNAME"),
	// 	os.Getenv("TEST_DB_PASSWORD"),
	// 	os.Getenv("TEST_DB_NAME"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
id SERIAL,
name TEXT NOT NULL,
CONSTRAINT products_pkey PRIMARY KEY (id)
)`