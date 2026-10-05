package main

import (
	"database/sql"
	"embed"
	"fmt"
	"jakegodsall/knowledge-graph/src/domain"
	"jakegodsall/knowledge-graph/src/repository"
	"jakegodsall/knowledge-graph/src/repository/sqlite"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DBPath string
}

func defaultConfig() Config {
	home, _ := os.UserHomeDir()
	return Config{
		DBPath: filepath.Join(home, ".ktree", "ktree.db"),
	}
}


//go:embed migrations
var migrations embed.FS

func main() {
	config := defaultConfig()

	if err := os.MkdirAll(filepath.Dir(config.DBPath), 0755); err != nil {
		fmt.Printf("could not create data directory: %v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("sqlite3", config.DBPath)
	if err != nil {
		fmt.Printf("could not open database: %v\n", err)
		os.Exit(1)
	}

	err = runMigrations(db)
	if err != nil {
		fmt.Printf("error running migrations: %v\n", err)
		os.Exit(1)
	}

	ktreeRepository := sqlite.NewKnowledgeTreeRepository(db)

	if len(os.Args) < 2 {
		fmt.Println("temp")
		os.Exit(1)
	}

	switch(os.Args[1]) {
	case "list":
		err := runList(ktreeRepository)
		if err != nil {
			fmt.Println("could not list trees")
			os.Exit(1)
		}
	case "show":
		runShow(os.Args[2:])
	case "create":
		err := runCreate(os.Args[2:], ktreeRepository)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func runList(repo repository.KnowledgeTreeRepository) error {
	trees, err := repo.GetAll()

	if err != nil {
		return err
	}

	for _, tree := range trees {
		fmt.Println(tree)
	}
	return nil
}

func runShow(args []string) {
	fmt.Println("run show")
}

func runCreate(args []string, repo repository.KnowledgeTreeRepository) error {
	if len(args) == 0 {
		return fmt.Errorf("No tree provided")
	}

	err := repo.Create(domain.NewKnowledgeTree(args[0]))

	if err != nil {
		return err
	}

	return nil
}

func runMigrations(db *sql.DB) error {
	dirs, err := migrations.ReadDir("migrations")

	if err != nil {
		return err
	}

	for _, entry := range dirs {
		if !strings.Contains(entry.Name(), ".up.sql") {
			continue
		}

		fileName := "migrations/" + entry.Name()
		content, err := migrations.ReadFile(fileName)

		if err != nil {
			return fmt.Errorf("could not read migration file %s: %w", fileName, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("could not execute migration %s: %w", fileName, err)
		}

	}
	return nil
}

func rollbackMigrations(db *sql.DB) error {
	dirs, err := migrations.ReadDir("migrations")

	if err != nil {
		return err
	}

	for i, j := 0, len(dirs)-1; i < j; i, j = i+1, j-1 {
		dirs[i], dirs[j] = dirs[j], dirs[i]
	}

	for _, entry := range dirs {
		if !strings.Contains(entry.Name(), ".down.sql") {
			continue
		}

		fileName := "migrations/" + entry.Name()
		content, err := migrations.ReadFile(fileName)

		if err != nil {
			return fmt.Errorf("could not read migration file %s: %w", fileName, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("could not execute migration %s: %w", fileName, err)
		}
	}

	return nil
}
