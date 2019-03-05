import Vue from 'vue'
import Router from 'vue-router'
import Login from './views/Login.vue'
import Container from './components/Container.vue'
import storage from './utils/storage'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      // beforeEnter: (to, from, next) => {
      //   if (!storage.getItem(storage.KEYS.tokenName)) {
      //     next({
      //       name: 'Login',
      //     })
      //     return
      //   }
      //   next()
      // },
      path: '/',
      component: Container,
      redirect: {name: 'serviceList'},
      children: [
        {
          name: 'serviceList',
          path: 'service',
          component: () => import('./views/service/ServiceList.vue'),
        },
        {
          name: 'serviceDetail',
          path: 'service/:name',
          component: () => import('./views/service/ServiceDetail.vue'),
        },
        {
          name: 'authList',
          path: 'auths',
          component: () => import('./views/auths/AuthList.vue'),
        },
        {
          name: 'authDetail',
          path: 'auths/:name',
          component: () => import('./views/auths/AuthDetail.vue'),
        },

        {
          name: 'health',
          path: 'health',
          component: () => import('./views/healthcheck/HealthCheck.vue'),
        },
      ],
    },
    // {
    //   path: '/login',
    //   name: 'Login',
    //   component: Login
    // },
    // {
    //   path: '/logout',
    //   name: 'logout',
    //   component: Login,
    //   beforeEnter(to, from, next) {
    //     storage.clear()
    //     next()
    //   },
    // },
    {
      path: '*',
      component: () => import('@/views/Notfound.vue'),
    },
  ]
})
