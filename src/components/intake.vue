<template>
<div class="modal is-active" @keypress.esc="$emit('close')">
    <div class="modal-background" @click="$emit('close')"></div>

    <div class="modal-content">
        <div>
            <div class="columns is-centered ">
                <div class="column is-full">
                    <div class="box has-text-dark">

                        <div class="content">
                            <p>
                                <h1>Photo Release Collection</h1>
                            </p>

                            <p>
                                Thank you for sharing your photos with our PyCascades 2022
                                Conference! We are looking forward to viewing your amazing
                                photography. Below is our release form which requires your signature
                                before we can share at the event.
                            </p>

                            <div class="field">
                                <label class="label">Full Name</label>
                                <div class="control">
                                    <input type="text" class="input" v-model="fullName" />
                                </div>
                                <p v-if="submitAttempted && !validateCompleted(fullName)" class="help is-danger">
                                    This field is required
                                </p>
                            </div>

                            <div class="field">
                                <label class="label">Email</label>
                                <div class="control">
                                    <input type="text" class="input" v-model="email" />
                                </div>
                                <p v-if="submitAttempted && !validateEmail(email)" class="help is-danger">
                                    This field is required
                                </p>
                            </div>

                            <p>
                                <h2>Photo Release Form</h2>
                            </p>

                            <p>
                                I, (the “Releasor”) grant permission and consent to Python Software Foundation (the “Releasee”) for the use of the following photograph(s) as identified below for presentation under any legal condition, including but not limited to: publicity, copyright purposes, illustration, advertising, and web content: Description: Photographs and/or video of the Releasor for PyCascades 2022 from February 2 to February 6, 2022.
                            </p>

                            <p>
                                <u>Payment</u>
                            </p>
                            <p>
                                I understand that there shall be no payment for this release.
                            </p>

                            <p>
                                <u>Royalties</u>
                            </p>
                            <p>
                                I understand that no royalty, fee, or other compensation shall become payable to me by reason of such use.
                            </p>

                            <p>
                                <u>Revocation</u>
                            </p>
                            <p>
                                I understand that I may revoke this authorization at any time by notifying the Releasee in writing. The revocation will not affect any actions taken before the receipt of this written notification. Images will be stored in a secure location and only authorized staff will have access to them. They will be kept as long as they are relevant and after that time destroyed or archived.
                            </p>

                            <p>
                                We, the Releasor and Releasee, understand and agree to the aforementioned terms and conditions.
                            </p>

                            <div class="field">
                                <label class="label">Signature</label>
                                <div class="control">
                                    <div class="signature-bordered">
                                        <signature ref="signature" w="600px" h="150px" />
                                    </div>
                                </div>
                                <p class="help">
                                    Use your mouse or finger to draw your signature.
                                </p>
                                <p v-if="submitAttempted && !signed" class="help is-danger">
                                    This field is required
                                </p>
                            </div>

                            <div class="field">
                                <label class="label">Do you have any questions?</label>
                                <div class="control">
                                    <textarea class="textarea has-fixed-size" rows="4"></textarea>
                                </div>
                            </div>

                            <div class="control">
                                <button class="button is-primary" @click="upload">Submit</button>
                            </div>

                        </div>

                    </div>

                </div>
            </div>
        </div>

    </div>

    <button class="modal-close is-large" @click="$emit('close')" aria-label="close"></button>
</div>
</template>

<script>
export default {
    data() {
        return {
            fullName: "",
            email: "",
            questions: "",
            signed: false,
            submitAttempted: false,
        }
    },
    mounted() {
        let self = this;
        document.addEventListener('keydown', (event) => {
            const e = event || window.event;

            if (e.keyCode === 27) { // Escape key
                self.$emit('close')
            }
        });
    },
    props: ["img"],
    methods: {
        resetSignature() {
            this.$refs.signature.clear()
            this.signed = false
        },
        isValidSignature() {
            return !this.$refs.signature.isEmpty()
        },
        validateEmail(field) {
            return this.validateCompleted(field) && field.includes("@") && field.includes(".") && !field.endsWith(".")
        },
        validateCompleted(field) {
            return field != ""
        },
        validateForm() {
            return !(this.fullName == "" || this.email == "" || !this.isValidSignature())
        },
        upload() {
            let self = this;
            this.submitAttempted = true
            if (!this.validateForm()) {
                return
            }

            let fd = new FormData()
            fd.append("photo", this.dataURLtoBlob(this.img))
            fd.append("signature", this.dataURLtoBlob(this.$refs.signature.save("image/png")))
            fd.append("fullName", this.fullName)
            fd.append("email", this.email)
            fd.append("questions", this.questions)
            fetch("/api/upload", {
                method: "POST",
                body: fd,
            }).then((d) => {
                self.$emit('close')
            })
        },
        dataURLtoBlob(dataurl) {
            let arr = dataurl.split(',')
            let mime = arr[0].match(/:(.*?);/)[1]
            let bstr = atob(arr[1])
            let n = bstr.length
            let u8arr = new Uint8Array(n)
            while (n--) {
                u8arr[n] = bstr.charCodeAt(n);
            }
            return new Blob([u8arr], {type: mime})
        },
    }
}
</script>

<style>
.signature-bordered {
    border: 1px solid #000;
}

</style>