<template>
    <div class="article-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="文章标题">
                    <el-input class="w-[280px]" v-model="queryParams.title" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item label="文章关键词">
                    <el-input class="w-[280px]" v-model="queryParams.keyWords" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item label="文章作者">
                    <el-input class="w-[280px]" v-model="queryParams.cid" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <!-- <el-form-item label="栏目作者">
                    <el-select class="w-[280px]" v-model="queryParams.cid">
                        <el-option label="全部" value />
                        <el-option v-for="item in optionsData.articleCate" :key="item.id" :label="item.name"
                            :value="item.id" />
                    </el-select>
                </el-form-item> -->
                <el-form-item label="文章状态">
                    <el-select class="w-[280px]" v-model="queryParams.status">
                        <el-option label="全部" :value="-1" />
                        <el-option label="显示" :value="1" />
                        <el-option label="隐藏" :value="0" />
                    </el-select>
                </el-form-item>
                <el-form-item label="是否推荐">
                    <el-select class="w-[280px]" v-model="queryParams.recommended">
                        <el-option label="全部" :value="-1" />
                        <el-option label="是" :value="1" />
                        <el-option label="否" :value="0" />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <router-link v-perms="['article:add', 'article:add/edit']" :to="{
                    path: getRoutePath('article:add/edit')
                }">
                    <el-button type="primary" class="mb-4">
                        <template #icon>
                            <icon name="el-icon-Plus" />
                        </template>
                        新增资讯
                    </el-button>
                </router-link>
            </div>
            <el-table size="large" v-loading="pager.loading" :data="pager.lists">
                <el-table-column label="序号" width="60">
                    <template #default="scope">
                        {{ scope.$index + 1 }}
                    </template>
                </el-table-column>
                <el-table-column label="ID" prop="id" min-width="80" />
                <el-table-column label="封面" min-width="100">
                    <template #default="{ row }">
                        <image-contain v-if="row.image" :src="row.image" :width="60" :height="45"
                            :preview-src-list="[row.image]" preview-teleported fit="contain" />
                    </template>
                </el-table-column>
                <el-table-column label="标题" prop="title" min-width="160" show-tooltip-when-overflow />
                <el-table-column label="文章关键词" prop="keyWords" min-width="100">
                    <template #default="{ row }">
                        <!-- 使用计算属性或方法展示name列表 -->
                        {{ getKeywordNames(row) }}
                    </template>
                </el-table-column>
                <el-table-column label="作者" prop="cid" min-width="60" />
                <el-table-column label="状态" min-width="100">
                    <template #default="{ row }">
                        <el-switch v-perms="['article:cate:change']" v-model="row.status" :active-value="1"
                            :inactive-value="0" @change="changeStatus(row)" />
                    </template>
                </el-table-column>
                <el-table-column label="推荐" min-width="100">
                    <template #default="{ row }">
                        <el-switch v-perms="['article:cate:change']" v-model="row.recommended" :active-value="1"
                            :inactive-value="0" @change="changeStatus(row)" />
                    </template>
                </el-table-column>
                <el-table-column label="发布时间" prop="publishTime" min-width="120" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button v-perms="['article:edit', 'article:add/edit']" type="primary" link>
                            <router-link :to="{
                                path: getRoutePath('article:add/edit'),
                                query: {
                                    id: row.id
                                }
                            }">
                                编辑
                            </router-link>
                        </el-button>
                        <el-button v-perms="['article:del']" type="danger" link @click="handleDelete(row.id)">
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
    </div>
</template>
<script lang="ts" setup name="articleLists">
import { articleLists, articleDelete, articleEdit, keyWordsLists } from '@/api/article'
import { usePaging } from '@/hooks/usePaging'
import { getRoutePath } from '@/router'
import feedback from '@/utils/feedback'
const queryParams = reactive({
    title: '',
    keyWords: '',
    cid: '',
    status: -1,
    recommended: -1
})
interface keyWords {
    id: number
    name: string
}
const optionsData = reactive<keyWords[]>([])
const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: articleLists,
    params: queryParams
})

const changeStatus = async (row: any) => {
    const data: any = { ...row }
    data.status = data.status === 1 ? 0 : 1
    try {
        await articleEdit(row)
        feedback.msgSuccess('修改成功')
        getLists()
    } catch (error) {
        getLists()
    }
}

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await articleDelete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}
const getKeywordNames = (row: any) => {
    const ids = row.keyWords.split(',').map((id: string) => parseInt(id, 10))
    console.log(optionsData)
    const names = ids
        .map((id: number) => {
            const option = optionsData.find((option) => option.id === id)
            return option ? option.name : null
        })
        .filter((name) => name !== null)
    console.log(names, 'asdasd')
    return names.join(', ')
}
const fetchData = async () => {
    try {
        const params = {
            pageNo: 1,
            pageSize: 100
        }
        const response = await keyWordsLists(params)
        if (response.lists && response.lists.length > 0) {
            optionsData.splice(0, optionsData.length, ...(response.lists as keyWords[]))
        }
    } catch (error) {
        console.error(error)
    }
}

onActivated(() => {
    getLists()
    fetchData()
})

getLists()
fetchData()
</script>
