import React from 'react';

const Sidebar = ({ clients, currentUser }) => {
  return (
    <div className="bg-gray-800 p-4 h-full text-white">
      <h2 className="text-xl font-bold mb-4 border-b border-gray-600 pb-2">Connected Clients</h2>
      <ul>
        {clients.map((client, index) => (
          <li 
            key={index} 
            className={`mb-2 p-2 rounded transition-all duration-300 
              ${client === currentUser ? 'bg-blue-500 text-white' : 'bg-gray-700 hover:bg-gray-600'}`}
          >
            {client} {client === currentUser ? "(You)":""}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Sidebar;
