<template>
    <div class="hero-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="英雄名称">
                    <el-input class="w-[280px]" v-model="queryParams.name" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item label="英雄状态">
                    <el-select class="w-[280px]" v-model="queryParams.status">
                        <el-option label="全部" :value="-1" />
                        <el-option label="显示" :value="1" />
                        <el-option label="隐藏" :value="0" />
                    </el-select>
                </el-form-item>
                <el-form-item label="英雄定位">
                    <el-select class="w-[280px]" v-model="queryParams.type">
                        <el-option v-for="item in typeList" :key="item.id" :label="item.name"
                            :value="item.id"></el-option>
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
                <router-link v-perms="['hero:add', 'hero:add/edit']" :to="{
                    path: getRoutePath('hero:add/edit'),
                }">
                    <el-button type="primary" class="mb-4">
                        <template #icon>
                            <icon name="el-icon-Plus" />
                        </template>
                        新增英雄
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
                <el-table-column label="图片" min-width="100">
                    <template #default="{ row }">
                        <image-contain v-if="row.image" :src="row.image" :width="60" :height="45"
                            :preview-src-list="[row.image]" preview-teleported fit="contain" />
                    </template>
                </el-table-column>
                <el-table-column label="英雄名称" prop="name" min-width="160" show-tooltip-when-overflow />
                <el-table-column label="英雄定位" prop="types" min-width="60">
                    <template #default="{ row }">
                        {{ typeList.find(item => item.id === row.types)?.name }}
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button v-perms="['hero:edit', 'hero:add/edit']" type="primary" link>
                            <router-link :to="{
                                path: getRoutePath('hero:add/edit'),
                                query: {
                                    id: row.id,
                                },
                            }">
                                编辑
                            </router-link>
                        </el-button>
                        <el-button v-perms="['hero:del']" type="danger" link @click="handleDelete(row.id)">
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
<script lang="ts" setup name="heroLists">
import { heroLists, heroDelete, heroEdit, heroTypeLists } from "@/api/information";
import { usePaging } from "@/hooks/usePaging";
import { getRoutePath } from "@/router";
import feedback from "@/utils/feedback";

interface heroType {
    id: number;
    name: string;
}

const queryParams = reactive({
    name: "",
    level: "",
    types: -1,
    status: -1,
    type: ''
});

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: heroLists,
    params: queryParams,
});

const typeList = reactive<heroType[]>([]);

const fetchData = async () => {
    try {
        const response = await heroTypeLists();
        console.log(response, 'response');
        if (response.lists && response.lists.length > 0) {
            typeList.splice(0, typeList.length, ...response.lists as heroType[]);
        }
    } catch (error) {
        console.error(error);
    }
};

const changeStatus = async (row: any) => {
    const data: any = { ...row };
    data.status = data.status === 1 ? 0 : 1;
    try {
        await heroEdit(row);
        feedback.msgSuccess("修改成功");
        getLists();
    } catch (error) {
        getLists();
    }
};

const handleDelete = async (id: number) => {
    await feedback.confirm("确定要删除？");
    await heroDelete({ id });
    feedback.msgSuccess("删除成功");
    getLists();
};

onActivated(() => {
    fetchData();
    getLists();
});

getLists();
fetchData();
</script>
