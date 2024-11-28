import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
// 导入自动导入和按需导入插件
import Components from 'unplugin-vue-components/vite'
import AutoImport from 'unplugin-auto-import/vite'
import { VantResolver } from 'unplugin-vue-components/resolvers'
import VueSetupExtend from 'vite-plugin-vue-setup-extend'
import obfuscator from 'rollup-plugin-obfuscator'

export default defineConfig({
  base: './', // 这里更改打包相对绝对路径
  // base: '/', // 这里更改打包相对绝对路径
  build: {
    outDir: 'dist', // 指定打包输出的目录为 dist
    rollupOptions: {
      output: {
        // 在这里修改静态资源路径
        chunkFileNames: 'static/assets/js/[name]-[hash].js',
        entryFileNames: 'static/assets/js/[name]-[hash].js',
        assetFileNames: 'static/assets/[ext]/[name]-[hash].[ext]'
      }
    }
  },
  esbuild: {
    drop: ['console', 'debugger'] // 打包去除
  },
  plugins: [
    vue(),
    VueSetupExtend(),
    // 配置自动导入
    AutoImport({
      imports: ['vue'],
      dts: 'src/auto-imports.d.ts'
    }),
    // 配置按需导入
    Components({
      resolvers: [VantResolver()],
      dts: 'src/components.d.ts'
    }),
    obfuscator({
      global: true,
      // options配置项实际为 javascript-obfuscator 选项，具体可查看https://cloud.tencent.com/developer/article/2354357
      options: {
        compact: true,
        controlFlowFlattening: false,
        controlFlowFlatteningThreshold: 0.75,
        numbersToExpressions: false,
        simplify: false,
        stringArrayShuffle: false,
        splitStrings: false,
        splitStringsChunkLength: 10,
        rotateUnicodeArray: false,
        deadCodeInjection: false,
        deadCodeInjectionThreshold: 0.4,
        debugProtection: true,
        debugProtectionInterval: 2000,
        disableConsoleOutput: true,
        domainLock: [],
        identifierNamesGenerator: 'hexadecimal',
        identifiersPrefix: '',
        inputFileName: '',
        log: true,
        renameGlobals: true,
        reservedNames: [],
        reservedStrings: [],
        seed: 0,
        selfDefending: true,
        sourceMap: false,
        sourceMapBaseUrl: '',
        sourceMapFileName: '',
        sourceMapMode: 'separate',
        stringArray: false,
        stringArrayEncoding: ['base64'],
        stringArrayThreshold: 0.75,
        target: 'browser',
        transformObjectKeys: false,
        unicodeEscapeSequence: false,

        domainLockRedirectUrl: 'about:blank',
        forceTransformStrings: [],
        identifierNamesCache: null,
        identifiersDictionary: [],
        ignoreImports: true,
        optionsPreset: 'default',
        renameProperties: false,
        renamePropertiesMode: 'safe',
        sourceMapSourcesMode: 'sources-content',

        stringArrayCallsTransform: false,
        stringArrayCallsTransformThreshold: 0.5,

        stringArrayIndexesType: ['hexadecimal-number'],
        stringArrayIndexShift: false,
        stringArrayRotate: false,
        stringArrayWrappersCount: 1,
        stringArrayWrappersChainedCalls: false,
        stringArrayWrappersParametersMaxCount: 2,
        stringArrayWrappersType: 'variable'
      }
    })
  ],
  css: {
    preprocessorOptions: {
      scss: {
        api: 'modern-compiler'
      }
    }
  },
  server: {
    open: true, // 在服务器启动时自动在浏览器中打开应用
    hmr: {
      overlay: true // 在页面上显示错误和警告的遮罩层
    }
    /* proxy: {
      '/api': {
        target: 'http://127.0.0.1:4523/m1/3895832-0-default', // 这里填写你要代理的接口域名
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    } */
  },
  resolve: {
    alias: {
      '@': path.resolve('./src')
    }
  }
  // resolve: {
  //   alias: {
  //     '@': fileURLToPath(new URL('./src', import.meta.url))
  //   }
  // }
})
