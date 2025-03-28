package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

type SavedGamesPath struct {
	Dir  string `json:"dir"`
	Name string `json:"name"`
}

type MoneyFile struct {
	OnlineBalance float64 `json:"OnlineBalance"`
}

var LOCAL_LOW_PATH = path.Join(os.Getenv("USERPROFILE"), "AppData", "LocalLow")
var SCHEDULE_BASE_PATH = path.Join(LOCAL_LOW_PATH, "TVGS", "Schedule I")
var SAVED_PATH = path.Join(SCHEDULE_BASE_PATH, "Saves")

func main() {
	savedGames, err := GetSavedGames()

	if err != nil {
		log.Panic(err)
	}

	if len(savedGames) == 0 {
		log.Panic("No saved games found")
	}

	selectedGame := savedGames[0]

	moneyFile := path.Join(selectedGame.Dir, "money.json")

	var moneyData MoneyFile
	err = formatMoneyFilePath(filepath.Clean(moneyFile), &moneyData)

	if err != nil {
		log.Panic(err)
	}

	if err := replaceMoneyValue(selectedGame.Dir, 5); err != nil {
		log.Panic(err)
	}

	fmt.Println(moneyData.OnlineBalance)
}

func replaceMoneyValue(dir string, newValue float64) error {
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
func formatMoneyFilePath(dir string, structPtr any) error {

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
