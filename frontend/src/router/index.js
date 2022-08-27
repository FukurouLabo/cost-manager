import {createMemoryHistory, createRouter} from 'vue-router'
import Index from '../views/index.vue'
import Setup from '../views/setup.vue'


const routes = [
  {
    path: '/',
    name: 'Index',
    component: Index
  },
  {
    path: '/setup',
    name: 'Setup',
    component: Setup
  }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

export default router
