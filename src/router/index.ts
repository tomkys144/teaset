import About from '@/views/About.vue'
import Home from '@/views/Home.vue'
import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    component: About,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((toRoute, fromRoute, next) => {
  let title: string = toRoute.name ? toRoute.name.toString() + ' | ' : ''
  title += 'Teaset'
  window.document.title = title

  next()
})

export default router
