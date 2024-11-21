// 存储每个六边形形状的动画属性和状态
const hexMovementProperties = {}
// 存储鼠标当前位置
let mousePosition = { x: -100, y: -100 }

// 全局状态变量
let currentActiveShape = 'block' // 当前激活的形状类型
let oldShape = '' // 上一个激活的形状类型
let switchStartTime = null // 形状切换开始的时间戳
let isFirstSwitch = true // 是否是首次切换形状
let isInitialized = false // 初始化标记，确保首次渲染正确
let switchPhase = 'none' // 添加切换阶段标记：'none', 'shrink', 'wait', 'grow'

// Web Worker 消息处理函数
self.onmessage = function (e) {
  let result

  switch (e.data.type) {
    case 'updateMousePosition': {
      // 更新鼠标位置信息
      mousePosition = e.data.data
      break
    }
    case 'shapeSwitch': {
      // 处理形状切换逻辑
      const { currentActiveShape: newShape, timestamp } = e.data.data
      if (newShape !== currentActiveShape) {
        if (!isInitialized) {
          // 首次切换前，确保当前形状完全显示
          Object.keys(hexMovementProperties).forEach((key) => {
            const prop = hexMovementProperties[key]
            if (prop.shapeType === currentActiveShape) {
              prop.currentWidth = prop.targetWidth
              prop.currentHeight = prop.targetHeight
            }
          })
          isInitialized = true
        }

        // 重置所有形状的动画状态
        Object.keys(hexMovementProperties).forEach((key) => {
          const prop = hexMovementProperties[key]
          // 重置旋转和位置相关的属性
          prop.rotation = 0
          prop.targetRotation = 0
          prop.currentY = prop.originalY
          prop.targetY = prop.originalY
          prop.hasMoved = false
          prop.isAnimating = false
          prop.animationPhase = 'none'

          // 设置目标尺寸
          if (prop.shapeType === oldShape) {
            prop.targetWidth = 0
            prop.targetHeight = 0
          }
        })

        oldShape = currentActiveShape
        currentActiveShape = newShape
        switchStartTime = timestamp
        switchPhase = 'shrink'
      }
      break
    }
    case 'calculateShapes': {
      // 计算所有形状的位置和状态
      let { canvas, imgConfig, timestamp } = e.data.data
      // 确保首次计算时有正确的时间戳
      if (!switchStartTime) {
        switchStartTime = timestamp
      }
      // 执行形状计算并返回结果
      result = calculateShapes(canvas, currentActiveShape, oldShape, imgConfig)
      self.postMessage({
        type: 'shapeResults',
        data: result
      })
      break
    }
  }
}

// 只保留这一个缓动函数
const easeInOutQuad = (t) => t * (2 - t)

// 检测鼠标是否悬停在指定形状上
const isHovering = (hexX, hexY, hexWidth, hexHeight) => {
  const mouseX = mousePosition.x
  const mouseY = mousePosition.y
  return mouseX >= hexX && mouseX <= hexX + hexWidth && mouseY >= hexY && mouseY <= hexY + hexHeight
}

