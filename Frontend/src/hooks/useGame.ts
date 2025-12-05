import { useEffect, useState } from "react";

type GameStatus = "Playing" | "Won" | "Lost";

export default function useGame(
    dailyword: string,
    maxTries: number,
): [string[], React.Dispatch<React.SetStateAction<string[]>>, number, GameStatus] {
    const [attempts, setAttempts] = useState(Array<string>(maxTries).fill(""));
    const [currentTry, setCurrentTry] = useState(0);

    console.log(maxTries);

    const gameStatus: GameStatus =
        currentTry > 0 && attempts[currentTry - 1] === dailyword
            ? "Won"
            : currentTry >= maxTries
              ? "Lost"
              : "Playing";

    useEffect(() => {
        const onKeyPress = (event: KeyboardEvent) => {
            switch (event.key) {
                case "Enter":
                    setCurrentTry(currentTry + 1);
                    break;
                default:
                    break;
            }
        };

        if (gameStatus == "Playing" && attempts[currentTry].length === dailyword.length)
            addEventListener("keydown", onKeyPress);

        return () => {
            removeEventListener("keydown", onKeyPress);
        };
    }, [attempts, currentTry, dailyword.length, gameStatus]);

    return [attempts, setAttempts, currentTry, gameStatus];
}
