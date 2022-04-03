module.exports = {
    devServer: {
        port: 8081, // CHANGE YOUR PORT HERE!
        https: true,
        proxy: 'https://localhost/',
    },
    configureWebpack: {
        devServer: {
            headers: { "Access-Control-Allow-Origin": "*" }
        }
    }
};