const { contextBridge, ipcRenderer } = require('electron');

contextBridge.exposeInMainWorld('api', {
  sayHello: (name) => ipcRenderer.invoke('say-hello', name),
  sendRecvNumbers: (numbers) => ipcRenderer.invoke('sendRecvNumbers-invoke', numbers),
  onNumbersReceived: (callback) => {
    const wrappedCallback = (_, numbers) => callback(numbers);
    ipcRenderer.on('numbers-received', wrappedCallback);
    return () => ipcRenderer.removeListener('numbers-received', wrappedCallback);
  },
  removeNumbersListener: (callback) => ipcRenderer.removeListener('numbers-received', callback)
});
