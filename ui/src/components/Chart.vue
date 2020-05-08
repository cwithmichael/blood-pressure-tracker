<template>
  <div id="history" class="small" style="margin-top:20px">
    <h2 class="text-center" style="font-family: 'Libre Baskerville', serif;">History</h2>
    <hr/>
    <div class="container-fluid text-center">
      <line-chart v-if="doneLoading" :chart-data="datacollection" :options="options"></line-chart>
      <p v-else class="text-center">Loading</p>
    </div>
    <br />
  </div>
</template>

<script>
import LineChart from "./LineChart.js";
import { mapState } from "vuex";
import moment from "moment";

export default {
  computed: {
    ...mapState(["readings"])
  },
  components: {
    LineChart
  },

  data() {
    return {
      datacollection: {},
      doneLoading: false,
      options: {
        scales: {
          xAxes: [
            {
              ticks: {
                display: false //this will remove only the label
              }
            }
          ]
        }
      }
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
      var systolic = [];
      var diastolic = [];
      var pulse = [];
      var dates = [];
      this.readings.forEach(reading => {
        systolic.push(reading.systolic);
        diastolic.push(reading.diastolic);
        pulse.push(reading.pulse);
        dates.push(moment(reading.createdDate).format("MMMM Do YYYY, h:mm:ss a"));
      });
      this.datacollection = {
        labels: dates,
        datasets: [
          {
            label: "Systolic",
            backgroundColor: "#f87979",
            borderColor: "#f87978",
            fill: false,
            data: systolic
          },
          {
            label: "Diastolic",
            backgroundColor: "#36a2eb",
            borderColor: "#36a2eb",
            fill: false,
            data: diastolic
          }
        ]
      };
    }
  }
};
</script>

<style>
.small {
  max-width: 600px;
  margin: auto;
}
</style>