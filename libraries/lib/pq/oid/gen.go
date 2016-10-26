// +build ignore

// Generate the table of OID values
// Run with 'go run gen.go'.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	_ "github.com/insionng/zenpress/libraries/lib/pq"
)

func main() {
	datname := os.Getenv("PGDATABASE")
	sslmode := os.Getenv("PGSSLMODE")

	if datname == "" {
		os.Setenv("PGDATABASE", "pqgotest")
	}

	if sslmode == "" {
		os.Setenv("PGSSLMODE", "disable")
	}

	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("gofmt")
	cmd.Stderr = os.Stderr
	w, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("types.go")
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stdout = f
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "// generated by 'go run gen.go'; do not edit")
	fmt.Fprintln(w, "\npackage oid")
	fmt.Fprintln(w, "const (")
	rows, err := db.Query(`
		SELECT typname, oid
		FROM pg_type WHERE oid < 10000
		ORDER BY oid;
	`)
	if err != nil {
		log.Fatal(err)
	}
	var name string
	var oid int
	for rows.Next() {
		err = rows.Scan(&name, &oid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "T_%s Oid = %d\n", name, oid)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, ")")
	w.Close()
	cmd.Wait()
}
