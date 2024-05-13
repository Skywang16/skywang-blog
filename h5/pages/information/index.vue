<template>
	<view class="main">
		<wwqy-header></wwqy-header>
		<view class="content">
			<scroll-view id="tab-bar" class="scroll-h" :scroll-x="true" :show-scrollbar="false" :scroll-into-view="scrollInto">
				<view
					v-for="(tab, index) in tabBars"
					:key="tab.id"
					:id="tab.id"
					class="uni-tab-item"
					:class="tabIndex == index ? 'uni-tab-item-active' : ''"
					:style="{ width: 100 / tabBars.length + '%' }"
					:data-current="index"
					@click="onTabTap"
				>
					<text class="uni-tab-item-title" :class="tabIndex == index ? 'uni-tab-item-title-active' : ''">{{ tab.name }}</text>
				</view>
			</scroll-view>
			<swiper :current="tabIndex" class="swiper-box" :duration="300" @change="onTabChange">
				<swiper-item class="swiper-item" v-for="tab in tabBars" :key="tab.name">
					<scroll-view scroll-y="true" style="height: 100%" @scrolltolower="onToButtom" :refresher-threshold="40">
						<view v-show="tabIndex === 0">
							<!-- 武器 -->
							<view class="content">
								<view class="list">
									<view class="weaponList" v-for="(weaponItem, index) in tableData" :key="index">
										<weapon :weaponItem="weaponItem"></weapon>
									</view>
								</view>
								<view class="loadOut">
									<view v-if="tableData.length > 0 || loadingStatus == 'no-more'">
										<uni-load-more iconType="circle" :status="loadingStatus" :contentText="contentText" />
									</view>
								</view>
							</view>
						</view>
						<!-- 角色 -->
						<view v-if="tabIndex === 1">
							<view class="content">
								<view class="list">
									<view class="weaponList" v-for="(heroItem, index) in tableData" :key="index" @click="goToDetail('hero')">
										<hero :heroItem="heroItem"></hero>
									</view>
								</view>
								<view class="loadOut">
									<view v-if="tableData.length > 0 || loadingStatus == 'no-more'">
										<uni-load-more iconType="circle" :status="loadingStatus" :contentText="contentText" />
									</view>
								</view>
							</view>
						</view>
						<!-- 地图 -->
						<view v-if="tabIndex === 2">
							<view class="content">
								<view class="list">
									<view class="weaponList" v-for="(mapItem, index) in tableData" :key="index" @click="goToDetail('map')">
										<maps :mapItem="mapItem"></maps>
									</view>
								</view>
								<view class="loadOut">
									<view v-if="tableData.length > 0 || loadingStatus == 'no-more'">
										<uni-load-more iconType="circle" :status="loadingStatus" :contentText="contentText" />
									</view>
								</view>
							</view>
						</view>
						<!-- 工具 -->
						<view v-if="tabIndex === 3">
							<view class="content">
								<view class="list">
									<view class="weaponList" v-for="(heroItem, index) in tableData" :key="index" @click="goToDetail('tool')">
										<hero :heroItem="heroItem"></hero>
									</view>
								</view>
								<view class="loadOut">
									<view v-if="tableData.length > 0 || loadingStatus == 'no-more'">
										<uni-load-more iconType="circle" :status="loadingStatus" :contentText="contentText" />
									</view>
								</view>
							</view>
						</view>
					</scroll-view>
				</swiper-item>
			</swiper>
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
import weapon from './components/weapon/index.vue';
import hero from './components/hero/index.vue';
import maps from './components/map/index.vue';
import tool from './components/weapon/index.vue';
import { weaponList, heroList, mapList } from '@/common/api/business.js';
export default {
	components: { wwqyHeader, wwqyFooter, weapon, hero, maps, tool },
	data() {
		return {
			page: 1,
			tabIndex: 0, // 当前选中的tab
			scrollInto: '', // tab位置
			tabBars: [
				{ id: 'a1', name: '武器' },
				{ id: 'a2', name: '英雄' },
				{ id: 'a3', name: '地图' },
				{ id: 'a4', name: '工具' }
			],
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
			tableData: []
		};
	},
	onLoad() {
		let index = uni.getStorageSync('switchTabKey');
		if (index === 0) {
			this.getWeaponList();
		} else if (index === 1) {
			this.getHeroList();
		} else if (index === 2) {
			this.getMapList();
		}
	},
	// 
	onShow() {
		this.tabIndex = uni.getStorageSync('switchTabKey') || 0;
	},
	// tab切换事件
	onTabItemTap() {
		let index = uni.getStorageSync('switchTabKey');
		if (index === 0) {
			this.getWeaponList();
		} else if (index === 1) {
			this.getHeroList();
		} else if (index === 2) {
			this.getMapList();
		}
	},
	methods: {
		/* 列表请求数据开始 */
		// 武器列表
		getWeaponList() {
			this.setLoadingOptions(true);
			let params = {
				pageNo: this.page,
				pageSize: 10,
				status: 1
			};
			weaponList(params).then((res) => {
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
		// 角色列表
		getHeroList() {
			this.setLoadingOptions(true);
			let params = {
				pageNo: this.page,
				pageSize: 10,
				status: 1
			};
			heroList(params).then((res) => {
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
		// 地图列表
		getMapList() {
			this.setLoadingOptions(true);
			let params = {
				pageNo: this.page,
				pageSize: 10,
				status: 1
			};
			mapList(params).then((res) => {
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

		/* tab事件开始 */
		// 点击头部标签
		onTabTap(e) {
			if (e !== undefined) {
				let index = e.target.dataset.current || e.currentTarget.dataset.current;
				// 当前下标改为点击的下标
				this.tabIndex = index;
			}
			this.page = 1;
			// 当前下标改为点击的下标
			// 当前tabbar选中Id改为点击的Id
			if (this.tabBars && this.tabBars[this.tabIndex]) {
				this.scrollInto = this.tabBars[this.tabIndex].id;
			}
		},
		// tab拖动切换事件
		onTabChange(e) {
			this.tableData = [];
			let index = e.detail.current;
			this.page = 1;
			// 当前下标改为点击的下标
			this.tabIndex = index;
			if (index === 1) {
				this.getHeroList();
			} else if (index === 0) {
				this.getWeaponList();
			} else if (index === 2) {
				this.getMapList();
			}
			uni.setStorageSync('switchTabKey', this.tabIndex);
			// 当前tabbar选中Id改为点击的Id
			this.scrollInto = this.tabBars[index].id;
		},
		/* tab事件结束 */

		// 页面跳转
		goToDetail(e) {
			if (e === 'hero') {
				uni.navigateTo({
					url: './components/hero/detail',
					animationType: 'fade-in',
					animationDuration: 300
				});
				/* uni.showToast({
					title: '功能开发中',
					icon: 'none',
					duration: 1000
				}); */
			}
		},
		// 滚动到底部
		onToButtom() {
			if (this.loadingStatus === 'no-more') {
				return;
			}
			this.page++;
			this.loadingStatus = 'loading';
			this.getWeaponList();
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

<style scoped lang="scss">
.content {
	.list {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-between;
	}
	.weaponList {
		width: 50%;
	}
	.loadOut {
		height: 100px;
	}
}
.uni-tab-item {
	display: inline-block;
	text-align: center;
	box-sizing: border-box;
}

.uni-tab-item-active {
	position: relative;
}

.uni-tab-item-active::after {
	content: '';
	position: absolute;
	left: 50%;
	transform: translateX(-50%);
	width: 40%;
	bottom: -12rpx;
	height: 3px;
	border-radius: 20rpx;
	background-color: $globalColor;
}

.uni-tab-item-title {
	color: #333333;
	font-size: 16px;
	height: 32px;
	line-height: 32px;
}

.uni-tab-item-title-active {
	color: $globalColor;
	font-size: 18px;
	font-weight: 600;
}
.scroll-h {
	background-color: rgb(235, 231, 226) !important;
	height: 40px;
	background-color: #fff;
	font-weight: 500;
	letter-spacing: 1px;
}
// 滑动区域高度
.swiper-box {
	background-size: 100%;
	background-size:cover;
	background-image: url('@/static/information/bj.webp');
	height: calc(100vh - 302rpx) !important;
}
</style>
