package main

import (
	"database/sql"
	"fmt"
	"jakegodsall/knowledge-graph/src/domain"
	"jakegodsall/knowledge-graph/src/repository"
	"jakegodsall/knowledge-graph/src/repository/sqlite"
	"os"
	"path/filepath"

	_ "embed"

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

var ktreeRepository repository.KnowledgeTreeRepository

//go:embed migrations/001_create_knowledge_trees_table.up.sql
var schema string

func main() {
	config := defaultConfig()

	if err := os.MkdirAll(filepath.Dir(config.DBPath), 0755); err != nil {
		fmt.Println("could not create data directory")
		os.Exit(1)
	}

	db, err := sql.Open("sqlite3", config.DBPath)
	if err != nil {
		fmt.Println("could not open database")
		os.Exit(1)
	}

	db.Exec(schema)

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
