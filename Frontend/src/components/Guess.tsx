import { useEffect, useRef } from "react";

export default function Guess({
    value,
    isCurrent,
    onchange,
    style,
    className,
}: {
    value: string;
    isCurrent: boolean;
    onchange: (event: React.ChangeEvent<HTMLInputElement>) => void;
    style?: React.CSSProperties;
    className?: string;
}) {
    const inputRef = useRef<HTMLInputElement>(null);

    useEffect(() => {
        if (isCurrent) {
            inputRef.current?.select();
        }
    }, [isCurrent]);

    return (
        <input
            value={value}
            maxLength={1}
            style={style}
            className={["guess", className].join(" ")}
            disabled={!isCurrent}
            onChange={onchange}
            ref={inputRef}
        />
    );
}
