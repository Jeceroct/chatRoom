import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/address',
    name: 'address',
    component: () => import('../views/address.vue')
  },
  {
    path: '/signup',
    name: 'signup',
    component: () => import('../views/signup.vue')
  },
  {
    path: '/login',
    name:'login',
    component: () => import('../views/login.vue')
  },
  {
    path: '/room',
    name:'room',
    component: () => import('../views/room.vue')
  },
  {
    path: '/connectionError',
    name:'connectionError',
    component: () => import('../views/connectionError.vue')
  },
  {
    path: '/',
    name:'starting',
    component: () => import('../views/starting.vue')
  },
  {
    path: '/:pathMatch(.*)*',
    name:'404',
    component: () => import('../views/starting.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
