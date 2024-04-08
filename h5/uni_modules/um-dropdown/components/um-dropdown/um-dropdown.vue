<template>
	<view class="option" :style="{ width: width, height: height }" ref="dropdown">
		<view class="option-select-title" @click="fnShowOptionList()">
			<!-- <input class="inp-select" :placeholder="placeholder" :value="value" disabled /> -->
			<text v-if="value == undefined || value == null || value == ''" class="inp-select">{{ placeholder }}</text>
			<text v-else class="inp-select">{{ value }}</text>
			<!-- 箭头图标 -->
			<view class="trans" :class="showOptionList ? 'trans-from' : ''">
				<um-icon name="down"></um-icon>
			</view>
		</view>
		<!-- 下拉列表 -->
		<view class="option-list" :style="[{ height: ListHeightVal }, addStyle(listStyle)]">
			<view class="option-list-padding">
				<block v-for="(item, index) in optionList" :key="index">
					<view class="option-item" :style="[addStyle(selectItemIdx == index ? selectedItemStyle : itemStyle)]" @click.stop="fnChangeOption(item, index)">
						{{ rangeKey ? item[rangeKey] : item }}
					</view>
				</block>
			</view>
		</view>
	</view>
</template>

<script>
/**
 * @param optionList 		{Array} 下拉列表数据
 * @example <um-dropdown  :optionList="optionList"></um-dropdown>/>
 */
