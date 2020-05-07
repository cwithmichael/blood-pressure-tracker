import Vue from 'vue'
import VueRouter from 'vue-router';
import App from './App.vue'
import Chart from './components/Chart.vue';
import ReadingsToday from './components/ReadingsToday.vue';

import store from './store'

Vue.config.productionTip = false
Vue.use(VueRouter)

const routes = [{
    path: '/today',
    component: ReadingsToday
  },
  {
    path: '/history',
    component: Chart
  },
  { path: '*', redirect: '/today' }
]

const router = new VueRouter({
  routes // short for `routes: routes`
})

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')