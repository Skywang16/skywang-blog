<template>
    <div class="maps-edit">
        <el-card class="!border-none" shadow="never">
            <el-page-header price="地图编辑" @back="$router.back()" />
        </el-card>
        <el-card class="mt-4 !border-none" shadow="never">
            <el-form ref="formRef" class="ls-form" :model="formData" label-width="85px" :rules="rules">
                <div class="xl:flex">
                    <div>
                        <el-form-item label="地图名称" prop="name">
                            <div class="w-80">
                                <el-input v-model="formData.name" placeholder="请输入地图名称"
                                    :autosize="{ minRows: 3, maxRows: 3 }" maxlength="20" show-word-limit clearable />
                            </div>
                        </el-form-item>
                        <el-form-item label="地图地点" prop="locations">
                            <div class="w-80">
                                <el-input v-model="formData.locations" placeholder="请输入地图地点"
                                    :autosize="{ minRows: 3, maxRows: 6 }" maxlength="8" show-word-limit clearable />
                            </div>
                        </el-form-item>
                        <el-form-item label="包点数量" prop="plantQuantity">
                            <div class="w-80">
                                <el-input-number v-model="formData.plantQuantity" :min="0" :max="3" />
                            </div>
                        </el-form-item>
                        <el-form-item label="地图机制" prop="mechanics">
                            <div class="w-80">
                                <el-input v-model="formData.mechanics" placeholder="请输入地图机制"
                                    :autosize="{ minRows: 3, maxRows: 6 }" maxlength="8" show-word-limit clearable />
                            </div>
                        </el-form-item>
                        <el-form-item label="地图图片" prop="image">
                            <div>
                                <div>
                                    <material-picker v-model="formData.image" :limit="1" />
                                </div>
                                <div class="form-tips">建议尺寸：240*180px</div>
                            </div>
                        </el-form-item>
                    </div>
                </div>
            </el-form>
        </el-card>
        <footer-btns>
            <el-button type="primary" @click="handleSave">保存</el-button>
        </footer-btns>
    </div>
</template>

<script lang="ts" setup name="mapsListsEdit">
import type { FormInstance } from 'element-plus'
import feedback from '@/utils/feedback'
import { mapsDetail, mapsEdit, mapsAdd } from '@/api/information'
import useMultipleTabs from '@/hooks/useMultipleTabs'

const route = useRoute()
const router = useRouter()


const formData = reactive({
    id: route.query.id,
    name: '',
    image: '',
    locations: '',
    plantQuantity: 1,
    mechanics: '',
    status: 1,
})

const { removeTab } = useMultipleTabs()
const formRef = shallowRef<FormInstance>()
const rules = reactive({
    name: [{ required: true, message: '请输入地图名称', trigger: 'blur' }],
    image: [{ required: true, message: '请输入地图图片', trigger: 'blur' }],
    locations: [{ required: true, message: '请输入地图地点', trigger: 'blur' }],
    plantQuantity: [{ required: true, message: '请输入包点数量', trigger: 'blur' }],
    mechanics: [{ required: true, message: '请输入地图机制', trigger: 'blur' }],
    series: [{ required: true, message: '请输入地图系列', trigger: 'blur' }],
})

const getDetails = async () => {
    const data = await mapsDetail({
        id: route.query.id
    })
    Object.keys(formData).forEach((key) => {
        //@ts-ignore
        formData[key] = data[key]
    })
}



const handleSave = async () => {
    await formRef.value?.validate()
    if (route.query.id) {
        await mapsEdit(formData)
    } else {
        await mapsAdd(formData)
    }
    feedback.msgSuccess('操作成功')
    removeTab()
    router.back()
}

route.query.id && getDetails()
</script>
