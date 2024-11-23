<template>
    <div style="height: 100%;">
        <a-row class="row" :gutter="16">
            <a-col :span="12">
                <div class="welcome-box">
                    <div class="welcome-text">
                        Hi!
                        <br>
                        欢迎访问
                    </div>
                    <div class="group-wrapper">
                        <div class="tags-group-icon-pair">
                            <div class="tags-group-icon">
                                <Vue3Lottie width="200px" height="200px" renderer="canvas" :animation-data="lo_json1" />
                            </div>
                            <div class="tags-group-icon">
                                <Vue3Lottie width="200px" height="200px" renderer="canvas" :animation-data="lo_json2" />
                            </div>
                        </div>
                        <div class="tags-group-icon-pair">
                            <div class="tags-group-icon">
                                <Vue3Lottie width="200px" height="200px" renderer="canvas" :animation-data="lo_json3" />
                            </div>
                            <div class="tags-group-icon">
                                <Vue3Lottie width="200px" height="200px" renderer="canvas" :animation-data="lo_json4" />
                            </div>
                        </div>
                        <div class="tags-group-icon-pair">
                            <div class="tags-group-icon">
                                <Vue3Lottie width="200px" height="200px" renderer="canvas" :animation-data="lo_json5" />
                            </div>
                            <div class="tags-group-icon">
                                <Vue3Lottie width="200px" height="200px" renderer="canvas" :animation-data="lo_json6" />
                            </div>
                        </div>
                        <div class="tags-group-icon-pair">
                            <div class="tags-group-icon">
                                <Vue3Lottie width="200px" height="200px" renderer="canvas" :animation-data="lo_json1" />
                            </div>
                            <div class="tags-group-icon">
                                <Vue3Lottie width="200px" height="200px" renderer="canvas" :animation-data="lo_json2" />
                            </div>
                        </div>
                        <div class="tags-group-icon-pair">
                            <div class="tags-group-icon">
                                <Vue3Lottie width="200px" height="200px" renderer="canvas" :animation-data="lo_json3" />
                            </div>
                            <div class="tags-group-icon">
                                <Vue3Lottie width="200px" height="200px" renderer="canvas" :animation-data="lo_json4" />
                            </div>
                        </div>
                    </div>
                </div>
            </a-col>
            <a-col :span="12">
                <div class="welcome-box2">
                    <div class="head-box">
                        <img src="@/assets/img/7b58525685c7fa8f58fac97575111bf.png" alt="">
                    </div>
                    <div class="welcome-text">
                        SkyWang
                    </div>
                    <div class="welcome-text" style="font-size: 24px;">
                        <div>
                            <div style="display: flex; justify-content: center">
                                关键词
                            </div>
                            <div style="display: flex; justify-content: center;height: 40px;">
                                {{ statistics.keyWordsCount }}
                            </div>
                        </div>
                        <div>
                            <div style="display: flex; justify-content: center">
                                文章数
                            </div>
                            <div style="display: flex; justify-content: center; height: 40px;">
                                {{ statistics.count }}
                            </div>
                        </div>
                    </div>
                    <div style="font-size: 20px; position: relative; bottom: 30px;">
                        <div style="display: flex; justify-content: center">
                            运行时间
                        </div>
                        <div style="display: flex; justify-content: center">
                            {{ daysToTargetDate }}
                        </div>
                    </div>
                </div>
            </a-col>
            <div :span="24" class="tip-box">
                <div class="tip-item" :style="{ backgroundColor: item.backgroundColor, color: item.fontColor }"
                    v-for="(item, index) in keyWordsList" :key="index" @click="selectItem(item)"
                    :class="{ 'selected': selectedItem === item }">
                    <div class="title">{{ item.name }}</div>
                </div>
            </div>
            <a-col :span="12">
                <div class="card">
                    <div class="articlesList-list" v-for="(item, index) in articlesList" :key="index"
                        @click="goToDetail(item)">
                        <div class="image-box">
                            <img :src="item.image" alt="">
                        </div>
                        <div class="title">{{ item.title }}</div>
                        <div class="introduction">{{ item.desc }}</div>
                        <div class="introduction">{{ item.publishTime }}</div>
                    </div>
                </div>
            </a-col>
            <a-col :span="12">
                <div class="card">
                    <div class="articlesList-list" v-for="(item, index) in articlesList2" :key="index"
                        @click="goToDetail(item)">
                        <div class="image-box">
                            <img :src="item.image" alt="">
                        </div>
                        <div class="title">{{ item.title }}</div>
                        <div class="introduction">{{ item.desc }}</div>
                        <div class="introduction">{{ item.publishTime }}</div>
                    </div>
                </div>
            </a-col>
        </a-row>
    </div>
