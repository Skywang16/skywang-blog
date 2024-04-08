<template>
    <div class="series-lists">
        <el-card class="!border-none" shadow="never">
            <div>
                <el-button v-perms="['weapons:series:add']" type="primary" @click="handleAdd">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <div class="mt-4">
                <div>
                    <el-table :data="pager.lists" size="large" v-loading="pager.loading">
                        <el-table-column prop="id" label="ID" min-width="100" />
                        <el-table-column prop="name" label="名称" min-width="150" />
                        <el-table-column prop="remark" label="备注" min-width="150" show-overflow-tooltip />
                        <el-table-column prop="sort" label="排序" min-width="100" />
                        <el-table-column label="是否启用" min-width="100">
                            <template #default="{ row }">
                                <el-switch v-perms="['weapons:cate:change']" v-model="row.isDisable" :active-value="0"
                                    :inactive-value="1" @change="changeStatus(row)" />
                            </template>
                        </el-table-column>
                        <el-table-column prop="createTime" label="创建时间" min-width="180" />
                        <el-table-column label="操作" width="190" fixed="right">
                            <template #default="{ row }">
                                <el-button v-perms="['weapons:series:edit']" link type="primary"
                                    @click="handleEdit(row)">
                                    编辑
                                </el-button>
                                <el-button v-perms="['weapons:series:del']" link type="danger"
                                    @click="handleDelete(row.id)">
                                    删除
                                </el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </div>
                <div class="flex justify-end mt-4">
                    <pagination v-model="pager" @change="getLists" />
                </div>
            </div>
        </el-card>
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
    </div>
</template>

<script lang="ts" setup name="series">
import { weaponsSeriesLists, weaponsSeriesDelete, weaponsSeriesEdit } from "@/api/information";
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const { pager, getLists } = usePaging({
    fetchFun: weaponsSeriesLists
})
const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.setFormData(data)
}

// 修改状态
const changeStatus = async (row: any) => {
    const data: any = { ...row };
    data.isDisable = data.isDisable === 1 ? 0 : 1;
    try {
        await weaponsSeriesEdit(row);
        feedback.msgSuccess("修改成功");
        getLists();
    } catch (error) {
        getLists();
    }
};

// 删除系列
const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await weaponsSeriesDelete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

getLists()
</script>
