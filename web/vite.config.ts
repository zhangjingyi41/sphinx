import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        vueDevTools(),
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        },
    },
    // 前端代理处理跨域问题
    server: {
        port: 5173,
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:8084',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, '')
            },
        }
    },
    build:{
        outDir: '../static',
        emptyOutDir: true,
        sourcemap: true,
        
    }
})
