<template>
  <div>
    <div>code : {{ code }}</div>
    <div>
      Description <input type="textarea" v-model="description">
      画像 <input type="file" @change="fileSelected">
    </div>
  <div>
    <img class="preview-image" :src="base64Image">
  </div>
    <button @click="update">編集する</button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
name: "Edit",
  data () {
    return {
      code: "",
      description: "",
      image: "",
      base64Image: "",
      fileInfo: ''
    }
  },
  created() {
  console.log(this.$route.params.id)
    axios.get('http://localhost/api/codeImage/' + this.$route.params.id)
      .then((res) => {
        console.log(res.data)
        this.code = res.data.Code
        this.description = res.data.Description
        this.image = res.data.Image
      })
  },
  methods: {
    fileSelected(event) {
      this.fileInfo = event.target.files[0]
      this.generateImageUrl(this.fileInfo);
    },
    generateImageUrl(file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.createImageObject(e.target.result)
          this.base64Image = e.target.result;
        }
        reader.readAsDataURL(file);
    },
    createImageObject(file) {
      const image = new Image();
      image.crossOrigin = 'Anonymous';

      image.onload = (e) => {
        const resizedBase64 = this.resize(image);
        const resizedImage = this.base64ToBlob(resizedBase64);
        const resizedImg = this.createObjectUrl(resizedImage);
        this.base64Image = resizedImg;
      };
      image.src = file;
    },
    base64ToBlob(base64) {
      const bin = atob(base64.replace(/^.*,/, ''));
      const buffer = new Uint8Array(bin.length);
      for (let i = 0; i < bin.length; i++) {
        buffer[i] = bin.charCodeAt(i);
      }
      return new Blob([buffer.buffer], {
        type: 'image/png'
      });
    },
    resize(image) {
      const canvas = document.createElement('canvas');
      canvas.width = 300;
      canvas.height = 300;
      canvas.getContext('2d').drawImage(image, 0, 0, 300, 300);
      return canvas.toDataURL('image/jpg');
    },
    createObjectUrl(resizedImage) {
      return window.URL.createObjectURL(resizedImage);
    },
    update() {
      // const formData = new FormData()
      // formData.append('file', this.fileInfo)
      let params = JSON.stringify({
        id: this.$route.params.id,
        description: this.description,
      })
      // formData.append('params',params)
      axios.post('http://localhost/api/codeImage/' + this.$route.params.id, params,{"headers": {"Content-Type": "application/json", "Accept": "application/json"}})
        .then((res) => {
          console.log(res)
        })
        .catch((err) => {
          console.log(err.response)
        })
    }
  }
}
</script>

<style scoped>

</style>
