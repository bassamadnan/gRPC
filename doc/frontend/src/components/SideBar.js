import React, { useEffect, useState } from 'react';
import { ClientState } from '../context/DocsProvider';
import { Client } from '../proto/docs_pb';

const Sidebar = () => {
    // client is DocumentServiceClient, name is client name
    const { client, name } = ClientState();
    const [clients, setClients] = useState([]);

    useEffect(() => {
        if (!client || !name) return;

        const clientMessage = new Client();
        clientMessage.setName(name);

        const stream = client.registerClient();

        // Send initial client information
        stream.write(clientMessage);

        stream.on('data', (response) => {
            console.log(`got data ${response}`);
            
            setClients(response.getClientsList());
        });

        stream.on('error', (err) => {
            console.error('Stream error:', err);
        });

        stream.on('end', () => {
            console.log('Stream ended');
        });

        // You can send periodic messages to keep the stream alive if needed
        const interval = setInterval(() => {
            stream.write(clientMessage);
        }, 5 * 1000); // every 5 seconds

        return () => {
            clearInterval(interval);
            stream.cancel();
        };
    }, [client, name]);

    return (
        <div className="bg-gray-800 p-4 h-full text-white">
            <h2 className="text-xl font-bold mb-4 border-b border-gray-600 pb-2">Connected Clients</h2>
            <ul>
                {clients.map((client, index) => (
                    <li 
                        key={index} 
                        className={`mb-2 p-2 rounded transition-all duration-300 
                            ${client === name ? 'bg-blue-500 text-white' : 'bg-gray-700 hover:bg-gray-600'}`}
                    >
                        {client} {client === name ? "(You)" : ""}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Sidebar;