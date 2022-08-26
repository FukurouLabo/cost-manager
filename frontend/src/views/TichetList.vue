<template>
  <div class="about">
    <h1>チケット一覧</h1>
    <div v-if="display">
      <p>{{message}}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
  return {
      display: false,
      message: ''
    }
  },
  mounted() {
    window.backend
      .basic()
      .then((res) => {
        this.display = true;
        this.message = res
      })
      .catch((err) => {
        console.log(err);
      });
  },
};
</script>