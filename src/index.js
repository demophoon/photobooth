import Vue from 'vue'
import WebCam from "vue-web-cam"
import VueSignature from "vue-signature"

import Root from "./components/root.vue"
import IntakeForm from "./components/intake.vue"
import Navigation from "./components/navigation.vue"
import Footer from "./components/footer.vue"

require('bulma/css/bulma.css')
require('./index.css')

Vue.use(WebCam)
Vue.component("signature", VueSignature)
Vue.component("intake-form", IntakeForm)
Vue.component("b-root", Root)
Vue.component("b-navigation", Navigation)
Vue.component("b-footer", Footer)

const vueContainer = document.createElement("div")
document.body.appendChild(vueContainer)


const app = new Vue({
    el: vueContainer,
    template: `<div class='has-text-white'>
        <b-navigation />
        <b-root />
        <b-footer />
    </div>`,
})