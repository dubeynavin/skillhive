import React, { useState } from "react";
import API from "../services/api";

function PostTask() {
  const [task, setTask] = useState({ title: "", description: "", contact: "" });

  const handleChange = (e) => setTask({ ...task, [e.target.name]: e.target.value });

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await API.post("/tasks", task);
      alert("Task Posted Successfully!");
    } catch (err) {
      alert("Failed to post task!");
    }
  };

  return (
    <div className="max-w-lg mx-auto mt-10 bg-white p-6 shadow-md rounded">
      <h2 className="text-2xl font-bold mb-4 text-center">Post a Task</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <input name="title" onChange={handleChange} placeholder="Task Title" className="w-full p-2 border rounded" />
        <textarea name="description" onChange={handleChange} placeholder="Task Description" className="w-full p-2 border rounded" />
        <input name="contact" onChange={handleChange} placeholder="Contact No or Email" className="w-full p-2 border rounded" />
        <button className="w-full bg-purple-600 text-white py-2 rounded">Submit</button>
      </form>
    </div>
  );
}

export default PostTask;
