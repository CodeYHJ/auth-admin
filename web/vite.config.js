import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';
import path from 'path'
console.log(path.resolve(__dirname, "./src/api"))
// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      "@api": path.resolve(__dirname, "./src/api")
    }
  },
  server:{
    proxy:{
      "/api":{
        target: 'http://127.0.0.1:1323',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  },
  plugins: [vue(), Components({
    resolvers: [AntDesignVueResolver()],
  }),]
})
