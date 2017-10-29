<template lang="html">
  <div>
    <h3>Mobility agent</h3>
    I want to come...
    <div class="input-field row">
      <i class="material-icons prefix">flight_takeoff</i>
      <input type="text" id="to" class="autocomplete">
      <label for="from">From</label>
    </div>
    <div class="input-field row">
      <i class="material-icons prefix">flight_land</i>
      <input type="text" id="to" class="autocomplete">
      <label for="to">To</label>
    </div>
    <button class="btn btn-large" v-on:click="proposalsVisible = true">
      <i class="material-icons left">search</i>
      Search
    </button>
    <div class="height: 100%; overflow-y: auto" v-if="proposalsVisible">
      <div class="card blue-grey darken-1">
        <div class="card-content white-text">
          <span class="card-title vertical-align">
            <img src="https://www.car2go.com/media/assets/patterns/static/img/car2go-logo.png" style="height: 60px; margin-top: -15px; margin-right: 15px" class="left">
            Car2Go
          </span>
          <br />
          <p>Current demand: HIGH</p>
        </div>
        <div class="card-action">
          <span class="white-text">approx. 5,85 €</span>
          <a href="#">Book now</a>
        </div>
      </div>
      <div class="card blue-grey darken-1">
        <div class="card-content white-text">
          <span class="card-title vertical-align">
            <img src="https://www.bahn.com/common/view/static/v8/img/db_em_rgb_100px.svg" style="height: 60px; margin-top: -15px; margin-right: 15px" class="left">
            DB
          </span>
          <br />
          <p>Estimated delay: 1min</p>
        </div>
        <div class="card-action">
          <span class="white-text">approx. 2,85 €</span>
          <a href="#">Book now</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="text/babel">
  export default {
    name: 'search',
    mounted() {
      $(function () {
        $('input.autocomplete').autocomplete({
          data: {
            "Hauptbahnhof": null,
            "Königsstraße": null,
            "Römerkastell": null,
            "Hackathon Stuttgart": null,
            "Bahnhofsstraße": null,
          },
          limit: 20, // The max amount of results that can be shown at once. Default: Infinity.
          onAutocomplete: function(val) {
            // Callback function when value is autcompleted.
          },
          minLength: 1, // The minimum length of the input for the autocomplete to start. Default: 1.
        });
      });
    },
    data() {
      return {
        proposalsVisible: false,
      };
    },
  }
</script>
