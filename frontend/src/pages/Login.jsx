import React, { useState } from "react";
import API from "../services/api";

function Login() {
  const [form, setForm] = useState({ email: "", password: "" });

  const handleChange = (e) => setForm({ ...form, [e.target.name]: e.target.value });

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await API.post("/auth/login", form);
      alert(res.data.message);
    } catch (err) {
      alert("Login failed!");
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 bg-white p-6 shadow-md rounded">
      <h2 className="text-2xl font-bold mb-4 text-center">Login</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <input name="email" onChange={handleChange} placeholder="Email" className="w-full p-2 border rounded" />
        <input name="password" type="password" onChange={handleChange} placeholder="Password" className="w-full p-2 border rounded" />
        <button className="w-full bg-green-600 text-white py-2 rounded">Login</button>
      </form>
    </div>
  );
}

export default Login;
