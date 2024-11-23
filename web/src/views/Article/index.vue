<template>
	<div>
		<div class="article-container" ref="mainRef" @scroll="handleScroll">
			<!-- 返回按钮 -->
			<!-- <div class="back-btn" @click="goBack">
				<i class="iconfont icon-back"></i>
				返回
			</div> -->
			<div class="article-content">
				<!-- 文章头部信息 -->
				<div class="article-header">
					<h1 class="title">{{ articleData.title }}</h1>
					<div class="article-meta">
						<span class="publish-time">
							<i class="iconfont icon-time"></i>
							{{ articleData.publishTime }}
						</span>
						<span class="category">
							作者：{{ articleData.cid }}
						</span>
						<span class="recommend" v-if="articleData.recommended">
							推荐文章
						</span>
						<!-- <span class="views">
						<i class="iconfont icon-eye"></i>
						{{ viewCount }} 阅读
					</span> -->
					</div>
					<div class="cover-image" v-if="articleData.image">
						<img :src="articleData.image" :alt="articleData.title">
					</div>
					<div class="desc">{{ articleData.desc }}</div>
				</div>

				<!-- 文章主体内容 -->
				<div class="article-body">
					<div class="content markdown-body" v-html="articleData.content"></div>
				</div>

				<!-- 文章底部 -->
				<div class="article-footer">
					<!-- <div class="interaction">
					<div class="like-btn" @click="handleLike" :class="{ active: isLiked }">
						<i class="iconfont icon-like"></i>
						<span>{{ likeCount }}</span>
					</div>
					<div class="share-btn">
						<i class="iconfont icon-share"></i>
						分享
					</div>
				</div> -->

					<div class="tags" v-if="articleData.keyWords">
						<div class="tags-title">
							<i class="iconfont icon-tag"></i>
							标签：
						</div>
						<span class="tag" v-for="tag in keyWordsList" :key="tag">
							{{ tag }}
						</span>
					</div>
				</div>

				<!-- 相关推荐 -->
				<div class="related-articles" v-if="relatedArticles.length > 0">
					<h3 class="section-title">相关推荐</h3>
					<div class="article-cards">
						<div class="article-card" v-for="article in relatedArticles" :key="article.id"
							@click="handleArticleClick(article.id)">
							<div class="card-cover">
								<img v-if="article.image" :src="article.image" :alt="article.title">
							</div>
							<div class="card-info">
								<h4>{{ article.title }}</h4>
								<p>{{ article.desc }}</p>
							</div>
						</div>
					</div>
				</div>
			</div>


		</div>
		<!-- 添加回到顶部按钮 -->
		<Transition name="fade-slide">
			<div class="back-to-top" v-if="showBackToTop" @click="scrollToTop">
				<div class="back-to-top-content">
					<img style="height: 30px;" src="@/assets/icons/up.png" alt="" />
					<span class="tooltip">返回顶部</span>
				</div>
			</div>
		</Transition>
	</div>

</template>

