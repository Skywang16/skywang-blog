<template>
	<text :class="['iconfont',customClass]" :style="{ 'color': color, 'font-size': size }">
		{{ icon }}
		<!-- &#xe6e0 -->
	</text>
</template>

<script>
	// 引入图标名称，以及对应的unicode
	import icons from './icons'
	export default {
		props: {
			// 图标类名
			name: {
				type: String
			},
			// 字体大小，注意加上单位，例如：12px
			size: {
				type: String
			},
			// 图标颜色
			color: {
				type: String,
				default: '#919399'
			},
			// 自定义class类名
			customClass: {
				type: String,
				default: ''
			}
		},
		computed: {
			// 通过图标名，查找对应的图标
			icon() {
				// 如果内置的图标中找不到对应的图标，就直接返回name值，因为用户可能传入的是unicode代码
				return icons['mu-icon-' + this.name] || this.name
			}
		},
	}
</script>

<style lang="scss" scoped>
	@font-face {
	  font-family: 'iconfont';
	  src: url('iconfont.ttf') format('truetype');
	  // src: url('~@/static/font/iconfont.ttf/iconfont.woff2') format('woff2'),
	  //      url('~@/static/font/iconfont.woff') format('woff'),
	  //      url('~@/static/font/iconfont.ttf') format('truetype');
	}
	
	.iconfont {
	  font-family: "iconfont" !important;
	  font-size: 16px;
	  font-style: normal;
	  -webkit-font-smoothing: antialiased;
	  -moz-osx-font-smoothing: grayscale;
	}
	
	// .iconfont:before {
	// 	// &#xe63b;
	// 	content: '\e63b';
	// }
</style>