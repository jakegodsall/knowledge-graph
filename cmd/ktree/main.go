package main

import (
	"database/sql"
	"fmt"
	"jakegodsall/knowledge-graph/src/repository"
	"jakegodsall/knowledge-graph/src/repository/sqlite"
	"os"
	"path/filepath"

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
