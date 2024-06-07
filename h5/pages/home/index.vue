<template>
	<view class="main">
		<view class="swiper-box">
			<swiper class="swiper" :indicator-dots="true" :autoplay="true" :interval="5000" :duration="500">
				<swiper-item v-for="(banner, index) in homeBanner" :key="index">
					<image :src="banner"></image>
				</swiper-item>
			</swiper>
		</view>
		<view class="body">
			<view class="recommended">
				<view class="pubTitle">
					<view class="en">Recommended</view>
					<view class="cn">推荐文章</view>
					<view class="line"></view>
				</view>
				<view class="card-body" v-if="newsInfoData.length > 0" v-for="(it, index) in newsInfoData" :key="index" @click="gotoDetail(it, index)">
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
		</view>
		<view class="comment">
			<view class="pubTitle">
				<view class="en">comment</view>
				<view class="cn">留言</view>
				<view class="line"></view>
			</view>
			<scroll-view scroll-x="true">
				<view class="card-body" v-if="commentsData.length > 0" v-for="(it, index) in commentsData" :key="index" @click="addComment(it)">
					<view class="content" v-if="!it.isAdd">
						<image :src="it.image" alt="" />
						<view class="new-title">
							<text>{{ it.title }}</text>
							<view class="new-time">{{ it.desc }}</view>
							<view class="new-foot">{{ it.desc }}</view>
						</view>
					</view>
					<view class="content" v-else>
						<image :src="it.image" alt="" />
						<uni-icons style="margin: auto; position: relative; left: 84rpx" color="#fd4453" size="80" type="plus" />
					</view>
				</view>
			</scroll-view>
		</view>
		<uni-popup ref="alertDialog" type="dialog">
			<uni-popup-dialog
				type="warn"
				cancelText="取消"
				confirmText="前往登录"
				title="通知"
				content="评论需要登录!是否前往登录。(评论功能未上线)"
				@confirm="dialogConfirm"
				@close="dialogClose"
			></uni-popup-dialog>
		</uni-popup>
		<uni-popup ref="popup" type="center">
			<view class="popup-content">
				<uni-easyinput class="input" type="textarea" v-model="introduction" placeholder="请输入评论" />
				<button class="btn" hover-class="btn-hover" @click="confirm">评论</button>
			</view>
		</uni-popup>
		<uni-popup ref="message" type="message">
			<uni-popup-message type="error" :message="messageText" :duration="2000"></uni-popup-message>
		</uni-popup>
		<loading :loading="showLoading" :firstContent="firstContent" :secondContent="secondContent" />
	</view>
</template>

