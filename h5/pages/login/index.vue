<template>
	<view class="login-index">
		<view class="title" style="text-align: center">用户登录</view>
		<view class="input-list">
			<view class="input-item">
				<view class="img-box">
					<image mode="widthFix" src="@/static/icon/auth-xm.png" alt="" />
				</view>
				<input class="uni-input" v-model="form.username" placeholder="账号" />
			</view>
			<view class="input-item input-code">
				<view class="img-box">
					<image mode="widthFix" src="@/static/icon/auth-yzm.png" alt="" />
				</view>
				<input class="uni-input" password v-model="form.password" placeholder="密码" />
			</view>
		</view>
		<view class="button">
			<button class="btn" type="primary" @click="login">登录</button>
			<view class="register" @click="register">立即注册</view>
		</view>
		<!-- <view class="seleteuser">
			<uni-data-checkbox style="width: 40rpx; margin-right: 10rpx" @change="changeRadio" v-model="hobby" multiple :localdata="hobbys" />
			<view>
				<span>已阅读并同意</span>
				<span @click="handleClick('1')" style="color: #2aa3f8">平台服务协议、</span>
				<span @click="handleClick('2')" style="color: #2aa3f8">用户隐私协议、</span>
				<span @click="handleClick('3')" style="color: #2aa3f8">数字证书协议</span>
			</view>
		</view> -->
		<uni-popup ref="message" type="message">
			<uni-popup-message type="error" :message="messageText" :duration="2000"></uni-popup-message>
		</uni-popup>
		<loading :loading="showLoading" :firstContent="firstContent" :secondContent="secondContent" />
	</view>
</template>

<script>
import { token, getUserInfo } from '@/common/api/login.js';
export default {
	data() {
		return {
			password: '',
			form: {
				uniteCode: '',
				username: '',
				password: ''
			},
			hobby: [0],
			hobbys: [
				{
					text: '',
					value: 1
				}
			],
			messageText: '错误信息',
			isCheck: true,
			showLoading: false, // 自定义遮罩显影
			firstContent: '', // 自定义遮罩第一行文本
			secondContent: '' // 自定义遮罩第二行文本
		};
	},
	onLoad() {},
	methods: {
		changeRadio(e) {
			this.isCheck = !this.isCheck;
		},
		register() {
			uni.navigateTo({
				url: './register',
				animationType: 'fade-in',
				animationDuration: 300
			});
		},
		login() {
			var that = this;
			if (!that.form.username) {
				that.messageText = '请填写账号';
				that.$refs.message.open();
				return;
			}
			if (!that.form.password) {
				that.messageText = '请填写密码';
				that.$refs.message.open();
				return;
			}
			let params = {
				username: that.form.username,
				password: that.form.password,
				terminal: 1
			};
			that.onGetToken(params);
		},
		onGetToken(params) {
			this.setLoadingOptions(true, '登录中');
			token(params).then((res) => {
				this.setLoadingOptions(false);
				let result = res.data;
				if (result.code === 200) {
					uni.setStorageSync('token', result.data.token);
					this.getUserInfo();
				} else {
					this.messageText = '登录账号或密码错误';
					this.$refs.message.open();
				}
			});
		},
		getUserInfo() {
			this.setLoadingOptions(true, '获取用户信息中...');
			getUserInfo().then((res) => {
				this.setLoadingOptions(false);
				let result = res.data;
				if (result.code === 200) {
					console.log(result.data);
					uni.setStorageSync('userData', result.data.user);
					this.toIndex();
				} else {
					this.setLoadingOptions(false);
				}
			});
		},
		toIndex() {
			uni.reLaunch({
				url: '/pages/home/index'
			});
		},
		// 设置遮罩数据
		setLoadingOptions(flag, firstText, secondText) {
			this.showLoading = flag || false;
			this.firstContent = firstText || '';
			this.secondContent = secondText || '';
		}
	}
};
</script>

<style lang="scss" scoped>
page {
	height: 100%;
}
.login-index {
	position: absolute;
	height: 100%;
	// height: 100vh;
	// margin-top: 160rpx;
	width: 100%;
	padding-top: 100rpx;
	background-color: #ffffff;
	box-sizing: border-box;
	border-top: 1px solid #efefef;
}
.title {
	font-size: 48rpx;
	font-weight: bold;
}
.input-list {
	width: 100%;
	display: flex;
	flex-direction: column;
	align-items: center;
	margin-top: 80rpx;
	.input-item {
		width: 640rpx;
		height: 80rpx;
		background: #f1f1f1;
		border-radius: 8rpx;
		display: flex;
		justify-content: center;
		padding-left: 30rpx;
		box-sizing: border-box;
		align-items: center;
		margin-bottom: 30rpx;
		.img-box {
			width: 60rpx;
			display: flex;
			image {
				width: 40rpx;
			}
		}
	}
}
.button {
	width: 100%;
	margin-top: 20rpx;
	margin-bottom: 20rpx;
}
.register {
	display: flex;
	justify-content: center;
	margin-top: 30rpx;
}
.btn {
	height: 80rpx;
	line-height: 80rpx;
	border-radius: 8rpx;
	width: 640rpx;
}
.button view {
	padding-bottom: 10rpx;
	color: #007aff;
}
.input-item {
	display: flex;
}
.unibum {
	min-width: 100rpx;
	text-align: center;
	// flex: 1;
	font-size: 26rpx;
	color: #fff;
	background-color: #0366f1;
	line-height: 50rpx;
	padding: 0 30rpx;
	margin-right: 20rpx;
	height: 50rpx;
	border-radius: 30rpx;
	border: 2rpx solid #0366f1;
}
.unibum2 {
	min-width: 40rpx;
	text-align: center;
	font-size: 32rpx;
	color: #fff;
	line-height: 50rpx;
	padding: 0 20rpx;
	height: 50rpx;
	border-radius: 30rpx;
}

.unibum:active {
	// background-color: red;
	// display:inline-block;
	transform: scale(0.8);
	// margin-top: 20px;
}

.on-disabled {
	pointer-events: none;
	color: #999;
	cursor: not-allowed;
}
.uni-input {
	flex: 2;
}
.loginImg {
	width: 100%;
	text-align: center;
	margin-bottom: 60rpx;
}
.loginImg image {
	width: 180rpx;
	height: 180rpx;
	border-radius: 50%;
	text-align: center;
}
.seleteuser {
	display: flex;
	font-size: 24rpx;
	color: #999999;
	justify-content: center;
	align-items: center;
}

.uni-data-checklist {
	flex: none;
}
.input-code {
	flex-direction: row !important;
	// justify-content: center;
	align-items: center;
}
/deep/ .uni-data-checklist .checklist-group .checklist-box {
	margin-right: 0;
}
</style>
