<template>
    <a-row :gutter="16">
        <a-col :span="4">
            <a-card style="background-color: rgba(0, 0, 0, 0.1);">
                <header>
                    <div class="header-box">
                        <div class="header"></div>
                    </div>
                    <div class="name-box">
                        <p>sky-wang</p>
                    </div>
                    <div class="count-box">
                        <div>
                            <a-row :gutter="24">
                                <a-col class="count-text" :span="8">文章</a-col>
                                <a-col class="count-text" :span="8">标签</a-col>
                                <a-col class="count-text" :span="8">分类</a-col>
                            </a-row>
                            <a-row :gutter="24">
                                <a-col class="count-text" :span="8">4</a-col>
                                <a-col class="count-text" :span="8">2</a-col>
                                <a-col class="count-text" :span="8">12</a-col>
                            </a-row>
                        </div>
                    </div>
                </header>
            </a-card>
        </a-col>
        <a-col class="articles-box" :span="19">
            <a-row :gutter="16">
                <a-col :span="16">
                    <div class="articles-list" v-for="(item, index) in articles" :key="index" @click="goToDetail(item)">
                        <div class="image-box">
                            <img src="@/assets/img/7b58525685c7fa8f58fac97575111bf.png" alt="">
                        </div>
                        <div class="title">{{ item.title }}</div>
                        <div class="introduction">{{ item.title }}</div>
                        <div class="introduction">{{ item.publishTime }}</div>
                    </div>
                </a-col>
                <a-col :span="6">
                    <a-card style="background-color: rgba(0, 0, 0, 0.1);">
                        <div class="comment-box">
                            <div class="comment-icon-box">
                                <img style="height: 40px;" src="@/assets/icons/msg.png" alt="">
                                <span>评论</span>
                            </div>

                        </div>
                    </a-card>
                </a-col>
            </a-row>

        </a-col>
    </a-row>
    <!-- <div class="main">
        <header>
            <div class="header-box">123</div>
        </header>
    </div> -->
</template>

<script>
import { reactive, toRefs, onBeforeMount, onMounted } from 'vue'
export default {
    name: 'wContent',
    setup() {
        console.log('1-开始创建组件-setup')
        const data = reactive({
            articles: [
                { title: '文章标题 1', summary: '文章摘要 1...' },
                { title: '文章标题 2', summary: '文章摘要 2...' },
                { title: '文章标题 3', summary: '文章摘要 3...' },
                { title: '文章标题 4', summary: '文章摘要 4...' },
            ],
        });
        const goToDetail = (item) => {
            console.log('点击文章', item)
        }

        onBeforeMount(() => {
            console.log('2.组件挂载页面之前执行----onBeforeMount')
        })
        onMounted(() => {
            console.log('3.-组件挂载到页面之后执行-------onMounted')
        })
        return {
            ...toRefs(data),
            goToDetail
        }
    },
}

</script>

<style scoped lang='scss'>
header {
    .header-box {
        display: flex;
        justify-content: center;
        margin-bottom: 5%;

        .header {
            height: 14vh;
            width: 14vh;
            border-radius: 50%;
            background-color: black;
        }
    }

    .name-box {
        display: flex;
        justify-content: center;

        p {
            font-size: 36px;
        }
    }

    .count-box {
        display: flex;
        justify-content: center;

        .count-text {
            font-size: 20px;
            display: flex;
            justify-content: center;

        }
    }


}

.articles-box::-webkit-scrollbar {
    /*滚动条整体样式*/
    width: 10px;
    /*高宽分别对应横竖滚动条的尺寸*/
    height: 1px;
}

.articles-box::-webkit-scrollbar-thumb {
    /*滚动条里面小方块*/
    border-radius: 5px;
    background: #929292;
}

.articles-box::-webkit-scrollbar-track {
    /*滚动条里面轨道*/
    border-radius: 5px;
    background: #ededed;
}

.articles-box {
    overflow-y: auto;
    height: 100vh;
}

.articles-list:hover {
    cursor: pointer;
    /* 鼠标悬停时保持手指形状 */
    /* 可选：增加圆角 */
}

.articles-list {
    padding: 0 0 20px 0;
    border-radius: 10px;
    margin: 0 0 30px 10px;
    background-color: #fff;
    box-shadow: 5px 5px 13px 0px rgba(117, 126, 136, 0.2);

    .image-box {
        border-radius: 10px;
        overflow: hidden;
        height: 350px;
        width: 100%;

        img {
            width: 100%;
        }
    }

    .title {
        margin: 20px;
        font-size: 36px;
    }

    .introduction {
        font-size: 26px;
        margin: 20px;
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

.comment-box {
    .comment-icon-box {
        background-color: #fff;
        padding: 10px;
        border-radius: 0 10px 10px 0;

        span {
            font-size: 18px;
            font-weight: 600;
            margin-left: 10px;
        }
    }
}
</style>