<script>
import { newsList, getUserInfo, website, keyWordsLists } from '@/common/api/business.js';
export default {
	data() {
		return {
			newsInfoData: [],
			commentsData: [],
			showLoading: false, // 自定义遮罩显影
			firstContent: '', // 自定义遮罩第一行文本
			secondContent: '', // 自定义遮罩第二行文本
			messageText: '错误信息',
			homeBanner: '',
			introduction: ''
		};
	},
	onLoad() {
		this.website();
		this.newsList();
		this.keyWordsList();
		this.commentsData.push({ isAdd: true });
	},
	onPullDownRefresh(){
		this.website();
		this.newsList();
		this.keyWordsList();
	},
	onShow() {},
	methods: {
		// 加载文章列表
		newsList() {
			let params = {
				pageNo: 1,
				pageSize: 3, // 只展示前三个
				title: '',
				keyWords: '',
				cid: '',
				status: -1,
				recommended: 1
			};
			newsList(params).then((res) => {
				uni.stopPullDownRefresh()
				let result = res.data;
				if (result.code === 200) {
					this.newsInfoData = result.data.lists;
					// this.commentsData = JSON.parse(JSON.stringify(result.data.lists));
					// if (result.data.lists.length > 0) {
					// 	this.commentsData.push({ isAdd: true });
					// }
					this.setLoadingOptions(false);
				} else {
					this.messageText = result.msg;
					this.$refs.message.open();
				}
			});
		},
		keyWordsList() {
			let params = {
				pageNo: 1,
				pageSize: 100
			};
			keyWordsLists(params).then((res) => {
				let result = res.data;
				if (result.code === 200) {
					let list = result.data.lists;
					let rollList = list.map((item) => {
						const size = Math.floor(Math.random() * 9) + 18; // 生成 18 到 26 之间的随机数
						const color = '#' + Math.floor(Math.random() * 16777215).toString(16); // 生成随机颜色
						return {
							id: item.id,
							text: item.name,
							size: size,
							color: color
						};
					});
					function shuffleArray(array) {
						for (let i = array.length - 1; i > 0; i--) {
							const j = Math.floor(Math.random() * (i + 1));
							[array[i], array[j]] = [array[j], array[i]];
						}
						return array;
					}
					// 克隆 rollList 数组
					let rollList2 = [...rollList];
					rollList2 = shuffleArray(rollList2);
					uni.setStorageSync('rollList', rollList);
					uni.setStorageSync('rollList2', rollList2);
				} else {
					this.messageText = result.msg;
					this.$refs.message.open();
				}
			});
		},
		// 加载网站配置
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
		// 跳转到文章详情
		gotoDetail(item) {
			uni.setStorageSync('detail', item)
			uni.navigateTo({
				url: '../blog/detail',
				animationType: 'fade-in',
				animationDuration: 300
			});
		},
		// 新增留言
		addComment(it) {
			let user = uni.getStorageSync('userData');
			if (user.id) {
				if (it.isAdd) {
					this.$refs.popup.open('');
				}
			} else {
				this.$refs.alertDialog.open();
			}
		},
		dialogConfirm() {
			uni.navigateTo({
				url: '../login/index',
				animationType: 'fade-in',
				animationDuration: 300
			});
		},
		dialogClose() {
			this.$refs.alertDialog.close();
		},
		// 确认评论
		confirm() {
			uni.showToast({
				title: '评论功能未开放',
				icon: 'none'
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
	background: url('/static/home/background.jpg') no-repeat;
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
		image {
			margin: auto;
			width: 60rpx;
		}
		display: flex;
	}
}

.body {
	.recommended {
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
				transform: translateY(-100rpx);
				color: $globalColor;
			}

			.line {
				width: 100rpx;
				height: 5rpx;
				background: $globalColor;
				opacity: 0.6;
				display: inline-block;
				transform: translateY(-100rpx);
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
						font-size: 32rpx;
						display: -webkit-box;
						-webkit-line-clamp: 1;
						-webkit-box-orient: vertical;
						overflow: hidden;
						text-overflow: ellipsis;
						width: 60vw;
						color: #565656;
					}
				}

				.new-foot {
					font-size: 26rpx;
					color: #999999;
					width: 60vw;
					text-overflow: -o-ellipsis-lastline;
					overflow: hidden;
					text-overflow: ellipsis;
					display: -webkit-box;
					-webkit-line-clamp: 2;
					line-clamp: 2;
					-webkit-box-orient: vertical;
				}
			}

			.parting-line {
				height: 2rpx;
				background-color: #e6e6e6;
				position: relative;
				top: 30rpx;
				border-radius: 10rpx;
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

.comment {
	padding-bottom: 100rpx;
	margin-top: 60rpx;
	.add-comment {
		height: 50rpx;
		width: 50rpx;
		float: right;
		position: relative;
		bottom: 130rpx;
		right: 30rpx;
	}
	scroll-view {
		white-space: nowrap;
		width: 100%;
	}
	.card-body {
		height: 200rpx;
		width: 70%;
		display: inline-flex;
		margin: 0 40rpx;
		border: 1px solid #ccc;
		border-radius: 20rpx;
		position: relative;
		padding: 20rpx;
		.content {
			display: flex;
			image {
				width: 100rpx;
				height: 100rpx;
				border-radius: 50%;
				margin-right: 10rpx;
			}
			.new-title {
				padding: 10rpx;
				font-size: 30rpx;
				color: $descColor;
				text {
					font-size: 30rpx;
					color: #565656;
				}
			}
			.new-time {
				font-size: 26rpx;
				color: #999999;
			}
			.new-foot {
				margin-top: 10rpx;
				font-size: 26rpx;
				color: #999999;
				width: 54vw;
				white-space: normal;
				text-overflow: -o-ellipsis-lastline;
				overflow: hidden;
				text-overflow: ellipsis;
				display: -webkit-box;
				-webkit-line-clamp: 3;
				line-clamp: 3;
				-webkit-box-orient: vertical;
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
			transform: translateY(-100rpx);
			color: $globalColor;
		}

		.line {
			width: 100rpx;
			height: 5rpx;
			background: $globalColor;
			opacity: 0.6;
			display: inline-block;
			transform: translateY(-100rpx);
		}
	}
}
.popup-content {
	background-color: #fff;
	padding: 40rpx;
	border-radius: 30rpx;
	width: 80vw;
	.input {
		margin-bottom: 30rpx;
	}
	.btn {
		width: 50%;
		background-color: $globalColor;
		color: white;
		border-radius: 10rpx;
		margin-top: 40rpx;
	}
	.btn-hover {
		position: relative;
		top: 3rpx;
		box-shadow: 0px 0px 8px #999 inset;
	}
}
/deep/.uni-highlight {
	color: $descColor !important;
}
/deep/.uni-icons.uniui-clear {
	font-size: 24px !important;
}
</style>
