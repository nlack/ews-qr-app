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
		<div style="height:80%;">
		<qrcode-vue id="qrcode" :value="qrcode" :size="windowSize" level="H" ></qrcode-vue>
	</div>
  </div>

</template>

<script>
import router from '@/router/index'
import QrcodeVue from 'qrcode.vue';




export default {
  name: 'show_participant',
	mounted: function() {
		console.log("DOM READY");
		document.querySelector('#qrcode').style = "border:50px solid #FFFFFF;";
	},
  methods: {
		logout: function () {
			localStorage.removeItem('participant_lastname');
			localStorage.removeItem('participant_firstname');
			localStorage.removeItem('participant_qrcode');
			router.push({name: "login_selection"});
		}
  },
	data: () => {
		let winH = window.innerHeight;
		let winW = window.innerWidth;
		let winSize = 0;
		if (winH < winW) {
			 winSize = (winH - 50 / 100 * 85) - 150;
		} else {
			winSize = (winW - 50 / 100 * 85) - 150;
		}

		return {
			qrcode: localStorage.getItem('participant_qrcode'),
			lastname: localStorage.getItem('participant_lastname'),
			firstname: localStorage.getItem('participant_firstname'),
			windowSize: winSize
		};
	},
	components: {
		QrcodeVue
	}
}
</script>

<style>
</style>
