<template>
	<view class="main">
		<view class="header">
			<wwqy-header></wwqy-header>
		</view>
		<view class="swiper-box">
			<swiper class="swiper" :indicator-dots="true" :autoplay="true" :interval="5000" :duration="500">
				<swiper-item v-for="(banner, index) in homeBanner" :key="index">
					<image :src="banner"></image>
				</swiper-item>
			</swiper>
		</view>
		<view class="nav-box">
			<uni-grid :column="4" :show-border="false" :square="false">
				<uni-grid-item>
					<view class="grid-item-box" @click="gotoPage(0)">
						<view class="img-box">
							<img src="/static/home/icon/gun.png" style="height: 60rpx; margin: auto" alt="" />
						</view>
						<text class="nav-text">武器介绍</text>
					</view>
				</uni-grid-item>
				<uni-grid-item>
					<view class="grid-item-box" @click="gotoPage(1)">
						<view class="img-box">
							<img src="/static/home/icon/hero.png" style="height: 60rpx; margin: auto" alt="" />
						</view>
						<text class="nav-text">英雄介绍</text>
					</view>
				</uni-grid-item>
				<uni-grid-item>
					<view class="grid-item-box" @click="gotoPage(2)">
						<view class="img-box">
							<img src="/static/home/icon/map.png" style="height: 60rpx; margin: auto" alt="" />
						</view>
						<text class="nav-text">地图介绍</text>
					</view>
				</uni-grid-item>
				<uni-grid-item>
					<view class="grid-item-box" @click="gotoPage(3)">
						<view class="img-box">
							<img src="/static/home/icon/news.png" style="height: 60rpx; margin: auto" alt="" />
						</view>
						<text class="nav-text">游戏工具</text>
					</view>
				</uni-grid-item>
			</uni-grid>
		</view>
		<view class="body">
			<view class="about">
				<view class="pubTitle">
					<view class="en">VALORANT</view>
					<view class="cn">无畏契约</view>
					<view class="line"></view>
				</view>
				<view class="content">
					<view class="row">《无畏契约》国服，又名valorant、瓦罗兰特，是《英雄联盟》开发商拳头游戏开发、腾讯代理、风靡全球的PC端战术射击力作。</view>
					<view class="row">
						以英雄角色为核心的 5v5 的战术射击免费网游，采用FPS经典“爆破模式”作为核心玩法，玩家将化身技能各异的战术英雄，使用不同类型的武器枪械参与战斗。
					</view>
					<view class="line"></view>
				</view>
			</view>
		</view>
		<view class="news">
			<view class="pubTitle">
				<view class="en">News</view>
				<view class="cn">游戏资讯</view>
			</view>
			<view class="card-body" v-if="newsInfoData.length > 0" v-for="(it, index) in newsInfoData" @click="gotoDetail(it, index)">
				<view class="content">
					<image mode="widthFix" :src="it.image" alt="" />
					<view class="new-title">
						<text>{{ it.title }}</text>
						<view class="new-foot">{{ it.desc }}</view>
					</view>
				</view>
				<view class="parting-line"></view>
			</view>
		</view>
		<view class="footer">
			<wwqy-footer></wwqy-footer>
		</view>
		<uni-popup ref="message" type="message">
			<uni-popup-message type="error" :message="messageText" :duration="2000"></uni-popup-message>
		</uni-popup>
		<loading :loading="showLoading" :firstContent="firstContent" :secondContent="secondContent" />
	</view>
</template>

