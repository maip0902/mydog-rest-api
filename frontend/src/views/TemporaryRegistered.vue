<template>
<div>
    <h1>仮登録完了</h1>
    <p>メールアドレスの確認のため、ご入力いただいたメールアドレスに本登録用のリンクを送信致しました。</p>
    <button class="main-btn" @click="goTop">ログイン画面に戻る</button>
</div>
</template>

<script>
export default {
  name: "TemporaryRegistered",
  methods: {
    goTop() {
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>

</style>