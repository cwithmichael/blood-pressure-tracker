import axios from "axios";

export default {
    getReadings (cb) {
        axios
      .get("http://localhost:8080/readings")
      .then(response => {
        cb(response.data);
      })
      .catch(error => {
        console.log(error);
        cb(error);
      })
    },
    addReading (cb, reading) {
        axios
        .post("http://localhost:8080/readings", {
          systolic: reading.systolic,
          diastolic: reading.diastolic,
          pulse: reading.pulse
        })
        .then((response) => {
            cb(response.data);
        })
        .catch(error => {
          console.log(error);
        });
    },
    deleteReading (cb, id) {
      axios
        .delete("http://localhost:8080/readings/" + id)
        .then(response => cb(response.data))
        .catch(error => console.log(error));
    },
 }