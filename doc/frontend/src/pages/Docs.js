import React from 'react';
import { ClientState } from '../context/DocsProvider';
import Sidebar from '../components/SideBar';
import Document from '../components/Document';

export default function Docs() {
  const { name, clients } = ClientState();
    if (!name) return <h1>Crash, sign in again..</h1>
  return (
    <div className="flex h-screen">
      <div className="w-1/5">
        <Sidebar clients={clients} currentUser={name} />
      </div>
      <div className="w-4/5">
        <Document />
      </div>
    </div>
  );
}