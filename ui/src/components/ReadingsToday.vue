
<template>
  <div id="readings-today">
    <h3
      class="text-center"
      style="margin-top:20px; font-family: 'Libre Baskerville', serif;"
    >Today's Readings</h3>
    <hr/>
    <h4
      class="text-center"
    >Average Systolic: {{ getAverageSystolic() }} | Average Diastolic: {{ getAverageDiastolic() }}</h4>
    <p v-if="loading" class="text-center">Loading...</p>
    <div v-else>
    <table class="table tale-sm table-bordered">
      <thead>
        <tr>
          <th scope="col">Systolic</th>
          <th scope="col">Diastolic</th>
          <th scope="col">Pulse</th>
          <th scope="col">Time</th>
          <th scope="col"></th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="reading in readingsForToday"
          v-bind:class="getReadingClass(reading)"
          v-bind:key="reading.id"
        >
          <td>{{ reading.systolic }}</td>
          <td>{{ reading.diastolic }}</td>
          <td>{{ reading.pulse }}</td>
          <td>{{ prettifyDate(reading.createdDate) }}</td>
          <td>
            <button v-on:click="deleteReading(reading)" class="btn btn-danger">
              Delete
              Reading
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    </div>
    <AddReading />
  </div>
</template>

<script>
import AddReading from "./AddReading.vue";
import { mapGetters, mapActions } from "vuex";
import moment from "moment";

export default {
  components: {
    AddReading
  },
  computed: {
    ...mapGetters(["readingsForToday"])
  },
  data() {
    return {
      readingAdded: false,
      averageSystolic: 0,
      averageDiastolic: 0,
      loading: true
    };
  },
  methods: {
    ...mapActions(["deleteReading"]),
    prettifyDate(date) {
      return moment(date).format("h:mm:ss a");
    },
    getReadingClass(reading) {
      if (reading.systolic > 120 || reading.diastolic > 80) {
        return "table-danger";
      }
      return "table-success";
    },
    getAverageSystolic() {
      var systolic = [];
      this.readingsForToday.forEach(reading => {
        systolic.push(reading.systolic);
      });
      var total = systolic.reduce((acc, c) => acc + c, 0);
      if (systolic.length > 0) return Math.floor(total / systolic.length);
      return 0;
    },
    getAverageDiastolic() {
      var diastolic = [];
      this.readingsForToday.forEach(reading => {
        diastolic.push(reading.diastolic);
      });
      var total = diastolic.reduce((acc, c) => acc + c, 0);
      if (diastolic.length > 0) return Math.floor(total / diastolic.length);
      return 0;
    }
  },
  created() {
      this.$store.dispatch("getAllReadings").then(() => {
        this.loading = false;
      });
    }
};
</script>