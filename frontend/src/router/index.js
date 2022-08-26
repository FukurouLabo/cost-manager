import {createMemoryHistory, createRouter} from 'vue-router'
import Home from '../views/Home.vue'
import About from '../views/About.vue'
import Test from '../views/TicketList.vue'


const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    // You can only use pre-loading to add routes, not the on-demand loading method.
    component: About
  },
  {
    path: '/test',
    name: 'Test',
    component: Test
  }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

export default router
