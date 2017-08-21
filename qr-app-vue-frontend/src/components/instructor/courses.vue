<template>
  <div id="courses">
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
  <!-- Main table element -->
  <b-table striped hover show-empty
           :items="items"
           :fields="fields"
           :filter="filter"
           :sort-by.sync="sortBy"
           :sort-desc.sync="sortDesc"
           @filtered="onFiltered"
  >
    <template slot="name" scope="row">{{row.value}}</template>
    <template slot="date" scope="row">{{row.value}}</template>
		<template slot="action" scope="row">
      <b-btn size="sm" @click.stop="showAddParticipants(row)">Hinzufügen</b-btn>
    </template>
  </b-table>

  <p>
    Sort By: {{ sortBy || 'n/a' }}, Direction: {{ sortDesc ? 'descending' : 'ascending' }}
  </p>

  <!-- Details modal -->
  <b-modal id="modal1" @hide="resetModal" ok-only>
    <h4 class="my-1 py-1" slot="modal-header">Index: {{ modalDetails.index }}</h4>
    <pre>{{ modalDetails.data }}</pre>
  </b-modal>

</div>
<b-button size="lg" variant="success" v-on:click="addCourse()">Kurs hinzufügen</b-button>
  </div>
</template>

<script>
import router from '@/router/index'
import axios from 'axios';

//TODO: add nested list for course participants

export default {
  name: 'courses',
	data: () => {
		let items = JSON.parse(localStorage.getItem('courses'));
		for (let i = 0; i < items.length; i++) {
			if (items[i].participants === null) {
				items[i]["par_length"] = 0;
			} else {
				items[i]["par_length"] = items[i].participants.length;
			}
		}

		return {
	    items: items,
	    fields: {
	      name:     { label: 'Kurs Name', sortable: true },
	      date:      { label: 'Datum', sortable: true, 'class': 'text-center'  },
				par_length: { label: "Teilnehmer"},
				action: { label: "Hinzufügen"}
	    },
	    currentPage: 1,
	    perPage: 5,
	    totalRows: items.length,
	    pageOptions: [{text:5,value:5},{text:10,value:10},{text:15,value:15}],
	    sortBy: null,
	    sortDesc: false,
	    filter: null,
	    modalDetails: { index:'', data:'' }
		};
  },
  methods: {
		logout: function () {
			localStorage.removeItem('ins_api_key');
			localStorage.removeItem('courses');
			localStorage.removeItem('clicked_course');
			router.push({name: "login_selection"});
		},
		resetModal: function () {
			this.modalDetails.data = '';
			this.modalDetails.index = '';
		},
		onFiltered: function (filteredItems) {
			// Trigger pagination to update the number of buttons/pages due to filtering
			this.totalRows = filteredItems.length;
			this.currentPage = 1;
		},
		showAddParticipants: function (row) {
			localStorage.setItem('clicked_course', JSON.stringify(row.index));
			router.push({ name: 'course_register' });
		},
		addCourse: function () {
			router.push({ name: 'course_add'});
		}
  }
}
</script>

<style>
</style>
