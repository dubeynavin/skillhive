import React, {useState} from "react";
import API from "../api/api";
import { useNavigate } from "react-router-dom";

export default function Login(){
  const [email,setEmail]=useState(""); const [password,setPassword]=useState("");
  const nav = useNavigate();
  const submit = async (e) => {
    e.preventDefault();
    const res = await API.post("/auth/login", {email, password});
    localStorage.setItem("token", res.data.token);
    nav("/");
  }
  return (
    <div style={{padding:30}}>
      <h2>Login</h2>
      <form onSubmit={submit}>
        <input placeholder="Email" value={email} onChange={e=>setEmail(e.target.value)} /><br/>
        <input placeholder="Password" type="password" value={password} onChange={e=>setPassword(e.target.value)} /><br/>
        <button>Login</button>
      </form>
      <p><a href="/forgot">Forgot password?</a></p>
    </div>
  )
}
