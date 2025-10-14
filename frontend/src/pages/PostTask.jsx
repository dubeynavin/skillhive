import React, {useState} from "react";
import API from "../api/api";

export default function PostTask(){
  const [title,setTitle]=useState(""); const [description,setDescription]=useState("");

  const submit = async (e) => {
    e.preventDefault();
    await API.post("/tasks", { title, description, posted_by: "demo-user" });
    alert("Task posted");
    setTitle(""); setDescription("");
  }

  return (
    <div style={{padding:30}}>
      <h2>Post a Task</h2>
      <form onSubmit={submit}>
        <input placeholder="Task Title" value={title} onChange={e=>setTitle(e.target.value)} /><br/>
        <textarea placeholder="Describe your requirement" value={description} onChange={e=>setDescription(e.target.value)} rows={6} cols={60} /><br/>
        <button>Post Task</button>
      </form>
    </div>
  )
}
