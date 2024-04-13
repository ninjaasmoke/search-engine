import SearchInputComponent from "../components/SearchInput.component";

import { useEffect, useState } from "react";
import "./App.css";
import { useNavigate } from "react-router-dom";

function App() {
  const navigate = useNavigate();
  const [searchText, setSearchText] = useState("");

  const handleSearch = async () => {
    // window.location.href = `/search?q=${searchText}`;
    navigate(`/search?q=${searchText}`);
  };

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === "Enter") {
        handleSearch();
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
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <div>
        <h1>an image search engine.</h1>
        <div className="searchArea">
          <SearchInputComponent
            searchText={searchText}
            handleChange={(text) => {
              setSearchText(text);
            }}
          />
        </div>
      </div>
      <div
        style={{
          padding: 20,
          // textAlign: "center",
          maxWidth: 680,
        }}
      >
        <p>
          A project developed by{" "}
          <a target="_blank" href="https://github.com/ninjaasmoke">
            Nithin Sai Kirumani Jagadish
          </a>{" "}
          as part of{" "}
          <a
            target="_blank"
            href="https://www.dcu.ie/engineeringandcomputing/mechanics-search"
          >
            CA6005 Mechanics of Search
          </a>{" "}
          project submitted to{" "}
          <a target="_blank" href="https://dcu.ie">
            Dublin City University
          </a>
          . The goal of this project is to develop an image based search engine
          using annotations and image features.
        </p>
        <p>
          This project uses images from{" "}
          <a target="_blank" href="https://unsplash.com">
            Unsplash
          </a>{" "}
          and the license to use the images can be found{" "}
          <a target="_blank" href="https://unsplash.com/license">
            here
          </a>
          .
        </p>
        <p>
          This search engine is <strong>NOT</strong> a competitor to Unsplash in
          any way and is developed for educational purposes only.
        </p>
        <p>
          <strong>ALL</strong> images you see in this website belong to Unsplash
          alone.
        </p>
        <p>
          In the off chance that anyone from Unsplash is viewing this, please{" "}
          <a
            target="_blank"
            href="mailto:nithin.kirumanijagadish2@mail.dcu.ie"
            style={{ textDecoration: "underline" }}
          >
            email me
          </a>{" "}
          if you have any concerns.
        </p>
      </div>
    </div>
  );
}

export default App;
