<script setup lang="ts">
import { ref,onMounted } from 'vue'
import { ElMessage, ElLoading } from 'element-plus'
import { Lock, Phone } from '@element-plus/icons-vue'
import http from './utils/axios'

// 响应式状态
const phoneNumber = ref('17837140378')
const password = ref('123456')
const showPassword = ref(false)
const loading = ref(false)
const phoneForm = ref()
const passwordForm = ref()
const isRegistering = ref(false) // 标记是否处于注册流程

// 手机号验证规则
const phoneRules = {
    phoneNumber: [
        { required: true, message: '请输入手机号码', trigger: 'blur' },
        { pattern: /^1[3-9]\d{9}$/, message: '请输入有效的手机号码', trigger: 'blur' }
    ]
}

// 密码验证规则
const passwordRules = {
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, message: '密码长度至少为6个字符', trigger: 'blur' }
    ]
}

// 验证手机号
async function verifyPhone() {
    if (!phoneForm.value) return

    await phoneForm.value.validate(async (valid: boolean) => {
        if (!valid) return

        const loadingInstance = ElLoading.service({
            lock: true,
            text: '验证手机号...',
            background: 'rgba(0, 0, 0, 0.7)'
        })

        loading.value = true

        try {
            const response = await http.post('/login', { phone: phoneNumber.value })
            // 如果手机号存在，显示密码输入框
            isRegistering.value = false // 确保不是注册模式
            showPassword.value = true
        } catch (error: any) {
            // 检查是否是手机号不存在的错误（code=10001）
            if (error && error.code === 10001) {
                // 切换到注册模式
                isRegistering.value = true
                ElMessage.info('该手机号未注册，请注册新账号')
            } else {
                // 其他错误，保持登录模式
                isRegistering.value = false
                showPassword.value = false
            }
        } finally {
            loading.value = false
            loadingInstance.close()
        }
    })
}

// 开始注册流程
async function startRegistration() {
    if (!phoneForm.value) return

    await phoneForm.value.validate(async (valid: boolean) => {
        if (!valid) return

        const loadingInstance = ElLoading.service({
            lock: true,
            text: '处理中...',
            background: 'rgba(0, 0, 0, 0.7)'
        })

        loading.value = true

        try {
            // 第一步注册请求，验证手机号是否可用
            const response = await http.post('/register', { phone: phoneNumber.value })
            // 如果请求成功，显示密码输入框
            showPassword.value = true
        } catch (error) {
            // 错误已在拦截器中处理
            showPassword.value = false
        } finally {
            loading.value = false
            loadingInstance.close()
        }
    })
}

// 完成注册
async function completeRegistration() {
    if (!passwordForm.value) return

    await passwordForm.value.validate(async (valid: boolean) => {
        if (!valid) return

        const loadingInstance = ElLoading.service({
            lock: true,
            text: '注册中...',
            background: 'rgba(0, 0, 0, 0.7)'
        })

        loading.value = true

        try {
            const response = await http.post('/register', {
                phone: phoneNumber.value,
                password: password.value
            })

            // 注册成功
            ElMessage.success('注册成功，即将登录...')
            
            // 注册成功后直接登录，获取redirectUrl
            const loginResponse = await http.post('/login', {
                phone: phoneNumber.value,
                password: password.value
            })
            
            const loginData = loginResponse.data
            if (loginData && loginData.data && loginData.data.redirectUrl) {
                setTimeout(() => {
                    window.location.href = loginData.data.redirectUrl
                }, 1000)
            }
        } catch (error) {
            // 错误已在拦截器中处理
        } finally {
            loading.value = false
            loadingInstance.close()
        }
    })
}

// 登录
async function login() {
    if (!passwordForm.value) return

    await passwordForm.value.validate(async (valid: boolean) => {
        if (!valid) return

        const loadingInstance = ElLoading.service({
            lock: true,
            text: '登录中...',
            background: 'rgba(0, 0, 0, 0.7)'
        })

        loading.value = true

        try {
            const response = await http.post('/login', {
                phone: phoneNumber.value,
                password: password.value
            })

            // 登录成功，跳转到后端返回的URL
            const data = response.data
            if (data && data.data && data.data.redirectUrl) {
                ElMessage.success('登录成功，即将跳转...')
                setTimeout(() => {
                    window.location.href = data.data.redirectUrl
                }, 1000)
            }
        } catch (error) {
            // 错误已在拦截器中处理
        } finally {
            loading.value = false
            loadingInstance.close()
        }
    })
}

