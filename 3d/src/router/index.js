import { createRouter, createWebHistory, createWebHashHistory  } from 'vue-router'

// const router = createRouter({
//   history: createWebHistory(import.meta.env.BASE_URL),
//   routes: [
//     {
//       path: '/',
//       name: 'home',
//       component: HomeView
//     },
//     {
//       path: '/about',
//       name: 'about',
//       // route level code-splitting
//       // this generates a separate chunk (About.[hash].js) for this route
//       // which is lazy-loaded when the route is visited.
//       component: () => import('../views/AboutView.vue')
//     }
//   ]
// })

//初始路由
const routes = [
  {
    path: '/Home',
    name: 'Home',
    component: () => import('../views/Home.vue')
  },
  /* {
    path: '/404',
    component: () => import('../views/404.vue')
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
// // 全局前置守卫
const whileList = ['/Home',]   //白名单

router.beforeEach((to, from, next) => {
  let hastoken = localStorage.getItem('token')
  if (hastoken) {
    // const menu = [
    //   {
    //     name: 'HomeView',
    //     url: '/HomeView',
    //     children: [
    //       {
    //         name: 'Product 1',
    //         url: '/news/product1',
    //       },
    //       {
    //         name: 'Product 2',
    //         url: '/news/product2',
    //         children: [
    //           {
    //             name: 'Sub Product 1',
    //             url: '/news/product2/subproduct1',
    //           },
    //         ],
    //       },
    //     ],
    //   },
    //   {
    //     name: 'AboutView',
    //     url: '/AboutView',
    //   }]
    //   const addroutes = generateRoutes(menu);
    //   console.log('addroutes', addroutes)
    //   addroutes.forEach((route) => {
    //     router.addRoute(route);
    //   });
      next();
     //console.log('获取所有路由', router.getRoutes())
  } else {
    if (whileList.includes(to.path)) {
      next();
    } else {
      next(`/Home?redirect=${to.path}`);
    }
  }
  //白名单 有值 或者登陆过存储了token信息可以跳转 否则就去登录页面
  // if (whileList.includes(to.path) || token) {
  //     const  menu  = [{url:'/news',name:'news'},{url:'/AboutView',name:'AboutView'}]
  //     const addroutes = generateRoutes(menu);
  //     addroutes.forEach((route) => {
  //       router.addRoute(route);
  //     });
  //     console.log('获取所有路由',router.getRoutes())
  //     next();//首页是静态不需要重新定向

  // } else {
  //   showToast('没有身份信息') 
  //     next({ path:'/'})
  //     // next({...to, replace: true});
  // }
})

// 根据菜单数据生成路由配置
// function generateRoutes(menuData) {
//   // let datas = {};
//  return   menuData.map(menu => {
//     let modules = import.meta.glob("../views/**/*.vue");
//     console.log('modules',modules[`../views${ menu.url}.vue`])
//     let datas = {
//       path: menu.url,
//       name: menu.name,
//       component: modules[`../views${ menu.url}.vue`]
//     }
//     // //如果存在子集进行递归解析
//     if (menu.children && menu.children.length) { 
//       datas.children = generateRoutes(menu.children);
//     }
//     return datas
//   });
// }

// function generateRoutes(menuData) {
//   let modules = import.meta.glob("../views/**/*.vue");
//   return menuData.map(item => {
//     if (item.children && item.children.length > 0) {
//       return {
//         path: item.url,
//         name: item.name,
//         children: generateRoutes(item.children), // 递归处理子菜单
//       };
//     } else {
//       return {
//         path: item.url,
//         name: item.name,
//         component: modules[`../views${item.url}.vue`]
//       };
//     }
//   });
// }


export default router
