module.exports = {
  devServer: {
    proxy: {
      '/readings': {
        target: 'http://localhost:8080',
        ws: true,
        changeOrigin: true
      }
    }
  }
}
