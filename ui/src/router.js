import Vue from 'vue'
import Router from 'vue-router'
import Container from './components/Container.vue'

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
            redirect: {name: 'testconfList'},
            children: [
                {
                    name: 'testconfList',
                    path: 'testconf',
                    component: () => import('./views/testconf/TestConfList.vue'),
                },
                {
                    name: 'testconfDetail',
                    path: 'testconf/:id',
                    component: () => import('./views/testconf/TestConfDetail.vue'),
                },


                {
                    name: 'statdata',
                    path: 'statdata',
                    component: () => import('./views/statdata/StatData.vue'),
                },
            ],
        },
        {
            path: '*',
            component: () => import('@/views/Notfound.vue'),
        },
    ]
})
