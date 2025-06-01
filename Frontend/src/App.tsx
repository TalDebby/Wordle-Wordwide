import { useEffect, useState } from "react";
import GameBoard from "./components/GameBoard";
import "./App.css";

type dailyWord = {
    dailyword: string;
};

function App() {
    const [dailyWord, setDailyWord] = useState("");
    useEffect(() => {
        const getDailyword = async () => {
            const data = await fetch(
                "http://localhost:8080/languages/English/words/dailyword",
            );
            return (await data.json()) as dailyWord;
        };

        getDailyword().then((dailyWord) => {
            setDailyWord(dailyWord.dailyword);
        });
    });

    return dailyWord === "" ? (
        <div>Loading...</div>
    ) : (
        <GameBoard dailyword={dailyWord}></GameBoard>
    );
}

export default App;
