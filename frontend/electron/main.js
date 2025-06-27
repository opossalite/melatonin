// electron/main.js
import { app, BrowserWindow } from 'electron';
import path from 'path';
import { fileURLToPath } from 'url';
import { spawn } from 'child_process';

//const isDev = !app.isPackaged;
const isDev = process.env.NODE_ENV === 'development';


// For __dirname support in ESM
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);


function createWindow() {
    let mainWindow = new BrowserWindow({
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

let rustBackend;

app.whenReady().then(() => {
    // 🟢 Start the Rust backend binary
    const binaryPath = path.join(__dirname, '../../backend/target/release/melatonin');
    rustBackend = spawn(binaryPath);

    rustBackend.stdout.on('data', (data) => {
        console.log(`[Rust stdout]: ${data}`);
    });

    rustBackend.stderr.on('data', (data) => {
        console.error(`[Rust stderr]: ${data}`);
    });

    rustBackend.on('close', (code) => {
        console.log(`Rust backend exited with code ${code}`);
    });

    createWindow();
}).catch(err => {
    console.error("Electron startup error:", err);
});

app.on('window-all-closed', () => {
    if (rustBackend) {
        rustBackend.kill(); // terminate backend
    }
    if (process.platform !== 'darwin') {
        app.quit();
    }
});


