import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import reportWebVitals from "./reportWebVitals";
import ClientProvider from "./context/DocsProvider";
import Docs from "./pages/Docs";

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <ClientProvider>
      <Router>
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="/docs" element={<Docs />} />
        </Routes>
      </Router>
    </ClientProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
