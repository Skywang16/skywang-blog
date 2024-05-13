<template>
    <div class="weapons-edit">
        <el-card class="!border-none" shadow="never">
            <el-page-header price="武器编辑" @back="$router.back()" />
        </el-card>
        <el-card class="mt-4 !border-none" shadow="never">
            <el-form ref="formRef" class="ls-form" :model="formData" label-width="85px" :rules="rules">
                <div class="xl:flex">
                    <div>
                        <el-form-item label="武器名称" prop="name">
                            <div class="w-80">
                                <el-input v-model="formData.name" placeholder="请输入武器名称"
                                    :autosize="{ minRows: 3, maxRows: 3 }" maxlength="20" show-word-limit clearable />
                            </div>
                        </el-form-item>
                        <el-form-item label="武器类型" prop="types">
                            <div class="w-80">
                                <el-select class="w-[320px]" v-model="formData.types">
                                    <el-option label="随身武器" :value="0" />
                                    <el-option label="冲锋枪" :value="1" />
                                    <el-option label="散弹枪" :value="2" />
                                    <el-option label="步枪" :value="3" />
                                    <el-option label="狙击步枪" :value="4" />
                                    <el-option label="机枪" :value="5" />
                                </el-select>
                            </div>
                        </el-form-item>
                        <el-form-item label="武器等级" prop="level">
                            <div class="w-80">
                                <el-select class="w-[320px]" v-model="formData.level">
                                    <el-option label="精选" :value="0" />
                                    <el-option label="奢华" :value="1" />
                                    <el-option label="尊爵" :value="2" />
                                    <el-option label="究极" :value="3" />
                                    <el-option label="限定" :value="4" />
                                </el-select>
                            </div>
                        </el-form-item>
                        <el-form-item label="武器价格" prop="price">
                            <div class="w-80">
                                <el-input v-model="formData.price" placeholder="请输入武器价格"
                                    :autosize="{ minRows: 3, maxRows: 6 }" maxlength="8" show-word-limit clearable />
                            </div>
                        </el-form-item>
                        <el-form-item label="武器系列" prop="series">
                            <el-select class="w-[320px]" v-model="formData.series" :value-key="'id'">
                                <el-option v-for="item in seriesList" :key="item.id" :label="item.name"
                                    :value="item"></el-option>
                            </el-select>
                        </el-form-item>
                        <el-form-item label="武器图片" prop="image">
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

<script lang="ts" setup name="weaponsListsEdit">
import type { FormInstance } from 'element-plus'
import feedback from '@/utils/feedback'
import { weaponsDetail, weaponsEdit, weaponsAdd, weaponsSeriesLists } from '@/api/information'
import useMultipleTabs from '@/hooks/useMultipleTabs'

const route = useRoute()
const router = useRouter()

interface WeaponSeries {
    id: number;
    name: string;
}

const formData = reactive({
    id: route.query.id,
    name: '',
    image: '',
    series: '',
    level: '',
    types: '',
    price: '',
    status: 1,
})

const { removeTab } = useMultipleTabs()
const formRef = shallowRef<FormInstance>()
const rules = reactive({
    name: [{ required: true, message: '请输入武器名称', trigger: 'blur' }],
    image: [{ required: true, message: '请输入武器图片', trigger: 'blur' }],
    level: [{ required: true, message: '请输入武器等级', trigger: 'blur' }],
    type: [{ required: true, message: '请输入武器类型', trigger: 'blur' }],
    price: [{ required: true, message: '请输入武器价格', trigger: 'blur' }],
    series: [{ required: true, message: '请输入武器系列', trigger: 'blur' }],
})

const getDetails = async () => {
    const data = await weaponsDetail({
        id: route.query.id
    })
    Object.keys(formData).forEach((key) => {
        //@ts-ignore
        formData[key] = data[key]
    })
}

const seriesList = reactive<WeaponSeries[]>([]);

const fetchData = async () => {
    try {
        const response = await weaponsSeriesLists();
        console.log(response, 'response')
        seriesList.splice(0, seriesList.length, ...response.lists as WeaponSeries[]);  // 将接口返回的数据填充到seriesList中
        console.log(seriesList, 'seriesListseriesListseriesList')
    } catch (error) {
        console.error(error);
    }
};

const handleSave = async () => {
    formData.level = Number(formData.level)
    formData.types = Number(formData.types)
    formData.series = String(formData.series.id)
    await formRef.value?.validate()
    if (route.query.id) {
        await weaponsEdit(formData)
    } else {
        await weaponsAdd(formData)
    }
    feedback.msgSuccess('操作成功')
    removeTab()
    router.back()
}

route.query.id && getDetails()
fetchData();
</script>
