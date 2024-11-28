<template>
	<div class="article-container" ref="mainRef" @scroll="handleScroll">
		<!-- 返回按钮 -->
		<!-- <div class="back-btn" @click="goBack">
				<i class="icon-font icon-back"></i>
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
				<div class="article-desc">
					<TypeWriter :text="articleData.desc" :speed="10" class="desc" />
				</div>
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
						<!-- 	<p>{{ article.desc }}</p> -->
						</div>
					</div>
				</div>
			</div>
		</div>

		<!-- 添加目录树组件 -->
		<div class="article-catalog" v-show="catalogItems.length">
			<div class="catalog-header">
				<div class="catalog-title">目录</div>
				<div class="catalog-toggle" @click="toggleCatalog" :class="{ collapsed: isCollapsed }">
					<img style="height: 20px;" src="@/assets/icons/list.png" alt="">
				</div>
			</div>
			<div class="catalog-list" :class="{ collapsed: isCollapsed }">
				<div v-for="(item, index) in catalogItems" :key="index" class="catalog-item" :class="[
					`level-${item.level}`,
					{ active: currentHeading === item.id }
				]" @click="scrollToHeading(item.id)">
					{{ item.text }}
				</div>
			</div>
		</div>
	</div>
	<!-- 添加回顶部按钮 -->
	<Transition name="fade-slide">
		<div class="back-to-top" v-if="showBackToTop" @click="scrollToTop">
			<div class="back-to-top-content">
				<img style="height: 30px;" src="@/assets/icons/up.png" alt="" />
				<span class="tooltip">返回顶部</span>
			</div>
		</div>
	</Transition>
</template>

<script>
import { reactive, toRefs, ref, onMounted, computed, watch, onUnmounted, nextTick } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { articleDetail, newsList } from '@/api/business'
import TypeWriter from '@/components/TypeWriter.vue'
import hljs from 'highlight.js'
import 'highlight.js/styles/atom-one-dark.css'

