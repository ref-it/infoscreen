import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'start',
      component: () => import('../views/StartPage.vue')
    },
    {
      path: '/ef/',
      name: 'ef-start',
      component: () => import('../views/erfurt/StartPage.vue')
    },
    {
      path: '/ef/canteens',
      name: 'ef-canteens',
      component: () => import('../views/erfurt/CanteensView.vue')
    },
    {
      path: '/il/',
      name: 'il-start',
      component: () => import('../views/ilmenau/StartPage.vue')
    },
    {
      path: '/il/helmholtz',
      name: 'il-helmholtz',
      component: () => import('../views/ilmenau/HelmholtzView.vue')
    },
    {
      path: '/il/canteens',
      name: 'il-canteens',
      component: () => import('../views/ilmenau/CanteensView.vue')
    }
  ]
})

export default router
