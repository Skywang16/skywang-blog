<template>
	<view class="content">
		<view class="pic">
			<image mode="widthFix" :src="weaponItem.image" alt="" />
		</view>
		<view class="text">
			<view class="title">{{ weaponItem.name }}</view>
			<view class="info">
				<view class="row">
					<view class="left">系 列：</view>
					<view class="right">{{ weaponItem.series }}</view>
				</view>
				<view class="row">
					<view class="left">价 格：</view>
					<view class="right">{{ weaponItem.price }}</view>
				</view>
				<view class="row">
					<view class="left">等 级：</text></view>
					<view class="right">{{ levelTexts[weaponItem.level] }}</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
export default {
	data() {
		return {
			levelTexts: {
				0: '精选',
				1: '奢华',
				2: '尊爵',
				3: '究极',
				4: '限定'
			}
		};
	},
	props: {
		weaponItem: {
			type: Object,
			default: null
		}
	},
	onLoad() {},
	methods: {}
};
</script>

<style lang="scss" scoped>
.content {
	margin: 10rpx 10rpx 0 10rpx;
	background-color: #fff;
	border-radius: 20rpx;
	box-shadow: 1px 10px 10px rgba(198, 196, 196, 0.1);
	padding: 20rpx;
	border: 1px solid #eee;
}

.content .pic {
	image {
		width: 100%;
		border-radius: 10rpx;
	}
}

.content .title {
	font-size: 36rpx;
	color: #333;
	padding-bottom: 20rpx;
	border-bottom: 1rpx solid #efefef;
	white-space: nowrap;
	text-overflow: ellipsis;
	overflow: hidden;
}

.content .row {
	font-size: 28rpx;
	color: #aaa;
	display: flex;
	line-height: 2em;
}

.content .row .left {
	font-size: 26rpx;
	white-space: nowrap;
}

.content .row .right {
	color: #666;
	white-space: nowrap;
}
</style>
