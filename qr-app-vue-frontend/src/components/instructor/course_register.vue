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
		<div>
			<video width="640" height="480" id="scanner"></video></br>
			<b-button v-bind:disabled="scanActive" variant="primary" v-on:click="runScanner(); scanActive=true">Starte Scanner</b-button>
		</div>

	</div>
</template>

<script>
import router from '@/router/index'
import axios from 'axios'
const Instascan = require('instascan-ngfar')


export default {
  name: 'courses',
	data: () => {
		return {
			scanActive: false
		}
	},
  methods: {
		logout: function () {
			localStorage.removeItem('ins_api_key');
			localStorage.removeItem('courses');
			router.push({name: "login_selection"});
		},
		runScanner: function () {
			let scanner = new Instascan.Scanner({
				video: document.getElementById('scanner'),
				backgroundScan: false
			});
			scanner.addListener('scan', (content) => {
				this.$notify(content, 'success');
				console.log("Found QR-Code Content: " + content);
			});
			Instascan.Camera.getCameras().then(function (cameras) {
			  if (cameras.length > 0) {
			    scanner.start(cameras[0]);
			  } else {
			    console.error('No cameras found.');
			  }
			}).catch(function (e) {
			  console.error(e);
			});
		}
  }
}
</script>

<style>
#scanner {
	border-style: dotted;
}
</style>
