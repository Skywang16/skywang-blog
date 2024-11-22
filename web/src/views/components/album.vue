<template>
    <div style="height: 100%;">
        <a-spin :spinning="loading">
            <a-row class="row" :gutter="16">
                <a-col :span="12">
                    <div class="welcome-box">
                        <div class="welcome-text">
                            Hi!
                            <br>
                            欢迎访问
                        </div>
                        <div class="group-wrapper">
                            <div v-for="(group, groupIndex) in cateList" :key="groupIndex" class="tags-group-icon-pair">
                                <div v-for="(item, index) in group" :key="index" class="tags-group-icon"
                                    @click="goToDetail(item)">
                                    <div class="group-name"
                                        :style="{ backgroundColor: item.backgroundColor, color: item.fontColor }">
                                        {{ item.name }}
                                    </div>
                                </div>
                            </div>
                            <div v-for="(group, groupIndex) in cateList" :key="groupIndex" class="tags-group-icon-pair">
                                <div v-for="(item, index) in group" :key="index" class="tags-group-icon"
                                    @click="goToDetail(item)">
                                    <div class="group-name"
                                        :style="{ backgroundColor: item.backgroundColor, color: item.fontColor }">
                                        {{ item.name }}
                                    </div>
                                </div>
                            </div>
                            <div v-for="(group, groupIndex) in cateList" :key="groupIndex" class="tags-group-icon-pair">
                                <div v-for="(item, index) in group" :key="index" class="tags-group-icon"
                                    @click="goToDetail(item)">
                                    <div class="group-name"
                                        :style="{ backgroundColor: item.backgroundColor, color: item.fontColor }">
                                        {{ item.name }}
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </a-col>
                <a-col :span="12">
                    <div class="welcome-box">
                        <div style="display: flex; justify-content: space-between; width: 100%;">
                            <div class="welcome-text">
                                统计
                                <br>
                                <br>
                                {{ statistics.cateCount }} 个分类
                                <br>
                                {{ statistics.albumCount }} 张照片
                                <br>
                                <br>
                                <br>
                                <span style="position: absolute;font-size: 35px; font-weight: 600; color: #363636;">
                                    最后操作时间 {{ statistics.mostRecentTime }}
                                </span>
                            </div>
                            <div class="welcome-text" style="writing-mode: vertical-rl;">
                                {{ statistics.sizeCount }}
                            </div>
                        </div>
                    </div>
                </a-col>
                <a-col :span="24" class="announcement-box">
                    <div style="cursor: pointer;" @click="goToDetail()">
                        <Vue3Lottie class="announcement-box-dog" width="60px" height="60px"
                            :animation-data="lo_jsonDog" />
                    </div>
                    <div class="line"></div>
                </a-col>
                <a-col class="articlesList-box" :span="12">
                    <div class="card">
                        <div class="articlesList-list" v-for="(item, index) in albumList" :key="index">
                            <div class="image-box">
                                <a-image style="height: 100%; " :src="item.uri" alt="" />
                            </div>
                            <div class="title">{{ item.name }}</div>
                            <div class="introduction">
                                <div class="tip-box" :style="getCategoryColors(item.cid) && {
                                    backgroundColor: getCategoryColors(item.cid).backgroundColor,
                                    color: getCategoryColors(item.cid).fontColor
                                }">
                                    {{ findCategoryName(item.cid) }}
                                </div>
                                <div>
                                    {{ item.size }}
                                </div>
                                <div>
                                    {{ item.createTime }}
                                </div>
                            </div>
                        </div>
                    </div>
                </a-col>
                <a-col class="articlesList-box" :span="12">
                    <div class="card">
                        <div class="articlesList-list" v-for="(item, index) in albumList2" :key="index">
                            <div class="image-box">
                                <a-image style="height: 100%; " :src="item.uri" alt="" />
                            </div>
                            <div class="title">{{ item.name }}</div>
                            <div class="introduction">
                                <div class="tip-box" :style="getCategoryColors(item.cid) && {
                                    backgroundColor: getCategoryColors(item.cid).backgroundColor,
                                    color: getCategoryColors(item.cid).fontColor
                                }">
                                    {{ findCategoryName(item.cid) }}
                                </div>
                                <div>
                                    {{ item.size }}
                                </div>
                                <div>
                                    {{ item.createTime }}
                                </div>
                            </div>
                        </div>
                    </div>
                </a-col>
            </a-row>
        </a-spin>
    </div>
</template>

