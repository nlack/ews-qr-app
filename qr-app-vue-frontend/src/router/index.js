import Vue from 'vue'
import Router from 'vue-router'
import Login_Selection from '@/components/login/login_selection'
import Login_Participant from '@/components/login/login_participant'
import Login_Instructor from '@/components/login/login_instructor'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'login_selection',
      component: Login_Selection
    },
    {
      path: '/participant',
      name: 'login_participant',
      component: Login_Participant
    },
    {
      path: '/instructor',
      name: 'login_instructor',
      component: Login_Instructor
    }
  ]
})
