import React from "react";
import { useParams } from "react-router-dom";

const teamMap = { /* same data as Team.jsx - in real app fetch from DB */ };

export default function Profile(){
  const {email} = useParams();
  const t = Object.values(teamMap).flat().find(x => x?.email === decodeURIComponent(email)) || {
    name:"Navneet", role:"Team Member", linkedin:"#", photo:"/assets/navneet.jpg", about:"Demo profile"
  };
  return (
    <div style={{padding:30}}>
      <div style={{display:"flex", gap:20}}>
        <img src={t.photo} alt={t.name} style={{width:220,height:220, objectFit:"cover", border:"4px solid #f59e0b"}}/>
        <div>
          <h2>{t.name}</h2>
          <p>{t.role}</p>
          <p>{t.about || "About info here."}</p>
          <p><a href={t.linkedin} target="_blank">LinkedIn</a></p>
        </div>
      </div>
    </div>
  )
}
