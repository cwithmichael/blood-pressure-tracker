import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from "vuex-persistedstate"
import bp from '../api/bp.js'
Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

const store = new Vuex.Store({
  plugins: [createPersistedState()],
  state: {
    readings: []
  },
  actions: {
    getAllReadings({
      commit
    }) {
      return new Promise((resolve) => {
        bp.getReadings(readings => {
          resolve(commit('setReadings', readings))
        })
      })
    },
    addReading({
      commit
    }, reading) {
      return new Promise((resolve) => {
        bp.addReading(reading => {
          resolve(commit('addReading', reading))
        }, reading)
      })
    },
    deleteReading({
      commit
    }, reading) {
      return new Promise((resolve) => {
        bp.deleteReading(() => {
          resolve(commit('deleteReading', reading.id))
        }, reading.id)
      })
    }
  },
  getters: {
    readingsForToday: state => {
      if (state.readings) {
        var sortedReadings = [...state.readings].sort((a,b) => a.readingDate - b.readingDate)
        return sortedReadings.filter(reading => {
          var now = new Date();
          var d = new Date(reading.readingDate * 1000);
          return now.getDate() === d.getDate();
        });
      } else {
        return []
      }
    },
    readingsExcludingToday: state => {
      if (state.readings) {
        return state.readings.filter(reading => {
          var now = new Date();
          var d = new Date(reading.readingDate * 1000);
          return now.getDate() !== d.getDate();
        });
      } else {
        return []
      }
    }
  },
  mutations: {
    setReadings(state, readings) {
      state.readings = readings;
    },
    addReading(state, reading) {
      if (state.readings && state.readings.length > 0)
        state.readings = [...state.readings, reading]
      else state.readings = [reading]
    },
    deleteReading(state, id) {
      state.readings = state.readings.filter(r => {
        return r.id != id;
      })
    }
  },
  strict: debug,
})

export default store;
