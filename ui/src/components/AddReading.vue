<template>
  <div id="add-reading" class="text-center">
    <form
      style="background-color: white;"
      v-if="userRequested"
      class="text-center"
      id="readings"
      @submit.prevent="processForm"
    >
      <br />
      <div class="row">
        <div class="col"></div>
        <div class="col text-center">
          <label for="systolic-input">Systolic Reading</label>
          <input
            v-model="systolic"
            type="text"
            class="form-control"
            id="systolic-input"
            placeholder="120"
            required
          />
          <label for="diastolic-input">Diastolic Reading</label>
          <input
            v-model="diastolic"
            type="text"
            class="form-control"
            id="diastolic-input"
            placeholder="70"
            required
          />
          <label for="pulse-input">Pulse (Optional)</label>
          <input v-model="pulse" type="text" class="form-control" id="pulse-input" placeholder="60" />
          <div class="text-center">
            <button type="submit" class="btn btn-primary">Add Reading</button>
          </div>
        </div>
        <div class="col"></div>
      </div>
    </form>
    <button
      v-if="!userRequested"
      class="btn btn-success"
      type="button"
      v-on:click="toggleUserRequest"
    >Add New Readings for Today</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      systolic: "",
      diastolic: "",
      pulse: "",
      userRequested: false,
    };
  },
  methods: {
    toggleUserRequest() {
      this.userRequested = !this.userRequested;
    },
    processForm() {
      if (
        !this.systolic ||
        this.systolic.length === 0 ||
        !this.diastolic ||
        this.diastolic.length === 0
      ) {
        return;
      }
      this.$store.dispatch("addReading", {
        systolic: this.systolic,
        diastolic: this.diastolic,
        pulse: this.pulse,
      });
    },
  },
};
</script>

<style>
input:focus:required:invalid {
  border: 2px solid red;
}
</style>