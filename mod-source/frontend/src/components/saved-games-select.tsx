import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { SavedGamesPath } from "@/types/game";

export function SavedGamesSelect({
  savedGamesPaths,
}: {
  savedGamesPaths: SavedGamesPath[];
}) {
  return (
    <Select>
      <SelectTrigger className="w-full">
        <SelectValue placeholder="Select a saved game" />
      </SelectTrigger>
      <SelectContent>
        {savedGamesPaths.map((savedGamesPath) => (
          <SelectItem key={savedGamesPath.dir} value={savedGamesPath.dir}>
            {savedGamesPath.name}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
}
