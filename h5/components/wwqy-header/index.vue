<template>
	<view class="main">
		<view class="header">
			<image class="logo" src="/static/logo.png" />
			<view @click="gotoMyInfo">
				<uni-icons type="person" size="30" style="color: #fd4453"></uni-icons>
			</view>
		</view>
		<view class="line"></view>
		<uni-popup ref="message" type="message">
			<uni-popup-message type="error" :message="messageText" :duration="2000"></uni-popup-message>
		</uni-popup>
	</view>
</template>

<script>
export default {
	data() {
		return {
			messageText: '错误信息'
		};
	},

	mounted() {},

	methods: {
		gotoMyInfo() {
			let user = uni.getStorageSync('userInfo');
			if (user.role === '0') {
				this.messageText = '您还未登录';
				this.$refs.message.open();
				setTimeout(() => {
					uni.navigateTo({
						url: '/pages/login/index?autoLogin=1',
						animationType: 'fade-in',
						animationDuration: 300
					});
				}, 2000);
			} else {
				uni.navigateTo({
					url: '/pages/myInfo/index',
					animationType: 'fade-in',
					animationDuration: 300
				});
			}
		}
	}
};
</script>

<style lang="scss" scoped>
.header {
	height: 120rpx;
	background-color: #000;
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 0 30rpx;

	.logo {
		height: 90rpx;
		width: 140rpx;
	}
}
.line {
	height: 1px;
	background-color: #fd4453;
}
</style>
