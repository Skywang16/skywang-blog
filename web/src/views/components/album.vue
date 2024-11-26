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
                                        :class="{ 'selected': allData.selectedCategoryId === item.id }"
                                        :style="{ backgroundColor: item.backgroundColor, color: item.fontColor }">
                                        {{ item.name }}
                                    </div>
                                </div>
                            </div>
                            <div v-for="(group, groupIndex) in cateList" :key="groupIndex" class="tags-group-icon-pair">
                                <div v-for="(item, index) in group" :key="index" class="tags-group-icon"
                                    @click="goToDetail(item)">
                                    <div class="group-name"
                                        :class="{ 'selected': allData.selectedCategoryId === item.id }"
                                        :style="{ backgroundColor: item.backgroundColor, color: item.fontColor }">
                                        {{ item.name }}
                                    </div>
                                </div>
                            </div>
                            <div v-for="(group, groupIndex) in cateList" :key="groupIndex" class="tags-group-icon-pair">
                                <div v-for="(item, index) in group" :key="index" class="tags-group-icon"
                                    @click="goToDetail(item)">
                                    <div class="group-name"
                                        :class="{ 'selected': allData.selectedCategoryId === item.id }"
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
import { reactive, toRefs, onMounted, watch } from 'vue'
import { cateList, albumList } from '@/api/business.js';
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
            statistics: {},
            selectedCategoryId: null
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
            if (allData.selectedCategoryId === id) {
                allData.selectedCategoryId = null;
                getAlbumList();
            } else {
                allData.selectedCategoryId = id;
                getAlbumList(id);
            }
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
        }
    },
}

</script>

<style scoped lang='scss'>
.card {
    border-radius: 15px;
    background-color: transparent;
}

.welcome-box {
    border-radius: 15px;
    background: #fff;
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
        cursor: pointer;
        background-color: #fff;
        padding: 3px;
        font-size: 30px;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
        width: 120px;
        height: 120px;
        border-radius: 30px;
        transition: all 0.3s ease;

        &:hover {
            transform: translateY(-5px);
            box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
        }

        .group-name {
            height: 100%;
            width: 100%;
            border-radius: 30px;
            display: flex;
            justify-content: center;
            align-items: center;
            transition: all 0.3s ease;
            border: 2px solid transparent;

            &:hover {
                opacity: 1 !important;
                transform: translateY(-2px);
                box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
            }

            &.selected {
                border: 2px solid #4a9eff;
                box-shadow: 0 0 10px rgba(74, 158, 255, 0.3);
                transform: scale(0.95);
            }
        }
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
    border-radius: 15px;
    height: 5vh;
    margin-bottom: 10px;
    margin-left: 9px;
    flex: 0 0 98.833333%;
    max-width: 98.833333%;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.articlesList-box {
    break-inside: avoid;
    margin-bottom: 20px;
}

.articlesList-list {
    background-color: #fff;
    border-radius: 15px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
    height: 50vh;
    padding-bottom: 20px;
    margin-bottom: 4vh;
    box-sizing: border-box;
    transition: all 0.3s ease;

    &:hover {
        transform: translateY(-5px);
        box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
    }

    .image-box {
        display: flex;
        border-radius: 15px;
        overflow: hidden;
        height: 80%;
        width: 100%;
        justify-content: center;
        align-items: center;

        img {
            width: 100%;
            height: 100%;
            object-fit: cover;
            transition: transform 0.5s ease;

            &:hover {
                transform: scale(1.05);
            }
        }
    }

    .title {
        margin: 15px 20px;
        font-size: 20px;
        color: #2c3e50;
        font-weight: 600;
        transition: color 0.3s ease;

    }

    .introduction {
        display: flex;
        justify-content: space-between;
        font-size: 16px;
        margin: 10px 20px;
        color: #babdba;

        .tip-box {
            padding: 5px 15px;
            border-radius: 20px;
            display: flex;
            justify-content: center;
            align-items: center;
            transition: all 0.3s ease;
        }
    }
}

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

@media screen and (max-width: 768px) {
    .welcome-box {
        display: none;

        .welcome-text {
            font-size: 28px;
            margin: 20px;
        }
    }

    .group-wrapper {
        width: 100%;
        margin: 0;
        padding: 10px;

        .tags-group-icon {
            width: 100px;
            height: 100px;
            font-size: 24px;
        }
    }

    .articlesList-list {
        height: auto;

        .title {
            font-size: 18px;
            margin: 10px 15px;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
        }

        .introduction {
            font-size: 14px;
        }
    }

    .a-row {
        .a-col {
            width: 100% !important;
            flex: 0 0 100% !important;
            max-width: 100% !important;
        }
    }

    .articlesList-box {
        width: 100% !important;
        flex: 0 0 100% !important;
        max-width: 100% !important;
    }

    .articlesList-list {
        height: auto;
        min-height: 300px;
        margin-bottom: 15px;

        .image-box {
            height: 200px;
        }

        .title {
            font-size: 18px;
            margin: 10px 15px;
        }

        .introduction {
            font-size: 14px;
            margin: 8px 15px;
            
            .tip-box {
                padding: 3px 10px;
            }
        }
    }

    .row {
        margin: 0 10px !important;
    }
}

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
</style>