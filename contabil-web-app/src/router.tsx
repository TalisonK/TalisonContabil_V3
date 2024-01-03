import React from "react";

import { BrowserRouter, Routes, Route, createBrowserRouter } from "react-router-dom";
import Login from "./pages/login/Login";
import Dashboard from "./pages/dashboard/dashboard";



export const RouterApp = () => {

    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" Component={Dashboard} />
            </Routes>
        </BrowserRouter>
    )

}