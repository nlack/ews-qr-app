import Vue from 'vue'
import Router from 'vue-router'
import Login_Selection from '@/components/login/login_selection'
import Login_Participant from '@/components/login/login_participant'
import Login_Instructor from '@/components/login/login_instructor'
import Show_Participant from '@/components/participant/show_participant'
import Course_List from '@/components/instructor/courses'
import Course_Register from '@/components/instructor/course_register'

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
    },
		{
			path: '/courses',
			name: 'courses',
			component: Course_List,
			beforeEnter: (to, from, next) => {
				let check = localStorage.getItem('ins_api_key');
				if (!check) {
					next({name: 'login_instructor'});
				} else {
					next();
				}
			}
		},
		{
			path: '/course_register',
			name: 'course_register',
			component: Course_Register,
			beforeEnter: (to, from, next) => {
				let check = localStorage.getItem('ins_api_key');
				if (!check) {
					next({name: 'login_instructor'});
				} else {
					next();
				}
			}
		},
		{
			path: '/participant_show',
			name: 'show_participant',
			component: Show_Participant,
			beforeEnter: (to, from, next) => {
				let check = localStorage.getItem('participant_qrcode');
				if (!check) {
					next({name: 'login_participant'});
				} else {
					next();
				}
			}
		}
  ]
})
