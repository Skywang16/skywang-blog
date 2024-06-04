<template>
    <div class="login flex flex-col">
        <div class="flex-1 flex items-center justify-center">
            <div class="login-card flex rounded-md">
                <div class="flex-1 h-full hidden md:inline-block">
                    <image-contain :src="config.webBackdrop" :width="400" height="100%" />
                </div>
                <div
                    class="login-form bg-body flex flex-col justify-center px-10 py-10 md:w-[400px] w-[375px] flex-none mx-auto"
                >
                    <div class="text-center text-3xl font-medium mb-8">{{ config.webName }}</div>
                    <el-form ref="formRef" :model="formData" size="large" :rules="rules">
                        <el-form-item prop="account">
                            <el-input
                                v-model.trim="formData.account"
                                placeholder="请输入账号"
                                @keyup.enter="handleEnter"
                            >
                                <template #prepend>
                                    <icon name="el-icon-User" />
                                </template>
                            </el-input>
                        </el-form-item>
                        <el-form-item prop="password">
                            <el-input
                                ref="passwordRef"
                                v-model="formData.password"
                                show-password
                                placeholder="请输入密码"
                                @keyup.enter="handleLogin"
                            >
                                <template #prepend>
                                    <icon name="el-icon-Lock" />
                                </template>
                            </el-input>
                        </el-form-item>
                    </el-form>
                    <div class="mb-5">
                        <el-checkbox v-model="remAccount" label="记住账号"></el-checkbox>
                    </div>
                    <!-- <el-button type="primary" size="large" :loading="isLock" @click="lockLogin">
                        登录
                    </el-button> -->
                    <div class="login-button-box">
                        <div class="login-button" href="#" data-title="登  录" @click="lockLogin"></div>
                    </div>
                </div>
            </div>
        </div>
        <layout-footer />
    </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref, shallowRef } from 'vue'
import type { InputInstance, FormInstance } from 'element-plus'
import LayoutFooter from '@/layout/components/footer.vue'
import useAppStore from '@/stores/modules/app'
import useUserStore from '@/stores/modules/user'
import cache from '@/utils/cache'
import { ACCOUNT_KEY } from '@/enums/cacheEnums'
import { PageEnum } from '@/enums/pageEnum'
import { useLockFn } from '@/hooks/useLockFn'
const passwordRef = shallowRef<InputInstance>()
const formRef = shallowRef<FormInstance>()
const appStore = useAppStore()
const userStore = useUserStore()
const route = useRoute()
const router = useRouter()
const remAccount = ref(false)
const config = computed(() => appStore.config)
const formData = reactive({
    account: '',
    password: ''
})
const rules = {
    account: [
        {
            required: true,
            message: '请输入账号',
            trigger: ['blur']
        }
    ],
    password: [
        {
            required: true,
            message: '请输入密码',
            trigger: ['blur']
        }
    ]
}
// 回车按键监听
const handleEnter = () => {
    if (!formData.password) {
        return passwordRef.value?.focus()
    }
    handleLogin()
}
// 登录处理
const handleLogin = async () => {
    await formRef.value?.validate()
    // 记住账号，缓存
    cache.set(ACCOUNT_KEY, {
        remember: remAccount.value,
        account: remAccount.value ? formData.account : ''
    })
    await userStore.login(formData)
    const {
        query: { redirect }
    } = route
    const path = typeof redirect === 'string' ? redirect : PageEnum.INDEX
    router.push(path)
}
const { isLock, lockFn: lockLogin } = useLockFn(handleLogin)

onMounted(() => {
    const value = cache.get(ACCOUNT_KEY)
    if (value?.remember) {
        remAccount.value = value.remember
        formData.account = value.account
    }
})
</script>

<style lang="scss" scoped>
.login {
    background-image: url('./images/login_bg.png');
    @apply min-h-screen bg-no-repeat bg-center bg-cover;
    .login-card {
        height: 400px;
    }
}
.login-button {
    position: relative;
    display: inline-block;
    padding: 0.6em;
    text-decoration: none;
    text-align: center;
    cursor: pointer;
    user-select: none;
    color: white;
    width: 100%;
    &::before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        bottom: 0;
        right: 0;
        background: linear-gradient(135deg, #6e8efb, #a777e3);
        border-radius: 4px;
        transition: box-shadow .5s ease, transform .2s ease;
        will-change: transform;
        box-shadow: 0 2px 5px rgba(0, 0, 0, .2);
        transform:
            translateY(var(--ty, 0)) rotateX(var(--rx, 0)) rotateY(var(--ry, 0)) translateZ(var(--tz, -12px));
    }

    &:hover::before {
        box-shadow: 0 5px 15px rgba(0, 0, 0, .3);
    }
    &::after {
        position: relative;
        display: inline-block;
        content: attr(data-title);
        transition: transform .2s ease;
        font-weight: bold;
        letter-spacing: .01em;
        will-change: transform;
        transform:
            translateY(var(--ty, 0)) rotateX(var(--rx, 0)) rotateY(var(--ry, 0));
    }
}

.login-button-box {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 40px;
    transform-style: preserve-3d;
    transform: perspective(800px);
}
</style>
