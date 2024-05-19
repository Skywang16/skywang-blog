<template>
	<!-- <div class="modal-box" v-if="loading">
		<div class="loading-modal" :style="secondContent === '' ? 'height: 100px' : 'height: 110px'">
			<div class="loader loader-three">
				<image class="loader-img" src="@/static/home/icon/hero.png"></image>
			</div>
			<div class="loader-text loader-first">{{ firstContent === '' ? '加载中' : firstContent }}</div>
			<div class="loader-text loader-second" v-if="secondContent !== ''">{{ secondContent }}</div>
		</div>
	</div> -->
</template>

<script>
	export default {
		props: {
			loading: {
				type: Boolean,
				default () {
					return false
				}
			},
			firstContent: {
				type: String,
				default () {
					return '登录成功'
				}
			},
			secondContent: {
				type: String,
				default () {
					return ''
				}
			}
		}
	}
</script>

<style>
	.modal-box{
		position: fixed;
		display: block;
		width: 10000px;
		height: 10000px;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		margin: auto;
		z-index: 100;
	}
	.loading-modal{
		display: inline-block;
		height: 110px;
		width: 150px;
		background-color: rgba(0, 0, 0, 0.5);
		border-radius: 5px;
		box-shadow: 0 0 0px 1000px rgba(0, 0, 0, 0.2);
		font-size: 14px;
		z-index: 9999;
		position: fixed;
		top: 0;
		right: 0;
		bottom: 70px;
		left: 0;
		margin: auto;
	}
	.loader-three {
		position: relative;
		width: 100px;
		height: 48px;
		margin-left: auto;
		margin-right: auto;
		margin-top: 15px;
	}

	.loader-three:after {
		position: absolute;
		top: 0px;
		left: 26px;
		content: "";
		width: 36px;
		height: 36px;
		border: solid transparent 3px;
		border-top: solid #fff 3px;
		border-left: solid #fff 3px;
		border-right: solid #fff 3px;
		border-radius: 50%;
		animation: loader3AnimationAfter 1.5s linear infinite;
	}
	.loader-text{
		color: #fff;
		text-align: center;
		line-height: 20px;
	}
	.loader-img{
		width: 30px;
		height: 30px;
		position: absolute;
		top: 6px;
		left: 32px;
	}

	@keyframes loader3Animation {
		0% {
			transform: rotate(0deg);
		}

		100% {
			transform: rotate(360deg);
		}
	}

	@keyframes loader3AnimationBefore {
		0% {
			transform: rotate(0deg);
		}

		100% {
			transform: rotate(360deg);
		}
	}

	@keyframes loader3AnimationAfter {
		0% {
			transform: rotate(360deg);
		}

		100% {
			transform: rotate(-360deg);
		}
	}
</style>
