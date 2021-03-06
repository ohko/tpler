// 路由
module.exports = [
   { path: '/', component: _ => import('../pages/dashboard.vue') },
   { path: '/admin', component: _ => import('../pages/dashboard.vue') },
   { path: '/admin/dashboard', component: _ => import('../pages/dashboard.vue') },
   { path: '/admin/login', name: "login", component: _ => import('../pages/login.vue') },
   { path: '/admin/logout', component: _ => import('../pages/logout.vue') },
   { path: '/admin/form', component: _ => import('../pages/form.vue') },
   { path: '/admin/table', component: _ => import('../pages/table.vue') },
   { path: '/admin/password', component: _ => import('../pages/password.vue') },
   { path: '/admin/success', component: _ => import('../pages/success.vue') },
   { path: '/admin/error', component: _ => import('../pages/error.vue') },
   { path: '/admin/user/list', component: _ => import('../user/list.vue') },
   { path: '/admin/user/add', component: _ => import('../user/add.vue') },
   { path: '/admin/user/edit/:id', component: _ => import('../user/edit.vue') },
   { path: '/admin/setting/list', component: _ => import('../setting/list.vue') },
   { path: '/admin/setting/add', component: _ => import('../setting/add.vue') },
   { path: '/admin/setting/edit/:key', component: _ => import('../setting/edit.vue') },
   { path: '*', component: _ => import('../pages/error.vue') },
]