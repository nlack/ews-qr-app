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
  <div class="row my-1">
    <div class="col-sm-8">
      <b-pagination :total-rows="totalRows" :per-page="perPage" v-model="currentPage" />
    </div>
  </div>
  <div class="my-1 row">
    <div class="col-md-6">
      <b-form-fieldset horizontal label="Rows per page" :label-cols="6">
        <b-form-select :options="pageOptions" v-model="perPage" />
      </b-form-fieldset>
    </div>
  </div>

  <!-- Main table element -->
  <b-table striped hover show-empty
           :items="items"
           :fields="fields"
           :current-page="currentPage"
           :per-page="perPage"
           :filter="filter"
           :sort-by.sync="sortBy"
           :sort-desc.sync="sortDesc"
           @filtered="onFiltered"
  >
    <template slot="name" scope="row">{{row.value}}</template>
    <template slot="date" scope="row">{{row.value}}</template>
		<template slot="action" scope="row">
      <b-btn size="sm" @click.stop="showAddParticipants()">Hinzufügen</b-btn>
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


  </div>
</template>

<script>
import router from '@/router/index'
import axios from 'axios';


export default {
  name: 'courses',
		data: () => {
		let items = JSON.parse(localStorage.getItem('courses'));
		for (let i = 0; i < items.length; i++) {
			items[i]["par_length"] = items[i].participants.length;
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
			router.push({name: "login_selection"});
		},
		resetModal() {
			this.modalDetails.data = '';
			this.modalDetails.index = '';
		},
		onFiltered(filteredItems) {
			// Trigger pagination to update the number of buttons/pages due to filtering
			this.totalRows = filteredItems.length;
			this.currentPage = 1;
		},
		showAddParticipants() {

		}
  }
}
</script>

<style>
</style>
