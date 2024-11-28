<template>
    <div style="height: 100%;">
        <div class="container">
            <div class="row">
                <div class="col col-xs-24 col-sm-24 col-md-12 col-lg-12">
                    <div class="welcome-box">
                        <div class="welcome-text">
                            Hi!
                            <br>
                            欢迎访问
                        </div>
                        <div class="group-wrapper">
                            <div class="tags-group-icon-pair">
                                <div class="tags-group-icon">
                                    <Vue3Lottie width="200px" height="200px" renderer="canvas"
                                        :animation-data="lo_json1" />
                                </div>
                                <div class="tags-group-icon">
                                    <Vue3Lottie width="200px" height="200px" renderer="canvas"
                                        :animation-data="lo_json2" />
                                </div>
                            </div>
                            <div class="tags-group-icon-pair">
                                <div class="tags-group-icon">
                                    <Vue3Lottie width="200px" height="200px" renderer="canvas"
                                        :animation-data="lo_json3" />
                                </div>
                                <div class="tags-group-icon">
                                    <Vue3Lottie width="200px" height="200px" renderer="canvas"
                                        :animation-data="lo_json4" />
                                </div>
                            </div>
                            <div class="tags-group-icon-pair">
                                <div class="tags-group-icon">
                                    <Vue3Lottie width="200px" height="200px" renderer="canvas"
                                        :animation-data="lo_json5" />
                                </div>
                                <div class="tags-group-icon">
                                    <Vue3Lottie width="200px" height="200px" renderer="canvas"
                                        :animation-data="lo_json6" />
                                </div>
                            </div>
                            <div class="tags-group-icon-pair">
                                <div class="tags-group-icon">
                                    <Vue3Lottie width="200px" height="200px" renderer="canvas"
                                        :animation-data="lo_json1" />
                                </div>
                                <div class="tags-group-icon">
                                    <Vue3Lottie width="200px" height="200px" renderer="canvas"
                                        :animation-data="lo_json2" />
                                </div>
                            </div>
                            <div class="tags-group-icon-pair">
                                <div class="tags-group-icon">
                                    <Vue3Lottie width="200px" height="200px" renderer="canvas"
                                        :animation-data="lo_json3" />
                                </div>
                                <div class="tags-group-icon">
                                    <Vue3Lottie width="200px" height="200px" renderer="canvas"
                                        :animation-data="lo_json4" />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col col-xs-24 col-sm-24 col-md-12 col-lg-12">
                    <div class="welcome-box2">
                        <div class="profile-section">
                            <div class="head-box">
                                <img src="@/assets/img/7b58525685c7fa8f58fac97575111bf.png" alt="profile">
                                <div class="head-box-overlay"></div>
                            </div>
                            <div class="name-text">SkyWang</div>
                        </div>
                        <div class="stats-section">
                            <div class="stat-item">
                                <div class="stat-content">
                                    <div class="stat-value">
                                        <CountUp :end-value="statistics.keyWordsCount" :duration="2500" />
                                    </div>
                                    <div class="stat-label">关键词</div>
                                </div>
                            </div>
                            <div class="stat-divider"></div>
                            <div class="stat-item">
                                <div class="stat-content">
                                    <div class="stat-value">
                                        <CountUp :end-value="statistics.count" />
                                    </div>
                                    <div class="stat-label">文章数</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="tip-box">
                    <div class="tip-item" :style="{ backgroundColor: item.backgroundColor, color: item.fontColor }"
                        v-for="(item, index) in keyWordsList" :key="index" @click="selectItem(item)"
                        :class="{ 'selected': selectedItem === item }">
                        <div class="title">{{ item.name }}</div>
                    </div>
                </div>
                <div class="col col-24">
                    <div class="card articles-container">
                        <div class="articlesList-list" v-for="(item, index) in articlesList" :key="index"
                            @click="goToDetail(item)">
                            <div class="image-box">
                                <img :src="item.image" alt="">
                            </div>
                            <div class="title">{{ item.title }}</div>
                            <!-- <div class="introduction">{{ item.desc }}</div> -->
                            <div class="meta-info">
                                <span class="publish-time">{{ item.publishTime }}</span>
                                <div class="tags-container">
                                    <span class="category-tag" v-for="(keyWordId, index) in item.keyWords.split(',')"
                                        :key="index">
                                        {{ getKeyWordName(keyWordId) }}
                                    </span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="footer-box">
                    <div class="footer-content">
                        <div class="footer-divider"></div>
                        <div class="footer-text">
                            <span>-- 已经到底啦 --</span>
                        </div>
                        <div class="runtime-info">
                            <span class="runtime-icon">⏳</span>
                            <span class="runtime-text">运行时间: {{ daysToTargetDate }}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
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
import { useRouter } from 'vue-router';
import CountUp from '@/components/CountUp.vue'
export default {
    name: 'wContent',
    props: {
        show: {
            type: Boolean,
            default: false
        }
    },
    components: {
        CountUp
    },
    setup(props) {
        const { show } = toRefs(props);
        const allData = reactive({
            articlesList: [],
            keyWordsList: [],
            statistics: {
                count: 0,
                keyWordsCount: 0
            },
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
                recommended: -1
            };
            newsList(parameter).then((res) => {
                if (res.code == 200) {
                    allData.articlesList = res.data.lists;
                    allData.statistics.count = Number(res.data.count) || 0;
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
                    allData.statistics.keyWordsCount = Number(res.data.count) || 0;

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

        const getKeyWordName = (keyWordsId) => {
            if (!keyWordsId || !allData.keyWordsList.length) return '';
            const keyWord = allData.keyWordsList.find(k => k.id === Number(keyWordsId));
            return keyWord ? keyWord.name : '';
        };

        return {
            ...toRefs(allData),
            goToDetail,
            selectItem,
            daysToTargetDate,
            getKeyWordName,
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

<style scoped lang="scss">
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
    transition: all 0.3s ease;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);

    .welcome-text {
        margin: 30px;
        line-height: 40px;
        font-size: 35px;
        font-weight: 600;
        color: #363636;
        background: linear-gradient(45deg, #363636, #666);
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        animation: fadeIn 1s ease-in;
    }

    .head-box {
        height: 150px;
        width: 150px;
        overflow: hidden;
        margin: 30px auto 0 auto;
        border-radius: 50%;
        border: 4px solid #fff;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
        transition: transform 0.3s ease;

        &:hover {
            transform: scale(1.05);
        }

        img {
            height: 150px;
            width: 150px;
            transition: transform 0.5s ease;

            &:hover {
                transform: scale(1.1);
            }
        }
    }
}

.welcome-box2 {
    border-radius: 15px;
    background-color: #fff;
    margin-bottom: 20px;
    overflow: hidden;
    height: 40vh;
    transition: all 0.3s ease;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    padding: 20px;
    display: flex;
    flex-direction: column;
    gap: 20px;
    background: linear-gradient(145deg, #ffffff, #f0f0f0);

    .profile-section {
        text-align: center;

        .head-box {
            position: relative;
            height: 120px;
            width: 120px;
            margin: 0 auto;
            border-radius: 50%;
            border: 4px solid #fff;
            box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
            overflow: hidden;
            transition: all 0.3s ease;

            &:hover {
                transform: scale(1.05);
                box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);

                .head-box-overlay {
                    opacity: 1;
                }

                img {
                    transform: scale(1.1);
                }
            }

            img {
                width: 100%;
                height: 100%;
                object-fit: cover;
                transition: transform 0.5s ease;
            }

            .head-box-overlay {
                position: absolute;
                top: 0;
                left: 0;
                right: 0;
                bottom: 0;
                background: linear-gradient(45deg,
                        rgba(74, 158, 255, 0.2),
                        rgba(130, 190, 255, 0.2));
                opacity: 0;
                transition: opacity 0.3s ease;
            }
        }

        .name-text {
            margin-top: 15px;
            font-size: 28px;
            font-weight: 600;
            background: linear-gradient(45deg, #2c3e50, #3498db);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            animation: fadeIn 0.5s ease;
        }
    }

    .stats-section {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 30px;
        padding: 20px 0;

        .stat-item {
            display: flex;
            align-items: center;
            gap: 10px;

            .stat-content {
                text-align: center;

                .stat-value {
                    font-size: 28px;
                    font-weight: 600;
                    color: #2c3e50;
                    margin-bottom: 5px;
                    transition: all 0.3s ease;

                    span {
                        display: inline-block;
                        background: linear-gradient(45deg, #2c3e50, #3498db);
                        -webkit-background-clip: text;
                        -webkit-text-fill-color: transparent;
                        animation: fadeIn 0.5s ease-out;
                    }
                }

                .stat-label {
                    font-size: 24px;
                    color: #666;
                    white-space: nowrap;
                }
            }
        }

        .stat-divider {
            width: 1px;
            height: 40px;
            background: linear-gradient(to bottom,
                    transparent,
                    #ddd,
                    transparent);
        }
    }
}

@media screen and (max-width: 768px) {
    .welcome-box2 {
        .stats-section {
            flex-direction: column;
            gap: 20px;

            .stat-divider {
                width: 80%;
                height: 1px;
            }
        }

        .runtime-section {
            flex-direction: column;
            gap: 10px;
        }
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
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);

    .tip-item {
        cursor: pointer !important;
        border-radius: 5px;
        margin: 5px 10px;
        padding: 0 15px;
        border: 2px solid transparent;
        transition: all 0.3s ease;

        &:hover {
            transform: translateY(-2px);
            box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
        }
    }

    .selected {
        border: 2px solid #4a9eff;
        box-shadow: 0 0 10px rgba(74, 158, 255, 0.3);
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

.articles-container {
    column-count: 2; // 设置两列
    column-gap: 20px; // 设置列间距
    background-color: transparent;

    // 添加媒体查询，在移动端下改为单列
    @media screen and (max-width: 768px) {
        column-count: 1; // 移动端单列显示
    }
}

.articlesList-list {
    break-inside: avoid;
    display: inline-block;
    width: 100%;
    margin-bottom: 4vh;
    background-color: #fff; // 给每个文章卡片添加白色背景
    border-radius: 15px;
    box-shadow: 2px 2px 10px 0px rgba(117, 126, 136, 0.2);
    transition: all 0.3s ease;

    &:hover {
        transform: translateY(-5px);
        box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
    }

    .image-box {
        border-radius: 10px;
        overflow: hidden;
        height: 60%;
        width: 100%;
        position: relative;
        overflow: hidden;

        img {
            width: 100%;
            transition: transform 0.5s ease;

            &:hover {
                transform: scale(1.05);
            }
        }
    }

    .title {
        margin: 20px;
        font-size: 26px;
        color: #2c3e50;
        font-weight: 600;
        transition: color 0.3s ease;

        &:hover {
            color: #4a9eff;
        }
    }

    .introduction {
        font-size: 18px;
        margin: 10px 20px;
        color: #babdba;
        // 超出两行省���
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

    .meta-info {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin: 10px 20px;

        .publish-time {
            font-size: 18px;
            color: #babdba;
        }

        .tags-container {
            display: flex;
            gap: 8px;
            flex-wrap: wrap;
            justify-content: flex-end;
        }

        .category-tag {
            display: inline-block;
            padding: 4px 12px;
            border-radius: 5px;
            font-size: 14px;
            color: #4a9eff;
            background-color: rgba(74, 158, 255, 0.1);
        }
    }
}

.footer-box {
    width: 100%;
    display: flex;
    justify-content: center;
    padding: 20px 0;
    margin-top: 20px;

    .footer-content {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 15px;
    }

    .runtime-info {
        display: flex;
        align-items: center;
        gap: 10px;
        color: #666;
        font-size: 16px;
        padding: 8px 16px;
        background: rgba(255, 255, 255, 0.8);
        border-radius: 20px;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
        transition: all 0.3s ease;


        .runtime-icon {
            font-size: 20px;
            opacity: 0.8;
        }

        .runtime-text {
            font-weight: 500;
        }
    }

    .footer-divider {
        width: 50px;
        height: 2px;
        background: linear-gradient(to right, transparent, #ddd, transparent);
    }

    .footer-text {
        color: #a3a3a3;
        font-size: 14px;

        span {
            position: relative;

            &::before,
            &::after {
                content: '';
                position: absolute;
                top: 50%;
                width: 15px;
                height: 1px;
                background: #ddd;
            }

            &::before {
                right: 100%;
                margin-right: 10px;
            }

            &::after {
                left: 100%;
                margin-left: 10px;
            }
        }
    }
}

/* 移动端适配 */
@media screen and (max-width: 768px) {
    .footer-box {
        padding: 15px 0;

        .runtime-info {
            font-size: 14px;
            padding: 6px 12px;
        }

        .footer-text {
            font-size: 12px;
        }
    }
}

// 添加一些动画
@keyframes fadeIn {
    from {
        opacity: 0;
        transform: translateY(20px);
    }

    to {
        opacity: 1;
        transform: translateY(0);
    }
}

// 添加滚动条美化
::-webkit-scrollbar {
    width: 8px;
    height: 8px;
}

::-webkit-scrollbar-track {
    background: #f1f1f1;
    border-radius: 4px;
}

::-webkit-scrollbar-thumb {
    background: #888;
    border-radius: 4px;

    &:hover {
        background: #555;
    }
}

/* 修改移动端适配部分 */
@media screen and (max-width: 768px) {
    .articles-container {
        column-count: 1;
    }

    .welcome-box {
        display: none;
    }

    .welcome-box2 {
        height: auto;
        margin-bottom: 10px;

        .stats-section {
            flex-direction: column;
            gap: 20px;

            .stat-divider {
                width: 80%;
                height: 1px;
            }
        }
    }

    .tip-box {
        height: auto;
        flex-wrap: wrap;
        padding: 5px;

        .tip-item {
            margin: 3px 5px;
            font-size: 14px;
        }
    }

    .articlesList-list {
        .title {
            font-size: 20px;
            margin: 15px;
        }

        .introduction {
            font-size: 16px;
            margin: 8px 15px;
        }
    }

    .footer-box {
        padding: 15px 0;

        .runtime-info {
            font-size: 14px;
            padding: 6px 12px;
        }

        .footer-text {
            font-size: 12px;
        }
    }
}

/* 添加新的网格系统样式 */
.container {
    width: 100%;
    padding-right: 15px;
    padding-left: 15px;
    margin-right: auto;
    margin-left: auto;
}

.row {
    display: flex;
    flex-wrap: wrap;
    margin-right: -15px;
    margin-left: -15px;
}

.col {
    position: relative;
    width: 100%;
    padding-right: 15px;
    padding-left: 15px;
}

/* 响应式列宽 */
.col-24 {
    flex: 0 0 100%;
    max-width: 100%;
}

@media (min-width: 768px) {
    .col-md-12 {
        flex: 0 0 50%;
        max-width: 50%;
    }
}

@media (min-width: 992px) {
    .col-lg-12 {
        flex: 0 0 50%;
        max-width: 50%;
    }
}

@keyframes countUp {
    0% {
        transform: translateY(20px);
        opacity: 0;
    }

    100% {
        transform: translateY(0);
        opacity: 1;
    }
}

@keyframes fadeIn {
    from {
        opacity: 0;
        transform: translateY(10px);
    }

    to {
        opacity: 1;
        transform: translateY(0);
    }
}
</style>