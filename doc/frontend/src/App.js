import React, { useState } from 'react';
import { GreeterClient } from './proto/docs_grpc_web_pb';
import { useNavigate } from 'react-router-dom';
import { ClientState } from './context/DocsProvider';

const client = new GreeterClient('http://localhost:5050', null, null);

export default function App() {
  const [inputName, setInputName] = useState('');
  const { setName, setClient, clients, setClients } = ClientState();
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    setName(inputName);
    setClient(client);
    setClients([...clients, inputName]);
    navigate('/docs');
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
      </div>
    </div>
  );
}
