<template>
  <div class="main" ref="mainRef" @scroll="handleScroll">
    <canvas ref="backgroundCanvas" class="background-canvas"></canvas>
    <header class="header">
      <div class="logo"></div>
      <div class="menu">
        <span @click="goHome" class="menuItem">
          <img src="@/assets/icons/home.png" alt="">
          首页
        </span>
        <span @click="goAlbum" class="menuItem">
          <img src="@/assets/icons/album.png" alt="">
          相册
        </span>
      </div>
    </header>
    <div class="page-container">
      <div class="threeDom-box" v-wLoading="{ loading: showLoading }" @dblclick.stop="handleDoubleClick">
        <div ref="threeDom" class="canvas-container"></div>
      </div>
      <div class="waves-box">
        <div>
          <svg class="waves" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
            viewBox="0 24 150 28" preserveAspectRatio="none" shape-rendering="auto">
            <defs>
              <path id="gentle-wave" d="M-160 44c30 0 58-18 88-18s 58 18 88 18 58-18 88-18 58 18 88 18 v44h-352z" />
            </defs>
            <g class="parallax">
              <use xlink:href="#gentle-wave" x="48" y="0" fill="rgba(255,255,255,0.7)" />
              <use xlink:href="#gentle-wave" x="48" y="3" fill="rgba(255,255,255,0.5)" />
              <use xlink:href="#gentle-wave" x="48" y="5" fill="rgba(255,255,255,0.3)" />
              <use xlink:href="#gentle-wave" x="48" y="7" fill="#fff" />
            </g>
          </svg>
        </div>
      </div>
      <div class="wContent">
        <wContent v-show="homeShow"></wContent>
        <album v-show="albumShow"></album>
      </div>
      <footer class="footer">
        <div class="footer-content">
          <p>
            <span class="miniapp"> 微信小程序：skyWang416</span>
            <span class="divider">|</span>
            <a href="https://beian.miit.gov.cn/" target="_blank" class="icp-info">浙ICP备2024096478号</a>
          </p>
        </div>
      </footer>
    </div>
    <Transition name="fade-slide">
      <div class="back-to-top" v-if="showBackToTop" @click="scrollToTop">
        <div class="back-to-top-content">
          <img style="height: 30px;" src="@/assets/icons/up.png" alt="" />
          <span class="tooltip">返回顶部</span>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script>
