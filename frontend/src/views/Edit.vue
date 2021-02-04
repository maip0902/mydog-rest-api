<template>
  <div>
    <div>code : {{ code }}</div>
    <div>
      <div class="edit-input">
        <label for="description" class="edit-label">Description</label>
        <input id="description" type="textarea" class="edit-field" v-model="description">
      </div>
      <input id="image" type="file" @change="fileSelected" style="display:none">
      <button @click="selectFile" class="main-btn">新しい画像を選択</button>
    </div>
    <div>
      <div>いまの画像</div>
      <img class="preview-image" :src="imageData">
    </div>
    <div v-if="isUploaded" class="new-image">
      <div>新しい画像</div>
      <img class="cancel-image" src="../assets/cancel.png" @click="cancelImage">
      <img class="preview-image" :src="base64Image">
      <!-- <button @click="cancelImage">この画像をキャンセル</button> -->
    </div>
    <button @click="update" class="main-btn edit-btn">編集する</button>
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
      fileInfo: '',
      imageData: "",
      isUploaded: false
    }
  },
  beforeCreate() {
    let params = JSON.stringify({
      token: this.$session.get('jwt')
    })
    axios.post("/api/authUser", params, {headers: {'Content-Type': 'application/json','Accept': 'application/json', 'Authentication': this.$session.get('jwt')}})
      .then((res) => {
        this.$session.set('user', res.data)
      })
      .catch((err) => {
        this.$session.destroy()
        this.$router.push('/login')
    })
  },
  created() {
    axios.get('/api/codeImage/' + this.$route.params.id)
      .then((res) => {
        this.code = res.data.Code
        this.description = res.data.Description
        this.image = res.data.Image
        axios.get('/api/codeImage/image/' + this.code)
      .then((res) => {
        this.imageData = 'data:image/png;base64,' + res.data.Image
      })
      })    
  },
  methods: {
    selectFile() {
      $('#image').click()
    },
    fileSelected(event) {
      this.isUploaded = true
      this.fileInfo = event.target.files[0]
      this.generateImageUrl(this.fileInfo);
    },
    generateImageUrl(file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.createImageObject(e.target.result)
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
        this.base64Image = resizedBase64;
        this.fileInfo = resizedBase64;
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
      return URL.createObjectURL(resizedImage);
    },
    cancelImage() {
      this.isUploaded = false
      this.fileInfo = ""
      $('#image').val(null)
    },
    update() {
      let params = JSON.stringify({
        code: this.code,
        id: this.$route.params.id,
        description: this.description,
        image: this.fileInfo.replace(/^data:\w+\/\w+;base64,/, '')
      })
      console.log(params);
      axios.post('/api/codeImage/' + this.$route.params.id, params,{"headers": {"Content-Type": "application/json", "Accept": "application/json"}})
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
