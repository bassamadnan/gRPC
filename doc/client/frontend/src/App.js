import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { ClientState } from './context/DocsProvider';

export default function App() {
  const [inputName, setInputName] = useState('');
  const [greeting, setGreeting] = useState('');
  const { setName, setClients } = ClientState();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const message = await window.api.sayHello(inputName);
      setGreeting(message);
      setName(inputName);
      setClients(prevClients => [...prevClients, inputName]);
      setTimeout(() => navigate('/docs')); // 1 second
    } catch (error) {
      console.error('gRPC Error:', error);
      setGreeting('Error communicating with server.');
    }
  };

  return (
    <div className="flex items-center justify-center h-screen bg-gray-900 text-white">
      <div className="p-8 bg-black rounded shadow-md">
        <h1 className="text-2xl font-bold mb-4 text-center">
          gRPC Collaborative Document
        </h1>
        <form onSubmit={handleSubmit} className="flex flex-col items-center">
          <input
            type="text"
            value={inputName}
            onChange={(e) => setInputName(e.target.value)}
            placeholder="Enter your name"
            className="border p-2 mb-4 w-full bg-gray-800 text-white"
            required
          />
          <button
            type="submit"
            className="bg-blue-500 text-white p-2 rounded w-full"
          >
            View Document
          </button>
        </form>
        {greeting && <p className="mt-4 text-lg text-center">{greeting}</p>}
      </div>
    </div>
  );
}