import Vue from 'vue'

import Cookies from 'js-cookie'

import 'normalize.css/normalize.css' // a modern alternative to CSS resets

import ElementUI from 'element-ui'
import './styles/element-variables.scss'

import '@/styles/index.scss' // global css
import '@/styles/admin.scss'

// 需要按需引入，先引入vue并引入element-ui
import AFTableColumn from 'af-table-column'
Vue.use(AFTableColumn)

import App from './App'
import store from './store'
import router from './router'
import permission from './directive/permission'
import clipboard from '@/directive/clipboard'

import { getDicts } from '@/api/admin/dict/data'
import { getItems, setItems } from '@/api/table'
import { getConfigKey } from '@/api/admin/sys-config'
import {
  parseTime,
  resetForm,
  addDateRange,
  selectDictLabel, /* download,*/
  selectItemsLabel,
  timeFormatter,
  dateFormatter,
  floatFormatter2,
  floatFormatter3,
  percentageFormatter, parseBoolean, boolFormatter
} from '@/utils/custom'

import './icons' // icon
import './permission' // permission control
import './utils/error-log' // error log

import Viser from 'viser-vue'
Vue.use(Viser)

import * as filters from './filters' // global filters

import Pagination from '@/components/Pagination'
import BasicLayout from '@/layout/BasicLayout'
import K2Dialog from '@/components/K2Dialog'
import DatetimeRanger from '@/components/DatetimeRanger'

// particle effect, see login/index.vue
import VueParticles from 'vue-particles'
Vue.use(VueParticles)

import '@/utils/dialog'

// 全局方法挂载
Vue.prototype.getDicts = getDicts
Vue.prototype.getItems = getItems
Vue.prototype.setItems = setItems
Vue.prototype.getConfigKey = getConfigKey
Vue.prototype.parseTime = parseTime
Vue.prototype.parseBoolean = parseBoolean
Vue.prototype.resetForm = resetForm
Vue.prototype.addDateRange = addDateRange
Vue.prototype.selectDictLabel = selectDictLabel
Vue.prototype.selectItemsLabel = selectItemsLabel
Vue.prototype.timeFormatter = timeFormatter
Vue.prototype.dateFormatter = dateFormatter
Vue.prototype.floatFormatter = floatFormatter2
Vue.prototype.floatFormatter3 = floatFormatter3
Vue.prototype.percentageFormatter = percentageFormatter
Vue.prototype.boolFormatter = boolFormatter

// Vue.prototype.download = download

// 全局组件挂载
Vue.component('Pagination', Pagination)
Vue.component('BasicLayout', BasicLayout)
Vue.component('K2Dialog', K2Dialog)
Vue.component('DatetimeRanger', DatetimeRanger)

Vue.prototype.msgSuccess = function(msg) {
  this.$message({ showClose: true, message: msg, type: 'success' })
}

Vue.prototype.msgError = function(msg) {
  this.$message({ showClose: true, message: msg, type: 'error' })
}

Vue.prototype.msgInfo = function(msg) {
  this.$message.info(msg)
}

Vue.use(permission)
Vue.use(clipboard)

Vue.use(ElementUI, {
  size: Cookies.get('size') || 'small' // set element-ui default size
})

import 'remixicon/fonts/remixicon.css'

console.info(`欢迎使用 K2`)

// register global utility filters
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

Vue.config.productionTip = false
ElementUI.Dialog.props.closeOnClickModal.default = false

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
