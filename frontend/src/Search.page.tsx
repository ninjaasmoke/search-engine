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

  const [searchResults, setSearchResults] = useState([] as SearchData[]);

  const fetchSearchResults = async () => {
    const response = await fetch(
      `${window.location.origin}/api/search?q=${searchText}`
      // `http://localhost:8080/api/search?q=${searchText}`
    );
    if (!response.ok) {
      console.error("Failed to fetch search results");
      return;
    }

    const data = await response.json() as SearchData[];

    setSearchResults(data);
  };

  useEffect(() => {
    fetchSearchResults();
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
    <div>
      <h1>Search</h1>
      <SearchInputComponent
        searchText={searchText}
        handleChange={(text) => {
          setSearchText(text);
        }}
      />
      {
        searchResults === null || searchResults.length === 0 ? (
          <div style={{
            marginTop: "1rem",
            color: "#666",
          }}>No results found</div>
        ) : (
          <SearchResults searchResults={searchResults} />
        )
      }
    </div>
  );
}

export default Search;
