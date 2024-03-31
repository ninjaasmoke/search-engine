import SearchInputComponent from "../components/SearchInput.component";
import SearchResults from "../components/SearchResults";

import { useNavigate, useSearchParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { SearchData } from "../types";

function Search() {
  const navigate = useNavigate();

  const [searchParams] = useSearchParams();
  const query = searchParams.get("q") || "";

  const [searchText, setSearchText] = useState(query);
  const [isLoading, setIsLoading] = useState(true);

  const [searchResults, setSearchResults] = useState([] as SearchData[]);

  const fetchSearchResults = async () => {
    const response = await fetch(
      `${window.location.origin}/api/search?q=${searchText}`
      // `http://localhost:8080/api/search?q=${searchText.trim()}`
    );
    if (!response.ok) {
      console.error("Failed to fetch search results");
      return;
    }

    const data = (await response.json()) as SearchData[];

    setSearchResults(data);
    setIsLoading(false);
  };

  useEffect(() => {
    fetchSearchResults();
    window.scrollTo({ top: 0 });
  }, []);

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === "Enter") {
        navigate(`/search?q=${searchText}`);
        fetchSearchResults();
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
        position: "relative",
      }}
    >
      <h1>Search</h1>
      <div className="searchArea">
        <SearchInputComponent
          searchText={searchText}
          handleChange={(text) => {
            setSearchText(text);
          }}
        />
      </div>
      {isLoading ? (
        <div
          style={{
            marginTop: "1rem",
            color: "#666",
          }}
        >
          Loading...
        </div>
      ) : searchResults === null || searchResults.length === 0 ? (
        <div
          style={{
            marginTop: "1rem",
            color: "#666",
          }}
        >
          No results found
        </div>
      ) : (
        <SearchResults searchResults={searchResults} />
      )}
    </div>
  );
}

export default Search;
