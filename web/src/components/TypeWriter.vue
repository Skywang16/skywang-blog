<template>
	<span :class="{ 'typing-finished': isFinished }">{{ displayText }}</span>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const props = defineProps({
	text: {
		type: String,
		required: true
	},
	speed: {
		type: Number,
		default: 50
	}
})

const displayText = ref('')
const isFinished = ref(false)

onMounted(() => {
	typeText()
})

const typeText = () => {
	let index = 0
	const timer = setInterval(() => {
		if (index < props.text.length) {
			displayText.value += props.text[index]
			index++
		} else {
			clearInterval(timer)
			isFinished.value = true
		}
	}, props.speed + Math.random() * 50)
}
</script>

<style scoped>
span {
	border-right: 2px solid #333;
	animation: blink 0.7s infinite;
}

span.typing-finished {
	border-right: none;
	animation: none;
}

@keyframes blink {
	50% {
		border-color: transparent;
	}
}
</style>