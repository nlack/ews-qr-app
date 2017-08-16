<template>
  <div id="course_register">
		<b-navbar toggleable type="light" variant="primary" toggle-breakpoint="md">
		  <b-nav-toggle target="nav_collapse"></b-nav-toggle>
		  <b-navbar-brand style="color: white;">Halle, Trainer</b-navbar-brand>
		  <b-collapse is-nav id="nav_collapse">
		    <b-nav is-nav-bar class="ml-auto">
					<b-button v-on:click="logout()"  variant="danger">Logout</b-button>
		    </b-nav>
		  </b-collapse>
		</b-navbar>
	</div>
</template>

<script>
import router from '@/router/index'
import axios from 'axios';


export default {
  name: 'courses',
  methods: {
		logout: function () {
			localStorage.removeItem('ins_api_key');
			localStorage.removeItem('courses');
			router.push({name: "login_selection"});
		}
  }
}
</script>

<style>
</style>
