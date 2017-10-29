<template>
  <div id="app">
    <nav class="light-blue lighten-1" role="navigation">
      <div class="nav-wrapper container"><a id="logo-container" href="#" class="brand-logo">Team 07</a>
        <ul class="right hide-on-med-and-down">
          <li><a href="#" v-on:click="goToStuttgart">Ab ins Ländle</a></li>
          <li><a href="#" v-on:click="goToNewYork">New York, New York!</a></li>
        </ul>
        <ul id="nav-mobile" class="side-nav">
          <li><a href="#" v-on:click="goToStuttgart">Go to Stuttgart</a></li>
        </ul>
        <a href="#" data-activates="nav-mobile" class="button-collapse"><i class="material-icons">menu</i></a>
      </div>
    </nav>
    <div class="row">
      <div class="col s10">
        <div class="heatmap" id="map-canvas" style="height: 85vh; width: 100%">
        </div>
      </div>
      <div class="col s2">
        <search></search>
      </div>
    </diV>
    <footer class="align-center">
      <form action="#">
        <p class="range-field">
          <input type="range" id="test5" min="0" max="23" v-model="myTime" />
        </p>
      </form>
    </footer>
  </div>
</template>

<script>
import _ from 'lodash';
import L from 'leaflet';
import math from 'mathjs';
import Simpleheat from 'simpleheat';
import LeafletHeat from 'leaflet.heat';

import stuttgart from './stuttgart';

import Search from './Search.vue';

const zonesCenters = {
  0: [40.805164, -73.955591],
  1: [40.799734, -73.941266],
  2: [40.791838, -73.965153],
  3: [40.85368, -73.949495],
  4: [40.776857, -73.976529],
  5: [40.771009, -73.960687],
  6: [40.757183, -73.990178],
  7: [40.751992, -73.973931]
};

const mostOutsidePoints = [
  [40.816033,-73.962023], [40.804079,-73.932326],
  [40.751822,-74.006140],	[40.741549,-73.975585],
]

const SAMPLING_COLUMNS = 100;
const SAMPLING_ROWS= 100;
const API_URL = "https://webbackend-webbackend.training.altemista.cloud";

export default {
  components: {
    Search,
  },
  data() {
    return {
      myTime: 0,
      myData: [],
      heat: null,
      map: null,
      demandForZones: [],
      samplingPoints: [],
    };
  },
  watch: {
    myTime() {
      this.getDemand();
      stuttgart.addToMap(this.map, this.myTime);
    }
  },
  methods: {
    goToStuttgart() {
      this.map.panTo(new L.LatLng(48.79208, 9.23218));
    },
    goToNewYork() {
      this.map.panTo(new L.LatLng(40.771009, -73.960687));
    },
    getDemand() {
      window.fetch(API_URL + '/Demands?h=' + this.myTime, {  })
        .then((response, reject) => {
          return response.json();
      }).then((json) => {
        this.demandForZones = json._demandsByZone;
        this.demandForZones = this.demandForZones.map((demand) => {
          return demand / 100;
        })
        this.recreateData();
        this.resetLayer();
      });
    },
    createMeasurementPoints() {
      this.samplingPoints = [];
      let dataRow = [mostOutsidePoints[0]];
      let deltaWithinRow = _.zipWith(mostOutsidePoints[0], mostOutsidePoints[1], (a, b) => {
        return (a - b) / SAMPLING_COLUMNS;
      });
      let deltaWithinColumn = _.zipWith(mostOutsidePoints[0], mostOutsidePoints[2], (a, b) => {
        return (a - b) / SAMPLING_ROWS;
      });

      let subtract = (a, b) => {return a - b};

      for (let i = 1; i < SAMPLING_COLUMNS; i++) {
        dataRow[i] = _.zipWith(dataRow[i - 1], deltaWithinRow, subtract);
      }

      for (let i = 1; i < SAMPLING_ROWS; i++) {
        this.samplingPoints = _.concat(this.samplingPoints, dataRow);
        dataRow = dataRow.map((point) => {
          return  _.zipWith(point, deltaWithinColumn, subtract);
        });
      }

      // store the zone as third information
      this.samplingPoints.map((point) => {
        return point.push(this.pointToIdZone({lat: point[0], long: point[1]}));
      });

    },
    recreateData() {
      this.myData = [];
      this.samplingPoints.forEach((point) => {
        this.myData.push([point[0], point[1], this.demandForZones[point[2]]]);
      });
    },
    resetLayer() {
      if(this.heat) {
        this.map.removeLayer(this.heat);
      }
      this.heat = L.heatLayer(
        this.myData, { radius: 20 * 10000 / SAMPLING_ROWS / SAMPLING_COLUMNS, blur: 40 }).addTo(this.map);
    },
    pointToIdZone({lat, long}) {
      let zone = _.minBy(_.map(zonesCenters, (zone, id) => {
        // squared distance
        return {
          id,
          distance: Math.pow(zone[0] - lat, 2) + Math.pow(zone[1] - long, 2),
        };
      }), 'distance').id;
      return zone;
    },
  },
  mounted() {
    let baseLayer = L.tileLayer(
      'http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '...',
        maxZoom: 18,
      },
    );

    this.map = new L.Map('map-canvas', {
      center: new L.LatLng(40.771009, -73.960687),
      zoom: 14,
      layers: [baseLayer],
    });
    this.createMeasurementPoints();
    this.getDemand();

    stuttgart.addToMap(this.map, this.myTime);
  }
}
</script>

<style lang="scss">
.leaflet-heatmap-layer {
opacity: .6;
}
</style>
