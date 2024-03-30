import SearchInputComponent from "../components/SearchInput.component";

import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [searchText, setSearchText] = useState("");

  const handleSearch = async () => {
    window.location.href = `/search?q=${searchText}`;
  };

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === "Enter") {
        handleSearch();
      }
    };

    document.addEventListener("keydown", handleKeyDown);

    return () => {
      document.removeEventListener("keydown", handleKeyDown);
    };
  }, [searchText]);

  return (
    <div>
      <h1>A simple image search engine</h1>
      <SearchInputComponent
        searchText={searchText}
        handleChange={setSearchText}
      />
    </div>
  );
}

export default App;
