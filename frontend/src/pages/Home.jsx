import React from "react";
import { Link } from "react-router-dom";

export default function Home(){
  return (
    <div style={{padding:30}}>
      <div style={{position:"relative", height:250, overflow:"hidden", borderRadius:8}}>
        {/* Embedded YouTube as background (muted autoplay) */}
        <iframe 
          title="bg"
          src="https://www.youtube.com/embed/dQw4w9WgXcQ?autoplay=1&mute=1&controls=0&loop=1&playlist=dQw4w9WgXcQ"
          style={{position:"absolute", top:0, left:0, width:"100%", height:"100%", pointerEvents:"none"}}
          frameBorder="0"
          allow="autoplay; encrypted-media"
        />
        <div style={{position:"relative", zIndex:2, color:"#fff", padding:20}}>
          <h1>Welcome to SkillHiveAI</h1>
          <p>Connect with trusted freelancers — post your task and get bids.</p>
        </div>
      </div>

      <section style={{marginTop:30}}>
        <h2>About SkillHiveAI</h2>
        <p>A simple freelancing demo app built with React + Go + MongoDB</p>
        <Link to="/team">Meet the Team</Link>
      </section>
    </div>
  )
}
