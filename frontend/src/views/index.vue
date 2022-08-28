<template>
  <Loading v-if="configFileExists === null" />
  <IssueList v-if="configFileExists === true" />
</template>

<script>
import Loading from "@/components/Loading.vue";
import IssueList from "@/components/IssueList.vue";

export default {
  name: "Index",
  components: {
    Loading,
    IssueList,
},
  data() {
    return {
      configFileExists: null,
    }
  },
  mounted() {
    // Configファイルの存在確認
    window.backend
      .fetchConfig()
      .then((res) => {
        console.log(res);
        this.configFileExists = true;
      })
      .catch((err) => {
        console.log(err);
        this.$router.push('/setup')
      })
  }
};
</script>
