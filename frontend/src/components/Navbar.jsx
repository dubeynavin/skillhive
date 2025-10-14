import React from "react";
import { Link } from "react-router-dom";

export default function Navbar(){
  const services = ["Web Development","Mobile App","UI/UX","Data Science","SEO","Content Writing"];
  return (
    <div style={{width:220, minHeight:"100vh", background:"#1f2937", color:"#fff", padding:20}}>
      <div style={{display:"flex", alignItems:"center", gap:10}}>
        <img src="/logo.png" alt="logo" style={{width:40, height:40}}/>
        <h2>SkillHiveAI</h2>
      </div>
      <hr style={{borderColor:"#374151"}}/>
      <nav>
        <Link to="/" style={{color:"#fff", display:"block", margin:"10px 0"}}>Home</Link>
        <Link to="/post-task" style={{color:"#fff", display:"block", margin:"10px 0"}}>Post Task</Link>
        <Link to="/team" style={{color:"#fff", display:"block", margin:"10px 0"}}>Team</Link>
        <Link to="/login" style={{color:"#fff", display:"block", margin:"10px 0"}}>Login / Register</Link>
        <h4 style={{marginTop:20}}>Services</h4>
        <ul>
          {services.map((s, i) => <li key={i} style={{margin:"6px 0"}}>{s}</li>)}
        </ul>
      </nav>
    </div>
  )
}
