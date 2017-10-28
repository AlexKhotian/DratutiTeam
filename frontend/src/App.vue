<template>
  <div id="app">
    <nav class="light-blue lighten-1" role="navigation">
      <div class="nav-wrapper container"><a id="logo-container" href="#" class="brand-logo">Logo</a>
        <ul class="right hide-on-med-and-down">
          <li><a href="#">Navbar Link</a></li>
        </ul>
        <ul id="nav-mobile" class="side-nav">
          <li><a href="#">Navbar Link</a></li>
        </ul>
        <a href="#" data-activates="nav-mobile" class="button-collapse"><i class="material-icons">menu</i></a>
      </div>
    </nav>
    <div class="container">
      <div class="row">
        <div class="col s9">
          <div class="heatmap" id="map-canvas" style="height: 80vh; width: 100%">
          </div>
        </div>
        <div class="col s3">
          <h3>Searching</h3>

        </div>
      </diV>
    </div>
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

const SAMPLING_COLUMNS = 30;
const SAMPLING_ROWS= 50;

export default {
  data() {
    return {
      myTime: 0,
      myData: [],
      heat: null,
      map: null,
      demandForZones: [0, 0.1, 0.4, 0.2, 0.8, 0.1, 1, 0],
      samplingPoints: [],
    };
  },
  watch: {
    myTime() {
      const fetch = () => window.fetch('/Demands?h=' + this.myTime, fetchDefaults())
        .then((response) => {
          response.json();
          console.log(response.json());
          this.recreateData();
          this.resetLayer();
      });
    }
  },
  methods: {
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
        this.myData, { radius: 40, blur: 50 }).addTo(this.map);
    },
    pointToIdZone({lat, long}) {
      return _.minBy(_.map(zonesCenters, (zone, id) => {
        // squared distance
        return {
          id,
          distance: Math.pow(zone[0] - lat, 2) + Math.pow(zone[1] - long, 2),
        };
      }), 'distance').id;
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
    this.recreateData();
    this.resetLayer();
  }
}
</script>

<style lang="scss">

</style>
