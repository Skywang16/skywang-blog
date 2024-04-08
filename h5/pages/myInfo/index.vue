<template>
	<view class="container">
		<view id="bg"></view>
		<view class="main">
			<view class="inform-header-box">
				<view class="inform-headerImg">
					<view>
						<img class="headImg" :src="user.avatar" alt="" srcset="" />
					</view>
				</view>
				<view class="inform-header">
					<view class="inform-title">{{user.nickname}}</view>
					<view class="inform-personal-identity">{{user.mobile}}</view>
				</view>
			</view>
		</view>
		<view class="foot-box">
			<view class="foot" @click="handleClick2">
				<span style="color: #3f6fff">《用户隐私协议》</span>
			</view>
			<view class="foot">
				<span style="color: #3f6fff">wang测试</span>
			</view>
		</view>
		<uni-popup ref="details" type="dialog" background-color="#fff">
			<view class="details1">
				<scroll-view scroll-y="true" :refresher-threshold="40" style="height: 100%; overflow: auto" @scroll="scroll">
					122123
				</scroll-view>
			</view>
			<view class="popConfig-btn">
				<button class="popupVerify" @click="popupVerify">确认</button>
			</view>
		</uni-popup>
	</view>
</template>

<script>
export default {
	data() {
		return {
			
		};
	},
	onLoad() {
		this.user = uni.getStorageSync('userInfo');
	},
	methods: {
		handleClick1() {
			this.fileType = '1';
			this.$refs.details.open('center');
		},
		handleClick2() {
			this.fileType = '2';
			this.$refs.details.open('center');
		},
		handleClick3() {
			this.fileType = '3';
			this.$refs.details.open('center');
		},
		popupVerify() {
			this.$refs.details.close();
		}
	}
};
</script>

<style scoped lang="scss">
page {
	height: 100%;
	background-color: #fff;
}

#bg {
	background-image: url('@/static/myInfo/bc.png');
	background-size: 100%;
	height: 700rpx;
	width: 100%;
}
.title {
	text-align: center;
	width: 100%;
	font-size: 20px;
	font-weight: 600;
	color: #fff;
	padding: 25px 0px;
}
.main {
	margin: -590rpx 30rpx 30rpx 30rpx;
}
.inform-header-box {
	display: flex;
	.inform-headerImg {
		margin-right: 20rpx;
	}
	.inform-header {
		margin-left: 10rpx;
		margin-top: 10px;
	}
	.inform-personal-identity {
		font-size: 28rpx;
		color: #666666;
		margin-top: 24rpx;
	}
	.headImg {
		border-radius: 50%;
		height: 100rpx;
		width: 100rpx;
	}
}

.inform-title {
	font-size: 40rpx;
	font-weight: 600;
	color: #333333;
	line-height: 20rpx;
}

.notice-bar-box {
	position: relative;
	margin: 49rpx 30rpx;
}
.grid-item-box {
	margin: 86rpx 0;
	display: flex;
	padding-bottom: 43rpx;

	img {
		height: 48rpx;
		margin: auto 0;
		margin-right: 24rpx;
		margin-left: -28rpx;
	}
	.text {
		width: 90%;
		padding: 43rpx 0;
		border-bottom: 1px solid #e6e6e6;
		margin: auto 0;
		font-size: 36rpx;
		color: #333333;
	}
	.text2 {
		padding: 43rpx 0;
		margin: auto 0;
		font-size: 36rpx;
		color: #333333;
	}
}
.foot-box {
	font-size: 14px;
	width: 100%;
	text-align: center;
	position: absolute;
	bottom: 30rpx;
}
.details1 {
	width: 80vw;
	height: 60vh;
	padding: 40rpx;
}
.popConfig-btn {
	display: flex;
	position: relative;
	bottom: 10px;
	width: 100%;
	justify-content: center;
	.popupVerify {
		width: 30%;
		height: 80rpx;
		line-height: 80rpx;
		background-color: #fff;
		border: 1px solid #5861d0ff;
		color: #5861d0ff;
		border-radius: 40rpx;
	}
}
/deep/.uni-popup__wrapper {
	border-radius: 30rpx;
}
.text-box {
	width: 100%;
	display: flex;
	border-bottom: 1px solid #e6e6e6;
	color: #333333;
}
.grid-item-box2 {
	margin: 20rpx auto;
	text {
		color: #3f6fff;
		font-size: 28rpx;
	}
}
</style>
