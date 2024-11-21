import { createRouter, createWebHashHistory } from 'vue-router';

// 初始路由
const routes = [
  {
    path: '/',
    redirect: '/Home'
  },
  {
    path: '/Home',
    name: 'Home',
    component: () => import('@/views/Home/Home.vue')
  },
  {
    path: '/article',
    name: 'Article',
    component: () => import('@/views/Article/index.vue')
  },

  /* {
    path: '/404',
    component: () => import('@/views/404.vue')
  }, */
  /* {
    path: '/:pathMatch(.*)',
    redirect: '/404'
  } */
];

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: routes,
  // 刷新时，滚动条位置还原
  scrollBehavior: () => ({ left: 0, top: 0 }),
});

// 全局前置守卫
const whileList = ['/Home', '/article']; // 白名单

router.beforeEach((to, from, next) => {
  let hastoken = localStorage.getItem('token');
  if (hastoken) {
    next();
  } else {
    if (whileList.includes(to.path)) {
      next();
    } else {
      next(`/Home?redirect=${to.path}`);
    }
  }
});

export default router;