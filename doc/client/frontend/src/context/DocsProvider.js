import { createContext, useContext, useState } from "react";

const ClientContext = createContext();

const ClientProvider = ({ children }) => {
  const [client, setClient] = useState();
  const [name, setName] = useState("");
  const [clients, setClients] = useState(["name1", "name2"])
  return (
    <ClientContext.Provider
      value={{
        client,
        setClient,
        name,
        setName,
        clients,
        setClients
      }}
    >
      {children}
    </ClientContext.Provider>
  );
};

export const ClientState = () => {
  return useContext(ClientContext);
};

export default ClientProvider;
