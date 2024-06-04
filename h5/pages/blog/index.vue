<template>
	<view class="main">
		<view class="content">
			<scroll-view scroll-y="true" style="height: 100%" @scrolltolower="onToButtom" :refresher-threshold="40">
				<view class="search-box">
					<img
						src="data:image/svg+xml;charset=utf-8;base64,PHN2ZyB0PSIxNzE1NjEyMDYzNDUxIiBjbGFzcz0iaWNvbiIgdmlld0JveD0iMCAwIDEwMjQgMTAyNCIgdmVyc2lvbj0iMS4xIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHAtaWQ9IjE1NDEiIHdpZHRoPSIxOCIgaGVpZ2h0PSIxOCI+PHBhdGggZD0iTTY0NC4yNjY2NjcgNjE0LjRsMTYyLjEzMzMzMyAxNjIuMTMzMzMzYzguNTMzMzMzIDguNTMzMzMzIDguNTMzMzMzIDIxLjMzMzMzMyAwIDI5Ljg2NjY2Ny04LjUzMzMzMyA4LjUzMzMzMy0yMS4zMzMzMzMgOC41MzMzMzMtMjkuODY2NjY3IDBsLTE2Mi4xMzMzMzMtMTYyLjEzMzMzM2MtNDYuOTMzMzMzIDM4LjQtMTAyLjQgNTkuNzMzMzMzLTE2Ni40IDU5LjczMzMzMy0xNDAuOCAwLTI1Ni0xMTUuMi0yNTYtMjU2czExNS4yLTI1NiAyNTYtMjU2YzExNS4yIDAgMjEzLjMzMzMzMyA3Ni44IDI0Ny40NjY2NjcgMTg3LjczMzMzMyA0LjI2NjY2NyAxMi44LTQuMjY2NjY3IDIxLjMzMzMzMy0xMi44IDI1LjYtMTIuOCA0LjI2NjY2Ny0yMS4zMzMzMzMtNC4yNjY2NjctMjUuNi0xMi44LTI1LjYtODkuNi0xMTAuOTMzMzMzLTE1My42LTIwNC44LTE1My42LTExOS40NjY2NjcgMC0yMTMuMzMzMzMzIDkzLjg2NjY2Ny0yMTMuMzMzMzM0IDIxMy4zMzMzMzRzOTMuODY2NjY3IDIxMy4zMzMzMzMgMjEzLjMzMzMzNCAyMTMuMzMzMzMzYzk4LjEzMzMzMyAwIDE3OS4yLTY0IDIwNC44LTE1My42IDQuMjY2NjY3LTEyLjggMTcuMDY2NjY3LTE3LjA2NjY2NyAyNS42LTEyLjggMTIuOCA0LjI2NjY2NyAxNy4wNjY2NjcgMTcuMDY2NjY3IDEyLjggMjUuNi0xMi44IDI5Ljg2NjY2Ny0yOS44NjY2NjcgNTkuNzMzMzMzLTUxLjIgODkuNnoiIHAtaWQ9IjE1NDIiPjwvcGF0aD48L3N2Zz4="
						alt=""
					/>
					<input v-model="searchValue" placeholder="搜索" />
					<uni-icons style="margin: auto" v-if="searchValue" size="24" type="close" @click="closeSearch" />
				</view>
				<view class="tag-cloud">
					<tag-cloud :rollList="rollList" @goSearch="search"></tag-cloud>
					<tag-cloud :rollList="rollList2" @goSearch="search"></tag-cloud>
				</view>
				<view class="content">
					<view class="list">
						<view class="articlesList" v-for="(item, index) in tableData" :key="index" @click="goToDetail(item)">
							<view class="image-box">
								<image mode="widthFix" :src="item.image" alt="" />
							</view>
							<view class="title">{{ item.title }}</view>
							<view class="introduction">{{ item.desc }}</view>
							<!-- <view class="icon">
								<img
									src="data:image/svg+xml;charset=utf-8;base64,PHN2ZyB0PSIxNzE1NTkzNTU3NjU3IiB2aWV3Qm94PSIwIDAgMTAyNCAxMDI0IiB2ZXJzaW9uPSIxLjEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgcC1pZD0iMzMwNiIgd2lkdGg9IjE4IiBoZWlnaHQ9IjE4Ij4KCQkJCQkJCQkJPHBhdGgKCQkJCQkJCQkJCWQ9Ik03NTAuOTMzMzMzIDY2NS42YzguNTMzMzMzLTguNTMzMzMzIDIxLjMzMzMzMy00LjI2NjY2NyAyOS44NjY2NjcgNC4yNjY2NjcgOC41MzMzMzMgOC41MzMzMzMgNC4yNjY2NjcgMjEuMzMzMzMzLTQuMjY2NjY3IDI5Ljg2NjY2Ni02OC4yNjY2NjcgNTUuNDY2NjY3LTE2Ni40IDg5LjYtMjY4LjggODkuNi0yMDAuNTMzMzMzIDAtMzYyLjY2NjY2Ny0xMjMuNzMzMzMzLTM2Mi42NjY2NjYtMjc3LjMzMzMzM1MzMTEuNDY2NjY3IDIzNC42NjY2NjcgNTEyIDIzNC42NjY2NjdzMzYyLjY2NjY2NyAxMjMuNzMzMzMzIDM2Mi42NjY2NjcgMjc3LjMzMzMzM2MwIDE3LjA2NjY2NyAwIDI5Ljg2NjY2Ny00LjI2NjY2NyA0Ni45MzMzMzMtNC4yNjY2NjcgMTIuOC0xMi44IDE3LjA2NjY2Ny0yNS42IDE3LjA2NjY2Ny0xMi44LTQuMjY2NjY3LTE3LjA2NjY2Ny0xMi44LTE3LjA2NjY2Ny0yNS42IDQuMjY2NjY3LTEyLjggNC4yNjY2NjctMjUuNiA0LjI2NjY2Ny0zOC40IDAtMTI4LTE0MC44LTIzNC42NjY2NjctMzIwLTIzNC42NjY2NjdTMTkyIDM4NCAxOTIgNTEyczE0MC44IDIzNC42NjY2NjcgMzIwIDIzNC42NjY2NjdjOTMuODY2NjY3IDAgMTc5LjItMjkuODY2NjY3IDIzOC45MzMzMzMtODEuMDY2NjY3eiIKCQkJCQkJCQkJCXAtaWQ9IjMzMDciCgkJCQkJCQkJCQlmaWxsPSIjYmFiZGJhIgoJCQkJCQkJCQk+PC9wYXRoPgoJCQkJCQkJCQk8cGF0aAoJCQkJCQkJCQkJZD0iTTUxMiA1NTQuNjY2NjY3YzI1LjYgMCA0Mi42NjY2NjctMTcuMDY2NjY3IDQyLjY2NjY2Ny00Mi42NjY2NjdzLTE3LjA2NjY2Ny00Mi42NjY2NjctNDIuNjY2NjY3LTQyLjY2NjY2Ny00Mi42NjY2NjcgMTcuMDY2NjY3LTQyLjY2NjY2NyA0Mi42NjY2NjcgMTcuMDY2NjY3IDQyLjY2NjY2NyA0Mi42NjY2NjcgNDIuNjY2NjY3eiBtMCA0Mi42NjY2NjZjLTQ2LjkzMzMzMyAwLTg1LjMzMzMzMy0zOC40LTg1LjMzMzMzMy04NS4zMzMzMzNzMzguNC04NS4zMzMzMzMgODUuMzMzMzMzLTg1LjMzMzMzMyA4NS4zMzMzMzMgMzguNCA4NS4zMzMzMzMgODUuMzMzMzMzLTM4LjQgODUuMzMzMzMzLTg1LjMzMzMzMyA4NS4zMzMzMzN6IgoJCQkJCQkJCQkJcC1pZD0iMzMwOCIKCQkJCQkJCQkJCWZpbGw9IiNiYWJkYmEiCgkJCQkJCQkJCT48L3BhdGg+CgkJCQkJCQkJPC9zdmc+"
									alt=""
								/>
								<text>123</text>
								<img
									src="data:image/svg+xml;charset=utf-8;base64,PHN2ZyB0PSIxNzE1NTkzOTE3NTE1IiBjbGFzcz0iaWNvbiIgdmlld0JveD0iMCAwIDEwMjQgMTAyNCIgdmVyc2lvbj0iMS4xIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHAtaWQ9IjEzODgiIHdpZHRoPSIxOCIgaGVpZ2h0PSIxOCI+CgkJCQkJCQkJCTxwYXRoCgkJCQkJCQkJCQlkPSJNNDg2LjQgNjgyLjY2NjY2N2M0LjI2NjY2NyAwIDQuMjY2NjY3IDAgMCAwSDcyNS4zMzMzMzNjNDYuOTMzMzMzIDAgODUuMzMzMzMzLTM4LjQgODUuMzMzMzM0LTg1LjMzMzMzNFYzNDEuMzMzMzMzYzAtNDYuOTMzMzMzLTM4LjQtODUuMzMzMzMzLTg1LjMzMzMzNC04NS4zMzMzMzNIMjk4LjY2NjY2N2MtNDYuOTMzMzMzIDAtODUuMzMzMzMzIDM4LjQtODUuMzMzMzM0IDg1LjMzMzMzM3YyNTZjMCA0Ni45MzMzMzMgMzguNCA4NS4zMzMzMzMgODUuMzMzMzM0IDg1LjMzMzMzNGgyMS4zMzMzMzNjMTIuOCAwIDIxLjMzMzMzMyA4LjUzMzMzMyAyMS4zMzMzMzMgMjEuMzMzMzMzcy04LjUzMzMzMyAyMS4zMzMzMzMtMjEuMzMzMzMzIDIxLjMzMzMzM0gyOTguNjY2NjY3Yy03Mi41MzMzMzMgMC0xMjgtNTUuNDY2NjY3LTEyOC0xMjhWMzQxLjMzMzMzM2MwLTcyLjUzMzMzMyA1NS40NjY2NjctMTI4IDEyOC0xMjhoNDI2LjY2NjY2NmM3Mi41MzMzMzMgMCAxMjggNTUuNDY2NjY3IDEyOCAxMjh2MjU2YzAgNzIuNTMzMzMzLTU1LjQ2NjY2NyAxMjgtMTI4IDEyOGgtMjMwLjRsLTE2Ni40IDgxLjA2NjY2N2MtOC41MzMzMzMgNC4yNjY2NjctMjEuMzMzMzMzIDAtMjkuODY2NjY2LTguNTMzMzMzLTQuMjY2NjY3LTguNTMzMzMzIDAtMjEuMzMzMzMzIDguNTMzMzMzLTI5Ljg2NjY2N2wxNzAuNjY2NjY3LTg1LjMzMzMzM2g4LjUzMzMzM3oiCgkJCQkJCQkJCQlwLWlkPSIxMzg5IgoJCQkJCQkJCQkJZmlsbD0iI2JhYmRiYSIKCQkJCQkJCQkJPjwvcGF0aD4KCQkJCQkJCQkJPHBhdGgKCQkJCQkJCQkJCWQ9Ik0zNDEuMzMzMzMzIDQ2OS4zMzMzMzNtLTQyLjY2NjY2NiAwYTQyLjY2NjY2NyA0Mi42NjY2NjcgMCAxIDAgODUuMzMzMzMzIDAgNDIuNjY2NjY3IDQyLjY2NjY2NyAwIDEgMC04NS4zMzMzMzMgMFoiCgkJCQkJCQkJCQlwLWlkPSIxMzkwIgoJCQkJCQkJCQkJZmlsbD0iI2JhYmRiYSIKCQkJCQkJCQkJPjwvcGF0aD4KCQkJCQkJCQkJPHBhdGgKCQkJCQkJCQkJCWQ9Ik01MTIgNDY5LjMzMzMzM20tNDIuNjY2NjY3IDBhNDIuNjY2NjY3IDQyLjY2NjY2NyAwIDEgMCA4NS4zMzMzMzQgMCA0Mi42NjY2NjcgNDIuNjY2NjY3IDAgMSAwLTg1LjMzMzMzNCAwWiIKCQkJCQkJCQkJCXAtaWQ9IjEzOTEiCgkJCQkJCQkJCQlmaWxsPSIjYmFiZGJhIgoJCQkJCQkJCQk+PC9wYXRoPgoJCQkJCQkJCQk8cGF0aAoJCQkJCQkJCQkJZD0iTTY4Mi42NjY2NjcgNDY5LjMzMzMzM20tNDIuNjY2NjY3IDBhNDIuNjY2NjY3IDQyLjY2NjY2NyAwIDEgMCA4NS4zMzMzMzMgMCA0Mi42NjY2NjcgNDIuNjY2NjY3IDAgMSAwLTg1LjMzMzMzMyAwWiIKCQkJCQkJCQkJCXAtaWQ9IjEzOTIiCgkJCQkJCQkJCQlmaWxsPSIjYmFiZGJhIgoJCQkJCQkJCQk+PC9wYXRoPgoJCQkJCQkJCTwvc3ZnPg=="
									alt=""
								/>
								<text>222</text>
							</view> -->
							<view class="introduction">{{ item.publishTime }}</view>
						</view>
						<view class="load-out">
							<view v-if="tableData.length > 0 || loadingStatus == 'no-more'">
								<uni-load-more iconType="circle" :status="loadingStatus" :contentText="contentText" />
							</view>
						</view>
					</view>
				</view>
			</scroll-view>
		</view>
		<view class="go-top" v-if="isShow">
			<uni-icons style="margin: auto" size="24" type="arrow-up" @click="goTop" />
		</view>
		<uni-popup ref="message" type="message">
			<uni-popup-message type="error" :message="messageText" :duration="2000"></uni-popup-message>
		</uni-popup>
		<loading :loading="showLoading" :firstContent="firstContent" :secondContent="secondContent" />
	</view>
