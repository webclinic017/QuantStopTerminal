module.exports = {
    /*devServer: {
        https: true,
        proxy: 'https://localhost/'
    }*/
    devServer: {
        proxy: {
            '^/api': {
                target: 'https://192.168.0.80/',
                ws: true,
                changeOrigin: true
            },

        }
    }
};