<script>
import wwqyHeader from '@/components/wwqy-header/index.vue';
import wwqyFooter from '@/components/wwqy-footer/index.vue';
import { newsList, getUserInfo, website } from '@/common/api/business.js';
export default {
	components: { wwqyHeader, wwqyFooter },
	data() {
		return {
			newsInfoData: [],
			showLoading: false, // 自定义遮罩显影
			firstContent: '', // 自定义遮罩第一行文本
			secondContent: '', // 自定义遮罩第二行文本
			messageText: '错误信息',
			homeBanner: ''
		};
	},

	onLoad() {
		this.website();
		this.newsList();
	},
	onShow() {},
	methods: {
		newsList() {
			let params = {
				pageNo: 1,
				pageSize: 3, // 只展示前三个
				title: '',
				cid: '',
				status: -1,
				recommended: 1
			};
			newsList(params).then((res) => {
				let result = res.data;
				if (result.code === 200) {
					this.newsInfoData = result.data.lists;
					this.setLoadingOptions(false);
				} else {
					that.messageText = result.msg;
					that.$refs.message.open();
				}
			});
		},
		website() {
			website().then((res) => {
				let result = res.data;
				if (result.code === 200) {
					this.homeBanner = result.data.homeBanner.split(',');
				} else {
					that.messageText = result.msg;
					that.$refs.message.open();
				}
			});
		},
		gotoPage(e) {
			uni.setStorageSync('switchTabKey', e);
			uni.switchTab({
				url: '/pages/information/index'
			});
		},
		gotoDetail(item) {
			uni.setStorageSync('newsDetail', item);
			uni.navigateTo({
				url: '../news/newsDetail',
				animationType: 'fade-in',
				animationDuration: 300
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
.swiper-box {
	padding: 30rpx;
	height: 410rpx;

	image {
		width: 100%;
		height: 350rpx;
		border-radius: 30rpx;
	}

	.swiper {
		padding-top: 30rpx;
		height: 410rpx;
		border-radius: 30rpx;
	}
}

/deep/.uni-swiper-dot-active {
	background-color: $globalColor !important;
	width: 30rpx;
	border-radius: 12rpx;
}

/deep/.uni-swiper-dots-horizontal {
	bottom: -1rpx;
}

.nav-box {
	.grid-item-box {
		margin: auto;
		padding: 10px 0px 10px 0px;
	}

	.img-box {
		display: flex;
	}
}

.body {
	background: url('/static/home/background.jpg') no-repeat;

	.about {
		padding: 60rpx 30rpx 60rpx;
		background-size: cover;

		.pubTitle {
			text-align: center;

			.en {
				font-size: 86rpx;
				font-weight: 900;
				text-transform: uppercase;
				color: $globalColor;
				opacity: 0.1;
				margin-bottom: 40rpx;
			}

			.cn {
				font-size: 56rpx;
				font-weight: 900;
				transform: translateY(-60rpx);
				color: $globalColor;
			}

			.line {
				width: 100rpx;
				height: 5rpx;
				background: $globalColor;
				opacity: 0.6;
				display: inline-block;
				transform: translateY(-40rpx);
			}
		}
	}

	.content {
		.row {
			line-height: 1.6em;
			text-indent: 2em;
			font-size: 32rpx;
			padding: 20rpx 0;
			border-bottom: 1rpx dashed #000;
			color: $focusColor;
		}

		.row:first-child {
			padding-top: 0;
		}

		.row:last-child {
			padding-bottom: 0;
			border-bottom: none;
		}
	}
}

.news {
	margin-top: 60rpx;

	.pubTitle {
		text-align: center;

		.en {
			font-size: 86rpx;
			font-weight: 900;
			text-transform: uppercase;
			color: $globalColor;
			opacity: 0.1;
			margin-bottom: 40rpx;
		}

		.cn {
			font-size: 56rpx;
			font-weight: 900;
			transform: translateY(-60rpx);
			color: $globalColor;
		}

		.line {
			width: 100rpx;
			height: 5rpx;
			background: $globalColor;
			opacity: 0.6;
			display: inline-block;
			transform: translateY(-40rpx);
		}
	}
	.card-body {
		position: relative;
		padding: 30rpx;
		.content {
			display: flex;

			image {
				margin: auto;
				width: 240rpx;
				margin-right: 10rpx;
			}

			.new-title {
				padding: 10rpx;
				font-size: 30rpx;
				color: $descColor;
				text {
					font-size: 30rpx;
					display: -webkit-box;
					-webkit-line-clamp: 2;
					-webkit-box-orient: vertical;
					overflow: hidden;
					text-overflow: ellipsis;
					width: 58vw;
					color: #565656;
				}
			}

			.new-foot {
				font-size: 26rpx;
				color: #999999;
				width: 60vw;
				overflow: hidden;
				white-space: nowrap;
				text-overflow: ellipsis;
			}
		}

		.parting-line {
			height: 2rpx;
			background-color: #e6e6e6;
			position: relative;
			top: 30rpx;
			right: 20rpx;
			border-radius: 10rpx;
		}
	}
}

/deep/.uni-highlight {
	color: $descColor !important;
}
</style>
