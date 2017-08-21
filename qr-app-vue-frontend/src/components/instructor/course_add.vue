
<template>
  <div id="course_add">
		<b-navbar toggleable type="light" variant="primary" toggle-breakpoint="md">
		  <b-nav-toggle target="nav_collapse"></b-nav-toggle>
		  <b-navbar-brand style="color: white;">Halle, Trainer</b-navbar-brand>
		  <b-collapse is-nav id="nav_collapse">
		    <b-nav is-nav-bar class="ml-auto">
					<b-button v-on:click="logout()"  variant="danger">Logout</b-button>
		    </b-nav>
		  </b-collapse>
		</b-navbar>
		<div class="container">
    <div class="row">
      <div class="col-md-12">
				<h5>Kursname:</h5>
					<b-form-input v-model="kursname"
		                  type="text"
		                  placeholder="Kursname eingeben"
											id="kname"
		    ></b-form-input>
				<h5>Datum</h5>
					<b-form-input v-model="date"
		                  type="text"
		                  placeholder="Beispiel: 2017-08-01"
											id="date"
		    ></b-form-input>
				<h5>Uhrzeit</h5>
					<b-form-input v-model="time"
		                  type="text"
		                  placeholder="Beispiel 20:15:00"
											id="time"
		    ></b-form-input>
      </div>
		</div>
		</div>
	</br>
		<b-button variant="warning" v-on:click="goBack()">Zurück</b-button>
		<b-button variant="success" v-on:click="addCourse()">Hinzufügen</b-button>
	</div>
</template>

<script>
import router from '@/router/index'
import axios from 'axios'

export default {
  name: 'course_add',
	data: () => {
		return {
			kursname: '',
			date: '',
			time: ''
		};
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
		addCourse: function () {
			let key = localStorage.getItem('ins_api_key')
			if (this.kursname && this.date && this.time) {
				axios.post(process.env.API_URL + '/courses/add', {
					"apikey": key,
					"name": this.kursname,
					"date": this.date + " " + this.time
				}, {validateStatus: function (status) {
					return status < 500;
				}})
				.then( (response) => {
					let resStatus = response.data.status;
					let resMsg = response.data.message;
					if (resStatus !== "success") {
						this.kursname = "";
						this.date = "";
						this.time = "";
						this.$notify("Hinzufügen fehlgeschlagen", 'error');
					} else {
						let apikey = localStorage.getItem('ins_api_key');
						axios.post(process.env.API_URL + '/courses', {
							"apikey": apikey
						}, {validateStatus: function (status) {
							return true;
						}})
						.then((response) => {
							let resStatus = response.data.status;
							let resData = response.data.data;
							if (resStatus !== "success") {
								// ERROR STATE
							} else {
								localStorage.setItem("courses", JSON.stringify(resData));
								this.$notify("Erfolgreich hinzugefügt.", "success");
								router.push({name: "courses"})
							}
						})
						.catch((err) => {
							console.error(err);
						});
					}
				})
				.catch( (err) => {
					console.error(err);
				})
			}
		}
  },
	components: {
	}
}
</script>

<style>

</style>
