module.exports = {
  devServer: {
    proxy: {
      '/readings': {
        target: 'http://0.0.0.0:9000',
        ws: true,
        changeOrigin: true
      }
    }
  }
}
