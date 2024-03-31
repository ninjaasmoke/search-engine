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
      <div style={{ padding: 20, fontSize: 12 }}>
        <p>
          developed by Nithin Sai Kirumani Jagadish as part of CA6005 Mechanics of Search project
        </p>
      </div>
    </div>
  );
}

export default App;
