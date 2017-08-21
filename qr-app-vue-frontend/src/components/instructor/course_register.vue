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
			<video height="40%" id="scanner"></video></br>
			<b-button variant="warning" v-on:click="goBack()">Zurück</b-button>
			<b-button v-bind:disabled="scanActive" variant="primary" v-on:click="runScanner(); scanActive=true">Starte Scanner</b-button>
		</div>

  <!-- Main table element -->
  <b-table striped hover show-empty
           :items="participants"
           :fields="fields"
  >
    <template slot="firstname" scope="row">{{row.value}}</template>
    <template slot="lastname" scope="row">{{row.value}}</template>
  </b-table>

  <!-- Details modal -->
  <b-modal id="modal1" @hide="resetModal" ok-only>
    <h4 class="my-1 py-1" slot="modal-header">Index: {{ modalDetails.index }}</h4>
    <pre>{{ modalDetails.data }}</pre>
  </b-modal>

	</div>
</template>

<script>
import router from '@/router/index'
import axios from 'axios'
const Instascan = require('instascan-ngfar')

//TODO: add register ajax call on successfully scanned qr-code
//TODO: list allready registered users, update on every register?

export default {
  name: 'course_register',
	data: () => {

		let index = localStorage.getItem('clicked_course');
		let courses =  JSON.parse(
			localStorage.getItem('courses')
		);

		let participants = courses[index].participants;

		return {
			scanActive: false,
	    participants: participants,
	    fields: {
	      "Firstname":     { label: 'Vorname', sortable: true },
	      "Lastname":      { label: 'Nachname', sortable: true}
	    },
	    modalDetails: { index:'', data:'' }
		}
	},
  methods: {
		logout: function () {
			localStorage.removeItem('ins_api_key');
			localStorage.removeItem('courses');
			router.push({name: "login_selection"});
		},
		goBack: function () {
			router.push({name: "courses"});
		},
		resetModal() {
			this.modalDetails.data = '';
			this.modalDetails.index = '';
		},
		runScanner: function () {
			let scanner = new Instascan.Scanner({
				video: document.getElementById('scanner'),
				backgroundScan: false
			});
			scanner.addListener('scan', (content) => {
				let key = localStorage.getItem('ins_api_key');
				let cId = Number(localStorage.getItem('clicked_course')) + 1;
				axios.put( process.env.API_URL + "/course/" + cId, {
					"apikey": key,
					"qrhash": content
				}, {validateStatus: function (status) {
					return true;
				}})
				.then( (response) => {
					let resStatus = response.data.status;
					let resMsg = response.data.message;

					if (resStatus !== "success") {
						if (resMsg === "Participant already added") {
							this.$notify("Teilnehmer schon hinzugefügt.", "warning");
						} else if (resMsg === "Participant has not payed"){
							this.$notify("Teilnehmer nicht zugelassen.", "error");
						} else {
							this.$notify(resMsg, "error");
						}
					} else {
						this.$notify("Teilnehmer erfolgreich hinzugefügt.", "info");
						axios.post(process.env.API_URL + '/courses', {
							"apikey": key
						}, {validateStatus: function (status) {
							return true;
						}})
						.then((response) => {
							let resStatus = response.data.status;
							let resData = "";
							if (response.data !== null) {
								resData = response.data.data;
							}
							if (resStatus !== "success") {
							} else {
								let course_index = Number(localStorage.getItem('clicked_course'));
								this.participants = resData[course_index].participants;
								localStorage.setItem("courses", JSON.stringify(resData));
							}
						})
						.catch((err) => {
							console.error(err);
						});
					}
				})
				.catch( (err) => {
					console.error(err);
				});
			});
			Instascan.Camera.getCameras().then(function (cameras) {
			  if (cameras.length > 0) {
			    scanner.start(cameras[0]);
			  } else {
			    console.error('No cameras found.');
					router.push({name: "courses"});
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
