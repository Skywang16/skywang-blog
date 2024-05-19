<template>
	<view class="container">
		<view class="main">
			<view class="card-body">
				<uni-forms ref="formData" :modelValue="formData">
					<uni-forms-item label="账号" name="username" required>
						<uni-easyinput :maxlength="12" v-model="formData.username" placeholder="请输入账号" />
					</uni-forms-item>
					<uni-forms-item label="头像" required>
						<uni-file-picker id="picker" fileMediatype="image" file-extname="png,jpg" limit="1" @select="select" @delete="deleteImg"></uni-file-picker>
					</uni-forms-item>
					<uni-forms-item label="昵称" name="nickname" required>
						<uni-easyinput :maxlength="9" v-model="formData.nickname" placeholder="请输入昵称" />
					</uni-forms-item>
					<uni-forms-item label="密码" name="password" required>
						<uni-easyinput :maxlength="9" type="password" v-model="formData.password" placeholder="请输入密码" />
					</uni-forms-item>
					<uni-forms-item label="确认密码" name="passwordConfirm" required>
						<uni-easyinput :maxlength="9" type="password" v-model="formData.passwordConfirm" placeholder="请再次输入密码" />
					</uni-forms-item>
				</uni-forms>
			</view>
		</view>
		<view class="footer">
			<button class="btn" type="primary" @click="submit">提交</button>
		</view>
		<uni-popup ref="message" type="message">
			<uni-popup-message type="error" :message="messageText" :duration="2500"></uni-popup-message>
		</uni-popup>
		<loading :loading="showLoading" :firstContent="firstContent" :secondContent="secondContent" />
	</view>
</template>

