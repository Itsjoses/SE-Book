import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Home from "./pages/home.tsx";
import Detail from "./pages/detail.tsx";
import Product from "./pages/product.tsx";
import AboutUs from "./pages/aboutUs.tsx";
import GetJson from "./pages/getJson.tsx";
const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/detail/:id",
    element: <Detail />,
  },
  {
    path: "/product",
    element: <Product />,
  },
  {
    path: "/about-us",
    element: <AboutUs />,
  },
  {
    path: "/getjson",
    element: <GetJson />,
  }
  
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
