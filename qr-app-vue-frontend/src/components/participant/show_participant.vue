<template>
  <div id="show_participant">
		<b-navbar toggleable type="light" variant="primary" toggle-breakpoint="md">
		  <b-nav-toggle target="nav_collapse"></b-nav-toggle>
		  <b-navbar-brand style="color: white;">Hallo, {{firstname}} {{lastname}}</b-navbar-brand>
		  <b-collapse is-nav id="nav_collapse">
		    <b-nav is-nav-bar class="ml-auto">
					<b-button v-on:click="logout()"  variant="danger">Logout</b-button>
		    </b-nav>
		  </b-collapse>
		</b-navbar>

		<img class="img-fluid" alt="Responsive image" v-bind:src="qrcode" />
  </div>

</template>

<script>
import router from '@/router/index'
export default {
  name: 'show_participant',
  methods: {
		logout: function () {
			localStorage.removeItem('participant_lastname');
			localStorage.removeItem('participant_firstname');
			localStorage.removeItem('participant_qrcode');
			router.push({name: "login_selection"});
		}
  },
	data: () => {
		return {
			qrcode: localStorage.getItem('participant_qrcode'),
			lastname: localStorage.getItem('participant_lastname'),
			firstname: localStorage.getItem('participant_firstname')
		}
	}
}
</script>

<style>
</style>
