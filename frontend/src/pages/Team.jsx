import React from "react";
import { Link } from "react-router-dom";

const team = [
  {name:"Avishkar", role:"CEO", email:"avishkar@example.com", linkedin:"https://linkedin.com/in/avishkar", photo:"/assets/avishkar.jpg"},
  {name:"Sadhana", role:"Manager", email:"sadhana@example.com", linkedin:"https://linkedin.com/in/sadhana", photo:"/assets/sadhana.jpg"},
  {name:"Vivek", role:"Team Member", email:"vivek@example.com", linkedin:"https://linkedin.com/in/vivek", photo:"/assets/vivek.jpg"},
  {name:"Navneet", role:"Team Member", email:"navneet@example.com", linkedin:"https://linkedin.com/in/navneet", photo:"/assets/navneet.jpg"},
  {name:"Atish", role:"Team Member", email:"atish@example.com", linkedin:"https://linkedin.com/in/atish", photo:"/assets/atish.jpg"}
];

export default function Team(){
  return (
    <div style={{padding:30}}>
      <h2>Our Team</h2>
      <div style={{display:"flex", gap:20}}>
        {team.map((t) => (
          <Link key={t.email} to={`/profile/${encodeURIComponent(t.email)}`} style={{textDecoration:"none", color:"#000"}}>
            <div style={{width:160, background:"#fff", padding:12, boxShadow:"0 1px 4px rgba(0,0,0,0.1)"}}>
              <img src={t.photo} alt={t.name} style={{width:140,height:100, objectFit:"cover", outline:"3px solid #ffd54f"}}/>
              <h4>{t.name}</h4>
              <small>{t.role}</small>
            </div>
          </Link>
        ))}
      </div>
    </div>
  )
}
