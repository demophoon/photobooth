import Vue from 'vue'
import WebCam from "vue-web-cam"

import Root from "./components/root.vue"

require('bulma/css/bulma.css')

Vue.use(WebCam)

const vueContainer = document.createElement("div")
document.body.appendChild(vueContainer)


const app = new Vue({
    el: vueContainer,
    data: function() {
        return {
            acceptedTerms: false,
            agreedToTerms: false,
        }
    },
    computed: {
        enableCamera() {
            return this.acceptedTerms && this.acceptedTerms
        }
    },
    components: {
        Root,
    },
    template: `<div class='has-text-white'>
        <Root v-if="enableCamera" />
        <div class="container" v-else>
            <p>
                terms and conditions placeholder text
            </p>
            <br>

            <input type="checkbox" id="checkbox" v-model="agreedToTerms">
            <label for="checkbox">Read release form</label>
            <br>
            <button class="button" @click="acceptedTerms = true" :disabled="!agreedToTerms">Agreed to release form</button>
        </div>
    </div>`,
})