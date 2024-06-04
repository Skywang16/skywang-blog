<template>
	<view ref="rolls" class="rolls" :style="{ '--trans-late': `-${totalWidth}px`, '--animation-duration': animationSpeed + 's' }">
		<!--实际数据组 GO-->
		<view
			class="roll"
			v-for="(item, index) in rollList"
			:key="index"
			@click="goSearch(item)"
			:style="{ fontSize: item.size + 'px', color: item.color, backgroundColor: item.backgroundColor || getColorOpacity(item.color) }"
		>
			{{ item.text }}
		</view>
		<!--填充数据组 GO-->
		<view
			class="roll"
			v-for="(item, index) in rollList"
			:key="item.text + index"
			@click="goSearch(item)"
			:style="{ fontSize: item.size + 'px', color: item.color, backgroundColor: item.backgroundColor || getColorOpacity(item.color) }"
		>
			{{ item.text }}
		</view>
	</view>
</template>

<script>
export default {
	props: {
		rollList: {
			type: Array,
			default() {
				return [];
			}
		},
		bcAlpha: {
			type: Number,
			default() {
				return 0.2;
			}
		},
		animationSpeed: {
			type: Number,
			default() {
				return 5;
			}
		}
	},
	data() {
		return {
			totalWidth: 0
		};
	},
	mounted() {
		setTimeout(() => {
			this.getElementWidth();
			console.log(this.totalWidth)
		}, 300);
	},
	methods: {
		getElementWidth() {
			// 创建 Promise 对象，用于异步获取元素宽度
			uni.createSelectorQuery()
				.in(this)
				.select('.rolls')
				.boundingClientRect((data) => {
					this.totalWidth = data.width / 2;
				})
				.exec();
		},
		goSearch(item) {
			this.$emit('goSearch', item);
		},
		getColorOpacity(color) {
			// 将传入的颜色转换成 RGBA 格式
			let rgbaColor = this.hexToRGBA(color, this.bcAlpha); // 0.2 为透明度，可以根据需要调整
			return rgbaColor;
		},
		// 十六转RGBA
		hexToRGBA(hex, alpha) {
			hex = hex.replace('#', '');
			let r = parseInt(hex.substring(0, 2), 16);
			let g = parseInt(hex.substring(2, 4), 16);
			let b = parseInt(hex.substring(4, 6), 16);
			return `rgba(${r},${g},${b},${alpha})`;
		}
	}
};
</script>

<style scoped lang="scss">
.rolls {
	animation: move var(--animation-duration) linear infinite;
	white-space: nowrap;
	margin-top: 20rpx;
	.roll {
		padding: 10rpx;
		overflow: hidden;
		border-radius: 12rpx;
		display: inline-block;
		margin-right: 20rpx;
	}
}

.rolls:hover {
	animation: move var(--animation-duration) linear infinite;
}

@keyframes move {
	0% {
		transform: translateX(0rpx);
	}

	100% {
		transform: translateX(var(--trans-late));
	}
}
</style>
