import SearchInputComponent from "../components/SearchInput.component";
import QuickSearches from "../components/QuickSearches";

import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import "./App.css";

function App() {
  const navigate = useNavigate();
  const [searchText, setSearchText] = useState("");

  const handleSearch = async () => {
    if (!searchText || searchText.length < 1) return;
    navigate(`/search?q=${searchText}`);
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
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <div>
        <h1>an image search engine.</h1>
        <div className="searchArea">
          <SearchInputComponent
            searchText={searchText}
            handleChange={(text) => {
              setSearchText(text);
            }}
          />
        </div>
      </div>

      <QuickSearches />

      <Link
        to={`/about`}
        style={{
          padding: "10px",
          marginBottom: "20px",
          fontSize: "0.8rem",
          textDecoration: "underline",
        }}
      >
        read about this
      </Link>
    </div>
  );
}

export default App;