</template>

<script>
import tagCloud from './components/tag-cloud.vue';
import { newsList, keyWordsLists } from '@/common/api/business.js';
export default {
	components: {
		tagCloud
	},
	data() {
		return {
			page: 1,
			// 加载信息
			loadingStatus: 'no-more',
			contentText: {
				contentdown: '上拉显示更多',
				contentrefresh: '加载中...',
				contentnomore: '没有更多数据了哦～'
			},
			showLoading: false, // 自定义遮罩显影
			firstContent: '', // 自定义遮罩第一行文本
			secondContent: '', // 自定义遮罩第二行文本
			messageText: '错误信息',
			tableData: [],
			searchValue: '', // 搜索值
			rollList: [],
			rollList2: [],
			animationData: {},
			isShow: false
		};
	},
	onLoad() {
		this.newsList();
		this.rollList = uni.getStorageSync('rollList');
		this.rollList2 = uni.getStorageSync('rollList2');
	},
	onPullDownRefresh() {
		if (this.searchValue){
			this.searchValue = ''
		} else {
			this.newsList();
		}
		this.rollList = uni.getStorageSync('rollList');
		this.rollList2 = uni.getStorageSync('rollList2');
	},
	onPageScroll(e) {
		if (e.scrollTop >= 800) {
			// 显示
			this.isShow = true;
		} else {
			// 隐藏
			this.isShow = false;
		}
	},
	beforeDestroy() {
		// 在组件销毁之前清除 debounce 定时器
		clearTimeout(this.debounceTimer);
		this.debounceTimer = null;
	},
	methods: {
		/* 列表请求数据开始 */
		newsList(id) {
			this.setLoadingOptions(true);
			let params = {
				title: this.searchValue,
				pageNo: this.page,
				pageSize: 10,
				keyWords: id || '',
				cid: '',
				status: -1,
				recommended: 1
			};
			newsList(params).then((res) => {
				uni.stopPullDownRefresh();
				let result = res.data;
				if (result.code === 200) {
					if (result.data.count === this.tableData.length) {
						// 数据已全部加载
						this.loadingStatus = 'no-more';
						uni.stopPullDownRefresh();
						this.setLoadingOptions(false);
						return;
					}
					// 拼接数组
					this.tableData = this.tableData.concat(result.data.lists);
					// 判断数组总量是否小于页面Size
					if (this.tableData.length == result.data.count) {
						this.loadingStatus = 'no-more';
					} else {
						this.loadingStatus = 'more';
					}
					uni.stopPullDownRefresh();
					this.setLoadingOptions(false);
				} else {
					this.messageText = result.msg;
					this.$refs.message.open();
				}
			});
		},

		/* 列表请求数据结束 */
		// 清除输入值
		closeSearch() {
			this.searchValue = null;
		},
		// 页面跳转
		goToDetail(item) {
			console.log(item);
			uni.setStorageSync('detail', item);
			uni.navigateTo({
				url: './detail',
				animationType: 'fade-in',
				animationDuration: 300
			});
			/* uni.showToast({
					title: '功能开发中',
					icon: 'none',
					duration: 1000
				}); */
		},
		// 滚动到底部
		onToButtom() {
			if (this.loadingStatus === 'no-more') {
				return;
			}
			this.page++;
			this.loadingStatus = 'loading';
			this.newsList();
		},
		search(item) {
			this.page = 1;
			this.tableData = [];
			console.log(item.id);
			this.newsList(item.id);
		},
		// 回到顶部
		goTop() {
			uni.pageScrollTo({
				scrollTop: 0, // 滚动到页面的目标位置  这个是滚动到顶部, 0
				duration: 500 // 滚动动画的时长
			});
		},
		// 设置遮罩数据
		setLoadingOptions(flag, firstText, secondText) {
			this.showLoading = flag || false;
			this.firstContent = firstText || '';
			this.secondContent = secondText || '';
		}
	},
	// 监听搜索栏
	watch: {
		searchValue: {
			handler(newValue) {
				clearTimeout(this.debounceTimer);
				this.debounceTimer = setTimeout(() => {
					this.page = 1;
					this.tableData = [];
					this.newsList();
				}, 200); // 设置防抖时间间隔，单位为200毫秒
			},
			immediate: false,
			deep: true
		}
	}
};
</script>

