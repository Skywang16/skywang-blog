<template>
	<view class="main">
		<view class="header">
			<view class="image-box">
				<view class="title">{{ data.title }}</view>
				<!-- <image mode="widthFix" src="/static/7b58525685c7fa8f58fac97575111bf.png" alt="" /> -->
				<view class="headImg">
					<image src="/static/headImg.jpg" alt="" />
					<view class="name">{{ data.cid }}</view>
					<view class="time">{{ data.publishTime }}</view>
				</view>
			</view>
		</view>
		<view class="content">
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
			<mp-html :content="data.content" lazy-load scroll-table selectable use-anchor :tag-style="tag_style"></mp-html>
			<!-- <view v-html="data.content">
				
			</view> -->
		</view>
		<view class="footer">
			<!-- <view class="next">
				上一篇：
				<text>无</text>
			</view> -->
			<view style="height: 20rpx"></view>
		</view>
	</view>
</template>

<script>
import MpHtml from '../../components/mp-html/mp-html.vue';
import MdHighlight from '../../components/cmder-MDParserHighlight/index.vue';
export default {
	components: {
		MpHtml,
		MdHighlight
	},
	data() {
		return {
			data: {},
			tag_style: {
				table: 'box-sizing: border-box; border-top: 1px solid #dfe2e5; border-left: 1px solid #dfe2e5;',
				th: 'border-right: 1px solid #dfe2e5; border-bottom: 1px solid #dfe2e5;',
				td: 'border-right: 1px solid #dfe2e5; border-bottom: 1px solid #dfe2e5;',
				li: 'margin: 5px 0;',
				h1: 'line-height: 50px;font-size:16px;',
				h2: 'line-height: 50px;font-size:16px',
				h3: 'line-height: 50px;font-size:16px',
				h4: 'line-height: 50px;font-size:16px'
			},
			container_style: 'font-size:14px;line-height: 26px;overflow-x: hidden;'
		};
	},
	onLoad() {
		this.data = uni.getStorageSync('detail');
		this.data.content = this.data.content
			.replace(/<h1([\s\w"=\/\.:;]+)((?:(style="[^"]+")))/gi, '<h1')
			.replace(/<h1([\s\w"=\/\.:;]+)((?:(class="[^"]+")))/gi, '<h1')
			.replace(/<h1>/gi, '<h1 class="h1">')

			.replace(/<h2([\s\w"=\/\.:;]+)((?:(style="[^"]+")))/gi, '<h2')
			.replace(/<h2([\s\w"=\/\.:;]+)((?:(class="[^"]+")))/gi, '<h2')
			.replace(/<h2>/gi, '<h2 class="h2">')

			.replace(/<h3([\s\w"=\/\.:;]+)((?:(style="[^"]+")))/gi, '<h3')
			.replace(/<h3([\s\w"=\/\.:;]+)((?:(class="[^"]+")))/gi, '<h3')
			.replace(/<h3>/gi, '<h3 class="h3">')

			.replace(/<h4([\s\w"=\/\.:;]+)((?:(style="[^"]+")))/gi, '<h4')
			.replace(/<h4([\s\w"=\/\.:;]+)((?:(class="[^"]+")))/gi, '<h4')
			.replace(/<h4>/gi, '<h4 class="h4">')

			.replace(/<h5([\s\w"=\/\.:;]+)((?:(style="[^"]+")))/gi, '<h5')
			.replace(/<h5([\s\w"=\/\.:;]+)((?:(class="[^"]+")))/gi, '<h5')
			.replace(/<h5>/gi, '<h5 class="h4">')

			.replace(/<h6([\s\w"=\/\.:;]+)((?:(style="[^"]+")))/gi, '<h6')
			.replace(/<h6([\s\w"=\/\.:;]+)((?:(class="[^"]+")))/gi, '<h6')
			.replace(/<h6>/gi, '<h6 class="h6">')

			.replace(/<code([\s\w"=\/\.:;]+)((?:(style="[^"]+")))/gi, '<code')
			.replace(/<code([\s\w"=\/\.:;]+)((?:(class="[^"]+")))/gi, '<code')
			.replace(/<code>/gi, '<code class="language-css">');
		/* .replace(/\<td/g,'<td style="border:1px solid #999;') */
		/* .replace(/\<th/g,'<th style="background-color: #d4d4d4;'); */
	}
};
</script>

<style scoped lang="scss">
.main {
	overflow: hidden;
	background-color: #f5f6f7;
}
.header {
	.image-box {
		background-image: url('@/static/7b58525685c7fa8f58fac97575111bf.png');
		background-size: 100%;
		height: 400rpx;
		padding: 30rpx;
		.title {
			height: 210rpx;
			width: 100%;
			font-size: 58rpx;
			line-height: 70rpx;
			margin: 100rpx 0 55rpx 0;
			overflow: hidden;
			text-overflow: ellipsis;
			-webkit-line-clamp: 3;
			display: -webkit-box;
			-webkit-box-orient: vertical;
		}
		.headImg {
			display: flex;
			image {
				border-radius: 15rpx;
				height: 90rpx;
				width: 90rpx;
			}
			.name {
				margin: 20rpx;
				color: #000;
				font-size: 32rpx;
			}
			.time {
				position: relative;
				left: 300rpx;
				margin: 20rpx 0 0 30rpx;
				color: #000;
				font-size: 32rpx;
			}
		}
	}
}
.content {
	.icon {
		position: relative;
		left: 67%;
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
	padding: 30rpx;
}
.footer {
	padding: 30rpx;
	.next {
		color: #999;
		margin-top: 40rpx;
		margin-bottom: 40rpx;
	}
}
</style>
