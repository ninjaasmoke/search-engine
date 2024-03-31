import { useEffect } from "react";
import { SearchData } from "../types";

import "./SearchResults.css";

type SearchResultsProps = {
  searchResults: SearchData[];
};

function cleanImageUrl(url: string, width: number) {
  const cleanedUrl = url.replace(/(\?|\&)(q|auto|fit|ixlib|ixid)=[^&]+/g, "");
  const replacedUrl = cleanedUrl.replace(/(\?|\&)w=[^&]+/g, "");
  const modifiedUrl =
    replacedUrl + (replacedUrl.includes("?") ? "&" : "?") + "w=" + width;
  return modifiedUrl;
}

function SearchResults({ searchResults }: SearchResultsProps) {
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
      scrollHere.scrollIntoView({ behavior: "smooth" });
    }
  }, [searchResults]);

  useEffect(() => {}, []);

  return (
    <>
      <div id="scrollHere" />
      <div className="gridContainerLg">
        <div className="gridItem">
          {column1Lg.map((result, index) => (
            <div key={`${result.title}-${index}`}>
              <img
                src={cleanImageUrl(result.url, 400)}
                alt={result.title}
                className="image"
              />
            </div>
          ))}
        </div>
        <div className="gridItem">
          {column2Lg.map((result, index) => (
            <div key={`${result.title}-${index}`}>
              <img
                src={cleanImageUrl(result.url, 400)}
                alt={result.title}
                className="image"
              />
            </div>
          ))}
        </div>
        <div className="gridItem">
          {column3Lg.map((result, index) => (
            <div key={`${result.title}-${index}`}>
              <img
                src={cleanImageUrl(result.url, 400)}
                alt={result.title}
                className="image"
              />
            </div>
          ))}
        </div>
      </div>
      <div className="gridContainerMd">
        <div className="gridItem">
          {column1Md.map((result, index) => (
            <div key={`${result.title}-${index}`}>
              <img
                src={cleanImageUrl(result.url, 400)}
                alt={result.title}
                className="image"
              />
            </div>
          ))}
        </div>
        <div className="gridItem">
          {column2Md.map((result, index) => (
            <div key={`${result.title}-${index}`}>
              <img
                src={cleanImageUrl(result.url, 400)}
                alt={result.title}
                className="image"
              />
            </div>
          ))}
        </div>
      </div>

      <div className="gridContainerSm">
        {searchResults.map((result, index) => (
          <div key={`${result.title}-${index}`}>
            <img
              src={cleanImageUrl(result.url, 400)}
              alt={result.title}
              className="image"
            />
          </div>
        ))}
      </div>
    </>
  );
}

export default SearchResults;
