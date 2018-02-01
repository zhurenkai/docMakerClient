import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import Authlogin from '@/components/auth/Login'
// import Workbench from '@/components/browser/Workbench'

Vue.use(Router)


const routes = [
    {
      path: '/',
      name: 'Hello',
      component: Hello,
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/login',
      name: 'Login',
      component: Authlogin
    },
    {
      path: '/settings',
      name: 'Settings',
      components: {
        default: require('../components/settings/Settings'),
          left_menu:require('../components/layouts/LeftMenu'),
          top_menu:require('../components/layouts/TopMenu'),
      },
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/workbench',
      name: 'Workbench',
      components: {
        default:require('../components/browser/Tabs'),
          left_menu:require('../components/layouts/LeftMenu'),
          top_menu:require('../components/layouts/TopMenu'),
    },
      meta: {
        requireAuth: true
      }
    }
  ];

const router = new Router({
  routes
});

router.beforeEach((to, from, next) => {

  if (to.matched.some(r => r.meta.requireAuth)) {
  let time_now =Math.round(new  Date().getTime()/1000)
  let expire_time = localStorage.getItem('expire_time')
  console.log(expire_time,time_now,123)

  if (expire_time && expire_time>time_now) {
    next();
  }
  else {
    next({
      path: '/login',
      query: {redirect: to.fullPath}
    })
  }
}
else {
  next();
}
})

export default router;
