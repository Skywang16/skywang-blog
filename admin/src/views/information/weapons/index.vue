<template>
    <div class="weapons-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="武器名称">
                    <el-input class="w-[280px]" v-model="queryParams.name" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item label="武器等级">
                    <el-select class="w-[320px]" v-model="queryParams.level" clearable @keyup.enter="resetPage">
                        <el-option label="精选" :value="0" />
                        <el-option label="奢华" :value="1" />
                        <el-option label="尊爵" :value="2" />
                        <el-option label="究极" :value="3" />
                        <el-option label="限定" :value="4" />
                    </el-select>
                </el-form-item>
                <!-- <el-form-item label="栏目作者">
                    <el-select class="w-[280px]" v-model="queryParams.level">
                        <el-option label="全部" value />
                        <el-option v-for="item in optionsData.weaponsCate" :key="item.id" :label="item.name"
                            :value="item.id" />
                    </el-select>
                </el-form-item> -->
                <el-form-item label="武器状态">
                    <el-select class="w-[280px]" v-model="queryParams.status">
                        <el-option label="全部" :value="-1" />
                        <el-option label="显示" :value="1" />
                        <el-option label="隐藏" :value="0" />
                    </el-select>
                </el-form-item>
                <el-form-item label="武器类型">
                    <el-select class="w-[280px]" v-model="queryParams.types">
                        <el-option label="全部" :value="-1" />
                        <el-option label="随身武器" :value="0" />
                        <el-option label="冲锋枪" :value="1" />
                        <el-option label="散弹枪" :value="2" />
                        <el-option label="步枪" :value="3" />
                        <el-option label="狙击步枪" :value="4" />
                        <el-option label="机枪" :value="5" />
                    </el-select>
                </el-form-item>
                <el-form-item label="武器系列">
                    <el-select class="w-[280px]" v-model="queryParams.series">
                        <el-option v-for="item in seriesList" :key="item.id" :label="item.name"
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
                <router-link v-perms="['weapons:add', 'weapons:add/edit']" :to="{
                    path: getRoutePath('weapons:add/edit'),
                }">
                    <el-button type="primary" class="mb-4">
                        <template #icon>
                            <icon name="el-icon-Plus" />
                        </template>
                        新增武器
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
                <el-table-column label="武器名称" prop="name" min-width="160" show-tooltip-when-overflow />
                <el-table-column label="武器等级" prop="level" min-width="60">
                    <template #default="{ row }">
                        <span v-if="row.level === 0">精选</span>
                        <span v-else-if="row.level === 1">奢华</span>
                        <span v-else-if="row.level === 2">尊爵</span>
                        <span v-else-if="row.level === 3">究极</span>
                        <span v-else-if="row.level === 4">限定</span>
                    </template>
                </el-table-column>
                <el-table-column label="武器类型" prop="types" min-width="60">
                    <template #default="{ row }">
                        <span v-if="row.types === 0">随身武器</span>
                        <span v-else-if="row.types === 1">冲锋枪</span>
                        <span v-else-if="row.types === 2">散弹枪</span>
                        <span v-else-if="row.types === 3">步枪</span>
                        <span v-else-if="row.types === 4">狙击步枪</span>
                        <span v-else-if="row.types === 5">机枪</span>
                    </template>
                </el-table-column>
                <el-table-column label="武器系列" min-width="60">
                    <template #default="{ row }">
                        {{ seriesList.find(item => item.id === row.series)?.name }}
                    </template>
                </el-table-column>
                <el-table-column label="状态" min-width="100">
                    <template #default="{ row }">
                        <el-switch v-perms="['weapons:cate:change']" v-model="row.status" :active-value="1"
                            :inactive-value="0" @change="changeStatus(row)" />
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button v-perms="['weapons:edit', 'weapons:add/edit']" type="primary" link>
                            <router-link :to="{
                                path: getRoutePath('weapons:add/edit'),
                                query: {
                                    id: row.id,
                                },
                            }">
                                编辑
                            </router-link>
                        </el-button>
                        <el-button v-perms="['weapons:del']" type="danger" link @click="handleDelete(row.id)">
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
<script lang="ts" setup name="weaponsLists">
import { weaponsLists, weaponsDelete, weaponsEdit, weaponsSeriesLists } from "@/api/information";
import { usePaging } from "@/hooks/usePaging";
import { getRoutePath } from "@/router";
import feedback from "@/utils/feedback";

interface WeaponSeries {
    id: number;
    name: string;
}

const queryParams = reactive({
    name: "",
    level: "",
    types: -1,
    status: -1,
    series: ''
});

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: weaponsLists,
    params: queryParams,
});

const seriesList = reactive<WeaponSeries[]>([]);

const fetchData = async () => {
    try {
        const response = await weaponsSeriesLists();
        console.log(response, 'response');
        if (response.lists && response.lists.length > 0) {
            seriesList.splice(0, seriesList.length, ...response.lists as WeaponSeries[]);
        }
    } catch (error) {
        console.error(error);
    }
};

const changeStatus = async (row: any) => {
    const data: any = { ...row };
    data.status = data.status === 1 ? 0 : 1;
    try {
        await weaponsEdit(row);
        feedback.msgSuccess("修改成功");
        getLists();
    } catch (error) {
        getLists();
    }
};

const handleDelete = async (id: number) => {
    await feedback.confirm("确定要删除？");
    await weaponsDelete({ id });
    feedback.msgSuccess("删除成功");
    getLists();
};

onActivated(() => {
    fetchData();  // 在页面加载时调用接口获取武器类型列表
    getLists();
});

getLists();
fetchData();
</script>
