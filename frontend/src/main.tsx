import React from 'react'
import ReactDOM from 'react-dom/client'
import {
  HashRouter,
  Route,
  Routes
} from "react-router-dom";
import App from './App.tsx'
import './index.css'
import Search from './Search.page.tsx';

// const router = createHashRouter([
//   {
//     path: "/",
//     element: <App />,
//   },
//   {
//     path: "/search",
//     element: <Search />,
//   }
// ]);


ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <HashRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/search" element={<Search />} />
      </Routes>
    </HashRouter>
  </React.StrictMode>,
)
