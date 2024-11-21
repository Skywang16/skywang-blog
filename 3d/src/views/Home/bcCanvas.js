// src/backgroundCanvas.js
import arrowSrc from '@/assets/bgImg/arrow.png'
import blockSrc from '@/assets/bgImg/block.png'
import circleSrc from '@/assets/bgImg/circle.png'
import tetrapodSrc from '@/assets/bgImg/tetrapod.png'
import triangleSrc from '@/assets/bgImg/triangle.png'
let currentActiveShape = 'block'
let oldShape = ''
const shapesConfig = [
  { src: arrowSrc, type: 'arrow' },
  { src: blockSrc, type: 'block' },
  { src: circleSrc, type: 'circle' },
  { src: tetrapodSrc, type: 'tetrapod' },
  { src: triangleSrc, type: 'triangle' }
]
const mousePosition = { x: -100, y: -100 }
let isMouseMoveThrottled = false
let backgroundWorker = null
let loadedImages = null

// 初始化worker
const initWorker = () => {
  if (window.Worker) {
    backgroundWorker = new Worker(new URL('@/workers/backgroundWorker.js', import.meta.url))
  }
}

// 新增加载图片的函数
const loadImages = () => {
  if (loadedImages) return Promise.resolve(loadedImages)

  const images = shapesConfig.map((shape) => {
    const img = new Image()
    img.src = shape.src
    img.shape = shape.type
    return img
  })

  return Promise.all(
    images.map(
      (img) =>
        new Promise((resolve) => {
          img.onload = () => resolve(img)
        })
    )
  ).then((loadedImgs) => {
    loadedImages = loadedImgs
    return loadedImgs
  })
}

// 背景初始化
export const setupCanvasBackground = (backgroundCanvas, imgType) => {
  let context
  const canvas = backgroundCanvas.value

  if (!imgType) {
    context = canvas.getContext('2d')
    if (!backgroundWorker) {
      initWorker()
    }
  } else {
    oldShape = currentActiveShape
    currentActiveShape = imgType
    if (backgroundWorker && currentActiveShape !== oldShape) {
      backgroundWorker.postMessage({
        type: 'shapeSwitch',
        data: {
          currentActiveShape: imgType,
          oldShape,
          timestamp: Date.now()
        }
      })
    }
    return // 如果只是切换图形，直接返回
  }

  // 只在第一次初始化时加载图片
  loadImages().then((images) => {
    canvas.width = window.innerWidth
    canvas.height = window.innerHeight

    if (!imgType) {
      animateBackgroundColor(context, images, canvas)
    }

    if (!document.querySelector('[data-mousemove-listener]')) {
      document.addEventListener('mousemove', mouseClickToBC)
      document.body.setAttribute('data-mousemove-listener', 'true')
    }
  })
}
let hue1 = 57
let hue2 = 148
// 绘制彩色背景
const renderBackground = (context, canvas) => {
  const gradient = context.createLinearGradient(0, 0, canvas.width, canvas.height)
  gradient.addColorStop(0, `hsl(${hue1}, 100%, 70%)`)
  gradient.addColorStop(1, `hsl(${hue2}, 100%, 70%)`)
  context.fillStyle = gradient
  context.fillRect(0, 0, canvas.width, canvas.height)
  hue1 += 0.1 // 适当调整颜色变化范围
  hue2 += 0.1
}
// 动画
const animateBackgroundColor = (context, images, canvas) => {
  let lastTime = 0
  const fps = 60
  const interval = 1000 / fps
  const offscreenCanvas = document.createElement('canvas')
  const offscreenContext = offscreenCanvas.getContext('2d')
  offscreenCanvas.width = canvas.width
  offscreenCanvas.height = canvas.height

  const loop = (timestamp) => {
    if (timestamp - lastTime >= interval) {
      lastTime = timestamp

      // 在离屏画布上渲染背景
      renderBackground(offscreenContext, offscreenCanvas)

      // 发送数据给worker计算
      if (backgroundWorker) {
        backgroundWorker.postMessage({
          type: 'calculateShapes',
          data: {
            canvas: {
              width: canvas.width,
              height: canvas.height
            },
            currentActiveShape,
            oldShape,
            imgConfig: images.map((img) => ({
              width: img.width,
              height: img.height,
              shape: img.shape
            })),
            timestamp: Date.now()
          }
        })
      }
    }
    requestAnimationFrame(loop)
  }

  // 监听worker返回的结果
  if (backgroundWorker) {
    backgroundWorker.onmessage = function (e) {
      if (e.data.type === 'shapeResults') {
        const shapes = e.data.data
        context.clearRect(0, 0, canvas.width, canvas.height)
        context.drawImage(offscreenCanvas, 0, 0)

        shapes.forEach((shape) => {
          const img = loadedImages.find((img) => img.shape === shape.type)
          if (shape.type === 'block') {
            context.drawImage(img, shape.x, shape.y, shape.width, shape.height)
          } else if (shape.type === 'tetrapod') {
            context.save()
            const hexRadius = shape.width / 2
            context.translate(shape.x + hexRadius, shape.y + shape.height / 2)
            context.rotate(shape.rotation)
            context.scale(shape.width / (img.width * 0.5), shape.height / (img.height * 0.5))
            context.drawImage(img, -hexRadius, -shape.height / 2, img.width * 0.5, img.height * 0.5)
            context.restore()
          }
        })
      }
    }
  }

  requestAnimationFrame(loop)
}
// 获取鼠标位置
export const mouseClickToBC = (event) => {
  if (!isMouseMoveThrottled) {
    isMouseMoveThrottled = true
    setTimeout(() => {
      const position = {
        x: event.clientX,
        y: event.clientY
      }
      mousePosition.x = position.x
      mousePosition.y = position.y

      // 发送鼠标位置到worker
      if (backgroundWorker) {
        backgroundWorker.postMessage({
          type: 'updateMousePosition',
          data: position
        })
      }

      isMouseMoveThrottled = false
    }, 16)
  }
}
