<template>
	<view class="main">
		<wwqy-header></wwqy-header>
		<view class="news">
			<view class="news-list">
				<scroll-view scroll-y="false" style="height: 100%" @scrolltolower="onToButtom" :refresher-threshold="40">
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
				</scroll-view>
			</view>
			<view class="loadOut">
				<view v-if="newsInfoData.length > 0 || loadingStatus == 'noMore'">
					<uni-load-more iconType="circle" :status="loadingStatus" :contentText="contentText" />
				</view>
			</view>
		</view>
		<wwqy-footer></wwqy-footer>
		<uni-popup ref="message" type="message">
			<uni-popup-message type="error" :message="messageText" :duration="2000"></uni-popup-message>
		</uni-popup>
		<loading :loading="showLoading" :firstContent="firstContent" :secondContent="secondContent" />
	</view>
</template>

<script>
import wwqyHeader from '@/components/wwqy-header/index.vue';
import wwqyFooter from '@/components/wwqy-footer/index.vue';
import { newsList } from '@/common/api/business.js';
export default {
	components: { wwqyHeader, wwqyFooter },
	data() {
		return {
			newsInfoData: [],
			loadingStatus: 'more',
			contentText: {
				contentdown: '上拉显示更多',
				contentrefresh: '加载中...',
				contentnomore: '没有更多数据了哦～'
			},
			showLoading: false, // 自定义遮罩显影
			firstContent: '', // 自定义遮罩第一行文本
			secondContent: '', // 自定义遮罩第二行文本
			messageText: '错误信息',
			page: 1
		};
	},

	onLoad() {
		this.newsList();
	},
	onShow() {},
	methods: {
		newsList() {
			this.setLoadingOptions(true);
			let params = {
				pageNo: this.page,
				pageSize: 10,
				title: '',
				cid: '',
				status: 1,
				recommended: -1
			};
			newsList(params).then((res) => {
				let result = res.data;
				if (result.code === 200) {
					if (result.data.count === this.newsInfoData.length) {
						// 数据已全部加载
						this.loadingStatus = 'noMore';
						uni.stopPullDownRefresh();
						this.setLoadingOptions(false);
						return;
					}
					// 拼接数组
					this.newsInfoData = this.newsInfoData.concat(result.data.lists);
					// 判断数组总量是否小于页面Size
					if (this.newsInfoData.length == result.data.count) {
						this.loadingStatus = 'noMore';
					} else {
						this.loadingStatus = 'more';
					}
					uni.stopPullDownRefresh();
					this.setLoadingOptions(false);
				} else {
					that.messageText = result.msg;
					that.$refs.message.open();
				}
			});
		},
		gotoDetail(item) {
			uni.setStorageSync('newsDetail', item);
			uni.navigateTo({
				url: './newsDetail',
				animationType: 'fade-in',
				animationDuration: 300
			});
		},
		// 滚动到底部
		onToButtom() {
			if (this.loadingStatus === 'no-more') {
				return;
			}
			this.page++;
			this.loadingStatus = 'loading';
			this.allNewsFn();
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
.card-body {
	position: relative;
	background-color: #fff;
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
</style>