<script>
import { register, token, getUserInfo, delImg } from '@/common/api/login.js';
export default {
	data() {
		return {
			formData: {
				file: undefined,
				nickname: '',
				password: '',
				passwordConfirm: '',
				username: ''
			},
			rules: {
				username: {
					rules: [
						{
							required: true,
							errorMessage: '请输入账号'
						},
						{
							validateFunction: function (rule, value, data, callback) {
								let va = /^\d+$/; //手机号码
								if (!va.test(value)) {
									callback('账号只能是数字');
								}
								return true;
							}
						}
					]
				},
				nickname: {
					rules: [
						{
							required: true,
							errorMessage: '请输入昵称'
						}
					]
				},
				password: {
					rules: [
						{
							required: true,
							errorMessage: '请输入密码'
						},
						{
							validateFunction: function (rule, value, data, callback) {
								let va = /^(?=.*\d)(?=.*[A-z])[\da-zA-Z]{1,9}$/; //手机号码
								if (!va.test(value)) {
									callback('密码只能由字母、数字组成，1-9位');
								}
								return true;
							}
						}
					]
				},
				passwordConfirm: {
					rules: [
						{
							required: true,
							errorMessage: '请再次输入密码'
						},
						{
							validateFunction: (rule, value, data, callback) => {
								// 异步需要返回 Promise 对象
								return new Promise((resolve, reject) => {
									setTimeout(() => {
										if (value === this.formData.password) {
											// 通过返回 resolve
											resolve();
										} else {
											// 不通过返回 reject(new Error('错误信息'))
											reject(new Error('两次密码不一致'));
										}
									}, 100);
								});
							}
						}
					]
				}
			},
			showLoading: false, // 自定义遮罩显影
			firstContent: '', // 自定义遮罩第一行文本
			secondContent: '', // 自定义遮罩第二行文本
			messageText: ''
		};
	},
	onReady() {
		// 如果需要兼容微信小程序，并且校验规则中含有方法等，只能通过setRules方法设置规则
		this.$refs.formData.setRules(this.rules);
	},
	onLoad(options) {},
	methods: {
		select(e) {
			const file = e.tempFiles[0].file;
			const fileType = file.type.toLowerCase();
			const fileSize = file.size / 1024 / 1024; // 转换为MB单位
			// 校验文件大小是否超过10MB
			if (fileSize > 1) {
				uni.showToast({
					title: '文件大小不得大于1MB',
					icon: 'none'
				});
				return;
			}
			this.formData.file = e;
		},
		//删除图片
		deleteImg() {
			this.formData.file = undefined;
		},
		// 表单校验
		submit() {
			this.$refs.formData
				.validate()
				.then((res) => {
					if (!this.formData.file) {
						uni.showToast({
							title: '请选择上传头像',
							icon: 'none'
						});
						return;
					} else {
						this.register();
					}
				})
				.catch((error) => {
					console.log('error', error);
				});
		},
		// 上传图片
		register() {
			this.setLoadingOptions(true, '注册中');
			//let BASE_URL = uni.getStorageSync('BASE_URL') ';
			let BASE_URL = 'http://localhost:8001/api/';
			let path = 'common/upload/image';
			// 创建随机数
			uni.uploadFile({
				url: BASE_URL + path,
				fileType: 'image',
				name: 'file',
				file: this.formData.file.tempFiles[0].file,
				formData: {
					cid: 1
				},
				success: (res) => {
					let result = JSON.parse(res.data);
					console.log(result);
					if (result.code === 200) {
						let avatar = result.data;
						this.addUser(avatar);
					} else {
						uni.showToast({
							title: '添加失败',
							icon: 'none'
						});
					}
				},
				fail: (err) => {
					uni.showToast({
						title: '添加失败',
						icon: 'none'
					});
				}
			});
		},
		delImg(avatar){
			let ids = []
			ids.push(avatar.id)
			let params = {
				ids: ids,
			};
			delImg(params).then((res) => {
			})
		},
		// 注册
		addUser(avatar) {
			let params = {
				avatar: avatar.path,
				deptId: 10,
				id: '',
				isDisable: 0,
				isMultipoint: 0,
				nickname: this.formData.nickname,
				password: this.formData.password,
				passwordConfirm: this.formData.passwordConfirm,
				role: 2,
				sort: 1,
				username: this.formData.username
			};
			register(params).then((res) => {
				this.setLoadingOptions(false);
				let result = res.data;
				if (result.code === 200) {
					this.token();
				} else {
					this.messageText = result.msg;
					this.$refs.message.open();
					this.delImg(avatar)
				}
			});
		},
		// 登录
		token() {
			this.setLoadingOptions(true, '注册成功，登陆中');
			let params = {
				terminal: 1,
				password: this.formData.password,
				username: this.formData.username
			};
			token(params).then((res) => {
				let result = res.data;
				if (result.code === 200) {
					uni.setStorageSync('token', result.data.token);
					this.getUserInfo();
				} else {
					this.messageText = '登录失败';
					this.$refs.message.open();
				}
			});
		},
		// 获取用户信息
		getUserInfo() {
			getUserInfo().then((res) => {
				this.setLoadingOptions(false);
				let result = res.data;
				if (result.code === 200) {
					uni.setStorageSync('userData', result.data.user);
					uni.reLaunch({
						url: '/pages/home/index'
					});
				} else {
					this.messageText = '获取用户信息失败';
					this.$refs.message.open();
				}
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

<style scoped lang="scss">
.container {
	background-color: #f5f6f7;
	height: 100%;
}

.header-bg {
	border-bottom: 22rpx solid #efefef;
	background: #fff;

	.title {
		margin: auto;
		height: 88rpx;
		padding: 10px 0px 0px 0px;
		font-weight: 600;
		font-size: 40rpx;
		font-family: PingFangHK-Semibold, PingFangHK;
		font-weight: 600;
		color: #333333;
		line-height: 56rpx;
	}
}

.card-body {
	width: 100%;
	background: #ffffff;
	box-shadow: 0rpx 18rpx 29rpx 0rpx rgba(78, 104, 156, 0.09);
}

/deep/.uni-forms-item__label {
	margin-left: 40rpx;
	margin-right: 10rpx;
	width: 200rpx;
	justify-content: flex-end !important;
}
/deep/.uni-easyinput {
	width: 90%;
}
/deep/.file-picker__box {
	width: 50% !important;
	padding-top: 50% !important;
}
/deep/.uni-progress-bar {
	height: 0px !important;
}
.footer {
	display: flex;
	position: fixed;
	bottom: 0px;
	background-color: #fff;
	height: 160rpx;
	width: 100%;

	.btn {
		margin: auto;
		height: 80rpx;
		line-height: 80rpx;
		border-radius: 8rpx;
		width: 640rpx;
	}
}
</style>