// 计算六边形的属性（位置、大小等）
function calculateBlock(hexKey, xPos, yOffset, hexWidth, hexHeight, shapeType, isInTransition) {
  if (!hexMovementProperties[hexKey]) {
    const isCurrentShape = shapeType === currentActiveShape
    hexMovementProperties[hexKey] = {
      currentY: yOffset,
      targetY: yOffset,
      currentWidth: isCurrentShape && !isFirstSwitch ? hexWidth : 0,
      currentHeight: isCurrentShape && !isFirstSwitch ? hexHeight : 0,
      targetWidth: isCurrentShape ? hexWidth : 0,
      targetHeight: isCurrentShape ? hexHeight : 0,
      centerX: xPos + hexWidth / 2,
      centerY: yOffset + hexHeight / 2,
      width: hexWidth,
      height: hexHeight,
      shapeType,
      isAnimating: false, // 是否正在动画中
      animationStartTime: 0, // 动画开始时间
      originalY: yOffset, // 原始Y位置
      animationPhase: 'none', // 动画阶段
      hasTriggered: false // 新增：是否已经触发过动画
    }
  }

  const hexProperty = hexMovementProperties[hexKey]

  // 更新目标值
  if (shapeType === currentActiveShape) {
    if (!isInitialized) {
      hexProperty.currentWidth = hexWidth
      hexProperty.currentHeight = hexHeight
    }
    hexProperty.targetWidth = hexWidth
    hexProperty.targetHeight = hexHeight
  } else if (shapeType === oldShape && isInTransition) {
    hexProperty.targetWidth = 0
    hexProperty.targetHeight = 0
  }

  // 修改悬停检测和动画逻辑
  if (
    isHovering(xPos, yOffset, hexWidth, hexHeight) &&
    !hexProperty.isAnimating &&
    !hexProperty.hasTriggered
  ) {
    hexProperty.isAnimating = true
    hexProperty.animationStartTime = Date.now()
    hexProperty.animationPhase = 'up'
    hexProperty.hasTriggered = true // 标记已触发
  }

  // 如果鼠标移开，重置触发标记，这样下次移入时可以再次触发
  if (!isHovering(xPos, yOffset, hexWidth, hexHeight)) {
    hexProperty.hasTriggered = false
  }

  // 处理动画逻辑
  if (hexProperty.isAnimating) {
    const upDuration = 180 // 上移动画时间
    const downDuration = 150 // 下移动画时间
    const moveDistance = 150 // 移动距离
    const elapsed = Date.now() - hexProperty.animationStartTime

    if (hexProperty.animationPhase === 'up') {
      if (elapsed < upDuration) {
        // 上移阶段 - 使用 easeInOutQuad 实现轻微的弹跳效果
        hexProperty.targetY = hexProperty.originalY - moveDistance
      } else {
        // 开始下移阶段
        hexProperty.animationStartTime = Date.now()
        hexProperty.animationPhase = 'down'
      }
    } else if (hexProperty.animationPhase === 'down') {
      if (elapsed < downDuration) {
        // 下移阶段 - 使用 easeInOutQuad 实现平滑的回落
        hexProperty.targetY = hexProperty.originalY - moveDistance + moveDistance
      } else {
        // 动画结束
        hexProperty.isAnimating = false
        hexProperty.animationPhase = 'none'
        hexProperty.targetY = hexProperty.originalY
      }
    }
  }

  // 提高位置更新的响应速度
  const progressY = easeInOutQuad(0.08) // 增加位置更新的速度
  const progressWidth = easeInOutQuad(0.1)
  const progressHeight = easeInOutQuad(0.1)

  hexProperty.currentY = hexProperty.currentY || yOffset
  hexProperty.currentWidth = hexProperty.currentWidth || 0
  hexProperty.currentHeight = hexProperty.currentHeight || 0

  // 使用更平滑的位置更新
  hexProperty.currentY += (hexProperty.targetY - hexProperty.currentY) * progressY
  hexProperty.currentWidth += (hexProperty.targetWidth - hexProperty.currentWidth) * progressWidth
  hexProperty.currentHeight +=
    (hexProperty.targetHeight - hexProperty.currentHeight) * progressHeight

  // 根据当前尺寸计算实际渲染位置，使缩放以中心点为基准
  const renderX = hexProperty.centerX - hexProperty.currentWidth / 2
  const renderY = hexProperty.currentY - (hexProperty.currentHeight - hexHeight) / 2

  return {
    x: renderX,
    y: renderY,
    width: hexProperty.currentWidth,
    height: hexProperty.currentHeight
  }
}

// 计算三叶草的属性（位置、大小、旋转等）
function calculateTetrapod(hexKey, xPos, yOffset, hexWidth, hexHeight, shapeType, isInTransition) {
  if (!hexMovementProperties[hexKey]) {
    const isCurrentShape = shapeType === currentActiveShape
    hexMovementProperties[hexKey] = {
      currentY: yOffset,
      targetY: yOffset,
      currentWidth: isCurrentShape && !isFirstSwitch ? hexWidth : 0,
      currentHeight: isCurrentShape && !isFirstSwitch ? hexHeight : 0,
      targetWidth: isCurrentShape ? hexWidth : 0,
      targetHeight: isCurrentShape ? hexHeight : 0,
      centerX: xPos + (hexWidth * 0.73) / 2,
      centerY: yOffset + (hexHeight * 0.73) / 2,
      width: hexWidth,
      height: hexHeight,
      rotation: 0,
      targetRotation: 0,
      shapeType,
      hasMoved: false,
      originalY: yOffset
    }
  }

  const hexProperty = hexMovementProperties[hexKey]

  // 更新目标值
  if (shapeType === currentActiveShape) {
    if (!isInitialized) {
      hexProperty.currentWidth = hexWidth
      hexProperty.currentHeight = hexHeight
    }
    hexProperty.targetWidth = hexWidth
    hexProperty.targetHeight = hexHeight
  } else if (shapeType === oldShape && isInTransition) {
    hexProperty.targetWidth = 0
    hexProperty.targetHeight = 0
  }

  if (isHovering(xPos, yOffset, hexWidth, hexHeight)) {
    if (!hexProperty.hasMoved) {
      hexProperty.targetRotation += -Math.PI / 1.5
      hexProperty.hasMoved = true
    }
  } else {
    hexProperty.hasMoved = false
  }

  const progressRotation = easeInOutQuad(0.1)
  const progressWidth = easeInOutQuad(0.1)
  const progressHeight = easeInOutQuad(0.1)

  // 确保所有值都被正确初始化
  hexProperty.currentY = hexProperty.currentY || yOffset
  hexProperty.currentWidth = hexProperty.currentWidth || 0
  hexProperty.currentHeight = hexProperty.currentHeight || 0

  hexProperty.currentWidth += (hexProperty.targetWidth - hexProperty.currentWidth) * progressWidth
  hexProperty.currentHeight +=
    (hexProperty.targetHeight - hexProperty.currentHeight) * progressHeight
  hexProperty.rotation += (hexProperty.targetRotation - hexProperty.rotation) * progressRotation

  // 修改渲染位置的计算方式
  const renderX = hexProperty.centerX - (hexProperty.currentWidth * 0.73) / 2
  const renderY = hexProperty.currentY - (hexProperty.currentHeight * 0.73) / 2

  return {
    x: renderX,
    y: renderY,
    width: hexProperty.currentWidth,
    height: hexProperty.currentHeight,
    rotation: hexProperty.rotation,
    rotationOriginX: hexProperty.centerX - renderX,
    rotationOriginY: hexProperty.centerY - renderY
  }
}

