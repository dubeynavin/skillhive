import React from "react";
import TeamCard from "../components/TeamCard";

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

function Home() {
  return (
    <div className="text-center">
      <h1 className="text-4xl font-bold mb-6">Welcome to Freelance Connect</h1>
      <p className="text-gray-600 mb-10">
        Connect, Collaborate, and Complete Freelance Projects with Ease.
      </p>

      <h2 className="text-2xl font-semibold mb-4">Our Leadership Team</h2>
      <div className="grid grid-cols-1 sm:grid-cols-3 gap-6 px-10">
        {team.map((member) => (
          <TeamCard key={member._id} member={member} />
        ))}
      </div>
    </div>
  );
}

export default Home;
