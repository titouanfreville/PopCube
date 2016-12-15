const electron = require('electron')

const app = electron.app

const Window = electron.BrowserWindow

const path = require('path')
const url = require('url')

let mainWindow

function createWindow () {
	mainWindow = new Window({width: 800, height: 600})

	mainWindow.loadURL(url.format({
		pathname: path.join(__dirname, 'index.html'),
		protocol: 'file',
		slashes: true
	}))

	mainWindow.webContents.openDevTools()

	mainWindow.on('closed', function () {
		mainWindow = null
	})
}

app.on('ready', createWindow)

app.on('window-all-closed', function () {
	if (process.platform !== 'darwin') {
		app.quit()
	}
})

app.on('activate', function () {
	if (mainWindow === null) {
		createWindow()
	}
})