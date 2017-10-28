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
import L from 'leaflet';
import Simpleheat from 'simpleheat';
import LeafletHeat from 'leaflet.heat';

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
      for (let i = -10; i < 10; i++) {
        for (let j = -10; j < 10; j++) {
          this.myData.push([48.79208 + i / 1000, 9.23218 + j / 1000, Math.random()]);
        }
      }
    },
    resetLayer() {
      if(this.heat) {
        this.map.removeLayer(this.heat);
      }
      this.heat = L.heatLayer(
        this.myData, { radius: 40 }).addTo(this.map);
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
      center: new L.LatLng(48.79208, 9.23218),
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
