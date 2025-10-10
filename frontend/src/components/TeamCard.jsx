import React from "react";
import { useNavigate } from "react-router-dom";

function TeamCard({ member }) {
  const navigate = useNavigate();

  return (
    <div
      onClick={() => navigate(`/profile/${member._id}`)}
      className="bg-white rounded-xl shadow-md p-4 hover:scale-105 cursor-pointer transition-all"
    >
      <img
        src={member.photo}
        alt={member.name}
        className="w-32 h-32 rounded-full mx-auto object-cover"
      />
      <h3 className="text-center mt-3 font-semibold">{member.name}</h3>
      <p className="text-center text-gray-500 text-sm">{member.role}</p>
    </div>
  );
}

export default TeamCard;
