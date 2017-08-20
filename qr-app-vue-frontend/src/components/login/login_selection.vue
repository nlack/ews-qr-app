
<template>
  <div id="login_selection">
	</br>
		<b-button variant="primary" style="width: 200px;" v-on:click="loadParticipantLogin" >Teilnehmer</b-button></br></br>
		<b-button variant="warning" style="width: 200px;" v-on:click="loadInstructorLogin" >Trainer</b-button>
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
</style>
