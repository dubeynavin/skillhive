import React, {useState} from "react";
import API from "../api/api";
import { useNavigate } from "react-router-dom";

export default function Register(){
  const [name,setName]=useState(""); const [email,setEmail]=useState(""); const [password,setPassword]=useState("");
  const nav = useNavigate();
  const submit = async (e) => {
    e.preventDefault();
    await API.post("/auth/register", {name, email, password});
    nav("/login");
  }
  return (
    <div style={{padding:30}}>
      <h2>Register</h2>
      <form onSubmit={submit}>
        <input placeholder="Name" value={name} onChange={e=>setName(e.target.value)} /><br/>
        <input placeholder="Email" value={email} onChange={e=>setEmail(e.target.value)} /><br/>
        <input placeholder="Password" type="password" value={password} onChange={e=>setPassword(e.target.value)} /><br/>
        <button>Register</button>
      </form>
    </div>
  )
}