// 处理表单提交
function handleSubmit() {
    if (!showPassword.value) {
        // 第一步：验证手机号或开始注册
        if (isRegistering.value) {
            startRegistration()
        } else {
            verifyPhone()
        }
    } else {
        // 第二步：登录或完成注册
        if (isRegistering.value) {
            completeRegistration()
        } else {
            login()
        }
    }
}

function handleOuterSizeLoginEvent(){
    // 处理外部登录事件
    const params = new URLSearchParams(window.location.search)
    const clientId = params.get('client_id')
    const redirectUri = params.get('redirect_uri')
    const responseType = params.get('response_type')
    const scope = params.get('scope')
    const state = params.get('state')

    if (clientId && redirectUri && responseType && scope && state) {
        // 这里可以根据需要处理这些参数
        console.log('外部登录参数:', { clientId, redirectUri, responseType, scope, state })
        alert(`外部登录参数:\nclient_id: ${clientId}\nredirect_uri: ${redirectUri}\nresponse_type: ${responseType}\nscope: ${scope}\nstate: ${state}`)
    }
}

onMounted(()=>{
    // 检查是否有第三方登录参数
    handleOuterSizeLoginEvent()
})

</script>

<template>
    <div class="login-container">
        <el-card class="login-box">
            <template #header>
                <div class="card-header">
                    <h2>{{ isRegistering ? '用户注册' : '用户登录' }}</h2>
                </div>
            </template>

            <div v-if="!showPassword">
                <el-form ref="phoneForm" :model="{ phoneNumber }" :rules="phoneRules" label-position="top"
                    @submit.prevent="handleSubmit">
                    <el-form-item label="手机号码" prop="phoneNumber">
                        <el-input v-model="phoneNumber" placeholder="请输入手机号码" maxlength="11" :prefix-icon="Phone"
                            clearable />
                    </el-form-item>

                    <el-button type="primary" :loading="loading" @click="handleSubmit" round class="submit-btn">
                        {{ isRegistering ? '注册' : '下一步' }}
                    </el-button>
                </el-form>
            </div>

            <div v-else>
                <el-form ref="passwordForm" :model="{ password }" :rules="passwordRules" label-position="top"
                    @submit.prevent="handleSubmit">
                    <el-alert type="success" :closable="false" class="mb-4">
                        {{ isRegistering ? '请设置账号密码' : '请输入账号密码' }}
                    </el-alert>

                    <el-form-item :label="isRegistering ? '设置密码' : '密码'" prop="password">
                        <el-input v-model="password" placeholder="请输入密码" show-password :prefix-icon="Lock" />
                    </el-form-item>

                    <div class="btn-group">
                        <el-button @click="showPassword = false">返回</el-button>
                        <el-button type="primary" :loading="loading" @click="handleSubmit" round>
                            {{ isRegistering ? '完成注册' : '登录' }}
                        </el-button>
                    </div>
                </el-form>
            </div>
        </el-card>
    </div>
</template>

<style scoped>
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    background-color: #f5f7fa;
}

.login-box {
    width: 100%;
    max-width: 420px;
}

.card-header {
    display: flex;
    justify-content: center;
    align-items: center;
}

.card-header h2 {
    font-size: 24px;
    font-weight: 500;
}

.submit-btn {
    width: 100%;
    margin-top: 20px;
}

.btn-group {
    display: flex;
    justify-content: space-between;
    margin-top: 20px;
}

.mb-4 {
    margin-bottom: 16px;
}

/* 增大表单元素的字体大小 */
:deep(.el-form-item__label) {
    font-size: 16px;
    font-weight: 500;
    margin-bottom: 8px;
}

:deep(.el-input__inner) {
    font-size: 16px;
    height: 48px;
    line-height: 48px;
}

:deep(.el-input__prefix) {
    font-size: 18px;
}

:deep(.el-button) {
    font-size: 16px;
    height: 48px;
    padding: 12px 20px;
}

:deep(.el-alert__content) {
    font-size: 15px;
}
</style>
