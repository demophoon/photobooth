<template>
<div>
    <div class="loading"></div>

    <div class="container py-3">
        <div class="columns is-centered">
            <div v-show="img" class="column is-half">
                <figure class="image is-1by1">
                    <img :src="img">
                </figure>
            </div>
            <div v-show="!img" class="column is-half">
                <span class="bordered">
                    <vue-web-cam
                        ref="webcam"
                        :device-id="deviceId"
                        @cameras="onCameras"
                        @camera-change="onCameraChange"
                    />
                </span>
            </div>
        </div>

        <div class="columns is-centered">
            <div class="column is-half has-text-centered">

                <span v-show="img">
                    <div class="content">
                        <p>
                            <button class="button is-info" @click="showIntakeForm">Share with PyCascades</button>
                            <button class="button is-danger" @click="reset">Snap Another</button>
                        </p>
                        <p>
                            If you decide to share your photo with PyCascades you
                            will allow us to add your photo to a collage that we put
                            together every year to show off our amazing community.
                        </p>
                        <p>
                            You may share this photo with your community regardless
                            of your choice to share or not share with PyCascades.
                        </p>
                    </div>
                </span>

                <span v-show="!img">
                    <button class="button" v-bind:class="{ 'is-loading': !captureEnable }" @click="onCapture">Capture</button>

                    <div class="select">
                        <select v-model="camera">
                            <option>-- Select Device --</option>
                            <option
                                v-for="device in devices"
                                :key="device.deviceId"
                                :value="device.deviceId"
                            >{{ deviceIDToLabel(device.deviceId) }}</option>
                        </select>
                    </div>
                </span>

            </div>
        </div>

    </div>

    <intake-form v-show="intakeFormVisible" :img="img" @close="hideIntakeForm" @resetBooth="reset" />

</div>
</template>

<script>
import ImageJS from "image-js"

function importAll(r) {
    return r.keys().map(r)
}
const OverlayImages = importAll(require.context("../img/overlays/", false, /\.png$/))
const eventName = require("../img/friday-night.png")
const imageSize = 1000

export default {
    data() {
        return {
            img: null,
            camera: null,
            deviceId: null,
            devices: [],
            deviceMap: {},
            intakeFormVisible: false,
            captureEnable: true,
            frames: [],
        };
    },
    computed: {
        device: function() {
            return this.devices.find(n => n.deviceId === this.deviceId);
        }
    },
    watch: {
        camera: function(id) {
            this.deviceId = id;
        },
        devices: function() {
            // Once we have a list select the first one
            const [first, ...tail] = this.devices;
            if (first) {
                this.camera = first.deviceId;
                this.deviceId = first.deviceId;
            }
        }
    },
    methods: {
        reset() {
            this.img = null
            this.hideIntakeForm()
        },
        showIntakeForm() {
            this.intakeFormVisible = true
        },
        hideIntakeForm() {
            this.intakeFormVisible = false
        },
        onCapture() {
            let self = this
            self.captureEnable = false
            let img = this.$refs.webcam.capture();

            var readyImage = function(src) {
                return new Promise((resolve, reject) => {
                    ImageJS.load(src).then((localImg) => {
                        if (localImg.height > localImg.width) {
                            localImg = localImg.resize({
                                width: imageSize,
                                preserveAspectRatio: true,
                            })
                        } else {
                            localImg = localImg.resize({
                                height: imageSize,
                                preserveAspectRatio: true,
                            })
                        }
                        let result = new Image()
                        result.onload = function() {
                            resolve(result)
                        }
                        result.src = localImg.toDataURL()
                    })
                })
            }

            let randomFrame = OverlayImages[Math.floor(Math.random() * OverlayImages.length)]
            Promise.all([
                readyImage(img),
                readyImage(randomFrame),
                readyImage(eventName),
            ]).then((images) => {
                let canvas = document.createElement("canvas")
                canvas.height = imageSize
                canvas.width = imageSize
                let ctx = canvas.getContext('2d')
                for (let i in images) {
                    let img = images[i]
                    let centered_x = (Math.max(imageSize, img.width) - imageSize) / -2
                    let centered_y = (Math.max(imageSize, img.height) - imageSize) / -2
                    ctx.drawImage(img, centered_x, centered_y)
                    self.img = canvas.toDataURL()
                    self.captureEnable = true
                }
            })
        },
        onStop() {
            this.$refs.webcam.stop();
        },
        onStart() {
            this.$refs.webcam.start();
        },
        onCameras(cameras) {
            this.devices = cameras;
            let self = this;
            navigator.mediaDevices.enumerateDevices().then((devices) => {
                devices.forEach((device) => {
                    self.deviceMap[device.deviceId] = device.label
                })
            })
        },
        onCameraChange(deviceId) {
            this.deviceId = deviceId;
            this.camera = deviceId;
        },
        deviceIDToLabel(deviceId) {
            return this.deviceMap[deviceId]
        },
    }
};
</script>

<style>
</style>