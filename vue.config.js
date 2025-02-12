const { defineConfig } = require('@vue/cli-service')

const fs = require('fs')
const path = require('path')
const configPath = path.resolve(__dirname, process.env.CONFIG_FILE_PATH)
const config = fs.readFileSync(configPath, 'utf-8')
const port = JSON.parse(config).GinPort

module.exports = defineConfig({
  publicPath: './',
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/api': {
        target: "http://localhost" + port,
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
})
