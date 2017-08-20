<template>
  <div id="login_instructor">
    <div class="title" >Trainer Login:</div><br>
    <input id="ins_user" v-model="ins_user" placeholder="Benutzer" class="text-input" type="text" name="" value="" autofocus><br>
    <input id="ins_password" v-model="ins_pass" placeholder="Passwort" class="password-input" type="password" name="" value=""><br>
	</br>
		<b-button variant="primary" style="width: 120px;" v-on:click="goBack()" >Zurück</b-button>
		<b-button variant="warning" style="width: 120px;" v-on:click="tryLogin()" >Login</b-button>
		<p v-if="error_ins_msg_01" >Bitte Benutzer / Passwort eingeben.</p>
  </div>
</template>

<script>
import router from '@/router/index'
import axios from 'axios';

export default {
  name: 'login_instructor',
	data: () => {
		return {
			ins_user: '',
			ins_pass: '',
			error_ins_msg_01: false,
		};
	},
  methods: {
    goBack: function() {
      router.push({name: "login_selection"})
    },
		tryLogin: function() {
			// Extract login credentials
			let user = this.ins_user;
			let password = this.ins_pass;

			// Clear input fields
			document.querySelector('#ins_user').value = "";
			document.querySelector('#ins_password').value = "";

			// Validate input
			if (user && password) {
				axios.post(process.env.API_URL + '/instructor', {
						"name": user,
						"password": password
				})
				.then((response) => {
					let resStatus = response.data.status;
					let resData = response.data.data;
					if (resStatus !== "success") {
						// ERROR STATE
						console.log("LOGIN FAILED");
						alert("WRONG LOGIN");
					} else {
						localStorage.setItem("ins_api_key", resData.apikey);
						axios.post(process.env.API_URL + '/courses', {
							"apikey": resData.apikey
						})
						.then((response) => {
							let resStatus = response.data.status;
							let resData = response.data.data;
							if (resStatus !== "success") {
								// ERROR STATE
								consoe.log("GET COURSE LIST FAILED");
								alert("API PERMISSION PROBLEM");
							} else {
								localStorage.setItem("courses", JSON.stringify(resData.courses));
      					router.push({name: "courses"})
							}
						})
						.catch((err) => {
							console.error(err);
						});
					}
				})
				.catch((err) => {
					console.error(err);
				});
			} else {
				this.error_ins_msg_01 = true;
			}
		}
  }
}
</script>

<style>
.title {
  font-size: 24pt;
}

#login_instructor {
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
