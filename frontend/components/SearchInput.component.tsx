import { useState, useEffect } from "react";

type SearchInputProps = {
  searchText: string;
  handleChange: (searchText: string) => void;
};

const placeholders = [
  "cats",
  "cats with sunglasses",
  "dogs",
  "dogs with hats",
  "mountains",
  "mountain sunrise",
  "beaches",
  "beach sunset",
];

function SearchInputComponent({ searchText, handleChange }: SearchInputProps) {
  const [currentPlaceholderIndex, setCurrentPlaceholderIndex] = useState(0);
  const [displayedPlaceholder, setDisplayedPlaceholder] = useState("search...");

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentPlaceholderIndex(
        (prevIndex) => (prevIndex + 1) % placeholders.length
      );
    }, 2000);

    return () => clearInterval(interval);
  }, []);

  useEffect(() => {
    let index = 0;
    const timeout = setTimeout(() => {
      const interval = setInterval(() => {
        setDisplayedPlaceholder(() => {
          if (index === placeholders[currentPlaceholderIndex].length) {
            clearInterval(interval);
            return `search... ${placeholders[currentPlaceholderIndex]}`;
          }
          index++;
          return `search... ${placeholders[currentPlaceholderIndex].slice(
            0,
            index
          )}`;
        });
      }, 250); // Adjust typing speed here

      return () => clearInterval(interval);
    }, 1000); // Delay before typing starts

    return () => clearTimeout(timeout);
  }, [currentPlaceholderIndex]);

  const clearSearchText = () => {
    handleChange("");
  };

  return (
    <div style={{ position: "relative" }}>
      <input
        type="text"
        placeholder={displayedPlaceholder}
        value={searchText}
        onChange={(e) => handleChange(e.target.value)}
      />
      {searchText && (
        <button
          style={{
            position: "absolute",
            right: "8px",
            top: "50%",
            transform: "translateY(-50%)",
            background: "transparent",
            border: "none",
            cursor: "pointer",
          }}
          onClick={clearSearchText}
        >
          clear
        </button>
      )}
    </div>
  );
}

export default SearchInputComponent;