export default {
	props: {
		// 菜单选择中时的样式
		itemStyle: {
			type: [String, Object],
			default: () => ({
				color: '#333',
				fontSize: '28rpx'
			})
		},
		// 菜单非选中时的样式
		selectedItemStyle: {
			type: [String, Object],
			default: () => ({
				color: '#2973F8',
				fontSize: '28rpx',
				background: '#F5F7FA'
			})
		},
		// 自定义列表样式
		listStyle: {
			type: [String, Object],
			default: ''
		},
		// 列表数据
		optionList: {
			type: Array,
			default: function () {
				return [];
			}
		},
		// 选择框的宽度
		width: {
			type: String,
			default: '100%'
		},
		// 选择框的高度
		height: {
			type: String,
			default: '100%'
		},
		// 如果数组包含对象时，需要显示的key值
		rangeKey: {
			type: String,
			default: ''
		},
		// 默认选中的下标
		defaultIndex: {
			type: [String, Number],
			default: ''
		},
		// 列表高度，若不设置会展示所有列表选项
		listHeight: {
			type: String,
			default: ''
		},
		placeholder: {
			type: String,
			default: '请选择'
		},
		
		selectItemData: {
			type: Object,
			default: {}
		},
	},
	data() {
		return {
			selectItem: '', // 选中的值
			selectItemIdx: null, // 选中的下标
			showOptionList: false // 显示下拉
		};
	},
	computed: {
		value() {
			this.selectItem=this.selectItemData.selectItem;
			this.selectItemIdx=this.selectItemData.selectItemIdx;
			if (!this.selectItem && this.defaultIndex && this.optionList.length) {
				this.selectItemIdx = this.defaultIndex;
				let _val = this.rangeKey ? this.optionList[this.defaultIndex][this.rangeKey] : this.optionList[this.defaultIndex];
				this.$emit('change', this.optionList[this.defaultIndex]);
				return _val;
			} else if (this.selectItem) {
				return this.rangeKey ? this.selectItem[this.rangeKey] : this.selectItem;
			} else {
				return '';
			}
		},
		ListHeightVal() {
			if (this.showOptionList) {
				if (this.listHeight) {
					return this.listHeight;
				} else {
					// 70是单行高度，24是列表上下内边距
					return this.optionList.length * 70 + 24 + 'rpx';
				}
			} else {
				return '0';
			}
		}
	},
	// #ifdef H5
	onLoad() {
		document.addEventListener('click', this.handleDocumentClick);
	},
	beforeDestroy() {
		document.removeEventListener('click', this.handleDocumentClick);
	},

	// #endif
	methods: {
		// 点击组件外部区域时自动收起菜单
		handleDocumentClick(e) {
			const dropdown = this.$refs.dropdown.$el;
			if (dropdown && !dropdown.contains(e.target)) {
				this.showOptionList = false;
			}
		},
		// 控制列表显示与隐藏
		fnShowOptionList() {
			this.$emit('click');
			if (this.optionList.length) {
				this.showOptionList = !this.showOptionList;
			} else {
				// 如果列表为空，发送一个事件
				this.$emit('openNull');
				console.log('mu-dropdown列表数据唯空无法展开');
			}
		},
		// 点击列表时触发
		fnChangeOption(item, index) {
			this.selectItem = item;
			this.selectItemIdx = index;
			this.showOptionList = false;
			this.$emit('change', item);
		},
		/**
		 * @description 样式转换
		 * 对象转字符串，或者字符串转对象
		 * @param {object | string} customStyle 需要转换的目标
		 * @param {String} target 转换的目的，object-转为对象，string-转为字符串
		 * @returns {object|string}
		 */
		addStyle(customStyle, target = 'object') {
			// 字符串转字符串，对象转对象情形，直接返回
			if (this.empty(customStyle) || (typeof customStyle === 'object' && target === 'object') || (target === 'string' && typeof customStyle === 'string')) {
				return customStyle;
			}
			// 字符串转对象
			if (target === 'object') {
				// 去除字符串样式中的两端空格(中间的空格不能去掉，比如padding: 20px 0如果去掉了就错了)，空格是无用的
				customStyle = this.trim(customStyle);
				// 根据";"将字符串转为数组形式
				const styleArray = customStyle.split(';');
				const style = {};
				// 历遍数组，拼接成对象
				for (let i = 0; i < styleArray.length; i++) {
					// 'font-size:20px;color:red;'，如此最后字符串有";"的话，会导致styleArray最后一个元素为空字符串，这里需要过滤
					if (styleArray[i]) {
						const item = styleArray[i].split(':');
						style[this.trim(item[0])] = this.trim(item[1]);
					}
				}
				return style;
			}
			// 这里为对象转字符串形式
			let string = '';
			for (const i in customStyle) {
				// 驼峰转为中划线的形式，否则css内联样式，无法识别驼峰样式属性名
				const key = i.replace(/([A-Z])/g, '-$1').toLowerCase();
				string += `${key}:${customStyle[i]};`;
			}
			// 去除两端空格
			return this.trim(string);
		},

		/**
		 * @description 去除空格
		 * @param String str 需要去除空格的字符串
		 * @param String pos both(左右)|left|right|all 默认both
		 */
		trim(str, pos = 'both') {
			str = String(str);
			if (pos == 'both') {
				return str.replace(/^\s+|\s+$/g, '');
			}
			if (pos == 'left') {
				return str.replace(/^\s*/, '');
			}
			if (pos == 'right') {
				return str.replace(/(\s*$)/g, '');
			}
			if (pos == 'all') {
				return str.replace(/\s+/g, '');
			}
			return str;
		},
		/**
		 * 判断是否为空
		 */
		empty(value) {
			switch (typeof value) {
				case 'undefined':
					return true;
				case 'string':
					if (value.replace(/(^[ \t\n\r]*)|([ \t\n\r]*$)/g, '').length == 0) return true;
					break;
				case 'boolean':
					if (!value) return true;
					break;
				case 'number':
					if (value === 0 || isNaN(value)) return true;
					break;
				case 'object':
					if (value === null || value.length === 0) return true;
					for (const i in value) {
						return false;
					}
					return true;
			}
			return false;
		}
	}
};
</script>

<style lang="scss">
// 去掉列表滚动条
::-webkit-scrollbar {
	display: none;
}

.mask {
	width: 100vw;
	height: 100vh;
	position: fixed;
	position: relative;
	background: #a3a3a3;
	opacity: 0.5;
	z-index: 9;
}

.option {
	width: 100%;
	height: 100%;
	position: relative;

	.option-select-title {
		height: 100%;
		padding: 0 20rpx;
		display: flex;
		justify-content: space-between;
		align-items: center;

		.inp-select {
			color: #828282;
			width: 100%;
		}

		.trans {
			transition: transform 0.2s;
		}

		.trans-from {
			transform: rotate(-180deg);
		}
	}

	.option-list {
		box-sizing: content-box;
		width: calc(100% - 20rpx);
		// height: 0;
		border-radius: 25rpx;
		background: #fff;
		overflow: scroll;
		position: absolute;
		top: calc(100% + 20rpx);
		left: 10rpx;
		right: 10rpx;
		z-index: 10;
		box-shadow: 0 0 15rpx #a3a3a3;
		transition: height 0.2s;

		.option-list-padding {
			padding: 12rpx 0;
		}

		.option-item {
			height: 70rpx;
			text-align: center;
			line-height: 70rpx;
		}
	}
}
</style>