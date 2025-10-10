import React from "react";
import { Link } from "react-router-dom";

function Navbar() {
  return (
    <nav className="bg-gray-900 text-white p-4 flex justify-between items-center">
      <h1 className="text-2xl font-bold">Freelance Connect</h1>
      <div className="space-x-4">
        <Link to="/" className="hover:text-blue-400">Home</Link>
        <Link to="/register" className="hover:text-blue-400">Register</Link>
        <Link to="/login" className="hover:text-blue-400">Login</Link>
        <Link to="/post-task" className="hover:text-blue-400">Post Task</Link>
      </div>
    </nav>
  );
}

export default Navbar;