</template>

<script>
import lo_json1 from '@/assets/lottie/1.json';
import lo_json2 from '@/assets/lottie/2.json';
import lo_json3 from '@/assets/lottie/3.json';
import lo_json4 from '@/assets/lottie/4.json';
import lo_json5 from '@/assets/lottie/5.json';
import lo_json6 from '@/assets/lottie/6.json';
import lo_jsonDog from '@/assets/lottie/dog.json';
import { reactive, toRefs, onUnmounted, onMounted, watch, ref, computed } from 'vue'
import { newsList, keyWordsLists } from '@/api/business.js';
import { message } from 'ant-design-vue';
import { useRouter } from 'vue-router';
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
        const allData = reactive({
            articlesList: [],
            articlesList2: [],
            keyWordsList: [],
            statistics: {},
            selectedItem: [],
        });
        const router = useRouter();
        const goToDetail = (item) => {
            // message.info('功能开发中...')
            router.push({
                path: '/article',
                query: {
                    id: item.id
                }
            });
        }
        const getNewsList = (keyWords) => {
            var parameter = {
                pageNo: 1,
                pageSize: 1000,
                title: '',
                keyWords: keyWords,
                cid: '',
                status: -1,
                recommended: 1
            };
            newsList(parameter).then((res) => {
                if (res.code == 200) {
                    const lists = res.data.lists;
                    const length = lists.length;
                    const midIndex = Math.floor(length / 2);
                    allData.articlesList2 = lists.slice(0, midIndex);
                    allData.articlesList = lists.slice(midIndex);
                    allData.statistics.count = res.data.count;
                } else {
                    console.log(res)
                }
            }).catch((err) => {
                console.log(err)
            });
        };
        const getKeyWordsLists = () => {
            var parameter = {
                pageNo: 1,
                pageSize: 1000,
            };
            keyWordsLists(parameter).then((res) => {
                if (res.code == 200) {
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
                    const lists = res.data.lists.map(item => ({
                        ...item,
                        backgroundColor: generateRandomColor().backgroundColor,
                        fontColor: generateRandomColor().fontColor
                    }));
                    allData.keyWordsList = lists;
                    allData.statistics.keyWordsCount = res.data.count;

                    // 添加缓存存储
                    localStorage.setItem('keyWordsList', JSON.stringify(lists));
                } else {
                    console.log(res)
                }
            }).catch((err) => {
                console.log(err)
            });
        };

        const selectItem = (item) => {
            if (allData.selectedItem === item) {
                // 如果当前项已经被选中，则取消选中
                allData.selectedItem = null;
            } else {
                // 否则，选中该项
                allData.selectedItem = item;
            }
            getNewsList(allData.selectedItem ? allData.selectedItem.id : '');

        };
        const targetDate = new Date(2024, 5, 12); // 注意：月份是从0开始计数的，所以6月是5
        const currentDate = ref(new Date());

        const daysToTargetDate = computed(() => {
            const diffInMs = targetDate - currentDate.value;
            const totalSeconds = Math.abs(diffInMs / 1000);
            const days = Math.floor(totalSeconds / (24 * 60 * 60));
            const hours = Math.floor((totalSeconds % (24 * 60 * 60)) / (60 * 60));
            const minutes = Math.floor((totalSeconds % (60 * 60)) / 60);
            const seconds = Math.floor(totalSeconds % 60);

            return `${days} 天 ${hours} 时 ${minutes} 分 ${seconds} 秒`;
        });

        let intervalId;

        onUnmounted(() => {
            clearInterval(intervalId);
        });
        watch(
            show,
            (val) => {
                if (val) {
                    // getNewsList();
                }
            },
            { deep: true, immediate: true },
        );
        onMounted(() => {
            getNewsList();
            getKeyWordsLists();
            intervalId = setInterval(() => {
                currentDate.value = new Date();
            }, 1000);
        })
        return {
            ...toRefs(allData),
            goToDetail,
            selectItem,
            daysToTargetDate,
            lo_json1,
            lo_json2,
            lo_json3,
            lo_json4,
            lo_json5,
            lo_json6,
            lo_jsonDog,
        }
    },
}

</script>

<style scoped lang='scss'>
.card {
    border-radius: 15px;
}

