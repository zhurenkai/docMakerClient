// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-default/index.css'
import VueResource from 'vue-resource'
import axios from './http'
import {getUri} from './config/config.js'
import Vuex from 'vuex'
import store from './store'
import VueClipboard from 'vue-clipboard2'
import helper from './helper'

Vue.use(helper)
// 剪切板
Vue.use(VueClipboard)
Vue.use(VueResource)
Vue.use(ElementUI)
Vue.prototype.axios = axios
// for English  (default chinese)
// version 2.2.0 or later
Vue.use(Vuex)
/* eslint-disable no-new */

new Vue({
  el: '#app',
  router,
  axios,
  store,
  template: '<App/>',
  components: { App },
  render: h => h(App)
})