// 主计算函数：计算所有形状的位置和状态
function calculateShapes(canvas, currentActiveShape, oldShape, imgConfig) {
  const shapes = []
  const now = Date.now()
  const shrinkDuration = 200
  const waitDuration = 200
  const growDuration = 200

  // 更新切换阶段逻辑
  if (switchPhase === 'shrink') {
    const elapsed = now - switchStartTime
    if (elapsed >= shrinkDuration) {
      switchPhase = 'wait'
      switchStartTime = now
    }
  } else if (switchPhase === 'wait') {
    const elapsed = now - switchStartTime
    if (elapsed >= waitDuration) {
      switchPhase = 'grow'
      switchStartTime = now
    }
  } else if (switchPhase === 'grow') {
    const elapsed = now - switchStartTime
    if (elapsed >= growDuration) {
      switchPhase = 'none'
    }
  }

  // 遍历配置，计算每个形状的位置和属性
  imgConfig.forEach((img) => {
    // 计算六边形的基本尺寸和间距
    const hexHeight = img.height * 0.5
    const hexWidth = img.width * 0.5
    const hexRadius = hexWidth / 2
    // 根据形状类型调整垂直和水平间距
    const hexVerticalSpacing = hexHeight * (img.shape === 'tetrapod' ? 0.73 : 0.75)
    const hexHorizontalSpacing = hexWidth * (img.shape === 'tetrapod' ? 0.73 : 1)
    // 计算起始偏移量
    const xOffsetBase = -img.width * (img.shape === 'tetrapod' ? 0.5 : 0.25)
    const yOffsetBase = -img.height * 0.5

    const maxY = Math.floor(canvas.height + hexHeight)
    const maxX = Math.floor(canvas.width + hexWidth)

    for (let y = 0; y < maxY; y += hexVerticalSpacing) {
      const isEvenRow = Math.floor(y / hexVerticalSpacing) % 2 === 0
      const xOffset = isEvenRow ? 0 : hexRadius * (img.shape === 'tetrapod' ? 0.73 : 1)

      for (let x = 0; x < maxX; x += hexHorizontalSpacing) {
        const yOffset = y + yOffsetBase
        const xPos = x + xOffset + xOffsetBase
        const hexKey = `${Math.round(xPos)}_${Math.round(yOffset)}`

        // 根据当前阶段决定是否渲染形状
        let shouldRenderShape = false
        if (switchPhase === 'shrink') {
          // 缩小阶段只渲染旧形状
          shouldRenderShape = img.shape === oldShape
        } else if (switchPhase === 'grow') {
          // 放大阶段只渲染新形状
          shouldRenderShape = img.shape === currentActiveShape
        } else if (switchPhase === 'none') {
          // 正常状态只渲染当前形状
          shouldRenderShape = img.shape === currentActiveShape
        }
        // wait 阶段不渲染任何形状

        if (shouldRenderShape) {
          const shapeData =
            img.shape === 'block'
              ? calculateBlock(
                  hexKey,
                  xPos,
                  yOffset,
                  hexWidth,
                  hexHeight,
                  img.shape,
                  switchPhase !== 'none'
                )
              : calculateTetrapod(
                  hexKey,
                  xPos,
                  yOffset,
                  hexWidth,
                  hexHeight,
                  img.shape,
                  switchPhase !== 'none'
                )

          if (shapeData.width > 0 || shapeData.height > 0) {
            shapes.push({
              type: img.shape,
              ...shapeData
            })
          }
        }
      }
    }
  })

  // 应用动画进度
  if (switchPhase !== 'none') {
    const elapsed = now - switchStartTime
    let progress = 1

    if (switchPhase === 'shrink') {
      progress = Math.max(0, 1 - elapsed / shrinkDuration)
      shapes.forEach((shape) => {
        shape.width *= progress
        shape.height *= progress
      })
    } else if (switchPhase === 'grow') {
      progress = Math.min(1, elapsed / growDuration)
      shapes.forEach((shape) => {
        shape.width *= progress
        shape.height *= progress
      })
    }
  }

  return shapes
}
