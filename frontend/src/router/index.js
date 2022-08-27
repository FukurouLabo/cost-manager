import {createMemoryHistory, createRouter} from 'vue-router'
import IssueList from '../views/IssueList.vue'


const routes = [
  {
    path: '/',
    name: 'Top',
    component: IssueList
  }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

export default router
