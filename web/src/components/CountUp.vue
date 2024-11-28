<template>
	<span ref="countRef">{{ displayValue }}</span>
</template>

<script>
import { ref, onMounted, watch } from 'vue'

export default {
	name: 'CountUp',
	props: {
		endValue: {
			type: Number,
			required: true
		},
		duration: {
			type: Number,
			default: 2000
		},
		decimals: {
			type: Number,
			default: 0
		}
	},
	setup(props) {
		const displayValue = ref(0)
		const countRef = ref(null)
		// 缓动函数
		const easeOutExpo = (t) => {
			return t === 1 ? 1 : 1 - Math.pow(2, -10 * t)
		}
		// 动画函数
		const animateValue = (start, end, duration) => {
			if (typeof end !== 'number' || isNaN(end)) {
				displayValue.value = 0
				return
			}
			// 获取当前时间
			const startTime = performance.now()
			const updateCount = (currentTime) => {
				const elapsed = currentTime - startTime
				const progress = Math.min(elapsed / duration, 1)
				const easedProgress = easeOutExpo(progress)
				const value = start + (end - start) * easedProgress
				displayValue.value = isNaN(value) ? 0 : Math.round(value)

				if (progress < 1) {
					requestAnimationFrame(updateCount)
				}
			}
			requestAnimationFrame(updateCount)
		}

		watch(() => props.endValue, (newVal) => {
			const startValue = isNaN(displayValue.value) ? 0 : displayValue.value
			const endValue = isNaN(newVal) ? 0 : newVal
			animateValue(startValue, endValue, props.duration)
		}, { immediate: true })

		return {
			displayValue,
			countRef
		}
	}
}
</script>

<style scoped>
span {
	display: inline-block;
	transition: transform 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
}

span:hover {
	transform: scale(1.1);
}
</style>