<template>
	<view class="container">
		<view class="header">
			<wwqy-header></wwqy-header>
		</view>
		<view class="main">
			<view class="inform-header-box">
				<view class="inform-headerImg">
					<view>
						<img class="headImg" :src="user.avatar" alt="" srcset="" />
					</view>
				</view>
				<view class="inform-header">
					<view class="inform-title">{{ user.nickname }}</view>
					<view class="inform-personal-identity">{{ user.mobile }}</view>
				</view>
			</view>
		</view>
		<view class="foot-box">
			<view>
				<span @click="handleClick1" class="foot" >《{{data.privacy.name}}》</span>
				<span @click="handleClick2" class="foot">《{{data.service.name}}》</span>
			</view>
			<view>
				<span class="foot">wang测试</span>
			</view>
		</view>
		<uni-popup ref="details1" type="dialog" background-color="#fff">
			<view class="details1">
				<scroll-view scroll-y="true" :refresher-threshold="40" style="height: 100%; overflow: auto" @scroll="scroll">
					<view v-html="data.privacy.content"></view>
				</scroll-view>
			</view>
			<view class="popConfig-btn">
				<button class="popupVerify" @click="popupVerify">确认</button>
			</view>
		</uni-popup>
		<uni-popup ref="details2" type="dialog" background-color="#fff">
			<view class="details1">
				<scroll-view scroll-y="true" :refresher-threshold="40" style="height: 100%; overflow: auto" @scroll="scroll">
					<view v-html="data.service.content"></view>
				</scroll-view>
			</view>
			<view class="popConfig-btn">
				<button class="popupVerify" @click="popupVerify">确认</button>
			</view>
		</uni-popup>
		<view class="footer"></view>
	</view>
</template>

<script>
import wwqyHeader from '@/components/wwqy-header/index.vue';
import wwqyFooter from '@/components/wwqy-footer/index.vue';
import { protocol } from '@/common/api/business.js';
export default {
	components: { wwqyHeader, wwqyFooter },
	data() {
		return {
			data: {}
		};
	},
	onLoad() {
		this.user = uni.getStorageSync('userInfo');
		this.protocol();
	},
	methods: {
		protocol() {
			protocol().then((res) => {
				let result = res.data;
				if (result.code === 200) {
					this.data = result.data;
				} else {
					that.messageText = result.msg;
					that.$refs.message.open();
				}
			});
		},
		handleClick1() {
			this.fileType = '1';
			this.$refs.details1.open('center');
		},
		handleClick2() {
			this.fileType = '2';
			this.$refs.details2.open('center');
		},
		popupVerify() {
			this.$refs.details1.close();
			this.$refs.details2.close();
		}
	}
};
</script>

<style scoped lang="scss">
page {
	height: 100%;
	background-color: #fff;
}
.title {
	text-align: center;
	width: 100%;
	font-size: 20px;
	font-weight: 600;
	color: #fff;
	padding: 25px 0px;
}

.inform-header-box {
	margin-top: 40rpx;
	margin-left: 30rpx;
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
	.foot {
		color: $descColor;
	}
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
