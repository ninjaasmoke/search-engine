import { Link } from "react-router-dom";
import Images from "../src/images/Images";

import "./QuickSearches.css";

function QuickSearches() {
  return (
    <div className="quickSearches">
      <QuickSearchImage
        src={Images.cat_with_sunglasses}
        alt="cat with sunglasses"
        link="/search?q=cat with sunglasses"
        title="cat with sunglasses"
      />
      <QuickSearchImage
        src={Images.flower_arrangement}
        alt="flower arrangement"
        link="/search?q=flower arrangement"
        title="flower arrangement"
      />
      <QuickSearchImage
        src={Images.sunset}
        alt="sunset"
        link="/search?q=sunset"
        title="sunset"
      />
      <QuickSearchImage
        src={Images.flowers_near_eiffel_tower}
        alt="flowers near eiffel tower"
        link="/search?q=flowers near eiffel tower"
        title="flowers near eiffel tower"
      />
    </div>
  );
}

const QuickSearchImage = ({
  src,
  alt,
  link,
  title,
}: {
  src: string;
  alt: string;
  link: string;
  title?: string;
}) => {
  return (
    <Link to={link}>
      <div
        style={{
          // background: "linear-gradient(0deg, #f3f3f3, #f3f3f3)",
          overflow: "hidden",
          margin: "10px",
          cursor: "pointer",

          boxShadow: "0 0 10px rgba(0, 0, 0, 0.2)",
          padding: "14px 10px",
        }}
      >
        <img
          className="quickSearchImage"
          src={src}
          alt={alt}
        />
        {title && <div style={{ marginTop: 8 }}>{title}</div>}
      </div>
    </Link>
  );
};

export default QuickSearches;
