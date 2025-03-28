import { useEffect, useState } from "react";
import { GetSavedGames } from "../wailsjs/go/main/App";
import "./App.css";
import { RootLayout } from "@/layouts/root-layout";
import { type SavedGamesPath } from "@/types/game";
import { SavedGamesSelect } from "./components/saved-games-select";

function App() {
  const [savedGamesPaths, setSavedGamesPaths] = useState<SavedGamesPath[]>();

  useEffect(() => {
    const fetchSavedGamesPath = async () => {
      try {
        const savedGamesPath = await GetSavedGames();
        console.log("savedGamesPath", savedGamesPath);
        setSavedGamesPaths(savedGamesPath);
      } catch (error) {
        console.error("Error fetching saved games path:", error);
      }
    };

    fetchSavedGamesPath();
  }, []);

  console.log("savedGamesPaths", savedGamesPaths);

  return (
    <RootLayout>
      <SavedGamesSelect savedGamesPaths={savedGamesPaths || []} />
    </RootLayout>
  );
}

export default App;
