<template>
	<view class="content">
		<view class="hero-img">
			<image mode="widthFix" src="@/static/information/01.png" alt="" />
		</view>
		<view class="skill-bar-box">
			<view class="line"></view>
			<view class="skill-bar">
				<view class="grid-item-box" v-for="(skill, index) in skills" :key="index" @click="changeVideo(skill)">
					<image mode="widthFix" :src="skill.image" alt="" srcset="" />
				</view>
			</view>
		</view>
		<view class="video-box">
			<video class="video" id="myVideo" :src="currentSkill.video" @error="videoErrorCallback" controls></video>
		</view>
		<view class="introduce-box">
			<view class="title">{{ currentSkill.title }}</view>
			<view class="detail">{{ currentSkill.description }}</view>
		</view>
	</view>
</template>

<script>
export default {
	data() {
		return {
			skills: [
				{
					title: 'C-剧震余波',
					description: '装备聚变炸药。按【射击】即可透过墙壁连续引发两次延时生效的爆破。该冲击会对作用范围内的所有人造成严重伤害。',
					video: 'https://media.w3.org/2010/05/sintel/trailer.mp4',
					image: '/static/information/01_c.png'
				},
				{
					title: 'Skill Title 2',
					description: 'Description for Skill 2',
					video: 'VIDEO_2_SRC_HERE',
					image: '/static/information/01_c.png'
				},
				{
					title: 'Skill Title 3',
					description: 'Description for Skill 3',
					video: 'VIDEO_3_SRC_HERE',
					image: '/static/information/01_c.png'
				},
				{
					title: 'Skill Title 4',
					description: 'Description for Skill 4',
					video: 'VIDEO_4_SRC_HERE',
					image: '/static/information/01_c.png'
				}
			],
			currentSkill: {}
		};
	},
	props: {},
	onLoad() {
		this.currentSkill = this.skills[0];
	},
	onReady: function (res) {
		this.videoContext = uni.createVideoContext('myVideo');
	},
	methods: {
		changeVideo(skill) {
			this.currentSkill = skill;
		},
		videoErrorCallback: function (e) {
			/* uni.showModal({
				content: e.target.errMsg,
				showCancel: false
			}); */
		},
		getRandomColor: function () {
			const rgb = [];
			for (let i = 0; i < 3; ++i) {
				let color = Math.floor(Math.random() * 256).toString(16);
				color = color.length == 1 ? '0' + color : color;
				rgb.push(color);
			}
			return '#' + rgb.join('');
		}
	}
};
</script>

<style scoped lang="scss">
.content {
	height: 100vh;
	background-image: url('@/static/information/bj.webp');
	background-size: 100% 100%;
	background-repeat: no-repeat;
}
.hero-img {
	display: flex;
	justify-content: center;
	image {
		width: 50vw;
	}
}
.line {
	height: 1px;
	background-color: #cccccc;
	position: relative;
	top: 110rpx;
}
.skill-bar-box {
	position: relative;
	.skill-bar {
		position: relative;
		display: flex;
		justify-content: space-evenly;
		margin: 30rpx;
		.grid-item-box {
			padding: 20rpx;
			margin: 20rpx;
			border: 1px solid #9c9c9c;
			border-radius: 15rpx;
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			image {
				width: 70rpx;
			}
		}
	}
}
.video-box {
	display: flex;
	/* .video {
		position: absolute;
		left: 14%;
		height: 220px;
		width: 76vw;
		border-radius: 8px;
		background-color: #ccc;
		-webkit-transform: translateY(5%);
		transform: translateY(5%);
	} */
	.video {
		box-shadow: 20rpx 20rpx #ccc;
		display: flex;
		margin: auto;
		height: 400rpx;
		width: 82vw;
		border-radius: 15rpx;
	}
}
.introduce-box {
	margin: 100rpx 30rpx 0 30rpx;
	.title {
		font-weight: 600;
		font-size: 20px;
		margin-bottom: 20rpx;
	}
	.detail {
		padding: 20rpx;
		line-height: 50rpx;
		color: #1b223d;
		font-size: 18px;
	}
}
</style>