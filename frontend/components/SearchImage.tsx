import { Link } from "react-router-dom";
import { SearchData } from "../types";
import React from "react";

type SearchImageProps = {
  imageData: SearchData;
};

function cleanImageUrl(url: string, width: number) {
  const cleanedUrl = url.replace(/(\?|\&)(q|auto|fit|ixlib|ixid)=[^&]+/g, "");
  const replacedUrl = cleanedUrl.replace(/(\?|\&)w=[^&]+/g, "");
  const modifiedUrl =
    replacedUrl + (replacedUrl.includes("?") ? "&" : "?") + "w=" + width;
  return modifiedUrl;
}

const SearchImage: React.FC<SearchImageProps> = ({ imageData }) => {
  return (
    <Link to={`/image/${imageData.id}`}        
    >
      <div>
        <img
          src={cleanImageUrl(imageData.url, 400)}
          alt={imageData.title}
          className="image"
        />
      </div>
    </Link>
  );
};

export default SearchImage;
