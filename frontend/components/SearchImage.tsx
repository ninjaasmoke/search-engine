import { Link } from "react-router-dom";
import { SearchData } from "../types";
import React from "react";
import { cleanImageUrl } from "../src/utils/cleanURL";

type SearchImageProps = {
  imageData: SearchData;
};

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
