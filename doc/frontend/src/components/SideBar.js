import React from 'react';

const Sidebar = ({ clients, currentUser }) => {
  return (
    <div className="bg-gray-100 p-4 h-full">
      <h2 className="text-xl font-bold mb-4">Connected Clients</h2>
      <ul>
        {clients.map((client, index) => (
          <li 
            key={index} 
            className={`mb-2 p-2 rounded ${client === currentUser ? 'bg-blue-200' : ''}`}
          >
            {client}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Sidebar;