module.exports = {
    devServer: {
        proxy: {
            '^/api': {
                target: 'https://localhost',
                ws: true,
                changeOrigin: true
            },
        }
    },
    configureWebpack: {
        devServer: {
            headers: { "Access-Control-Allow-Origin": "*" }
        }
    }
};