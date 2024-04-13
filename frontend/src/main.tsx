import { HashRouter, Route, Routes } from "react-router-dom";
import ReactDOM from "react-dom/client";
import React from "react";

import ImagePage from "./Image.page.tsx";
import Search from "./Search.page.tsx";
import App from "./App.tsx";

import "./index.css";
import AboutPage from "./About.page.tsx";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <HashRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/search" element={<Search />} />
        <Route path="/image/:id" element={<ImagePage />} />
        <Route path="/about" element={<AboutPage />} />
        <Route path="*" element={<App />} />
      </Routes>
    </HashRouter>
  </React.StrictMode>
);
