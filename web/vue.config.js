/*
const fs = require('fs')

module.exports = {
    devServer: {
        https: {
            key: fs.readFileSync('C:\\Users\\ethan\\AppData\\Roaming\\QuantstopTerminal\\tls\\key.pem'),
            cert: fs.readFileSync('C:\\Users\\ethan\\AppData\\Roaming\\QuantstopTerminal\\tls\\cert.pem'),
        },
        public: 'https://localhost:8081/'
    }
}*/

/*
module.exports = {
    devServer: {
        open: process.platform === 'win32',
        host: '0.0.0.0',
        port: 8085, // CHANGE YOUR PORT HERE!
        https: true,
        hotOnly: false,
    },
}*/

module.exports = {
    devServer: {
        port: 8081, // CHANGE YOUR PORT HERE!
        https: false,
    },
    configureWebpack: {
        devServer: {
            headers: { "Access-Control-Allow-Origin": "*" }
        }
    }
};