<template>
  <div id="login_participant">
    <div class="title" >Mitglieder Login:</div><br>
    <input id="par_user" v-model="par_user" placeholder="Benutzer" class="text-input" type="text" name="" value="" autofocus><br>
    <input id="par_pass" v-model="par_pass" placeholder="Passwort" class="password-input" type="password" name="" value=""><br>
	</br>
		<b-button variant="primary" style="width: 120px;" v-on:click="goBack()" >Zurück</b-button>
		<b-button variant="warning" style="width: 120px;" v-on:click="tryLogin()" >Login</b-button>
		<p v-if="error_par_msg_01" >Bitte Benutzer / Passwort eingeben.</p>
  </div>
</template>

<script>
import router from '@/router/index'
import axios from 'axios';

export default {
  name: 'login_participant',
	data: () => {
		return {
			par_user: '',
			par_pass: '',
			error_par_msg_01: false,
		};
	},
  methods: {
    goBack: function() {
      router.push({name: "login_selection"})
    },
		tryLogin: function() {
			// Extract login credentials
			let user = this.par_user;
			let password = this.par_pass;

			// Clear input fields
			document.querySelector('#par_user').value = "";
			document.querySelector('#par_pass').value = "";

			// Validate input
			if (user && password) {
				axios.post( process.env.API_URL + '/participant', {
					"name": user,
					"password": password
				}, {validateStatus: function (status) {
					return true;
				}})
				.then((response) => {
					let resStatus = response.data.status;
					let resData = response.data.data;
					if (resStatus !== "success") {
						// ERROR STATE
						this.$notify("Login fehlgeschlagen.", "warning");
						console.log("LOGIN FAILED");
					} else {
						// SUCCESS STATE
						localStorage.setItem('participant_lastname', resData.lastname)
						localStorage.setItem('participant_firstname', resData.firstname)
						localStorage.setItem('participant_qrcode', resData.qrhash)
						// GOTO ROUTE
      			router.push({name: "show_participant"})
						this.$notify("Login erfolgreich.", "info");
					}
				})
				.catch((err) => {
					console.error(err);
				});
			} else {
				this.error_par_msg_01 = true;
			}
		}
  }
}
</script>

<style>
.title {
  font-size: 24pt;
}

#login_participant {
  margin: auto;
}

.text-input {
  text-align: center;
  font-size: 18pt;
  background: transparent;
  border: 1;
  border-color: black;
  border-radius: 1;
  box-shadow: none;
  color: black;
  height: 50px;
  -webkit-transition: none;
  transition: none;
  width: 250px;
  margin-bottom: 5px;
}

.password-input {
  text-align: center;
  font-size: 18pt;
  background: transparent;
  border: 1;
  border-color: black;
  border-radius: 10;
  box-shadow: none;
  color: black;
  height: 50px;
  -webkit-transition: none;
  transition: none;
  width: 250px;
  margin-top: 5px;
}

.backButton {
  margin-bottom: 10px;
  background-color: #4CAF50;
  border: none;
  color: white;
  width: 125px;
  height: 55px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 18px;
  text-transform: uppercase;
}

.backButton:hover {
  color: black;
  transition-duration: 0.5s;
}

.loginButton {
  margin-top: 10px;
  background-color: #008CBA;
  border: none;
  color: white;
  width: 125px;
  height: 55px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 18px;
  text-transform: uppercase;
}

.loginButton:hover {
  color: black;
  transition-duration: 0.4s;
}
</style>
