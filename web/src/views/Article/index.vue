<template>
	<div class="article-container">
		<canvas ref="backgroundCanvas" class="background-canvas"></canvas>
	</div>
</template>

<script>
import { reactive, toRefs, onBeforeMount, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router';
import { setupCanvasBackground, mouseClickToBC } from '@/views/Home/bcCanvas';
export default {
	name: 'wContent',
	props: {
		show: {
			type: Boolean,
			default: false
		}
	},
	setup(props) {
		const { show } = toRefs(props);
		const backgroundCanvas = ref(null);// 定义引用，用于存储背景画布元素
		const allData = reactive({
			cateList: [],
			albumList: [],
			albumList2: [],
			colorSchemes: {},
			statistics: {}
		});

		const getCateList = () => {
			var parameter = {
			};
			cateList(parameter).then((res) => {
				if (res.code == 200) {
					// 生成随机颜色的函数
					const generateRandomColor = () => {
						// 生成清新的颜色  
						const randomHue = Math.floor(Math.random() * 360); // 随机色调  
						const saturation = '60%'; // 饱和度  
						const lightness = '90%'; // 初始亮度  
						// 使用 HSL 来生成背景颜色  
						const backgroundColor = `hsl(${randomHue}, ${saturation}, ${lightness})`;
						// 背景颜色的亮度降低 20% 来生成字体颜色  
						const darkerLightness = Math.max(0, parseInt(lightness) - 50) + '%';
						const fontColor = `hsl(${randomHue}, ${saturation}, ${darkerLightness})`;
						return {
							backgroundColor,
							fontColor
						};
					};
					const chunkArray = (array, size) => {
						const result = [];
						for (let i = 0; i < array.length; i += size) {
							result.push(array.slice(i, i + size));
						}
						return result;
					}
					const lists = res.data.map(item => ({
						...item,
						backgroundColor: generateRandomColor().backgroundColor,
						fontColor: generateRandomColor().fontColor
					}));
					allData.statistics.cateCount = res.data.length;
					allData.cateList = chunkArray(lists, 2);
				} else {
					console.log(res)
				}
			}).catch((err) => {
				console.log(err)
			});
		};
		onMounted(() => {
			// setupCanvasBackground(backgroundCanvas, '', mouseClickToBC);
		})
		return {
			...toRefs(allData),
			backgroundCanvas
		}
	},
}

</script>

<style scoped lang='scss'>
.article-container {
	display: flex;
  overflow-y: auto;
  justify-content: center;
  height: 100vh;
}
</style>