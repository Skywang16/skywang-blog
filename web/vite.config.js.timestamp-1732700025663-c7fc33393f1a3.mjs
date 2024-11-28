// vite.config.js
import { defineConfig } from "file:///F:/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/web/node_modules/vite/dist/node/index.js";
import vue from "file:///F:/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/web/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import path from "path";
import Components from "file:///F:/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/web/node_modules/unplugin-vue-components/dist/vite.js";
import AutoImport from "file:///F:/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/web/node_modules/unplugin-auto-import/dist/vite.js";
import { VantResolver } from "file:///F:/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/web/node_modules/unplugin-vue-components/dist/resolvers.js";
import VueSetupExtend from "file:///F:/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/web/node_modules/vite-plugin-vue-setup-extend/dist/index.mjs";
import obfuscator from "file:///F:/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/web/node_modules/rollup-plugin-obfuscator/dist/rollup-plugin-obfuscator.js";
var vite_config_default = defineConfig({
  base: "./",
  // 这里更改打包相对绝对路径
  // base: '/', // 这里更改打包相对绝对路径
  build: {
    outDir: "dist",
    // 指定打包输出的目录为 dist
    rollupOptions: {
      output: {
        // 在这里修改静态资源路径
        chunkFileNames: "static/assets/js/[name]-[hash].js",
        entryFileNames: "static/assets/js/[name]-[hash].js",
        assetFileNames: "static/assets/[ext]/[name]-[hash].[ext]"
      }
    }
  },
  esbuild: {
    drop: ["console", "debugger"]
    // 打包去除
  },
  plugins: [
    vue(),
    VueSetupExtend(),
    // 配置自动导入
    AutoImport({
      imports: ["vue"],
      dts: "src/auto-imports.d.ts"
    }),
    // 配置按需导入
    Components({
      resolvers: [VantResolver()],
      dts: "src/components.d.ts"
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
        debugProtectionInterval: 2e3,
        disableConsoleOutput: true,
        domainLock: [],
        identifierNamesGenerator: "hexadecimal",
        identifiersPrefix: "",
        inputFileName: "",
        log: true,
        renameGlobals: true,
        reservedNames: [],
        reservedStrings: [],
        seed: 0,
        selfDefending: true,
        sourceMap: false,
        sourceMapBaseUrl: "",
        sourceMapFileName: "",
        sourceMapMode: "separate",
        stringArray: false,
        stringArrayEncoding: ["base64"],
        stringArrayThreshold: 0.75,
        target: "browser",
        transformObjectKeys: false,
        unicodeEscapeSequence: false,
        domainLockRedirectUrl: "about:blank",
        forceTransformStrings: [],
        identifierNamesCache: null,
        identifiersDictionary: [],
        ignoreImports: true,
        optionsPreset: "default",
        renameProperties: false,
        renamePropertiesMode: "safe",
        sourceMapSourcesMode: "sources-content",
        stringArrayCallsTransform: false,
        stringArrayCallsTransformThreshold: 0.5,
        stringArrayIndexesType: ["hexadecimal-number"],
        stringArrayIndexShift: false,
        stringArrayRotate: false,
        stringArrayWrappersCount: 1,
        stringArrayWrappersChainedCalls: false,
        stringArrayWrappersParametersMaxCount: 2,
        stringArrayWrappersType: "variable"
      }
    })
  ],
  css: {
    preprocessorOptions: {
      scss: {
        api: "modern-compiler"
      }
    }
  },
  server: {
    open: true,
    // 在服务器启动时自动在浏览器中打开应用
    hmr: {
      overlay: true
      // 在页面上显示错误和警告的遮罩层
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
      "@": path.resolve("./src")
    }
  }
  // resolve: {
  //   alias: {
  //     '@': fileURLToPath(new URL('./src', import.meta.url))
  //   }
  // }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcuanMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJGOlxcXFxcdTRFMkFcdTRFQkFcdTUzNUFcdTVCQTJcXFxcd2ViXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCJGOlxcXFxcdTRFMkFcdTRFQkFcdTUzNUFcdTVCQTJcXFxcd2ViXFxcXHZpdGUuY29uZmlnLmpzXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ltcG9ydF9tZXRhX3VybCA9IFwiZmlsZTovLy9GOi8lRTQlQjglQUElRTQlQkElQkElRTUlOEQlOUElRTUlQUUlQTIvd2ViL3ZpdGUuY29uZmlnLmpzXCI7aW1wb3J0IHsgZGVmaW5lQ29uZmlnIH0gZnJvbSAndml0ZSdcclxuaW1wb3J0IHZ1ZSBmcm9tICdAdml0ZWpzL3BsdWdpbi12dWUnXHJcbmltcG9ydCBwYXRoIGZyb20gJ3BhdGgnXHJcbi8vIFx1NUJGQ1x1NTE2NVx1ODFFQVx1NTJBOFx1NUJGQ1x1NTE2NVx1NTQ4Q1x1NjMwOVx1OTcwMFx1NUJGQ1x1NTE2NVx1NjNEMlx1NEVGNlxyXG5pbXBvcnQgQ29tcG9uZW50cyBmcm9tICd1bnBsdWdpbi12dWUtY29tcG9uZW50cy92aXRlJ1xyXG5pbXBvcnQgQXV0b0ltcG9ydCBmcm9tICd1bnBsdWdpbi1hdXRvLWltcG9ydC92aXRlJ1xyXG5pbXBvcnQgeyBWYW50UmVzb2x2ZXIgfSBmcm9tICd1bnBsdWdpbi12dWUtY29tcG9uZW50cy9yZXNvbHZlcnMnXHJcbmltcG9ydCBWdWVTZXR1cEV4dGVuZCBmcm9tICd2aXRlLXBsdWdpbi12dWUtc2V0dXAtZXh0ZW5kJ1xyXG5pbXBvcnQgb2JmdXNjYXRvciBmcm9tICdyb2xsdXAtcGx1Z2luLW9iZnVzY2F0b3InXHJcblxyXG5leHBvcnQgZGVmYXVsdCBkZWZpbmVDb25maWcoe1xyXG4gIGJhc2U6ICcuLycsIC8vIFx1OEZEOVx1OTFDQ1x1NjZGNFx1NjUzOVx1NjI1M1x1NTMwNVx1NzZGOFx1NUJGOVx1N0VERFx1NUJGOVx1OERFRlx1NUY4NFxyXG4gIC8vIGJhc2U6ICcvJywgLy8gXHU4RkQ5XHU5MUNDXHU2NkY0XHU2NTM5XHU2MjUzXHU1MzA1XHU3NkY4XHU1QkY5XHU3RUREXHU1QkY5XHU4REVGXHU1Rjg0XHJcbiAgYnVpbGQ6IHtcclxuICAgIG91dERpcjogJ2Rpc3QnLCAvLyBcdTYzMDdcdTVCOUFcdTYyNTNcdTUzMDVcdThGOTNcdTUxRkFcdTc2ODRcdTc2RUVcdTVGNTVcdTRFM0EgZGlzdFxyXG4gICAgcm9sbHVwT3B0aW9uczoge1xyXG4gICAgICBvdXRwdXQ6IHtcclxuICAgICAgICAvLyBcdTU3MjhcdThGRDlcdTkxQ0NcdTRGRUVcdTY1MzlcdTk3NTlcdTYwMDFcdThENDRcdTZFOTBcdThERUZcdTVGODRcclxuICAgICAgICBjaHVua0ZpbGVOYW1lczogJ3N0YXRpYy9hc3NldHMvanMvW25hbWVdLVtoYXNoXS5qcycsXHJcbiAgICAgICAgZW50cnlGaWxlTmFtZXM6ICdzdGF0aWMvYXNzZXRzL2pzL1tuYW1lXS1baGFzaF0uanMnLFxyXG4gICAgICAgIGFzc2V0RmlsZU5hbWVzOiAnc3RhdGljL2Fzc2V0cy9bZXh0XS9bbmFtZV0tW2hhc2hdLltleHRdJ1xyXG4gICAgICB9XHJcbiAgICB9XHJcbiAgfSxcclxuICBlc2J1aWxkOiB7XHJcbiAgICBkcm9wOiBbJ2NvbnNvbGUnLCAnZGVidWdnZXInXSAvLyBcdTYyNTNcdTUzMDVcdTUzQkJcdTk2NjRcclxuICB9LFxyXG4gIHBsdWdpbnM6IFtcclxuICAgIHZ1ZSgpLFxyXG4gICAgVnVlU2V0dXBFeHRlbmQoKSxcclxuICAgIC8vIFx1OTE0RFx1N0Y2RVx1ODFFQVx1NTJBOFx1NUJGQ1x1NTE2NVxyXG4gICAgQXV0b0ltcG9ydCh7XHJcbiAgICAgIGltcG9ydHM6IFsndnVlJ10sXHJcbiAgICAgIGR0czogJ3NyYy9hdXRvLWltcG9ydHMuZC50cydcclxuICAgIH0pLFxyXG4gICAgLy8gXHU5MTREXHU3RjZFXHU2MzA5XHU5NzAwXHU1QkZDXHU1MTY1XHJcbiAgICBDb21wb25lbnRzKHtcclxuICAgICAgcmVzb2x2ZXJzOiBbVmFudFJlc29sdmVyKCldLFxyXG4gICAgICBkdHM6ICdzcmMvY29tcG9uZW50cy5kLnRzJ1xyXG4gICAgfSksXHJcbiAgICBvYmZ1c2NhdG9yKHtcclxuICAgICAgZ2xvYmFsOiB0cnVlLFxyXG4gICAgICAvLyBvcHRpb25zXHU5MTREXHU3RjZFXHU5ODc5XHU1QjlFXHU5NjQ1XHU0RTNBIGphdmFzY3JpcHQtb2JmdXNjYXRvciBcdTkwMDlcdTk4NzlcdUZGMENcdTUxNzdcdTRGNTNcdTUzRUZcdTY3RTVcdTc3MEJodHRwczovL2Nsb3VkLnRlbmNlbnQuY29tL2RldmVsb3Blci9hcnRpY2xlLzIzNTQzNTdcclxuICAgICAgb3B0aW9uczoge1xyXG4gICAgICAgIGNvbXBhY3Q6IHRydWUsXHJcbiAgICAgICAgY29udHJvbEZsb3dGbGF0dGVuaW5nOiBmYWxzZSxcclxuICAgICAgICBjb250cm9sRmxvd0ZsYXR0ZW5pbmdUaHJlc2hvbGQ6IDAuNzUsXHJcbiAgICAgICAgbnVtYmVyc1RvRXhwcmVzc2lvbnM6IGZhbHNlLFxyXG4gICAgICAgIHNpbXBsaWZ5OiBmYWxzZSxcclxuICAgICAgICBzdHJpbmdBcnJheVNodWZmbGU6IGZhbHNlLFxyXG4gICAgICAgIHNwbGl0U3RyaW5nczogZmFsc2UsXHJcbiAgICAgICAgc3BsaXRTdHJpbmdzQ2h1bmtMZW5ndGg6IDEwLFxyXG4gICAgICAgIHJvdGF0ZVVuaWNvZGVBcnJheTogZmFsc2UsXHJcbiAgICAgICAgZGVhZENvZGVJbmplY3Rpb246IGZhbHNlLFxyXG4gICAgICAgIGRlYWRDb2RlSW5qZWN0aW9uVGhyZXNob2xkOiAwLjQsXHJcbiAgICAgICAgZGVidWdQcm90ZWN0aW9uOiB0cnVlLFxyXG4gICAgICAgIGRlYnVnUHJvdGVjdGlvbkludGVydmFsOiAyMDAwLFxyXG4gICAgICAgIGRpc2FibGVDb25zb2xlT3V0cHV0OiB0cnVlLFxyXG4gICAgICAgIGRvbWFpbkxvY2s6IFtdLFxyXG4gICAgICAgIGlkZW50aWZpZXJOYW1lc0dlbmVyYXRvcjogJ2hleGFkZWNpbWFsJyxcclxuICAgICAgICBpZGVudGlmaWVyc1ByZWZpeDogJycsXHJcbiAgICAgICAgaW5wdXRGaWxlTmFtZTogJycsXHJcbiAgICAgICAgbG9nOiB0cnVlLFxyXG4gICAgICAgIHJlbmFtZUdsb2JhbHM6IHRydWUsXHJcbiAgICAgICAgcmVzZXJ2ZWROYW1lczogW10sXHJcbiAgICAgICAgcmVzZXJ2ZWRTdHJpbmdzOiBbXSxcclxuICAgICAgICBzZWVkOiAwLFxyXG4gICAgICAgIHNlbGZEZWZlbmRpbmc6IHRydWUsXHJcbiAgICAgICAgc291cmNlTWFwOiBmYWxzZSxcclxuICAgICAgICBzb3VyY2VNYXBCYXNlVXJsOiAnJyxcclxuICAgICAgICBzb3VyY2VNYXBGaWxlTmFtZTogJycsXHJcbiAgICAgICAgc291cmNlTWFwTW9kZTogJ3NlcGFyYXRlJyxcclxuICAgICAgICBzdHJpbmdBcnJheTogZmFsc2UsXHJcbiAgICAgICAgc3RyaW5nQXJyYXlFbmNvZGluZzogWydiYXNlNjQnXSxcclxuICAgICAgICBzdHJpbmdBcnJheVRocmVzaG9sZDogMC43NSxcclxuICAgICAgICB0YXJnZXQ6ICdicm93c2VyJyxcclxuICAgICAgICB0cmFuc2Zvcm1PYmplY3RLZXlzOiBmYWxzZSxcclxuICAgICAgICB1bmljb2RlRXNjYXBlU2VxdWVuY2U6IGZhbHNlLFxyXG5cclxuICAgICAgICBkb21haW5Mb2NrUmVkaXJlY3RVcmw6ICdhYm91dDpibGFuaycsXHJcbiAgICAgICAgZm9yY2VUcmFuc2Zvcm1TdHJpbmdzOiBbXSxcclxuICAgICAgICBpZGVudGlmaWVyTmFtZXNDYWNoZTogbnVsbCxcclxuICAgICAgICBpZGVudGlmaWVyc0RpY3Rpb25hcnk6IFtdLFxyXG4gICAgICAgIGlnbm9yZUltcG9ydHM6IHRydWUsXHJcbiAgICAgICAgb3B0aW9uc1ByZXNldDogJ2RlZmF1bHQnLFxyXG4gICAgICAgIHJlbmFtZVByb3BlcnRpZXM6IGZhbHNlLFxyXG4gICAgICAgIHJlbmFtZVByb3BlcnRpZXNNb2RlOiAnc2FmZScsXHJcbiAgICAgICAgc291cmNlTWFwU291cmNlc01vZGU6ICdzb3VyY2VzLWNvbnRlbnQnLFxyXG5cclxuICAgICAgICBzdHJpbmdBcnJheUNhbGxzVHJhbnNmb3JtOiBmYWxzZSxcclxuICAgICAgICBzdHJpbmdBcnJheUNhbGxzVHJhbnNmb3JtVGhyZXNob2xkOiAwLjUsXHJcblxyXG4gICAgICAgIHN0cmluZ0FycmF5SW5kZXhlc1R5cGU6IFsnaGV4YWRlY2ltYWwtbnVtYmVyJ10sXHJcbiAgICAgICAgc3RyaW5nQXJyYXlJbmRleFNoaWZ0OiBmYWxzZSxcclxuICAgICAgICBzdHJpbmdBcnJheVJvdGF0ZTogZmFsc2UsXHJcbiAgICAgICAgc3RyaW5nQXJyYXlXcmFwcGVyc0NvdW50OiAxLFxyXG4gICAgICAgIHN0cmluZ0FycmF5V3JhcHBlcnNDaGFpbmVkQ2FsbHM6IGZhbHNlLFxyXG4gICAgICAgIHN0cmluZ0FycmF5V3JhcHBlcnNQYXJhbWV0ZXJzTWF4Q291bnQ6IDIsXHJcbiAgICAgICAgc3RyaW5nQXJyYXlXcmFwcGVyc1R5cGU6ICd2YXJpYWJsZSdcclxuICAgICAgfVxyXG4gICAgfSlcclxuICBdLFxyXG4gIGNzczoge1xyXG4gICAgcHJlcHJvY2Vzc29yT3B0aW9uczoge1xyXG4gICAgICBzY3NzOiB7XHJcbiAgICAgICAgYXBpOiAnbW9kZXJuLWNvbXBpbGVyJ1xyXG4gICAgICB9XHJcbiAgICB9XHJcbiAgfSxcclxuICBzZXJ2ZXI6IHtcclxuICAgIG9wZW46IHRydWUsIC8vIFx1NTcyOFx1NjcwRFx1NTJBMVx1NTY2OFx1NTQyRlx1NTJBOFx1NjVGNlx1ODFFQVx1NTJBOFx1NTcyOFx1NkQ0Rlx1ODlDOFx1NTY2OFx1NEUyRFx1NjI1M1x1NUYwMFx1NUU5NFx1NzUyOFxyXG4gICAgaG1yOiB7XHJcbiAgICAgIG92ZXJsYXk6IHRydWUgLy8gXHU1NzI4XHU5ODc1XHU5NzYyXHU0RTBBXHU2NjNFXHU3OTNBXHU5NTE5XHU4QkVGXHU1NDhDXHU4QjY2XHU1NDRBXHU3Njg0XHU5MDZFXHU3RjY5XHU1QzQyXHJcbiAgICB9XHJcbiAgICAvKiBwcm94eToge1xyXG4gICAgICAnL2FwaSc6IHtcclxuICAgICAgICB0YXJnZXQ6ICdodHRwOi8vMTI3LjAuMC4xOjQ1MjMvbTEvMzg5NTgzMi0wLWRlZmF1bHQnLCAvLyBcdThGRDlcdTkxQ0NcdTU4NkJcdTUxOTlcdTRGNjBcdTg5ODFcdTRFRTNcdTc0MDZcdTc2ODRcdTYzQTVcdTUzRTNcdTU3REZcdTU0MERcclxuICAgICAgICBjaGFuZ2VPcmlnaW46IHRydWUsXHJcbiAgICAgICAgcmV3cml0ZTogKHBhdGgpID0+IHBhdGgucmVwbGFjZSgvXlxcL2FwaS8sICcnKVxyXG4gICAgICB9XHJcbiAgICB9ICovXHJcbiAgfSxcclxuICByZXNvbHZlOiB7XHJcbiAgICBhbGlhczoge1xyXG4gICAgICAnQCc6IHBhdGgucmVzb2x2ZSgnLi9zcmMnKVxyXG4gICAgfVxyXG4gIH1cclxuICAvLyByZXNvbHZlOiB7XHJcbiAgLy8gICBhbGlhczoge1xyXG4gIC8vICAgICAnQCc6IGZpbGVVUkxUb1BhdGgobmV3IFVSTCgnLi9zcmMnLCBpbXBvcnQubWV0YS51cmwpKVxyXG4gIC8vICAgfVxyXG4gIC8vIH1cclxufSlcclxuIl0sCiAgIm1hcHBpbmdzIjogIjtBQUF5UCxTQUFTLG9CQUFvQjtBQUN0UixPQUFPLFNBQVM7QUFDaEIsT0FBTyxVQUFVO0FBRWpCLE9BQU8sZ0JBQWdCO0FBQ3ZCLE9BQU8sZ0JBQWdCO0FBQ3ZCLFNBQVMsb0JBQW9CO0FBQzdCLE9BQU8sb0JBQW9CO0FBQzNCLE9BQU8sZ0JBQWdCO0FBRXZCLElBQU8sc0JBQVEsYUFBYTtBQUFBLEVBQzFCLE1BQU07QUFBQTtBQUFBO0FBQUEsRUFFTixPQUFPO0FB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
