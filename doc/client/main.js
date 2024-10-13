const { app, BrowserWindow, ipcMain } = require('electron');
const path = require('path');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = path.join(__dirname, '../docs.proto');

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});
const docsProto = grpc.loadPackageDefinition(packageDefinition).docs;

const client = new docsProto.DocumentService(
  'localhost:5050',
  grpc.credentials.createInsecure()
);

function createWindow() {
  const win = new BrowserWindow({   
    width: 800,
    height: 600,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'), // Optional for secure IPC, does not run without this
      nodeIntegration: false, // Enable if not using preload
      contextIsolation: true, // Disable if not using preload
    },
  });

  win.webContents.openDevTools();
  win.loadURL('http://localhost:3000'); // React dev server
}

app.whenReady().then(() => {
  createWindow();

  app.on('activate', function () {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

ipcMain.handle('say-hello', async (event, name) => {
  return new Promise((resolve, reject) => {
    client.SayHello({ name }, (error, response) => {
      if (error) {
        reject(error);
      } else {
        resolve(response.message);
      }
    });
  });
});


ipcMain.handle('sendRecvNumbers-invoke', async (event, numbers) => {
  return new Promise((resolve, reject) => {
    const call = client.SendRecvNumbers();
    const receivedNumbers = [];
    
    call.on('data', (data) => {
      console.log("Received from server:", data.number);
      receivedNumbers.push(data.number);
    });

    call.on('end', () => {
      console.log("Stream ended");
      console.log("All received numbers:", receivedNumbers);
      event.sender.send('numbers-received', receivedNumbers);
      resolve('Stream completed');
    });

    call.on('error', (error) => {
      console.error("Stream error:", error);
      reject(error);
    });

    numbers.forEach((num) => {
      call.write({ number: num });
    });

    call.end();
  });
});
app.on('window-all-closed', function () {
  if (process.platform !== 'darwin') app.quit();
});
