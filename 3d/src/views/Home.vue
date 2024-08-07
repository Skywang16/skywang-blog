<template>
  <div class="main">
    <canvas ref="backgroundCanvas" class="background-canvas"></canvas>
    <div ref="screenDom" class="canvas-container"></div>
    <div class="header">
      <div class="logo"></div>
      <div class="menu">
        <span @click="aaaaaa" class="menuItem">首页</span>
        <span @click="bbbbbb" class="menuItem">分类</span>
        <span @click="cccccc" class="menuItem">数据归档</span>
        <span class="menuItem" @mouseover="showMobBox = true" @mouseleave="showMobBox = false">
          设置
        </span>
        <div v-if="showMobBox" class="mob-box">微信搜索</div>
      </div>
    </div>
    <div class="pages" ref="pages">
      <div class="page">
      </div>
      <div class="page2">
        <wContent></wContent>
      </div>
    </div>
  </div>
</template>

<script>
import * as THREE from "three";
import { ref, onMounted } from "vue";
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader";
import { DRACOLoader } from "three/examples/jsm/loaders/DRACOLoader";
import { gsap } from "gsap";
import wContent from "./wContent.vue"
import { setupCanvasBackground, mouseClickToBC, mouseLeaveToBC } from './bcCanvas';

export default {
  components: {
    wContent
  },
  setup() {
    const showMobBox = ref(false);
    let screenDom = ref(null);// 定义引用，用于存储屏幕 DOM 元素
    let scene = new THREE.Scene();// 创建一个新的 Three.js 场景
    let camera;// 声明相机变量
    let renderer;// 声明渲染器变量
    let light;// 声明光源变量
    /* 绘制背景START */
    let backgroundCanvas = ref(null);// 定义引用，用于存储背景画布元素
    let pages = ref(null);
    /* 绘制背景END */
    let gltf1 = ref(null);
    // 初始化 ThreeJS
    const initThreeJS = () => {
      setupCamera(); // 设置相机
      setupRenderer(); // 设置渲染器
      setupLight(); // 设置光源
      loadModels(); // 加载模型
      animateModels(); // 动画模型
      //  setupCanvasBackground(backgroundCanvas, '', mouseClickToBC, mouseLeaveToBC);
    };
    let page = 0;
    let timeline2 = gsap.timeline();
    let isScroll = true;
    window.addEventListener("mousewheel", (e) => {
      if (e.wheelDelta < 0) {
        page++;
        if (isScroll) {
          isScroll = false
          setTimeout(() => {
            setupCanvasBackground(backgroundCanvas, '', mouseClickToBC, mouseLeaveToBC);
          }, 1000);
        }

        if (page > 1) {
          console.log(page, "到底了");
          page = 1
        }
        console.log(page)
      }
      if (e.wheelDelta > 0) {
        page--;
        if (page < 0) {
          page = 0;
        }
      }
      if (!timeline2.isActive()) {
        timeline2.to(camera.position, {
          duration: 1,
          y: page * -8,
        });
        gsap.to(pages.value, {
          duration: 1,
          y: -page * window.innerHeight,
        });
      }
    });
    // 设置相机
    const setupCamera = () => {
      camera = new THREE.PerspectiveCamera(45, window.innerWidth / window.innerHeight, 0.1, 10000);
      camera.position.z = 10;
    };
    // 设置渲染器
    const setupRenderer = () => {
      renderer = new THREE.WebGLRenderer({ antialias: true, });
      renderer.setSize(window.innerWidth, window.innerHeight);

      screenDom.value.appendChild(renderer.domElement);
    };
    // 设置光源
    const setupLight = () => {
      light = new THREE.DirectionalLight(0xffffff, 1);
      light.position.set(0, 0, 1);
      scene.add(light);
    };

    /* 模型事件----------------------------------------------------------------------START */
    // 加载模型
    const loadModels = () => {
      const dracoLoader = new DRACOLoader();
      dracoLoader.setDecoderPath("./draco/gltf/");
      const loader = new GLTFLoader().setDRACOLoader(dracoLoader);
      loader.load("./model/wang.glb", (gltf) => {
        console.log("模型加载成功:", gltf);
        setupModel(gltf, { scale: 4, position: { x: 0, y: -0.3, z: 0 } });
        animateModelRotation(gltf);
      }, undefined, (error) => {
        console.error("模型加载失败:", error);
      });
    };
    // 设置模型的缩放和位置
    const setupModel = (gltf, { scale, position }) => {
      gltf.scene.scale.set(scale, scale, scale);
      gltf.scene.position.set(position.x, position.y, position.z);
      scene.add(gltf.scene);
    };
    const animateModels = () => {
      requestAnimationFrame(animateModels);
      renderer.render(scene, camera);
    };
    // 动画模型
    const animateModelRotation = (gltf) => {
      gsap.to(gltf.scene.rotation, {
        duration: 5000,
        y: "+=360",
        repeat: -1,
        ease: "none",
      });
      gltf1.value = gltf
    };
    /* 模型事件----------------------------------------------------------------------END */

    /* 公共事件 ----------------------------------------------------------------------START*/
    // 事件触发函数
    const aaaaaa = () => {
      console.log('1-开始创建组件-setup')
      setupCanvasBackground(backgroundCanvas, 'block', mouseClickToBC, mouseLeaveToBC);
    };
    const bbbbbb = () => {
      setupCanvasBackground(backgroundCanvas, 'tetrapod', mouseClickToBC, mouseLeaveToBC);
    };
    const cccccc = () => {
      setupCanvasBackground(backgroundCanvas, 'circle', mouseClickToBC, mouseLeaveToBC);
    };

    // 挂载
    onMounted(() => {
      initThreeJS();
    });

    return {
      initThreeJS,
      screenDom,
      scene,
      camera,
      renderer,
      light,
      showMobBox,
      backgroundCanvas,
      aaaaaa,
      bbbbbb,
      cccccc,
      pages
    };
  }
}


