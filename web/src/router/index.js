import { createRouter, createWebHistory } from 'vue-router'

// 初始路由
const routes = [
  {
    path: '/',
    redirect: '/Home'
  },
  {
    path: '/Home',
    name: 'Home',
    component: () => import('@/views/Home/Home.vue'),
  },
  {
    path: '/article',
    name: 'Article',
    component: () => import('@/views/Article/index.vue'),
  }
  // ...其他路由
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
  scrollBehavior: () => ({ left: 0, top: 0 })
})

// 全局前置守卫
const whileList = ['/Home', '/article']

router.beforeEach((to, from, next) => {
  let hastoken = localStorage.getItem('token')
  if (hastoken) {
    next()
  } else {
    if (whileList.includes(to.path)) {
      next()
    } else {
      next(`/Home?redirect=${to.path}`)
    }
  }
})

export { router }
