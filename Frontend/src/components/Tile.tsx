import { useCallback, useEffect, useState } from "react";
import Guess from "./Guess";

const getGuessStatus = (word: string, currentLetter: string, index: number) => {
    if (word[index] == currentLetter) return "success";

    if (word.includes(currentLetter)) return "missplaced";

    return "wrong";
};

export default function Tile({
    word,
    reveal,
    isCurrentTile,
    onChange,
}: {
    word: string;
    reveal: boolean;
    isCurrentTile: boolean;
    onChange?: (newGuess: string) => void;
}) {
    const [fullGuess, setFullGuess] = useState("");
    const [currentLetterIndex, setCurrentLetterIndex] = useState(0);

    const handleChange = useCallback(
        (newFullGuess: string) => {
            setFullGuess(newFullGuess);
            onChange?.(newFullGuess);
        },
        [onChange],
    );

    useEffect(() => {
        const onKeyPress = (event: KeyboardEvent) => {
            switch (event.key) {
                case "Backspace":
                    setCurrentLetterIndex(
                        currentLetterIndex > 0 ? currentLetterIndex - 1 : 0,
                    );
                    handleChange(fullGuess.slice(0, fullGuess.length - 1) ?? "");
                    break;
                default:
                    break;
            }
        };

        if (isCurrentTile) addEventListener("keydown", onKeyPress);

        return () => {
            removeEventListener("keydown", onKeyPress);
        };
    }, [fullGuess, currentLetterIndex, isCurrentTile, handleChange]);

    const letters = word.split("");
    return (
        <div className="tile">
            {letters.map((_, index) => (
                <Guess
                    value={fullGuess[index] ?? ""}
                    key={index}
                    isCurrent={isCurrentTile && currentLetterIndex == index}
                    onchange={(event) => {
                        handleChange(fullGuess + event.target.value.toLowerCase());
                        setCurrentLetterIndex(currentLetterIndex + 1);
                    }}
                    style={{ animationDelay: `${index}s` }}
                    className={
                        reveal ? getGuessStatus(word, fullGuess[index], index) : ""
                    }
                ></Guess>
            ))}
        </div>
    );
}
