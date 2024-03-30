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
  const column1 = searchResults.slice(0, Math.ceil(searchResults.length / 3));
  const column2 = searchResults.slice(
    column1.length,
    column1.length + Math.ceil(searchResults.length / 3)
  );
  const column3 = searchResults.slice(
    column1.length + column2.length,
    searchResults.length
  );
  return (
    <div className="gridContainer">
      <div className="gridItem">
        {column1.map((result, index) => (
          <div key={`${result.title}-${index}`}>
            <img
              src={cleanImageUrl(result.url, 300)}
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
              src={cleanImageUrl(result.url, 300)}
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
              src={cleanImageUrl(result.url, 300)}
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
