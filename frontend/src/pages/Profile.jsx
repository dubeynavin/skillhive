import React from "react";
import { useParams } from "react-router-dom";

function Profile() {
  const { id } = useParams();

  // Team Data
  const team = [
    {
      _id: "1",
      name: "Avishkar",
      role: "CEO",
      photo: "https://i.pravatar.cc/150?img=5",
      linkedin: "https://linkedin.com/in/avishkar",
      bio: "CEO & Visionary Leader of Freelance Connect",
    },
    {
      _id: "2",
      name: "Sadhana Pandey",
      role: "Manager",
      photo: "https://i.pravatar.cc/150?img=6",
      linkedin: "https://linkedin.com/in/sadhanapandey",
      bio: "Project Manager | Ensuring smooth collaboration and delivery.",
    },
    {
      _id: "3",
      name: "Vivek Pandey",
      role: "Developer",
      photo: "https://i.pravatar.cc/150?img=7",
      linkedin: "https://linkedin.com/in/thevivek",
      bio: "Full Stack Developer | GoLang & React Specialist",
    },
    {
      _id: "4",
      name: "Atish Kumar",
      role: "Frontend Developer",
      photo: "https://i.pravatar.cc/150?img=8",
      linkedin: "https://linkedin.com/in/atishkumar",
      bio: "Frontend Developer | Crafting UI experiences with ReactJS.",
    },
    {
      _id: "5",
      name: "Navneet Dubey",
      role: "Backend Engineer",
      photo: "https://i.pravatar.cc/150?img=9",
      linkedin: "https://linkedin.com/in/navneetdubey",
      bio: "Backend Developer | Building scalable APIs in GoLang.",
    },
  ];

  const profile = team.find((member) => member._id === id);

  if (!profile) {
    return (
      <div className="text-center mt-10">
        <h2 className="text-2xl font-bold">Profile Not Found</h2>
      </div>
    );
  }

  return (
    <div className="max-w-lg mx-auto mt-10 bg-white p-6 shadow-md rounded text-center">
      <img
        src={profile.photo}
        className="w-32 h-32 rounded-full mx-auto"
        alt={profile.name}
      />
      <h2 className="text-2xl font-bold mt-3">{profile.name}</h2>
      <p className="text-gray-500">{profile.role}</p>
      <p className="text-gray-600 mt-2">{profile.bio}</p>
      <a
        href={profile.linkedin}
        target="_blank"
        rel="noreferrer"
        className="text-blue-600 underline mt-3 inline-block"
      >
        LinkedIn Profile
      </a>
    </div>
  );
}

export default Profile;
