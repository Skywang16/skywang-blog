<template>
	<view class="container">
		<view id="bg">
			<view class="main">
				<view class="card-body" style="height: calc(100vh - 90rpx); overflow-y: auto">
					<view class="card-title">{{ result.title }}</view>
					<view class="card-details">
						<view style="width: 340rpx">
							来源：
							<span class="details-text">{{ result.cid }}</span>
						</view>
						<view class="">
							<span class="details-text">
								{{ result.publishTime }}
							</span>
						</view>
					</view>
					<view class="card-content">
						<view v-html="result.content"></view>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
export default {
	data() {
		return {
			result: {}
		};
	},
	onLoad() {
		this.result = uni.getStorageSync('newsDetail');
		console.log(this.result)
	},
	methods: {}
};
</script>

<style scoped lang="scss">
.container {
	background-color: #fd4453;
	min-height: 100vh;
}

#bg {
	background: linear-gradient(180deg, #000 0%, rgba(68, 151, 247, 0) 100%);
	height: 954rpx;
	width: 100%;
}

.card-body {
	position: relative;
	top: 44rpx;
	margin: auto;
	width: 690rpx;
	background: #fff;
	box-shadow: 0rpx 18rpx 29rpx 0rpx rgba(78, 104, 156, 0.09);
	border-radius: 20rpx;
	.card-title {
		padding-top: 30rpx;
		text-align: center;
		font-size: 44rpx;
		font-family: PingFangHK-Semibold, PingFangHK;
		font-weight: 600;
		color: #333333;
		line-height: 48rpx;
	}

	.card-details {
		display: flex;
		justify-content: space-between;
		margin: 40rpx;
		font-size: 30rpx;
		font-family: PingFangHK-Regular, PingFangHK;
		color: #999999;
		.details-text {
			font-size: 30rpx;
			font-family: PingFangHK-Regular, PingFangHK;
			color: #999999;
		}
	}
}

.card-content {
	font-size: 32rpx;
	margin: 10rpx 40rpx;
}
</style>
