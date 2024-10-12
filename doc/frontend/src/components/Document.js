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

    if (client && name) {
      sayHello();
    }
  }, [client, name]);

  return (
    <div className="p-4">
      <h2 className="text-2xl font-bold mb-4">Collaborative Document</h2>
      {greeting && (
        <p className="text-lg mb-4">Server says: {greeting}</p>
      )}
      {error && (
        <p className="text-lg text-red-500 mb-4">{error}</p>
      )}
      <textarea 
        className="w-full h-64 p-2 border rounded"
        placeholder="Start typing your document here..."
      />
    </div>
  );
};

export default Document;