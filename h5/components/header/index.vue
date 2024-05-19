<template>
	<view class="header">
		<view class="row">Blog</view>
	</view>
</template>

<script>
export default {
	data() {
		return {
			year: new Date().getFullYear()
		};
	},

	mounted() {},

	methods: {}
};
</script>

<style lang="scss" scoped>
.header{
	height: 100rpx;
	background-color: #000,
}
.row{
	color: #fff;
}
</style>
