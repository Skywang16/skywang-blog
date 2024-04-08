<template>
    <div class="article-edit">
        <el-card class="!border-none" shadow="never">
            <el-page-header content="文章编辑" @back="$router.back()" />
        </el-card>
        <el-card class="mt-4 !border-none" shadow="never">
            <el-form ref="formRef" class="ls-form" :model="formData" label-width="85px" :rules="rules">
                <div class="xl:flex">
                    <div>
                        <el-form-item label="文章标题" prop="title">
                            <div class="w-80">
                                <el-input v-model="formData.title" placeholder="请输入文章标题" type="textarea"
                                    :autosize="{ minRows: 3, maxRows: 3 }" maxlength="64" show-word-limit clearable />
                            </div>
                        </el-form-item>
                        <el-form-item label="文章简介" prop="desc">
                            <div class="w-80">
                                <el-input v-model="formData.desc" placeholder="请输入文章简介" type="textarea"
                                    :autosize="{ minRows: 3, maxRows: 6 }" :maxlength="200" show-word-limit clearable />
                            </div>
                        </el-form-item>
                        <el-form-item label="文章封面" prop="image">
                            <div>
                                <div>
                                    <material-picker v-model="formData.image" :limit="1" />
                                </div>
                                <div class="form-tips">建议尺寸：240*180px</div>
                            </div>
                        </el-form-item>
                        <el-form-item label="作者" prop="cid">
                            <div class="w-80">
                                <el-input v-model="formData.cid" placeholder="请输入作者名称" />
                            </div>
                        </el-form-item>
                        <el-form-item label="排序" prop="sort">
                            <div>
                                <el-input-number v-model="formData.sort" :min="0" :max="9999" />
                                <div class="form-tips">默认为0， 数值越大越排前</div>
                            </div>
                        </el-form-item>
                        <el-form-item label="发布时间" prop="publishTime">
                            <div class="w-80">
                                <el-input v-model="formData.publishTime" placeholder="请输入发布时间" />
                            </div>
                        </el-form-item>
                        <el-form-item label="文章状态" required prop="status">
                            <el-radio-group v-model="formData.status">
                                <el-radio :label="1">显示</el-radio>
                                <el-radio :label="0">隐藏</el-radio>
                            </el-radio-group>
                        </el-form-item>
                        <el-form-item label="文章推荐" required prop="recommended">
                            <el-radio-group v-model="formData.recommended">
                                <el-radio :label="1">是</el-radio>
                                <el-radio :label="0">否</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </div>
                    <div class="xl:ml-20">
                        <el-form-item label="文章内容" required prop="content">
                            <editor v-model="formData.content" :height="667" :width="375" />
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

<script lang="ts" setup name="articleListsEdit">
import type { FormInstance } from 'element-plus'
import feedback from '@/utils/feedback'
import { articleDetail, articleEdit, articleAdd } from '@/api/article'
import useMultipleTabs from '@/hooks/useMultipleTabs'

const route = useRoute()
const router = useRouter()
const formData = reactive({
    id: route.query.id,
    title: '',
    image: '',
    cid: '',
    desc: '',
    content: '',
    sort: 0,
    publishTime: '',
    status: 0,
    recommended: 0
})

const { removeTab } = useMultipleTabs()
const formRef = shallowRef<FormInstance>()
const rules = reactive({
    title: [{ required: true, message: '请输入文章标题', trigger: 'blur' }],
    image: [{ required: true, message: '请输入文章封面', trigger: 'blur' }],
    cid: [{ required: true, message: '请输入文章作者', trigger: 'blur' }],
    desc: [{ required: true, message: '请输入简介', trigger: 'blur' }],
    content: [{ required: true, message: '请输入文章内容', trigger: 'blur' }],
    sort: [{ required: true, message: '请输入文章排序', trigger: 'blur' }],
    publishTime: [{ required: true, message: '请输入文章发布时间', trigger: 'blur' }],
})

const getDetails = async () => {
    const data = await articleDetail({
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
        await articleEdit(formData)
    } else {
        await articleAdd(formData)
    }
    feedback.msgSuccess('操作成功')
    removeTab()
    router.back()
}

route.query.id && getDetails()
</script>