<script>
import { reactive, toRefs, ref, onMounted, computed, watch, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { articleDetail, newsList } from '@/api/business'

export default {
	name: 'ArticleDetail',
	setup() {
		const router = useRouter()
		const route = useRoute()
		const viewCount = ref(1234)
		const likeCount = ref(88)
		const isLiked = ref(false)

		const allData = reactive({
			articleData: {},
			relatedArticles: [],
			keyWordsMap: {}
		});

		const keyWordsList = computed(() => {
			if (!allData.articleData.keyWords) return [];
			return allData.articleData.keyWords.split(',')
				.filter(id => id)
				.map(id => allData.keyWordsMap[id] || id);
		});

		const handleLike = () => {
			if (!isLiked.value) {
				likeCount.value++
				isLiked.value = true
			}
		}

		const goBack = () => {
			router.back()
		}

		const getArticleDetail = (id) => {
			articleDetail({ id }).then(res => {
				allData.articleData = res.data
				if (res.data.keyWords) {
					getRelatedArticles(res.data.keyWords)
				}
			})
		}

		const getRelatedArticles = (keyWords) => {
			const firstKeyWord = keyWords.split(',')[0]

			const parameter = {
				pageNo: 1,
				pageSize: 3,  // 只取3篇相关文章
				keyWords: firstKeyWord,
				status: 1,
				recommended: 1
			}

			newsList(parameter).then(res => {
				if (res.code === 200) {
					// 过滤掉当前文章
					allData.relatedArticles = res.data.lists.filter(item => item.id !== allData.articleData.id)
				}
			})
		}

		const initKeyWordsMap = () => {
			const cachedKeyWordsList = localStorage.getItem('keyWordsList');
			if (cachedKeyWordsList) {
				const keyWordsList = JSON.parse(cachedKeyWordsList);
				// 创建id到name的映射
				const mapping = {};
				keyWordsList.forEach(item => {
					mapping[item.id] = item.name;
				});
				allData.keyWordsMap = mapping;
			}
		};

		const handleArticleClick = (id) => {
			router.push(`/article?id=${id}`);
			getArticleDetail(id);
			window.scrollTo(0, 0);
		}

		watch(
			() => route.query.id,
			(newId) => {
				if (newId) {
					getArticleDetail(newId);
				}
			}
		);

		onMounted(() => {
			initKeyWordsMap();
			const id = route.query.id
			if (id) {
				getArticleDetail(id)
			}

			// 添加全局滚动监听
			window.addEventListener('scroll', handleScroll);
		});

		onUnmounted(() => {
			window.removeEventListener('scroll', handleScroll);
		});

		const mainRef = ref(null);
		const showBackToTop = ref(false);

		// 修改滚动监听方法
		const handleScroll = () => {
			showBackToTop.value = window.scrollY > 300;
		};

		// 修改返回顶部方法
		const scrollToTop = () => {
			window.scrollTo({
				top: 0,
				behavior: 'smooth'
			});
		};

		return {
			...toRefs(allData),
			keyWordsList,
			viewCount,
			likeCount,
			isLiked,
			handleLike,
			goBack,
			handleArticleClick,
			mainRef,
			showBackToTop,
			scrollToTop,
			handleScroll,
		}
	},
}
</script>
<style lang="scss">
// 滚动条样式
::-webkit-scrollbar {
	width: 8px;
	height: 8px;
}

::-webkit-scrollbar-track {
	background: #f1f1f1;
	border-radius: 4px;
}

::-webkit-scrollbar-thumb {
	background: #c1c1c1;
	border-radius: 4px;

	&:hover {
		background: #a8a8a8;
	}
}
</style>
<style scoped lang='scss'>
.article-container {
	max-width: 960px;
	margin: 0 auto;
	padding: 20px;
	position: relative;

	&::before {
		content: '';
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background: rgba(255, 255, 255, 0.8);
		z-index: 0;
		pointer-events: none;
	}

	.article-content {
		position: relative;
		z-index: 1;
		background: rgba(255, 255, 255, 0.95);
		box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
		will-change: transform;
		transform: translateZ(0);
		perspective: 1000;
		backface-visibility: hidden;
		border-radius: 20px;
		box-shadow:
			0 10px 30px rgba(31, 38, 135, 0.1),
			0 1px 2px rgba(0, 0, 0, 0.05);
		padding: 40px;
		margin-top: 20px;
		border: 1px solid rgba(255, 255, 255, 0.3);

		.article-header {
			margin-bottom: 40px;
			position: relative;

			.title {
				font-size: 38px;
				background: linear-gradient(135deg, #2c3e50 0%, #3498db 100%);
				-webkit-background-clip: text;
				-webkit-text-fill-color: transparent;
				letter-spacing: -0.5px;
			}

			.desc {
				font-size: 16px;
				color: #666;
				line-height: 1.8;
				margin: 25px 0;
				padding: 20px;
				background: rgba(240, 242, 245, 0.8);
				border-radius: 12px;
				border-left: 5px solid #1890ff;
				box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.05);
			}

			.article-meta {
				display: flex;
				align-items: center;
				gap: 20px;
				color: #8c8c8c;
				font-size: 14px;

				span {
					display: flex;
					align-items: center;
					gap: 6px;
					background: rgba(0, 0, 0, 0.03);
					padding: 6px 12px;
					border-radius: 20px;
					transition: all 0.3s ease;

					&:hover {
						background: rgba(0, 0, 0, 0.06);
					}

					i {
						font-size: 16px;
					}
				}

				.recommend {
					background: rgba(255, 77, 79, 0.1);
					color: #ff4d4f;
					font-weight: 500;
				}
			}

			.cover-image {
				margin: 24px 0;
				border-radius: 12px;
				overflow: hidden;
				box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);

				img {
					width: 100%;
					height: auto;
					object-fit: cover;
					transition: transform 0.3s;

				}
			}
		}

		.article-body {
			.content {
				font-size: 16px;
				line-height: 1.9;
				color: #2c3e50;

				::v-deep(img) {
					max-width: 100% !important;
					height: auto !important;
					border-radius: 12px;
					box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
					transition: transform 0.3s ease;
					display: block;
					margin: 20px auto;
					object-fit: contain;

					@media screen and (max-width: 768px) {
						border-radius: 8px;
						margin: 15px auto;
					}

					&:hover {
						transform: scale(1.02);
					}
				}

				::v-deep(p) {
					img {
						margin: 0 auto;
					}
				}

				::v-deep(h1, h2, h3, h4) {
					margin: 24px 0 16px;
					font-weight: 600;
					color: #1a1a1a;
				}

				::v-deep(p) {
					margin-bottom: 16px;
				}

				::v-deep(pre) {
					background: #282c34;
					padding: 20px;
					border-radius: 12px;
					box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
					margin: 16px 0;
					overflow-x: auto;

					code {
						color: #abb2bf; // 代码文本颜色

						// 代码中的关键字
						.hljs-keyword {
							color: #c678dd;
						}

						// 字符串
						.hljs-string {
							color: #98c379;
						}

						// 数字
						.hljs-number {
							color: #d19a66;
						}

						// 函数名
						.hljs-function {
							color: #61afef;
						}

						// 注释
						.hljs-comment {
							color: #5c6370;
							font-style: italic;
						}

						// 变量名
						.hljs-variable {
							color: #e06c75;
						}
					}
				}

				::v-deep(code):not(pre code) {
					background: rgba(0, 0, 0, 0.06);
					padding: 2px 6px;
					border-radius: 4px;
					font-size: 0.9em;
					color: #476582;
				}

				::v-deep(table) {
					width: 100%;
					margin: 16px 0;
					border-collapse: collapse;
					border-spacing: 0;
					
					th, td {
						padding: 12px;
						border: 1px solid #e8e8e8;
						text-align: left;
					}
					
					th {
						background-color: #fafafa;
						font-weight: 600;
						color: #1a1a1a;
					}
					
					tr {
						&:hover {
							background-color: #fafafa;
						}
					}
					
					tbody tr:nth-child(even) {
						background-color: #fafcff;
					}
				}
			}
		}

		.article-footer {
			margin-top: 40px;
			padding-top: 20px;
			border-top: 1px solid #f0f0f0;

			.interaction {
				display: flex;
				justify-content: center;
				gap: 40px;
				margin-bottom: 30px;

				.like-btn,
				.share-btn {
					display: flex;
					align-items: center;
					gap: 8px;
					padding: 8px 24px;
					border-radius: 24px;
					cursor: pointer;
					transition: all 0.3s;

					i {
						font-size: 20px;
					}
				}

				.like-btn {
					background: #f5f5f5;
					color: #595959;

					&.active {
						background: #ff4d4f;
						color: white;
					}

					&:hover {
						transform: scale(1.05);
					}
				}

				.share-btn {
					background: #e6f7ff;
					color: #1890ff;

					&:hover {
						background: #1890ff;
						color: white;
					}
				}
			}

			.tags {
				display: flex;
				flex-wrap: wrap;
				gap: 12px;
				align-items: center;

				.tags-title {
					color: #8c8c8c;
					display: flex;
					align-items: center;
					gap: 4px;
				}

				.tag {
					padding: 8px 18px;
					background: rgba(24, 144, 255, 0.08);
					color: #1890ff;
					border-radius: 20px;
					font-weight: 500;
					transition: all 0.3s ease;

					/* &:hover {
						background: #1890ff;
						color: white;
						transform: translateY(-2px);
						box-shadow: 0 4px 12px rgba(24, 144, 255, 0.2);
					} */
				}
			}
		}

		.related-articles {
			margin-top: 60px;

			.section-title {
				font-size: 24px;
				font-weight: 600;
				margin-bottom: 24px;
				color: #1a1a1a;
			}

			.article-cards {
				display: grid;
				grid-template-columns: repeat(3, 1fr);
				gap: 20px;

				.article-card {
					background: rgba(255, 255, 255, 0.9);
					backdrop-filter: blur(10px);
					border: 1px solid rgba(255, 255, 255, 0.2);
					border-radius: 12px;
					overflow: hidden;
					box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
					transition: all 0.4s ease;
					cursor: pointer;

					&:hover {
						transform: translateY(-8px);
						box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);

						.card-cover img {
							transform: scale(1.08);
						}
					}

					.card-cover {
						height: 160px;
						background: #f5f5f5;
						overflow: hidden;

						img {
							width: 100%;
							height: 100%;
							object-fit: cover;
							transition: transform 0.6s ease;

							&:hover {
								transform: scale(1.05);
							}
						}
					}

					.card-info {
						padding: 16px;

						h4 {
							font-size: 16px;
							font-weight: 600;
							margin-bottom: 8px;
							color: #1a1a1a;
							background: linear-gradient(120deg, #2c3e50 0%, #3498db 100%);
							-webkit-background-clip: text;
							-webkit-text-fill-color: transparent;
						}

						p {
							font-size: 14px;
							color: #8c8c8c;
							line-height: 1.5;
						}
					}
				}
			}
		}
	}

	&::after {
		content: '';
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background-image:
			radial-gradient(circle at 80% 10%, rgba(94, 114, 235, 0.1) 0%, transparent 20%),
			radial-gradient(circle at 20% 90%, rgba(249, 146, 146, 0.1) 0%, transparent 20%);
		pointer-events: none;
		z-index: 0;
	}
}

