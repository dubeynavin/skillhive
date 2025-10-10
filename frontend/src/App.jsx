import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import Home from "./pages/Home";
import Register from "./pages/Register";
import Login from "./pages/Login";
import PostTask from "./pages/PostTask";
import Profile from "./pages/Profile";

function App() {
  return (
    <Router>
      <Navbar />
      <div className="p-6">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/post-task" element={<PostTask />} />
          <Route path="/profile/:id" element={<Profile />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
