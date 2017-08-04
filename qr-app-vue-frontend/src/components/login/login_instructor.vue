<template>
  <div id="login_instructor">
    <div class="title" >Trainer Login:</div><br>
    <input placeholder="Benutzer" class="text-input" type="text" name="" value="" autofocus><br>
    <input placeholder="Passwort" class="password-input" type="password" name="" value=""><br>
    <button class="backButton" type="button" v-on:click="goBack()" >Zurück</button>
    <button class="loginButton" type="button" v-on:click="" >Login</button>
  </div>
</template>

<script>
import router from '@/router/index'
export default {
  name: 'login_instructor',
  methods: {
    goBack: function() {
      router.push({name: "login_selection"})
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
