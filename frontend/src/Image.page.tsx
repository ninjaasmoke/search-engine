import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { SearchData } from "../types";
import { API_URL_DEV, API_URL_PROD } from "../constants";
import { cleanImageUrl } from "./utils/cleanURL";

const ImagePage = () => {
  const params = useParams();

  const id = params.id;

  const [image, setImage] = useState<SearchData | null>(null);

  const fetchData = async () => {
    const response = await fetch(
      import.meta.env.MODE == "production"
        ? `${API_URL_PROD}imageData/${id}`
        : `${API_URL_DEV}imageData/${id}`
    );

    if (!response.ok) {
      console.error("Failed to fetch image data");
      return;
    }

    const data = (await response.json()) as SearchData;

    setImage(data);
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div className="imagePage">
      {image && (
        <>
          <div style={{ width: "100%" }}>
            <h2>Image Data</h2>
            <h3>Title</h3>
            <p>{image.title}</p>
            <div className="imageTags">
              <div>
                <h3>Related Image Tags</h3>
                <ul>
                  {image.related_image_tags.map((tag, idx) => (
                    <li key={tag + idx}>{tag}</li>
                  ))}
                </ul>
              </div>
              <div>
                <h3>Annotated Image Tags</h3>
                <ul>
                  {image.annotated_image_tags.map((tag, idx) => (
                    <li key={tag + idx}>{tag}</li>
                  ))}
                </ul>
              </div>
            </div>
          </div>
          <a href={image.url} target="_blank">
            <img src={cleanImageUrl(image.url, 400)} alt={image.title} />
          </a>
        </>
      )}
    </div>
  );
};

export default ImagePage;