import * as THREE from "three";
import { ref, onMounted } from "vue";
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader";
import { DRACOLoader } from "three/examples/jsm/loaders/DRACOLoader";
import wContent from "../components/wContent.vue"
import { gsap } from "gsap";
import { setupCanvasBackground, mouseClickToBC } from './bcCanvas';
import album from "../components/album.vue"
export default {
  components: {
    wContent,
    album,
  },
  name: 'HomeView',
  setup() {
    const showMobBox = ref(false);
    let threeDom = ref(null);// 定义引用，用于存储屏幕 DOM 元素
    let scene = new THREE.Scene();// 创建一个新的 Three.js 场景
    let camera;// 声明相机变量
    let renderer;// 声明渲染器变量
    let light;// 声明光源变量
    let gltf;
    /* 绘制背景START */
    const backgroundCanvas = ref(null);// 定义引用，用于存储背景画布元素
    const pages = ref(null);
    // 自定义遮罩相关
    const showLoading = ref(false); // 自定义遮罩显影
    const homeShow = ref(true);
    const albumShow = ref(false);
    let zoomIn = true; // 用于切换放大和缩小
    /* 绘制背景END */
    const mainRef = ref(null);
    const showBackToTop = ref(false);
    // 初始化 ThreeJS
    const initThreeJS = () => {
      setupCamera(); // 设置相机
      setupRenderer(); // 设置渲染器
      setupLight(); // 设置光源
      loadModels(); // 加载模型
      animateModels(); // 动画模型
      setupCanvasBackground(backgroundCanvas, '', mouseClickToBC);
    };
    // 设置相机
    const setupCamera = () => {
      camera = new THREE.PerspectiveCamera(45, window.innerWidth / window.innerHeight, 0.1, 10000);
      camera.position.z = 10;
    };
    // 缩放相机
    const handleDoubleClick = () => {
      const zoomAmount = 6; // 双击时的缩放量
      const targetZoom = zoomIn ? camera.position.z - zoomAmount : camera.position.z + zoomAmount;
      gsap.to(camera.position, {
        z: targetZoom,
        duration: 1, // 动画持续时间
        ease: "power2.inOut" // 缓动效果
      });
      zoomIn = !zoomIn; // 切换放大和缩小状态
    };
    // 设置渲染器
    const setupRenderer = () => {
      renderer = new THREE.WebGLRenderer({ antialias: true });
      const container = threeDom.value;
      renderer.setSize(window.innerWidth, container.clientHeight); // 使用 canvas-container 的尺寸
      // renderer.setSize(window.innerWidth, window.innerHeight);
      renderer.setClearColor(0x000000, 0); // 设置背景颜色为透明（或其他颜色）
      container.appendChild(renderer.domElement);
      // container.addEventListener('wheel', zoomCamera); // 添加滚轮事件监听器
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
      showLoading.value = true; // 开始加载时显示加载状态
      const dracoLoader = new DRACOLoader();
      dracoLoader.setDecoderPath("./draco/gltf/");
      const loader = new GLTFLoader().setDRACOLoader(dracoLoader);
      loader.load("./model/wang.glb",
        (gltf) => {
          console.log("模型加载成功:", gltf);
          setupModel(gltf, { scale: 3, position: { x: 0, y: 0.2, z: 0 } });
          animateModelRotation(gltf);
          showLoading.value = false; // 加载完成后隐藏加载状态
        },
        undefined,
        (error) => {
          console.error("模型加载失败:", error);
        }
      );
    };
    // 设置模型的缩放和位置
    const setupModel = (gltf, { scale, position }) => {
      gltf.scene.scale.set(scale, scale, scale);
      gltf.scene.position.set(position.x, position.y, position.z);
      scene.add(gltf.scene);
    };
    const animateModels = () => {
      if (renderer && camera) {
        requestAnimationFrame(animateModels);
        renderer.render(scene, camera);
      }
    };
    // 动画模型
    const animateModelRotation = (gltf1) => {
      // 初始化自动旋转动画
      let autoRotateTween;
      const startAutoRotate = () => {
        autoRotateTween = gsap.to(gltf1.scene.rotation, {
          duration: 5000,
          y: "+=360",
          repeat: -1,
          ease: "none",
        });
      };
      startAutoRotate(); // 初始开始自动旋转
      // 保存当前使用的 glTF 对象
      gltf = gltf1;
      let isMouseDown = false;
      let startX = 0,
        startY = 0;

      // 将mousedown事件仍然绑定在threeDom上，因为我们只想在3D模型区域开始拖动
      threeDom.value.addEventListener("mousedown", (e) => {
        isMouseDown = true;
        startX = e.clientX;
        startY = e.clientY;
        // 暂停自动旋转
        autoRotateTween.pause();
      });

      // 将mouseup和mousemove事件绑定到window上
      window.addEventListener("mouseup", () => {
        if (!gltf) return;
        isMouseDown = false;
        // 获取当前的旋转角度
        const currentRotationY = gltf.scene.rotation.y;
        // 重新启动自动旋转动画并从当前角度开始
        autoRotateTween.kill(); // 杀掉旧的动画
        autoRotateTween = gsap.to(gltf.scene.rotation, {
          duration: 5000,
          y: currentRotationY + 360,
          repeat: -1,
          ease: "none",
        });
      });

      window.addEventListener("mousemove", (e) => {
        if (!isMouseDown || !gltf) return;
        // 计算鼠标移动的偏移量
        let offsetX = (startX - e.clientX) / window.innerWidth * 2;
        let offsetY = (startY - e.clientY) / window.innerHeight * 2;
        // 更新对象的旋转
        gsap.set(gltf.scene.rotation, {
          x: gltf.scene.rotation.x + offsetY,
          y: gltf.scene.rotation.y + offsetX,
        });
        // 更新开始位置
        startX = e.clientX;
        startY = e.clientY;
      });
    };
    /* 模型事件----------------------------------------------------------------------END */

    /* 数据请求 ----------------------------------------------------------------------START*/
    /* 数据请求----------------------------------------------------------------------END */

    /* 公共事件 ----------------------------------------------------------------------START*/
    // 事件触发函数
    const goHome = () => {
      homeShow.value = true;
      albumShow.value = false;
      setupCanvasBackground(backgroundCanvas, 'block', mouseClickToBC);
    };
    const goAlbum = () => {
      homeShow.value = false;
      albumShow.value = true;
      setupCanvasBackground(backgroundCanvas, 'tetrapod', mouseClickToBC);
    };

    // 滚动监听方法
    const handleScroll = () => {
      if (mainRef.value) {
        showBackToTop.value = mainRef.value.scrollTop > 300;
      }
    };

    // 返回顶部方法
    const scrollToTop = () => {
      if (mainRef.value) {
        mainRef.value.scrollTo({
          top: 0,
          behavior: 'smooth'
        });
      }
    };

    onMounted(() => {
      initThreeJS();
    });

    return {
      threeDom,
      showMobBox,
      backgroundCanvas,
      goHome,
      goAlbum,
      handleDoubleClick,
      pages,
      showLoading,
      homeShow,
      albumShow,
      mainRef,
      showBackToTop,
      scrollToTop,
      handleScroll,
    };
  }
}
</script>