.welcome-box {
    border-radius: 15px;
    background-color: #fff;
    margin-bottom: 20px;
    display: flex;
    overflow: hidden;
    height: 40vh;

    .welcome-text {
        margin: 30px;
        line-height: 40px;
        font-size: 35px;
        font-weight: 600;
        color: #363636;
    }

    .head-box {
        height: 150px;
        width: 150px;
        overflow: hidden;
        margin: 30px auto 0 auto;
        border-radius: 50%;

        img {
            height: 150px;
            width: 150px;
        }
    }
}

.welcome-box2 {
    border-radius: 15px;
    background-color: #fff;
    margin-bottom: 20px;
    overflow: hidden;
    height: 40vh;

    .head-box {
        height: 150px;
        width: 150px;
        overflow: hidden;
        margin: 30px auto 0 auto;
        border-radius: 50%;

        img {
            height: 150px;
            width: 150px;
        }
    }

    .welcome-text {

        margin-top: 10px;
        display: flex;
        justify-content: space-evenly;
        line-height: 40px;
        font-size: 35px;
        font-weight: 600;
        color: #363636;
    }
}

.group-wrapper {
    margin-top: 0;
    display: flex;
    flex-wrap: wrap;
    animation: rowup 30s linear infinite;
    width: 320px;
    margin-left: auto;

    img {
        width: 150px;
    }

    .tags-group-icon-pair {
        display: flex;
        margin-left: 1rem;
    }

    .tags-group-icon {
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
        box-shadow: 5px 5px 13px 0px rgba(117, 126, 136, 0.2);
        width: 120px;
        height: 120px;
        border-radius: 30px;
    }

    .tags-group-icon:nth-child(even) {
        margin-left: 2.5rem;
    }

    .tags-group-icon:nth-child(even) {
        margin-top: 2rem;
    }
}

@keyframes rowup {
    0% {
        transform: translateY(0);
    }

    100% {
        transform: translateY(-120%);
    }
}

.articlesList-list:hover {
    cursor: pointer !important;
    border-color: transparent;
    box-shadow: 0 1px 2px -2px rgba(0, 0, 0, 0.16), 0 3px 6px 0 rgba(0, 0, 0, 0.12), 0 5px 12px 4px rgba(0, 0, 0, 0.09);
}

.tip-box {
    display: flex;
    background-color: #fff;
    border-radius: 5px;
    height: 5vh;
    margin-bottom: 10px;
    margin-left: 9px;
    flex: 0 0 98.833333%;
    max-width: 98.833333%;
    align-items: center;

    .tip-item {
        cursor: pointer !important;
        border-radius: 5px;
        margin: 5px 10px;
        padding: 0 15px;
        border: 2px solid transparent;
    }

    .selected {
        border: 2px solid #7cbbf6;
        /* 选中的边框样式 */
    }

    .line {
        position: absolute;
        bottom: 8px;
        left: 0;
        height: 2px;
        width: 100%;
        background-color: rgb(127, 127, 127);
    }

    /* 定义动画 */
    /* 修改后的动画 */
    @keyframes moveAndFlip {
        0% {
            left: 1px;
            /* 从左侧开始 */
            transform: scaleX(1);
            /* 保持正常方向 */
        }

        50% {
            left: calc(100% - 60px);
            /* 移动到右侧 */
            transform: scaleX(1);
            /* 不翻转，保持正常方向 */
        }

        51% {
            transform: scaleX(-1);
            /* 到达右侧后立刻翻转 */
        }

        100% {
            left: 1px;
            /* 回到左侧 */
            transform: scaleX(-1);
            /* 保持翻转状态返回 */
        }
    }
}

.articlesList-list {
    background-color: #fff;
    border-radius: 10px;
    box-shadow: 2px 2px 10px 0px rgba(117, 126, 136, 0.2);
    height: 40vh;
    margin-bottom: 4vh;
    box-sizing: border-box;

    .image-box {
        border-radius: 10px;
        overflow: hidden;
        height: 60%;
        width: 100%;

        img {
            width: 100%;
        }
    }

    .title {
        margin: 20px;
        font-size: 26px;
    }

    .introduction {
        font-size: 18px;
        margin: 10px 20px;
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
        margin-left: 20px;
        display: flex;

        img {
            width: 45px;
            height: 40px;
        }

        text {
            margin-right: 10px;
            line-height: 38px;
            color: #babdba;
        }
    }
}
</style>