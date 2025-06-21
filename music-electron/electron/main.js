// electron/main.js
import { app, BrowserWindow } from 'electron';
import path from 'path';
import { fileURLToPath } from 'url';

//const isDev = !app.isPackaged;
const isDev = process.env.NODE_ENV === 'development';


// For __dirname support in ESM
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

let mainWindow;

function createWindow() {
  mainWindow = new BrowserWindow({
    width: 900,
    height: 600,
    webPreferences: {
      contextIsolation: true
    },
  });

  if (isDev) {
    mainWindow.loadURL('http://localhost:5173');
    console.log("DEBUGGING: ELECTRON LOADING HTTP");
  } else {
    mainWindow.loadFile(path.join(__dirname, '../build/index.html'));
    console.log("DEBUGGING: ELECTRON LOADING INDEX.HTML");
  }
}

app.whenReady().then(createWindow);

