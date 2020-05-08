import axios from "axios";

export default {
    getReadings (cb) {
        axios
      .get("http://0.0.0.0:9000/readings")
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
        .post("http://0.0.0.0:9000/readings", {
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
        .delete("http://0.0.0.0:9000/readings/" + id)
        .then(response => cb(response.data))
        .catch(error => console.log(error));
    },
 }