<style scoped lang="scss">
.main {
	background-color: #000;
}

.search-box {
	padding: 0 16rpx;
	background-color: #fff;
	border-radius: 15rpx;
	margin: 10rpx auto;
	width: 88%;
	display: flex;
	height: 75rpx;
	img {
		position: absolute;
		width: 60rpx;
		height: 75rpx;
		margin-right: 10rpx;
	}
	input {
		margin-left: 60rpx;
		height: 75rpx;
		font-size: 36rpx;
		color: #999;
		width: 90vw;
	}
}

.tag-cloud {
	height: 240rpx;
	background-color: #f5f6f7;
	border-top-left-radius: 40rpx;
	border-top-right-radius: 40rpx;
	display: flex;
	flex-wrap: wrap;
	width: 100%; /* 根据需要调整宽度  */
	overflow: hidden;
}

.list {
	background-color: #f5f6f7;
	.articlesList {
		padding: 0 0 30rpx 0;
		border-radius: 20rpx;
		margin: 0 30rpx 30rpx 30rpx;
		background-color: #fff;
		box-shadow: 5rpx 5rpx 13rpx 0rpx rgba(117, 126, 136, 0.2);
		.image-box {
			border-radius: 20rpx;
			overflow: hidden;
			height: 250rpx;
			width: 100%;
			image {
				width: 100%;
			}
		}
		.title {
			margin: 20rpx;
			font-size: 36rpx;
		}
		.introduction {
			margin: 20rpx;
			color: #babdba;
			// 超出两行省略
			text-overflow: -o-ellipsis-lastline;
			overflow: hidden;
			text-overflow: ellipsis;
			display: -webkit-box;
			-webkit-line-clamp: 2;
			line-clamp: 2;
			-webkit-box-orient: vertical;
		}
		.icon {
			margin-left: 20rpx;
			display: flex;
			img {
				width: 45rpx;
				height: 40rpx;
			}
			text {
				margin-right: 10rpx;
				line-height: 38rpx;
				color: #babdba;
			}
		}
	}
	.load-out {
		width: 100%;
	}
}
.go-top {
	display: flex;
	height: 90rpx;
	width: 90rpx;
	border-radius: 50%;
	background-color: rgba(0, 0, 0, 0.2);
	position: fixed;
	bottom: 130rpx;
	right: 60rpx;
}
// 滑动区域高度
.swiper-box {
	height: calc(100vh - 302rpx) !important;
}
</style>