<script>
import { reactive, toRefs, onBeforeMount, onMounted, watch } from 'vue'
import { cateList, albumList } from '@/api/business.js';
import lo_jsonDog from '@/assets/lottie/dog.json';
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
            loading: true,
            cateList: [],
            albumList: [],
            albumList2: [],
            colorSchemes: {},
            statistics: {}
        });
        const findCategoryName = (cid) => {
            const category = allData.cateList.flat().find(cate => cate.id === cid);
            return category ? category.name : '未找到分类';
        };
        const getCateList = () => {
            allData.loading = true;
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
            }).finally(() => {
                allData.loading = false;
            });
        };
        const getAlbumList = (cid) => {
            var parameter = {
                pageNo: 1,
                pageSize: 1000,
                name: '',
                type: 10,
                cid: cid,
            };
            albumList(parameter).then((res) => {
                if (res.code == 200) {
                    const lists = res.data.lists;
                    const length = lists.length;
                    const midIndex = Math.floor(length / 2);
                    allData.albumList2 = lists.slice(0, midIndex);
                    allData.albumList = lists.slice(midIndex);
                    allData.statistics.albumCount = res.data.count;
                    const sumSizes = (lists) => {
                        let totalSize = 0;
                        for (const item of lists) {
                            const sizeInKB = parseSize(item.size);
                            totalSize += sizeInKB;
                        }
                        return totalSize;
                    };
                    allData.statistics.sizeCount = `${sumSizes(res.data.lists).toFixed(2)} KB`;
                    allData.statistics.mostRecentTime = findMostRecentTime(res.data.lists, 'updateTime').updateTime;
                } else {
                    console.log(res)
                }
            }).catch((err) => {
                console.log(err)
            });
        };
        const goToDetail = (item) => {
            let id = item?.id ? item.id : ''
            getAlbumList(id)
        };
        const getCategoryColors = (cid) => {
            const category = allData.cateList.flat().find(item => item.id === cid);
            if (category) {
                return {
                    backgroundColor: category.backgroundColor,
                    fontColor: category.fontColor
                };
            }
            return {
                backgroundColor: '#f0f0f0',
                fontColor: '#666666'
            };
        };
        /* tool */
        // 定义一个函数来处理单位换算并累加大小
        const parseSize = (sizeStr) => {
            const match = sizeStr.match(/^(\d+(\.\d+)?)\s*(KB|MB)$/i);
            if (!match) {
                throw new Error(`Invalid size format: ${sizeStr}`);
            }
            const value = parseFloat(match[1]);
            const unit = match[3].toUpperCase();

            if (unit === 'KB') {
                return value;
            } else if (unit === 'MB') {
                return value * 1024; // 1 MB = 1024 KB
            } else {
                throw new Error(`Unsupported unit: ${unit}`);
            }
        };
        const findMostRecentTime = (list, timeKey = 'updateTime') => {
            let mostRecentItem = null;
            let mostRecentTime = null;
            list.forEach(item => {
                const itemTime = new Date(item[timeKey]);
                if (!mostRecentTime || itemTime > mostRecentTime) {
                    mostRecentTime = itemTime;
                    mostRecentItem = item;
                }
            });
            return mostRecentItem;
        };
        watch(
            show,
            (val) => {
                if (val) {
                    // getCateList();
                    // getAlbumList();
                }
            },
            { deep: true, immediate: true },
        );
        onMounted(() => {
            getCateList();
            getAlbumList();
        })
        return {
            ...toRefs(allData),
            findCategoryName,
            getCategoryColors,
            goToDetail,
            allData,
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
    background: radial-gradient(at 10px 0, #a974d2, #7baff8, #a1f0f0, #e1f4ff, #f6ffff, #ffffff);
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
        cursor: pointer !important;
        background-color: #fff;
        padding: 3px;
        font-size: 30px;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
        box-shadow: 5px 5px 13px 0px rgba(117, 126, 136, 0.2);
        width: 120px;
        height: 120px;
        border-radius: 30px;
        transition: box-shadow 0.3s, font-size 0.3s;

        .group-name {
            height: 100%;
            width: 100%;
            border-radius: 30px;
            display: flex;
            justify-content: center;
            align-items: center;
        }
    }

    .tags-group-icon:hover {
        border: 2px solid #7cbbf6;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        font-size: 36px;
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
        transform: translateY(-80%);
    }
}

.main::-webkit-scrollbar {
    /*滚动条整体样式*/
    width: 10px;
    /*高宽分别对应横竖滚动条的尺寸*/
    height: 1px;
}

.main::-webkit-scrollbar-thumb {
    /*滚动条里面小方块*/
    border-radius: 5px;
    background: #929292;
}

.main::-webkit-scrollbar-track {
    /*滚动条里面轨道*/
    border-radius: 5px;
    background: #fff;
}

.articlesList-list:hover {
    cursor: pointer !important;
    border-color: transparent;
    box-shadow: 0 1px 2px -2px rgba(0, 0, 0, 0.16), 0 3px 6px 0 rgba(0, 0, 0, 0.12), 0 5px 12px 4px rgba(0, 0, 0, 0.09);
}

.announcement-box {
    background-color: #fff;
    border-radius: 5px;
    height: 5vh;
    margin-bottom: 10px;
    margin-left: 9px;
    flex: 0 0 98.833333%;
    max-width: 98.833333%;

    .announcement-box-dog {
        position: absolute;
        bottom: 0;
        animation: moveAndFlip 10s infinite linear;
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
        display: flex;
        border-radius: 10px;
        overflow: hidden;
        height: 80%;
        width: 100%;
        justify-content: center;
        align-items: center;
    }

    .title {
        margin-left: 20px;
        font-size: 20px;
    }

    .introduction {
        display: flex;
        justify-content: space-between;
        font-size: 18px;
        margin: 0 20px;
        color: #babdba;

        .tip-box {
            width: 70px;
            height: 30px;
            border-radius: 10px;
            display: flex;
            justify-content: center;
            align-items: center;
        }
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