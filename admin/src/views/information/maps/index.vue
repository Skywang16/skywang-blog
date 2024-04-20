<template>
    <div class="maps-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="地图名称">
                    <el-input class="w-[280px]" v-model="queryParams.name" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item label="地图状态">
                    <el-select class="w-[280px]" v-model="queryParams.status">
                        <el-option label="全部" :value="-1" />
                        <el-option label="显示" :value="1" />
                        <el-option label="隐藏" :value="0" />
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
                <router-link v-perms="['maps:add', 'maps:add/edit']" :to="{
                    path: getRoutePath('maps:add/edit'),
                }">
                    <el-button type="primary" class="mb-4">
                        <template #icon>
                            <icon name="el-icon-Plus" />
                        </template>
                        新增地图
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
                <el-table-column label="地图名称" prop="name" min-width="160" show-tooltip-when-overflow />
                <el-table-column label="地图地点" prop="locations" min-width="60" />
                <el-table-column label="包点数量" prop="plantQuantity" min-width="60" />
                <el-table-column label="地图机制" prop="mechanics" min-width="60" />
                <el-table-column label="状态" min-width="100">
                    <template #default="{ row }">
                        <el-switch v-perms="['maps:cate:change']" v-model="row.status" :active-value="1"
                            :inactive-value="0" @change="changeStatus(row)" />
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button v-perms="['maps:edit', 'maps:add/edit']" type="primary" link>
                            <router-link :to="{
                                path: getRoutePath('maps:add/edit'),
                                query: {
                                    id: row.id,
                                },
                            }">
                                编辑
                            </router-link>
                        </el-button>
                        <el-button v-perms="['maps:del']" type="danger" link @click="handleDelete(row.id)">
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
<script lang="ts" setup name="mapsLists">
import { mapsLists, mapsDelete, mapsEdit } from "@/api/information";
import { usePaging } from "@/hooks/usePaging";
import { getRoutePath } from "@/router";
import feedback from "@/utils/feedback";

const queryParams = reactive({
    name: "",
    status: -1,
});

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: mapsLists,
    params: queryParams,
});


const changeStatus = async (row: any) => {
    const data: any = { ...row };
    data.status = data.status === 1 ? 0 : 1;
    try {
        await mapsEdit(row);
        feedback.msgSuccess("修改成功");
        getLists();
    } catch (error) {
        getLists();
    }
};

const handleDelete = async (id: number) => {
    await feedback.confirm("确定要删除？");
    await mapsDelete({ id });
    feedback.msgSuccess("删除成功");
    getLists();
};

onActivated(() => {
    getLists();
});

getLists();
</script>
