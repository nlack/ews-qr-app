
<template>
  <div id="login_selection">
  <button class="participant" v-on:click="loadParticipantLogin" >Teilnehmer</button> </br>
  <button class="instructor" v-on:click="loadInstructorLogin" >Trainer</button>
  </div>
</template>

<script>

import router from '@/router/index'
export default {
  name: 'login_selection',
  methods: {
    loadParticipantLogin: function() {
      router.push({name: "login_participant"})
    },
    loadInstructorLogin: function() {
      router.push({name: "login_instructor"})
    }
  },
  components: {
    'login_participant': {
      template: 'login_participant'
    },
    'login_instructor': {
      template: 'login_instructor'
    }
  }
}
</script>

<style>
#login_selection {
  margin: auto;
}

.participant {
  margin-bottom: 10px;
  background-color: #4CAF50;
  border: none;
  color: white;
  padding: 15px 15px;
  width: 200px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 24px;
  text-transform: uppercase;
}

.participant:hover {
  color: black;
  transition-duration: 0.5s;
}

.instructor {
  margin-top: 10px;
  background-color: #008CBA;
  border: none;
  color: white;
  padding: 15px 15px;
  width: 200px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 24px;
  text-transform: uppercase;
}

.instructor:hover {
  color: black;
  transition-duration: 0.4s;
}
</style>
