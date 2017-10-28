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
          <input type="range" id="test5" min="-100" max="100" v-model="myTime" />
        </p>
      </form>
    </footer>
  </div>
</template>

<script>
import _ from 'lodash';
import L from 'leaflet';
import Simpleheat from 'simpleheat';
import LeafletHeat from 'leaflet.heat';

import Search from './Search.vue';

const zonesCenters = [
  [40.805164, -73.955591], [40.799734, -73.941266],
  [40.791838, -73.965153], [40.85368, -73.949495],
  [40.776857, -73.976529], [40.771009, -73.960687],
  [40.757183, -73.990178], [40.751992, -73.973931]
];


export default {
  data() {
    return {
      myTime: 0,
      myData: [],
      heat: null,
      map: null,
    };
  },
  watch: {
    myTime() {
      this.recreateData();
      this.resetLayer();
    }
  },
  methods: {
    recreateData() {
      this.myData = [];
      for (let i = 0; i < 50; i++) {
        for (let j = 0; j < 50; j++) {
          let data = _.clone(zonesCenters[0]);
          data[0] -= 0.05/50 * i;
          data[1] -= 0.05/50 * j;
          const zoneId = this.pointToIdZone({lat: data[0], long: data[1]});
          data[2] = zoneId / 8;
          this.myData.push(data);
        }
      }
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

    this.recreateData();
    this.resetLayer();
  }
}
</script>

<style lang="scss">

</style>
