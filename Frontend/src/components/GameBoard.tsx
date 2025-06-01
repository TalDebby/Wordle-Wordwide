import useGame from "../hooks/useGame";
import Tile from "./Tile";

const MAX_TRIES = 5;

export default function GameBoard({ dailyword }: { dailyword: string }) {
    const [attempts, setAttempts, currentTry, gameStatus] = useGame(dailyword, MAX_TRIES);

    return (
        <div className="main-board">
            {attempts.map((_, index) => (
                <Tile
                    key={index}
                    word={dailyword}
                    reveal={index < currentTry}
                    isCurrentTile={currentTry == index && gameStatus == "Playing"}
                    onChange={
                        currentTry == index
                            ? (guess) => {
                                  const newTries = [...attempts];
                                  newTries[index] = guess;
                                  setAttempts(newTries);
                              }
                            : undefined
                    }
                ></Tile>
            ))}
        </div>
    );
}
