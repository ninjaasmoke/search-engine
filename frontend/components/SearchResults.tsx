import { useEffect } from "react";
import { SearchData } from "../types";

import SearchImage from "./SearchImage";

import "./SearchResults.css";
import { useNavigate } from "react-router-dom";

type SearchResultsProps = {
  searchResults: SearchData[];
  correctedQuery?: string;
  originalQuery?: string;
};

function SearchResults({
  searchResults,
  correctedQuery,
  originalQuery,
}: SearchResultsProps) {
  if (searchResults.length === 0) {
    return <div>No results found</div>;
  }

  // split the search results into 3 even columns
  const column1Lg: SearchData[] = [];
  const column2Lg: SearchData[] = [];
  const column3Lg: SearchData[] = [];

  searchResults.forEach((result, index) => {
    if (index % 3 === 0) {
      column1Lg.push(result);
    } else if (index % 3 === 1) {
      column2Lg.push(result);
    } else {
      column3Lg.push(result);
    }
  });

  // split the search results into 2 even columns
  const column1Md: SearchData[] = [];
  const column2Md: SearchData[] = [];

  searchResults.forEach((result, index) => {
    if (index % 2 === 0) {
      column1Md.push(result);
    } else {
      column2Md.push(result);
    }
  });

  useEffect(() => {
    const scrollHere = document.getElementById("scrollHere");
    if (scrollHere) {
      // scrollHere.scrollIntoView({ behavior: "smooth" });
    }
  }, [searchResults, correctedQuery]);

  const navigate = useNavigate();

  const forceReloadLink = (url: string) => {
    navigate(url, { replace: true });
    window.location.reload();
  };

  return (
    <>
      <div id="scrollHere" />
      {correctedQuery && (
        <div>
          showing results for:{" "}
          <span
            onClick={() => forceReloadLink(`/search?q=${correctedQuery}`)}
            style={{
              cursor: "pointer",
              color: "blue",
            }}
            onMouseOver={(e) => {
              e.currentTarget.style.textDecoration = "underline";
            }}
            onMouseOut={(e) => {
              e.currentTarget.style.textDecoration = "none";
            }}
          >
            {correctedQuery}
          </span>
          <br />
          show original results:{" "}
          <span
            onClick={() =>
              forceReloadLink(`/search?q=${originalQuery}&skipSpellCheck=true`)
            }
            style={{
              cursor: "pointer",
              color: "blue",
            }}
            onMouseOver={(e) => {
              e.currentTarget.style.textDecoration = "underline";
            }}
            onMouseOut={(e) => {
              e.currentTarget.style.textDecoration = "none";
            }}
          >
            {originalQuery}
          </span>
        </div>
      )}
      <div className="gridContainerLg">
        <div className="gridItem">
          {column1Lg.map((result, index) => (
            <SearchImage key={`${result.id}-${index}`} imageData={result} />
          ))}
        </div>
        <div className="gridItem">
          {column2Lg.map((result, index) => (
            <SearchImage key={`${result.id}-${index}`} imageData={result} />
          ))}
        </div>
        <div className="gridItem">
          {column3Lg.map((result, index) => (
            <SearchImage key={`${result.id}-${index}`} imageData={result} />
          ))}
        </div>
      </div>
      <div className="gridContainerMd">
        <div className="gridItem">
          {column1Md.map((result, index) => (
            <SearchImage key={`${result.id}-${index}`} imageData={result} />
          ))}
        </div>
        <div className="gridItem">
          {column2Md.map((result, index) => (
            <SearchImage key={`${result.id}-${index}`} imageData={result} />
          ))}
        </div>
      </div>

      <div className="gridContainerSm">
        {searchResults.map((result, index) => (
          <SearchImage key={`${result.id}-${index}`} imageData={result} />
        ))}
      </div>
    </>
  );
}

export default SearchResults;