</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
}

body {
  background-color: #000;
}

.canvas-container {
  width: 100%;
  height: 100%;

}

.header {
  background-color: rgba(0, 0, 0, 0.10);
  z-index: 10000;
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 80px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header .logo {
  height: 100px;
  width: 300px;
  background-image: url("@/assets/wanglogo.svg");
  background-size: 100%;
  background-position: center;
  background-repeat: no-repeat;
}


.menu {
  justify-content: space-between;
  align-items: center;
  margin-right: 50px;
}

.menuItem {
  padding: 0 10px;
  text-decoration: none;
  color: #000;
  font-weight: 600;
  font-size: 18px;
}

.menuItem:hover {
  cursor: pointer;
  /* 鼠标悬停时保持手指形状 */
  /* 可选：增加圆角 */
}

.titles {
  display: flex;
  flex-direction: column;
  position: fixed;
  top: 0;
  left: 0;
}

.titles .title {
  width: 100vw;
  height: 100vh;
  font-weight: bold;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  color: #fff;
  padding: 15%;
  box-sizing: border-box;
  font-size: 50px;
}

.content {
  /* background-color: #fff; */
  overflow: -moz-scrollbars-none;
  /* Firefox */
  -ms-overflow-style: none;
  /* IE 10+ */
  scrollbar-width: none;
  z-index: 99;
  /*  border-radius: 10px; */
}

.content-box {
  margin: auto;
  height: 85vh;
  width: calc(96vw - 12px);
}

.switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 34px;
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #303136;
  -webkit-transition: .4s;
  transition: .4s;
  border-radius: 34px;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background: #303136;
  box-shadow: inset -3px -2px 5px -2px #8983f7, inset -10px -4px 0 0 #a3dafb;
  border-radius: 50%;
  -webkit-transition: .4s;
  transition: .4s;
}

input:checked+.slider {
  background-color: #ccc;
}

input:checked+.slider:before {
  -webkit-transform: translateX(26px);
  -ms-transform: translateX(26px);
  transform: translateX(26px);
  background: linear-gradient(40deg, #ff0080, #ff8c00 70%);

}

/* 可选的样式，用于指示滑块被按下时的阴影效果 */
.slider:active:before {
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15) inset;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 3s ease-in-out;
  /* 添加 'ease-in-out' */
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.fade-enter-to,
.fade-leave-from {
  opacity: 1;
}

.background-canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 3;
  /* 确保背景在所有内容的后面 */
}

.hexagon {
  transition: transform 0.5s ease-in-out;
}

.pages {
  display: flex;
  flex-direction: column;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 3;
}

.pages .page {
  width: 100vw;
  height: 100vh;
  flex-direction: column;
}

.pages .page2 {
  background-color: transparent;
  padding: 100px 0 40px 40px;
}
</style>
