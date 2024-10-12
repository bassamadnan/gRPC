import React, { useState, useEffect, useRef } from 'react';
import { ClientState } from '../context/DocsProvider';
import { HelloRequest } from '../proto/docs_pb';

const Document = () => {
  const { name, client } = ClientState();
  const [greeting, setGreeting] = useState('');
  const [error, setError] = useState('');
  const hasSaidHello = useRef(false);

  useEffect(() => {
    const sayHello = () => {
      if (hasSaidHello.current) return;
      hasSaidHello.current = true;

      const request = new HelloRequest();
      request.setName(name);

      client.sayHello(request, {}, (err, response) => {
        if (err) {
          console.error('Error:', err);
          setError(`Error occurred while fetching greeting: ${err.message}`);
        } else {
          setGreeting(response.getMessage());
        }
      });
    };
    sayHello()
  });

  const handleTyping = (event) => {
    console.log('Character typed at position:', event.target.selectionStart);
  };

  return (
    <div className="flex flex-col h-screen p-4 bg-gray-900 text-white">
      <h2 className="text-2xl font-bold mb-4 border-b border-gray-600 pb-2">Collaborative Document</h2>
      {greeting && (
        <p className="text-lg mb-4 bg-green-700 p-2 rounded">Server says: {greeting}</p>
      )}
      {error && (
        <p className="text-lg text-red-400 mb-4">{error}</p>
      )}
      <textarea 
        className="flex-1 w-full p-3 border border-gray-700 rounded resize-none bg-gray-800 text-white focus:outline-none focus:border-blue-500"
        placeholder="Start typing your document here..."
        onInput={handleTyping}
        style={{
          minHeight: '300px'
        }}
      />
    </div>
  );
};

export default Document;