export default {
	name: 'ArticleDetail',
	components: {
		TypeWriter
	},
	setup() {
		// 路由相关
		const router = useRouter()
		const route = useRoute()

		// 文章基础数据
		const viewCount = ref(1234)
		const likeCount = ref(88)
		const isLiked = ref(false)
		const loading = ref(false)

		// 文章详情数据
		const allData = reactive({
			articleData: {},
			relatedArticles: [],
			keyWordsMap: {}
		});

		// DOM引用
		const mainRef = ref(null);

		// 目录相关状态
		const catalogItems = ref([]);
		const currentHeading = ref('');
		const isCollapsed = ref(false);
		const showBackToTop = ref(false);

		// Observer实例
		let headingObserver = null;
		let imageObserver = null;

		// 清理函数集合
		const cleanupFunctions = new Set();

		// 计算属性 - 关键词列表
		const keyWordsList = computed(() => {
			if (!allData.articleData.keyWords) return [];
			return allData.articleData.keyWords.split(',')
				.filter(id => id)
				.map(id => allData.keyWordsMap[id] || id);
		});

		// 节流函数
		function throttle(fn, delay) {
			let timer = null;
			let lastTime = 0;
			return function (...args) {
				const now = Date.now();
				if (now - lastTime >= delay) {
					fn.apply(this, args);
					lastTime = now;
				} else if (!timer) {
					timer = setTimeout(() => {
						fn.apply(this, args);
						lastTime = now;
						timer = null;
					}, delay);
				}
			}
		}

		// 防抖函数
		const debounce = (fn, delay) => {
			let timer = null;
			return function (...args) {
				if (timer) clearTimeout(timer);
				timer = setTimeout(() => {
					fn.apply(this, args);
					timer = null;
				}, delay);
			};
		};

		// 处理文章点赞
		const handleLike = () => {
			if (!isLiked.value) {
				likeCount.value++
				isLiked.value = true
			}
		}

		// 返回上一页
		const goBack = () => {
			router.back()
		}

		// 获取文章详情
		const getArticleDetail = (id) => {
			loading.value = true
			articleDetail({ id }).then(res => {
				allData.articleData = res.data
				// 直接设置文档标题
				document.title = res.data.title || '文章详情'
				if (res.data.keyWords) {
					getRelatedArticles(res.data.keyWords)
				}
			}).finally(() => {
				loading.value = false
			})
		}

		// 获取相关文章推荐
		const getRelatedArticles = (keyWords) => {
			const firstKeyWord = keyWords.split(',')[0]
			const parameter = {
				pageNo: 1,
				pageSize: 3,
				keyWords: firstKeyWord,
				status: 1,
				recommended: -1
			}
			newsList(parameter).then(res => {
				if (res.code === 200) {
					allData.relatedArticles = res.data.lists.filter(item => item.id !== allData.articleData.id)
				}
			}).finally(() => {
				loading.value = false
			})
		}

		// 初始化关键词映射表
		const initKeyWordsMap = () => {
			const cachedKeyWordsList = localStorage.getItem('keyWordsList');
			if (cachedKeyWordsList) {
				const keyWordsList = JSON.parse(cachedKeyWordsList);
				const mapping = {};
				keyWordsList.forEach(item => {
					mapping[item.id] = item.name;
				});
				allData.keyWordsMap = mapping;
			}
		};

		// 处理文章点击事件
		const handleArticleClick = (id) => {
			router.push(`/article?id=${id}`);
			getArticleDetail(id);
			window.scrollTo(0, 0);
		}

		// 处理页面滚动事件
		const handleScroll = throttle(() => {
			showBackToTop.value = window.scrollY > 300;
			updateCurrentHeading();
			handleFastScroll();
		}, 100);

		// 滚动到页面顶部
		const scrollToTop = () => {
			window.scrollTo({
				top: 0,
				behavior: 'smooth'
			});
		};

		// 设置标题观察器
		const setupHeadingObserver = () => {
			const observer = new IntersectionObserver(
				(entries) => {
					entries.forEach(entry => {
						if (entry.isIntersecting) {
							currentHeading.value = entry.target.id;
						}
					});
				},
				{
					rootMargin: '-100px 0px -90% 0px'
				}
			);
			catalogItems.value.forEach(item => {
				const element = document.getElementById(item.id);
				if (element) {
					observer.observe(element);
				}
			});
			return observer;
		};

		// 生成文章目录
		const generateCatalog = () => {
			nextTick(() => {
				const content = document.querySelector('.content');
				if (!content) return;

				// 等待内容渲染完成
				setTimeout(() => {
					const headings = content.querySelectorAll('h1, h2, h3, h4, h5, h6');
					const items = [];

					headings.forEach((heading, index) => {
						const id = `heading-${index}`;
						heading.id = id;
						items.push({
							id,
							text: heading.textContent.trim(),
							level: parseInt(heading.tagName.charAt(1))
						});
					});

					catalogItems.value = items;

					// 重新设置观察器
					if (headingObserver) {
						headingObserver.disconnect();
					}
					headingObserver = setupHeadingObserver();

					// 默认选中第一个标题
					if (items.length > 0) {
						currentHeading.value = items[0].id;
					}
				}, 100); // 给予一定延迟确保内容已渲染
			});
		};

		// 滚动到指定标题位置
		const scrollToHeading = (id) => {
			const element = document.getElementById(id);
			if (element) {
				element.scrollIntoView({ behavior: 'smooth' });
			}
		};

		// 更新当前目录项
		const updateCurrentHeading = debounce(() => {
			const headings = catalogItems.value.map(item => document.getElementById(item.id));
			const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
			if (scrollTop < 100) {
				currentHeading.value = headings[0]?.id || '';
				return;
			}
			for (let i = headings.length - 1; i >= 0; i--) {
				const heading = headings[i];
				if (heading && heading.getBoundingClientRect().top <= 100) {
					currentHeading.value = heading.id;
					break;
				}
			}
		}, 100);

		// 处理快速滚动
		const handleFastScroll = () => {
			const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
			if (scrollTop < 100) {
				const firstHeading = catalogItems.value[0]?.id;
				if (firstHeading) {
					currentHeading.value = firstHeading;
				}
			}
		};

		// 切换目录展开/收起状态
		const toggleCatalog = () => {
			isCollapsed.value = !isCollapsed.value;
		};

		// 设置图片懒加载
		const setupLazyLoading = () => {
			const imageObserver = new IntersectionObserver(
				(entries) => {
					entries.forEach(entry => {
						if (entry.isIntersecting) {
							const img = entry.target;
							img.src = img.dataset.src;
							imageObserver.unobserve(img);
						}
					});
				},
				{
					rootMargin: '50px 0px'
				}
			);
			const images = document.querySelectorAll('.content img');
			images.forEach(img => {
				if (img.src) {
					img.dataset.src = img.src;
					img.src = 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7';
					imageObserver.observe(img);
				}
			});
			return imageObserver;
		};

		// 添加清理函数
		const addCleanup = (fn) => {
			cleanupFunctions.add(fn);
		};

		// 监听路由变化
		watch(() => route.query.id, (newId) => {
			if (newId) {
				getArticleDetail(newId);
			}
		});

		// 监听文章内容变化
		watch(() => allData.articleData.content, () => {
			nextTick(() => {
				generateCatalog();
				setupCodeCopy();
			});
		}, { immediate: true });

		// 组件挂载
		onMounted(() => {
			initKeyWordsMap();
			const id = route.query.id
			if (id) {
				getArticleDetail(id)
			}
			window.addEventListener('scroll', handleScroll, { passive: true });
			generateCatalog();
			imageObserver = setupLazyLoading();

			const scrollHandler = handleScroll;
			window.addEventListener('scroll', scrollHandler, { passive: true });
			addCleanup(() => window.removeEventListener('scroll', scrollHandler));
			const observer = setupHeadingObserver();
			addCleanup(() => observer.disconnect());
			const imgObserver = setupLazyLoading();
			addCleanup(() => imgObserver.disconnect());
			setupCodeCopy();
		});

		// 组件卸载
		onUnmounted(() => {
			window.removeEventListener('scroll', handleScroll);
			if (headingObserver) {
				headingObserver.disconnect();
				headingObserver = null;
			}
			if (imageObserver) {
				imageObserver.disconnect();
				imageObserver = null;
			}
			cleanupFunctions.forEach(cleanup => cleanup());
			cleanupFunctions.clear();
		});

		// 设置代码复制功能
		const setupCodeCopy = () => {
			nextTick(() => {
				const codeBlocks = document.querySelectorAll('pre code');
				codeBlocks.forEach(block => {
					// 应用代码高亮
					hljs.highlightElement(block);

					// 获取语言类名并转换为驼峰格式
					const language = block.className.match(/language-(\w+)/)?.[1] || '';
					// 为每个代码块创建复制按钮和语言标识
					const codeWrapper = document.createElement('div');
					codeWrapper.className = 'code-wrapper';

					const languageTag = document.createElement('span');
					languageTag.className = 'language-tag';
					languageTag.textContent = language;

					const copyButton = document.createElement('button');
					copyButton.className = 'code-copy-btn';
					copyButton.innerHTML = '复制';

					// 包装代码块
					block.parentNode.insertBefore(codeWrapper, block);
					codeWrapper.appendChild(languageTag);
					codeWrapper.appendChild(block);
					codeWrapper.appendChild(copyButton);

					// 添加点击事件
					copyButton.addEventListener('click', () => {
						const code = block.textContent;
						navigator.clipboard.writeText(code).then(() => {
							copyButton.innerHTML = '已复制';
							setTimeout(() => {
								copyButton.innerHTML = '复制';
							}, 2000);
						}).catch(err => {
							console.error('复制失败:', err);
							copyButton.innerHTML = '复制失败';
							setTimeout(() => {
								copyButton.innerHTML = '复制';
							}, 2000);
						});
					});
				});
			});
		};

		return {
			loading,
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
			catalogItems,
			currentHeading,
			scrollToHeading,
			isCollapsed,
			toggleCatalog,
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
				background-clip: text;
				-webkit-text-fill-color: transparent;
				letter-spacing: -0.5px;
			}

			.article-desc {
				.desc {
					display: block;
					font-size: 16px;
					color: #4a5568;
					line-height: 1.8;
					margin: 25px 0;
					padding: 25px 30px;
					background: #f8fafc;
					border-radius: 20px;
					position: relative;
					transition: all 0.3s ease;
					border: 1px solid #e2e8f0;

					// 顶部装饰条
					&::before {
						content: '';
						position: absolute;
						top: 0;
						left: 50%;
						transform: translateX(-50%);
						width: 70px;
						height: 4px;
						background: linear-gradient(90deg, #60a5fa, #3b82f6);
						border-radius: 0 0 4px 4px;
					}
				}
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
					border-radius: 10px;
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

				// 基础元素样式
				::v-deep {

					// 图片样式
					img {
						max-width: 100%;
						height: auto;
						display: block;
						margin: 20px auto;
						border-radius: 12px;
						box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
						transition: transform 0.3s ease;
						object-fit: contain;

						&:hover {
							transform: scale(1.02);
						}

						@media screen and (max-width: 768px) {
							border-radius: 8px;
							margin: 15px auto;
						}
					}

					// 段落样式
					p {
						margin-bottom: 16px;
						line-height: 1.8;

						&[style*="text-indent"] {
							text-indent: 2em;
						}

						&+p {
							margin-top: 1em;
						}
					}

					// 标题样式
					h1,
					h2,
					h3,
					h4,
					h5,
					h6 {
						margin: 1.6em 0 0.8em;
						line-height: 1.4;
						color: #2c3e50;
						font-weight: 600;
					}

					h1 {
						font-size: 2em;
					}

					h2 {
						font-size: 1.75em;
					}

					h3 {
						font-size: 1.5em;
					}

					h4 {
						font-size: 1.25em;
					}

					h5 {
						font-size: 1.1em;
					}

					h6 {
						font-size: 1em;
					}

					// 列表样式
					ul,
					ol {
						padding-left: 2em;
						margin: 1em 0;

						li {
							margin-bottom: 0.5em;
							line-height: 1.8;
						}
					}

					// 行内代码样式
					code:not(pre code) {
						background: rgba(0, 0, 0, 0.06);
						padding: 2px 6px;
						border-radius: 4px;
						font-size: 0.9em;
						color: #476582;
						font-family: Consolas, Monaco, 'Andale Mono', monospace;
					}

					// 引用块样式
					blockquote {
						margin: 20px 0;
						padding: 20px;
						background: rgba(245, 247, 250, 0.8);
						border-left: 4px solid #1890ff;
						border-radius: 8px;
						color: #5c6b77;
						font-size: 16px;
						line-height: 1.8;
						position: relative;

						&::before {
							content: '"';
							position: absolute;
							top: -20px;
							left: 10px;
							font-size: 60px;
							color: rgba(24, 144, 255, 0.1);
							font-family: Georgia, serif;
						}

						p {
							margin: 0;
							position: relative;
							z-index: 1;

							&+p {
								margin-top: 16px;
							}
						}

						br {
							display: block;
							margin: 0.5em 0;
							content: "";
						}
					}

					// 文本样式
					strong,
					b {
						font-weight: 600;
						color: #2c3e50;
					}

					em,
					i {
						font-style: italic;
					}

					del,
					s {
						text-decoration: line-through;
						color: #999;
					}

					u {
						text-decoration: underline;
						text-underline-offset: 2px;
					}

					mark {
						background-color: #ffeaa7;
						padding: 0 4px;
						border-radius: 2px;
					}

					// 代码块样式
					pre {
						background: #282c34;
						border-radius: 12px;
						box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
						margin: 16px 0;
						overflow-x: auto;
						position: relative;

						code {
							color: #abb2bf;
							font-family: Consolas, Monaco, 'Andale Mono', monospace;
							font-size: 14px;
							line-height: 1.6;
							tab-size: 4;

							// 代码高亮
							.hljs {
								&-keyword {
									color: #c678dd;
								}

								&-string {
									color: #98c379;
								}

								&-number {
									color: #d19a66;
								}

								&-function {
									color: #61afef;
								}

								&-comment {
									color: #5c6370;
									font-style: italic;
								}

								&-variable {
									color: #e06c75;
								}

								&-attr {
									color: #d19a66;
								}

								&-built_in {
									color: #56b6c2;
								}

								&-title {
									color: #61afef;
								}

								&-type {
									color: #e5c07b;
								}
							}
						}

						// 复制按钮
						.code-copy-btn {
							position: absolute;
							right: 10px;
							top: 10px;
							padding: 6px 12px;
							background: rgba(255, 255, 255, 0.1);
							border: 1px solid rgba(255, 255, 255, 0.2);
							border-radius: 4px;
							color: #abb2bf;
							font-size: 12px;
							cursor: pointer;
							transition: all 0.3s ease;
							opacity: 0;

							&:hover {
								background: rgba(255, 255, 255, 0.2);
							}
						}

						&:hover .code-copy-btn {
							opacity: 1;
						}
					}

					// 代码语言标识
					pre[class*="language-"] {
						&::before {
							content: attr(class);
							position: absolute;
							top: 0;
							left: 20px;
							font-size: 12px;
							color: #abb2bf;
							background: rgba(255, 255, 255, 0.1);
							padding: 3px 10px;
							border-radius: 0 0 4px 4px;
							text-transform: uppercase;
						}
					}

					// 代码块容器样式
					.code-wrapper {
						position: relative;
						margin: 16px 0;
						background: #282c34;
						border-radius: 12px;
						box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);

						pre {
							margin: 0;
							padding: 40px 20px 20px;
							background: transparent;
							box-shadow: none;

							code {
								padding: 0;
								background: transparent;
								font-family: 'Fira Code', Consolas, Monaco, 'Andale Mono', monospace;
								font-size: 14px;
								line-height: 1.6;
								tab-size: 4;
							}
						}

						// 语言标签样式
						.language-tag {
							position: absolute;
							top: -10px;
							left: 10px;
							background: rgba(255, 255, 255, 0.1);
							color: #abb2bf;
							padding: 0 10px;
							font-size: 12px;
							border-radius: 5px;
							font-family: 'Fira Code', monospace;
						}

						// 复制按钮样式
						.code-copy-btn {
							position: absolute;
							right: 10px;
							top: 10px;
							padding: 6px 12px;
							background: rgba(255, 255, 255, 0.1);
							border: 1px solid rgba(255, 255, 255, 0.2);
							border-radius: 4px;
							color: #abb2bf;
							font-size: 12px;
							cursor: pointer;
							transition: all 0.3s ease;
							opacity: 0;

							&:hover {
								background: rgba(255, 255, 255, 0.2);
							}
						}

						&:hover .code-copy-btn {
							opacity: 1;
						}
					}

					// highlight.js 的自定义主题覆盖
					.hljs {
						background: transparent;
						color: #abb2bf;

						&-keyword {
							color: #c678dd;
						}

						&-string {
							color: #98c379;
						}

						&-number {
							color: #d19a66;
						}

						&-function {
							color: #61afef;
						}

						&-comment {
							color: #5c6370;
							font-style: italic;
						}

						&-variable {
							color: #e06c75;
						}

						&-attr {
							color: #d19a66;
						}

						&-built_in {
							color: #56b6c2;
						}

						&-title {
							color: #61afef;
						}

						&-type {
							color: #e5c07b;
						}
					}

					// 添加表格样式
					table {
						width: 100%;
						margin: 16px 0;
						border-collapse: collapse;
						border-spacing: 0;
						border-radius: 15px;
						overflow: hidden;
						background: #fff;
						box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);

						th, td {
							padding: 12px 16px;
							text-align: left;
							border: 1px solid #f0f0f0;
						}

						th {
							background: #fafafa;
							font-weight: 600;
							color: #1a1a1a;
						}

						tr {
							&:nth-child(even) {
								background: #fafafa;
							}

							&:hover {
								background: #f5f5f5;
							}
						}

						// 响应式处理
						@media screen and (max-width: 768px) {
							display: block;
							overflow-x: auto;
							-webkit-overflow-scrolling: touch;
							// 添加左右滑动提示
							&::before {
								content: '← 左右滑动查看更多 →';
								display: block;
								text-align: center;
								font-size: 12px;
								color: #999;
								padding: 8px 0;
							}

							th, td {
								padding: 8px 12px;
								min-width: 120px; // 确保单元格有最小宽度
								white-space: nowrap; // 防止内容换行
							}
						}

						// 超小屏幕适配
						@media screen and (max-width: 480px) {
							font-size: 14px;
							th, td {
								padding: 6px 8px;
								min-width: 100px;
							}
						}
					}

					// 表格容器
					.table-container {
						margin: 16px 0;
						overflow-x: auto;
						border-radius: 8px;
						background: #fff;
						box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
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
							background-clip: text;
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
	animation: float 3s ease-in-out infinite;

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
}


// 添加全局滚动优化
html {
	scroll-behavior: smooth;
	-webkit-overflow-scrolling: touch; // 优化iOS滚动
	overscroll-behavior-y: none; // 防止过度滚动
}

/* 添加目录树样式 */
.article-catalog {
	position: fixed;
	right: 20px;
	top: 100px;
	width: 250px;
	max-height: calc(100vh - 300px);
	overflow-y: auto;
	background: rgba(255, 255, 255, 0.95);
	border-radius: 12px;
	padding: 20px;
	box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
	transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
	transform-origin: right top;
	backdrop-filter: blur(10px);

	.catalog-header {
		position: sticky;
		top: -20px;
		background: rgba(255, 255, 255, 0.95);
		z-index: 2;
		padding-top: 20px;
		margin-top: -20px;
		backdrop-filter: blur(10px);
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 12px;
		padding-bottom: 8px;
		border-bottom: 1px solid #eee;

		.catalog-title {
			font-size: 16px;
			font-weight: 600;
			transition: all 0.3s;
			white-space: nowrap;
			overflow: hidden;
		}

		.catalog-toggle {
			cursor: pointer;
			width: 24px;
			height: 24px;
			display: flex;
			align-items: center;
			justify-content: center;
			border-radius: 4px;
			transition: all 0.3s ease;

			img {
				transition: transform 0.3s ease;
				transform: rotate(0deg);
			}

			&.collapsed img {
				transform: rotate(135deg);
			}
		}
	}

	.catalog-list {
		transition: all 0.3s ease;
		opacity: 1;
		max-height: calc(100vh - 400px);
		overflow: hidden;

		&.collapsed {
			max-height: 0;
			opacity: 0;
			margin: 0;
			padding: 0;
		}

		.catalog-item {
			padding: 6px 12px;
			margin: 4px 0;
			cursor: pointer;
			border-radius: 6px;
			transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
			font-size: 14px;
			position: relative;
			overflow: hidden;

			&::before {
				content: '';
				position: absolute;
				left: 0;
				top: 0;
				height: 100%;
				width: 0;
				background: rgba(24, 144, 255, 0.1);
				border-radius: 6px;
				transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1);
				z-index: -1;
			}

			&:hover::before {
				width: 100%;
			}

			&:hover {
				color: #1890ff;
			}

			&.active {
				color: #1890ff;
				font-weight: 500;

				&::before {
					width: 100%;
				}
			}

			&.level-1 {
				padding-left: 12px;
				font-weight: 600;
			}

			&.level-2 {
				padding-left: 24px;
			}

			&.level-3 {
				padding-left: 36px;
			}

			&.level-4 {
				padding-left: 48px;
			}

			&.level-5 {
				padding-left: 60px;
			}

			&.level-6 {
				padding-left: 72px;
			}
		}
	}

	&:hover {
		box-shadow: 0 4px 20px rgba(0, 0, 0, 0.12);
	}
}

// 添加展开/收起图标的动画
.icon-collapse {
	transform: rotate(0deg);
}

.icon-expand {
	transform: rotate(180deg);
}

// 调整文章内容宽度,为目录留出空
.article-container {
	max-width: calc(960px + 300px);
}

// 响应式处理
@media screen and (max-width: 1200px) {
	.article-catalog {
		display: none !important;
	}

	.article-container {
		max-width: 960px;
	}
}
</style>
