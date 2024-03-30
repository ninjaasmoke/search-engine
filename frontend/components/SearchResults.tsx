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
  const column1: SearchData[] = [];
  const column2: SearchData[] = [];
  const column3: SearchData[] = [];

  searchResults.forEach((result, index) => {
    if (index % 3 === 0) {
      column1.push(result);
    } else if (index % 3 === 1) {
      column2.push(result);
    } else {
      column3.push(result);
    }
  });

  return (
    <div className="gridContainer">
      <div className="gridItem">
        {column1.map((result, index) => (
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
        {column2.map((result, index) => (
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
        {column3.map((result, index) => (
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
  );
}

export default SearchResults;
