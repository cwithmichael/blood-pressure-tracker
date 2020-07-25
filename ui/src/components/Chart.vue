<template>
  <div id="history" class="small" style="margin-top:20px">
    <h2 class="text-center" style="font-family: 'Libre Baskerville', serif;">History</h2>
    <hr />
    <div class="container-fluid text-center">
      <chartist v-if="doneLoading" type="Line" :data="datacollection" :chartOptions="chartOptions"></chartist>
      <p v-else class="text-center">Loading</p>
    </div>
    <br />
  </div>
</template>

<script>
import { mapState } from "vuex";
import moment from "moment";

export default {
  computed: {
    ...mapState(["readings"]),
  },
  components: {},

  data() {
    return {
      datacollection: {},
      doneLoading: false,
      chartOptions: { lineSmooth: false },
    };
  },
  created() {
    this.$store.dispatch("getAllReadings").then(() => {
      this.fillData();
      this.doneLoading = true;
    });
  },
  methods: {
    fillData() {
      var dates = [];
      var systolic = [];
      var diastolic = [];
      if (this.readings) {
        var sorted = [...this.readings].sort(function (a, b) {
          return a.readingDate - b.readingDate;
        });
        sorted.forEach((reading) => {
          systolic.push(reading.systolic);
          diastolic.push(reading.diastolic);
          dates.push(
            moment(reading.readingDate * 1000).format("MM/DD/YY hh:mm a")
          );
        });
        this.datacollection = {
          labels: dates,
          series: [systolic, diastolic],
        };
      }
    },
  },
};
</script>

<style>
.small {
  max-width: 600px;
  margin: auto;
}
</style>
