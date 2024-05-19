<template>
	<view ref="rolls" class="rolls" :style="{ '--trans-late': `-${totalWidth}px` }">
		<!--实际数据组 GO-->
		<view
			class="roll"
			v-for="(item, index) in rollList"
			:key="index"
			@click="goSearch(item)"
			:style="{ fontSize: item.size + 'px', color: item.color, backgroundColor: item.backgroundColor }"
		>
			{{ item.text }}
		</view>
		<!--填充数据组 GO-->
		<view
			class="roll"
			v-for="(item, index) in rollList"
			:key="item.text + index"
			@click="goSearch(item)"
			:style="{ fontSize: item.size + 'px', color: item.color, backgroundColor: item.backgroundColor }"
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
		}
	},
	data() {
		return {
			totalWidth: 0,
			aaa: false
		};
	},
	async mounted() {
		this.totalWidth = await this.getElementHeight();
		console.log(this.totalWidth);
	},
	methods: {
		getElementHeight() {
			return new Promise((reslove) => {
				let info = uni.createSelectorQuery().in(this).select('.rolls');
				info.boundingClientRect(function (data) {}).exec((res) => reslove(res[0].width / 2));
			});
		},
		goSearch(it) {
			let options = {
				...it
			};
			this.$emit('goSearch', options.id);
			console.log(it);
		}
	}
};
</script>

<style scoped lang="scss">
.rolls {
	animation: move 7s linear infinite;
	white-space: nowrap;
	margin-top: 20rpx;
	.roll {
		padding: 10rpx;
		overflow: hidden;
		border-radius: 12rpx;
		display: inline-block;
		margin-right: 20rpx;
		image {
			width: 300rpx;
			height: 300rpx;
		}
	}
}

.rolls:hover {
	animation: move 7s linear infinite;
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
