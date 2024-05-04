import SearchInputComponent from "../components/SearchInput.component";
import SearchResults from "../components/SearchResults";

import { useNavigate, useSearchParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { SearchResponse } from "../types";
import { API_URL_DEV, API_URL_PROD } from "../constants";

function Search() {
  const navigate = useNavigate();

  const [searchParams] = useSearchParams();
  const query = searchParams.get("q") || "";
  const skipSpellCheck = searchParams.get("skipSpellCheck") === "true";

  const [searchText, setSearchText] = useState(query);
  const [isLoading, setIsLoading] = useState(true);

  const [searchResults, setSearchResults] = useState({} as SearchResponse);

  const fetchSearchResults = async () => {
    const response = await fetch(
      // import.meta.env.MODE == "production"
      //   ? `${API_URL_PROD}search?q=${searchText.trim()}`
      //   : `${API_URL_DEV}search?q=${searchText.trim()}`
      import.meta.env.MODE == "production"
        ? `${API_URL_PROD}search${skipSpellCheck ? "/noCheck" : ''}?q=${searchText.trim()}`
        : `${API_URL_DEV}search${skipSpellCheck ? "/noCheck" : ''}?q=${searchText.trim()}` 
    );
    if (!response.ok) {
      console.error("Failed to fetch search results");
      return;
    }

    const data = (await response.json()) as SearchResponse;

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
        if (!searchText || searchText.length < 1) return;
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
      ) : searchResults === null || searchResults.Documents == null || searchResults.Documents.length === 0 ? (
        <div
          style={{
            marginTop: "1rem",
            color: "#666",
          }}
        >
          No results found
        </div>
      ) : (
        <SearchResults
          searchResults={searchResults.Documents}
          correctedQuery={searchResults.CorrectedQuery}
          originalQuery={searchResults.Query}
        />
      )}
    </div>
  );
}

export default Search;
