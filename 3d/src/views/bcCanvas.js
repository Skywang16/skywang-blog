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
let hue1 = 0
let hue2 = 60
const mousePosition = { x: -100, y: -100 }
const hexMovementProperties = {}
let isOldShapeActive = false // 标志位，表示是否处于oldShape的活跃期
let oldShapeActiveStartTime = 0 // 记录oldShape活跃期开始的时间戳
// 背景初始化
export const setupCanvasBackground = (backgroundCanvas, imgType) => {
  let context
  const canvas = backgroundCanvas.value
  if (!imgType) {
    context = canvas.getContext('2d')
  } else {
    oldShape = currentActiveShape
    currentActiveShape = imgType
    isOldShapeActive = false
    oldShapeActiveStartTime = null
  }
  // 加载所有图案
  const images = shapesConfig.map((shape) => {
    const img = new Image()
    img.src = shape.src
    img.shape = shape.type // 保存形状类型
    return img
  })
  Promise.all(
    images.map(
      (img) =>
        new Promise((resolve) => {
          img.onload = resolve
        })
    )
  ).then(() => {
    canvas.width = window.innerWidth
    canvas.height = window.innerHeight
    animateBackgroundColor(context, images, canvas)

    document.addEventListener('mousemove', mouseClickToBC)
    document.addEventListener('mouseleave', mouseLeaveToBC)
  })
}
// 绘制彩色背景
const renderBackground = (context, canvas) => {
  const gradient = context.createLinearGradient(0, 0, canvas.width, canvas.height)
  gradient.addColorStop(0, `hsl(${hue1}, 100%, 60%)`) // 更高的饱和度和亮度
  gradient.addColorStop(1, `hsl(${hue2}, 100%, 60%)`) // 更高的饱和度和亮度
  context.fillStyle = gradient
  context.fillRect(0, 0, canvas.width, canvas.height)
  hue1 += 0.05 // 适当调整颜色变化范围
  hue2 += 0.05
}
const animateBackgroundColor = (context, images, canvas) => {
  requestAnimationFrame(function loop() {
    renderBackground(context, canvas)
    renderHexagons(context, images, canvas)
    requestAnimationFrame(loop)
  })
}
// 渲染图形
const renderHexagons = (context, images, canvas) => {
  // 渲染逻辑：绘制当前图形，当切换时同时绘制新的图形并且保留绘制老图形0.6秒，以便渲染切换动画
  const now = Date.now() // 获取当前时间戳
  images.forEach((img) => {
    /* ----------------------------------------------------block----------------------------------------------------------- */
    if (img.shape === currentActiveShape && currentActiveShape === 'block') {
      // console.log('1')
      renderBlock(context, canvas, img)
    }
    if (img.shape === oldShape && oldShape === 'block') {
      if (!isOldShapeActive) {
        isOldShapeActive = true
        oldShapeActiveStartTime = now
        // console.log('2')
        renderBlock(context, canvas, img)
      } else if (now - oldShapeActiveStartTime < 600) {
        // console.log('2')
        renderBlock(context, canvas, img)
      }
    }
    /* ----------------------------------------------------tetrapod----------------------------------------------------------- */
    if (img.shape === currentActiveShape && currentActiveShape === 'tetrapod') {
      // console.log('3')
      renderTetrapod(context, canvas, img)
    }
    if (img.shape === oldShape && oldShape === 'tetrapod') {
      if (!isOldShapeActive) {
        isOldShapeActive = true
        oldShapeActiveStartTime = now
        // console.log('4')
        renderTetrapod(context, canvas, img)
      } else if (now - oldShapeActiveStartTime < 600) {
        // console.log('4')
        renderTetrapod(context, canvas, img)
      }
    }
  })
}
// 方块渲染逻辑
const renderBlock = (context, canvas, img) => {
  const hexHeight = img.height * 0.5 // 六边形高度
  const hexWidth = img.width * 0.5 // 六边形宽度
  const hexRadius = hexWidth / 2 // 六边形半径
  const hexVerticalSpacing = hexHeight * 0.75 // 内径
  const hexHorizontalSpacing = hexWidth // 间距
  const xOffsetBase = -img.width * 0.25 // 左移
  const yOffsetBase = -img.height * 0.5 // 上移
  // 遍历绘制六边形
  for (let y = 0; y < canvas.height + hexHeight; y += hexVerticalSpacing) {
    for (let x = 0; x < canvas.width + hexWidth; x += hexHorizontalSpacing) {
      const xOffset = ((y / hexVerticalSpacing) % 2 === 0 ? 0 : hexRadius) + xOffsetBase
      const yOffset = y + yOffsetBase
      const hexKey = `${Math.round(x + xOffset)}_${Math.round(yOffset)}`
      // 初始化六边形属性
      if (!hexMovementProperties[hexKey]) {
        hexMovementProperties[hexKey] = {
          currentY: yOffset,
          targetY: yOffset,
          currentWidth: 0,
          currentHeight: 0,
          targetWidth: hexWidth,
          targetHeight: hexHeight,
          timer: null, // 用于存储计时器
          switchTimer: null,
          shapeType: img.shape
        }
      }
      const hexProperty = hexMovementProperties[hexKey]
      const isCurrentActive = hexProperty.shapeType === currentActiveShape
      // 处理鼠标悬停逻辑
      if (isHovering(x + xOffset, yOffset, hexWidth, hexHeight)) {
        hexProperty.targetY = yOffset - 100 // 上移100px
        // 清除恢复计时器
        if (hexProperty.timer) {
          clearTimeout(hexProperty.timer)
          hexProperty.timer = null // 重置计时器
        }
      } else {
        if (!hexProperty.timer) {
          hexProperty.timer = setTimeout(() => {
            hexProperty.targetY = yOffset // 恢复到原位置
            hexProperty.timer = null // 清除计时器
          }, 200)
        }
      }
      // 处理当前图案缩小和新图案放大的逻辑
      if (isCurrentActive && hexProperty.currentWidth >= 0 && hexProperty.currentHeight >= 0) {
        hexProperty.switchTimer = setTimeout(() => {
          hexProperty.targetWidth = hexWidth // 目标宽度
          hexProperty.targetHeight = hexHeight // 目标高度
          hexProperty.switchTimer = null // 清除计时器
        }, 400)
      } else {
        // 如果当前不是选中状态，先缩小到0
        if (hexProperty.currentWidth > 0 && hexProperty.currentHeight > 0) {
          hexProperty.targetWidth = 0
          hexProperty.targetHeight = 0
        }
      }
      // 使用缓动函数计算动画进度
      const progressY = easeInOutQuad(0.06)
      const progressWidth = easeInOutQuad(0.03)
      const progressHeight = easeInOutQuad(0.03)
      hexProperty.currentY += (hexProperty.targetY - hexProperty.currentY) * progressY
      hexProperty.currentWidth +=
        (hexProperty.targetWidth - hexProperty.currentWidth) * progressWidth
      hexProperty.currentHeight +=
        (hexProperty.targetHeight - hexProperty.currentHeight) * progressHeight
      if (hexProperty.currentWidth > 0 && hexProperty.currentHeight > 0) {
        context.drawImage(
          img,
          x + xOffset,
          hexProperty.currentY,
          hexProperty.currentWidth,
          hexProperty.currentHeight
        )
      }
    }
  }
}
// 三叶渲染逻辑
const renderTetrapod = (context, canvas, img) => {
  const hexHeight = img.height * 0.5 // 六边形高度
  const hexWidth = img.width * 0.5 // 六边形宽度
  const hexRadius = hexWidth / 2 // 六边形半径
  const hexVerticalSpacing = hexHeight * 0.73 // 内径
  const hexHorizontalSpacing = hexWidth // 间距
  const xOffsetBase = -img.width * 0.5 // 左移
  const yOffsetBase = -img.height * 0.5 // 上移
  // 遍历绘制六边形
  for (let y = 0; y < canvas.height + hexHeight; y += hexVerticalSpacing) {
    for (let x = 0; x < canvas.width + hexWidth; x += hexHorizontalSpacing * 0.73) {
      const xOffset = ((y / hexVerticalSpacing) % 2 === 0 ? 0 : hexRadius * 0.73) + xOffsetBase
      const yOffset = y + yOffsetBase
      const hexKey = `${Math.round(x + xOffset)}_${Math.round(yOffset)}`
      // 初始化六边形属性
      if (!hexMovementProperties[hexKey]) {
        hexMovementProperties[hexKey] = {
          currentY: yOffset,
          targetY: yOffset,
          currentWidth: 0,
          currentHeight: 0,
          targetWidth: 0,
          targetHeight: 0,
          rotation: 0, // 当前旋转角度
          targetRotation: 0, // 目标旋转角度
          shapeType: img.shape,
          hasMoved: false, // 标记六边形是否已旋转
          switchTimer: null // 标记是否正在缩小
        }
      }
      const hexProperty = hexMovementProperties[hexKey]
      const isCurrentActive = hexProperty.shapeType === currentActiveShape
      // 处理当前图案缩小和新图案放大的逻辑
      if (isCurrentActive && hexProperty.currentWidth >= 0 && hexProperty.currentHeight >= 0) {
        hexProperty.switchTimer = setTimeout(() => {
          hexProperty.targetWidth = hexWidth // 目标宽度
          hexProperty.targetHeight = hexHeight // 目标高度
          hexProperty.switchTimer = null // 清除计时器
        }, 400)
      } else {
        // 如果当前不是选中状态，先缩小到0
        if (hexProperty.currentWidth > 0 && hexProperty.currentHeight > 0) {
          hexProperty.targetWidth = 0
          hexProperty.targetHeight = 0
        }
      }
      // 处理鼠标悬停逻辑
      if (isHovering(x + xOffset, yOffset, hexWidth, hexHeight)) {
        // 一旦鼠标悬停，立即设置目标旋转角度为120度
        if (!hexProperty.hasMoved) {
          hexProperty.targetRotation += -Math.PI / 1.5 // 逆时针旋转120度
          hexProperty.hasMoved = true // 标记为旋转
        }
      } else {
        hexProperty.hasMoved = false // 标记为未旋转
      }
      const progressRotation = easeInOutQuad(0.06)
      const progressWidth = easeInOutQuad(0.03)
      const progressHeight = easeInOutQuad(0.03)
      hexProperty.currentWidth +=
        (hexProperty.targetWidth - hexProperty.currentWidth) * progressWidth
      hexProperty.currentHeight +=
        (hexProperty.targetHeight - hexProperty.currentHeight) * progressHeight
      hexProperty.rotation += (hexProperty.targetRotation - hexProperty.rotation) * progressRotation // 更新旋转角度
      // 绘制六边形
      if (hexProperty.currentWidth > 0 && hexProperty.currentHeight > 0) {
        context.save() // 保存当前状态
        context.translate(x + xOffset + hexRadius, hexProperty.currentY + hexHeight / 2) // 移动到六边形中心
        context.rotate(hexProperty.rotation) // 旋转
        // 使用当前宽度和高度进行缩放
        context.scale(hexProperty.currentWidth / hexWidth, hexProperty.currentHeight / hexHeight)
        context.drawImage(
          img,
          -hexRadius, // 调整x坐标，使图像居中
          -hexHeight / 2, // 调整y坐标，使图像居中
          hexWidth, // 使用原始宽度
          hexHeight // 使用原始高度
        )
        context.restore() // 恢复到之前的状态
      }
    }
  }
}
// 缓动函数
const easeInOutQuad = (t) => t * (2 - t)
// 判断鼠标是否在六边形内
const isHovering = (hexX, hexY, hexWidth, hexHeight) => {
  const mouseX = mousePosition.x
  const mouseY = mousePosition.y
  return mouseX >= hexX && mouseX <= hexX + hexWidth && mouseY >= hexY && mouseY <= hexY + hexHeight
}
// 获取鼠标位置
export const mouseClickToBC = (event) => {
  mousePosition.x = event.clientX
  mousePosition.y = event.clientY
}
// 鼠标移动时，更新鼠标位置
export const mouseLeaveToBC = () => {
  mousePosition.x = -100 // 设置到无效值
  mousePosition.y = -100
}
