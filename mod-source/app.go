package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

var LOCAL_LOW_PATH = path.Join(os.Getenv("USERPROFILE"), "AppData", "LocalLow")
var SCHEDULE_BASE_PATH = path.Join(LOCAL_LOW_PATH, "TVGS", "Schedule I")
var SAVED_PATH = path.Join(SCHEDULE_BASE_PATH, "Saves")

type SavedGamesPath struct {
	Dir  string `json:"dir"`
	Name string `json:"name"`
}

type MoneyFile struct {
	OnlineBalance float64 `json:"OnlineBalance"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetMoneyValue(dir string) (float64, error) {
	moneyFile := filepath.Join(dir, "money.json")

	file, err := os.ReadFile(moneyFile)
	if err != nil {
		return 0, err
	}

	var moneyData MoneyFile
	err = json.Unmarshal(file, &moneyData)
	if err != nil {
		return 0, err
	}

	return moneyData.OnlineBalance, nil
}

func (a *App) ReplaceMoneyValue(dir string, newValue float64) error {
	moneyFile := filepath.Join(dir, "money.json")

	file, err := os.ReadFile(moneyFile)
	if err != nil {
		return err
	}

	var moneyData map[string]interface{}
	err = json.Unmarshal(file, &moneyData)
	if err != nil {
		return err
	}

	if _, exists := moneyData["OnlineBalance"]; exists {
		moneyData["OnlineBalance"] = newValue
	} else {
		return fmt.Errorf("key 'OnlineBalance' not found in JSON file")
	}

	newFile, err := json.MarshalIndent(moneyData, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(moneyFile, newFile, 0644)
	if err != nil {
		return err
	}

	return nil
}
func FormatMoneyFilePath(dir string, structPtr any) error {

	file, err := os.ReadFile(dir)
	if err != nil {
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, structPtr)
	if err != nil {
		return err
	}

	return nil
}
func GetSavedGames() ([]SavedGamesPath, error) {

	var savedGames []SavedGamesPath

	DirEntry, err := os.ReadDir(SAVED_PATH)

	if err != nil {
		return nil, err
	}

	for _, entry := range DirEntry {
		if entry.IsDir() {
			gamesEntryPath := path.Join(SAVED_PATH, entry.Name())

			gamesEntry, err := os.ReadDir(gamesEntryPath)
			if err != nil {
				log.Panic(err)
			}

			for _, gameEntry := range gamesEntry {
				if gameEntry.IsDir() {
					gameEntryPath := path.Join(gamesEntryPath, gameEntry.Name())

					savedGames = append(savedGames, SavedGamesPath{
						Dir:  gameEntryPath,
						Name: gameEntry.Name(),
					})
				}
			}
		}
	}

	return savedGames, nil
}
