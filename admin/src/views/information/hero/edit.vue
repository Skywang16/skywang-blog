<template>
    <div class="hero-edit">
        <el-card class="!border-none" shadow="never">
            <el-page-header price="角色编辑" @back="$router.back()" />
        </el-card>
        <el-card class="mt-4 !border-none" shadow="never">
            <el-form ref="formRef" class="ls-form" :model="formData" label-width="85px" :rules="rules">
                <div class="xl:flex">
                    <div>
                        <el-form-item label="角色名称" prop="name">
                            <div class="w-80">
                                <el-input v-model="formData.name" placeholder="请输入角色名称"
                                    :autosize="{ minRows: 3, maxRows: 3 }" maxlength="20" show-word-limit clearable />
                            </div>
                        </el-form-item>
                        <el-form-item label="角色定位" prop="type">
                            <el-select class="w-[320px]" v-model="formData.type" :value-key="'id'">
                                <el-option v-for="item in typeList" :key="item.id" :label="item.name"
                                    :value="item"></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="角色图片" prop="image">
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

<script lang="ts" setup name="heroListsEdit">
import type { FormInstance } from 'element-plus'
import feedback from '@/utils/feedback'
import { heroDetail, heroEdit, heroAdd, heroTypeLists } from '@/api/information'
import useMultipleTabs from '@/hooks/useMultipleTabs'

const route = useRoute()
const router = useRouter()

interface WeaponType {
    id: number;
    name: string;
}

const formData = reactive({
    id: route.query.id,
    name: '',
    image: '',
    type: '',
    level: '',
    price: '',
    status: 1,
})

const { removeTab } = useMultipleTabs()
const formRef = shallowRef<FormInstance>()
const rules = reactive({
    name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
    image: [{ required: true, message: '请输入角色图片', trigger: 'blur' }],
    type: [{ required: true, message: '请输入角色定位', trigger: 'blur' }],
})

const getDetails = async () => {
    const data = await heroDetail({
        id: route.query.id
    })
    Object.keys(formData).forEach((key) => {
        //@ts-ignore
        formData[key] = data[key]
    })
}

const typeList = reactive<WeaponType[]>([]);

const fetchData = async () => {
    try {
        const response = await heroTypeLists();
        console.log(response, 'response')
        typeList.splice(0, typeList.length, ...response.lists as WeaponType[]);  // 将接口返回的数据填充到typeList中
        console.log(typeList, 'typeListtypeListtypeList')
    } catch (error) {
        console.error(error);
    }
};

const handleSave = async () => {
    formData.level = Number(formData.level)
    formData.types = Number(formData.types)
    formData.type = String(formData.type)
    await formRef.value?.validate()
    if (route.query.id) {
        await heroEdit(formData)
    } else {
        await heroAdd(formData)
    }
    feedback.msgSuccess('操作成功')
    removeTab()
    router.back()
}

route.query.id && getDetails()
fetchData();
</script>