<style scoped lang='scss'>
.main {
  display: flex;
  overflow-y: auto;
  justify-content: center;
  height: 100vh;
  user-select: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
}

::-webkit-scrollbar {
  display: none;
}

.canvas-container {
  height: 550px;
  width: 100%;
  display: flex;
  justify-content: center;
  z-index: 2;
}

.page-container {
  position: relative;
  z-index: 2;

  .threeDom-box {
    background-color: rgb(130, 112, 202);
    user-select: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
  }

  .waves-box {
    position: relative;
    bottom: 100px;
    text-align: center;
    color: white;

    &::after {
      content: '';
      position: absolute;
      bottom: -100px;
      left: 0;
      width: 100%;
      height: 100px;
      background: linear-gradient(to bottom, #fff, transparent);
      pointer-events: none;
    }
  }

  .waves {
    position: relative;
    width: 100%;
    height: 10vh;
    margin-bottom: -7px;
    min-height: 100px;
    max-height: 150px;
  }

  .content {
    position: relative;
    height: 20vh;
    text-align: center;
    background-color: white;
  }

  .content a {
    margin: 0 5px;
    font-size: 12px;
    color: #333;
  }

  .content a:hover {
    color: #000;
  }

  /* Animation */
  .parallax>use {
    animation: move-forever 25s cubic-bezier(0.55, 0.5, 0.45, 0.5) infinite;
  }

  .parallax>use:nth-child(1) {
    animation-delay: -2s;
    animation-duration: 7s;
  }

  .parallax>use:nth-child(2) {
    animation-delay: -3s;
    animation-duration: 10s;
  }

  .parallax>use:nth-child(3) {
    animation-delay: -4s;
    animation-duration: 13s;
  }

  .parallax>use:nth-child(4) {
    animation-delay: -5s;
    animation-duration: 20s;
  }

  @media (max-width: 768px) {
    .waves {
      height: 40px;
      min-height: 40px;
    }

    .content {
      height: 30vh;
    }

    h1 {
      font-size: 24px;
    }
  }

  .wContent {
    position: relative;
    bottom: 50px;
    width: 100%;
    padding: 0 15%;
  }
}

@keyframes move-forever {
  0% {
    transform: translate3d(-90px, 0, 0);
  }

  100% {
    transform: translate3d(85px, 0, 0);
  }
}

.header {
  background-color: transparent;
  z-index: 4;
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 60px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: rgba(255, 255, 255, 0.8);
}

.header .logo {
  height: 100px;
  width: 200px;
  background-image: url("@/assets/wanglogo.svg");
  background-size: 100%;
  background-position: center;
  background-repeat: no-repeat;
}

.footer {
  background: linear-gradient(135deg, #005bea 0%, #00c6fb 100%);
  color: #fff;
  padding: 12px 0;
  position: relative;
  width: 100%;
  z-index: 4;

  .footer-content {
    text-align: center;

    p {
      margin: 0;
      color: #ecf0f1;
      font-size: 0.9rem;
      display: flex;
      align-items: center;
      justify-content: center;
      gap: 15px;

      .miniapp {
        i {
          margin-right: 5px;
        }
      }

      .divider {
        color: rgba(255, 255, 255, 0.5);
      }

      .icp-info {
        color: #ecf0f1;
        text-decoration: none;
        transition: all 0.3s ease;

        &:hover {
          color: #3498db;
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .footer {
    padding: 8px 0;

    .footer-content p {
      flex-direction: column;
      gap: 5px;

      .divider {
        display: none;
      }
    }
  }
}

.menu {
  justify-content: space-between;
  align-items: center;
  margin-right: 20px;
}

.menuItem {
  padding: 0 10px;
  text-decoration: none;
  color: #000;
  font-weight: 600;
  font-size: 18px;
  cursor: pointer !important;

  img {
    margin-right: 5px;
    height: 25px;
  }
}

.background-canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1;
  /* 确保背景在有内容的后面 */
}

.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.3s ease;
}

.fade-slide-enter-from,
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(20px);
}

.back-to-top {
  position: fixed;
  right: 40px;
  bottom: 100px;
  width: 50px;
  height: 50px;
  background: linear-gradient(135deg, #00c6fb 0%, #005bea 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer !important;
  z-index: 999;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.15);
  transition: all 0.3s ease;

  .back-to-top-content {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 100%;

    .tooltip {
      position: absolute;
      right: 140%;
      background: rgba(0, 0, 0, 0.8);
      color: white;
      padding: 5px 10px;
      border-radius: 4px;
      font-size: 14px;
      opacity: 0;
      visibility: hidden;
      transition: opacity 0.3s ease;
      white-space: nowrap;
      pointer-events: none;

      &::after {
        content: '';
        position: absolute;
        right: -6px;
        top: 50%;
        transform: translateY(-50%);
        border-left: 6px solid rgba(0, 0, 0, 0.8);
        border-top: 6px solid transparent;
        border-bottom: 6px solid transparent;
      }
    }
  }

  &:hover {
    transform: scale(1.05);
    box-shadow: 0 6px 20px rgba(0, 98, 234, 0.4);
    background: linear-gradient(135deg, #00a1fb 0%, #0046ea 100%);

    .tooltip {
      opacity: 1;
      visibility: visible;
    }
  }

  &:active {
    transform: scale(0.95);
  }
}

@keyframes rocketShake {

  0%,
  100% {
    transform: rotate(-45deg) translate(0, 0);
  }

  25% {
    transform: rotate(-40deg) translate(-2px, -2px);
  }

  75% {
    transform: rotate(-50deg) translate(2px, 2px);
  }
}

/* 移动端适配 */
@media (max-width: 768px) {
  .back-to-top {
    width: 40px;
    height: 40px;

    .back-to-top-content {
      i {
        font-size: 20px;
      }
    }

    .tooltip {
      display: none; // 移动端不显示提示文字
    }
  }
}
</style>