// 响应式设计
@media screen and (max-width: 768px) {
	.article-container {
		padding: 16px;

		.back-btn {
			position: static;
			margin-bottom: 20px;
		}

		.article-content {
			.article-header {
				.title {
					font-size: 26px;
				}
			}

			.article-body {
				.content {
					font-size: 15px;
				}
			}

			.related-articles {
				.article-cards {

					grid-template-columns: 1fr;
				}
			}
		}
	}
}

/* 添加回到顶部按钮样式 */
.fade-slide-enter-active,
.fade-slide-leave-active {
	transition: all 0.3s ease;
}

.fade-slide-enter-from,
.fade-slide-leave-to {
	opacity: 0;
	transform: translateY(20px);
}

.back-to-top {
	position: fixed;
	right: 40px;
	bottom: 100px;
	width: 50px;
	height: 50px;
	background: linear-gradient(135deg, #00c6fb 0%, #005bea 100%);
	border-radius: 50%;
	display: flex;
	align-items: center;
	justify-content: center;
	cursor: pointer !important;
	z-index: 999;
	box-shadow: 0 4px 15px rgba(0, 0, 0, 0.15);
	transition: all 0.3s ease;

	.back-to-top-content {
		position: relative;
		display: flex;
		align-items: center;
		justify-content: center;
		width: 100%;
		height: 100%;

		.tooltip {
			position: absolute;
			right: 140%;
			background: rgba(0, 0, 0, 0.8);
			color: white;
			padding: 5px 10px;
			border-radius: 4px;
			font-size: 14px;
			opacity: 0;
			visibility: hidden;
			transition: opacity 0.3s ease;
			white-space: nowrap;
			pointer-events: none;

			&::after {
				content: '';
				position: absolute;
				right: -6px;
				top: 50%;
				transform: translateY(-50%);
				border-left: 6px solid rgba(0, 0, 0, 0.8);
				border-top: 6px solid transparent;
				border-bottom: 6px solid transparent;
			}
		}
	}

	&:hover {
		transform: scale(1.05);
		box-shadow: 0 6px 20px rgba(0, 98, 234, 0.4);
		background: linear-gradient(135deg, #00a1fb 0%, #0046ea 100%);

		.tooltip {
			opacity: 1;
			visibility: visible;
		}
	}

	&:active {
		transform: scale(0.95);
	}
}

/* 移动端适配 */
@media (max-width: 768px) {
	.back-to-top {
		width: 40px;
		height: 40px;

		.back-to-top-content {
			i {
				font-size: 20px;
			}
		}

		.tooltip {
			display: none; // 移动端不显示提示文字
		}
	}
}

.back-btn {
	position: fixed;
	left: 20px;
	top: 20px;
	padding: 10px 20px;
	background: rgba(255, 255, 255, 0.9);
	border-radius: 25px;
	box-shadow: 0 4px 15px rgba(0, 0, 0, 0.08);
	cursor: pointer;
	transition: all 0.3s ease;
	display: flex;
	align-items: center;
	gap: 8px;
	z-index: 10;
	border: 1px solid rgba(255, 255, 255, 0.2);
	backdrop-filter: blur(8px);
	color: #1890ff;
	font-weight: 500;

	i {
		font-size: 18px;
		transition: transform 0.3s ease;
	}

	&:hover {
		transform: translateX(-5px);
		background: rgba(255, 255, 255, 0.95);
		box-shadow: 0 6px 20px rgba(24, 144, 255, 0.15);
		color: #0960bd;

		i {
			transform: translateX(-3px);
		}
	}

	&:active {
		transform: translateX(-2px) scale(0.98);
	}

	// 添加一个小动画效果
	@keyframes float {
		0% {
			transform: translateY(0px);
		}

		50% {
			transform: translateY(-2px);
		}

		100% {
			transform: translateY(0px);
		}
	}

	animation: float 3s ease-in-out infinite;
}

// 移动端适配
@media screen and (max-width: 768px) {
	.back-btn {
		position: static;
		margin: 10px 0 20px;
		padding: 8px 16px;
		font-size: 14px;

		i {
			font-size: 16px;
		}

		&:hover {
			transform: none;
		}

		animation: none;
	}
}

// 添加全局滚动优化
html {
	scroll-behavior: smooth;
	-webkit-overflow-scrolling: touch; // 优化iOS滚动
	overscroll-behavior-y: none; // 防止过度滚动
}
</style>