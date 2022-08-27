<template>
  <div class="setup">
    <div class="header">
      <h1>Setup Page</h1>
      <div>
        <p>
          Setup your config file.
        </p>
      </div>
    </div>
    <div class="form">
      <div class="input-block">
        <label for="mail">
          <p class="title">mail</p>
          <p class="description">
            Jiraアカウントで使用しているメールアドレスを入力してください。
          </p>
        </label>
        <input id="mail" v-model="mail" type="email">
      </div>
      <div class="input-block">
        <label for="token">
          <p class="title">token</p>
          <p class="description">
            <a href="https://id.atlassian.com/login?continue=https%3A%2F%2Fid.atlassian.com%2Fmanage-profile%2Fsecurity%2Fapi-tokens">
              こちら
            </a>
            でトークンを作成し、下記フィールドにペーストしてください。<br>
            ※ Labelの指定はありません。
          </p>
        </label>
        <input id="token" v-model="token" type="text">
      </div>
      <div class="submit-block">
        <button class="submit" @click="createConfigFile">
          Submit
        </button>
      </div>
    </div>
  </div>
</template>

<script>
// import Loading from "@/components/Loading.vue";

export default {
  name: "Setup",
  data() {
    return {
      mail: "",
      token: "",
    }
  },
  methods: {
    createConfigFile() {
      // Configファイル生成
      window.backend
        .createConfigFile(this.mail, this.token)
        .then((res) => {
          console.log(res);
          // TODO:バックエンドの実装が完了したらコメントアウト撤去
          // this.$router.push('/')
        })
        .catch((err) => {
          console.log(err);
        })
    }
  }
};
</script>

<style lang="scss" scoped>
.setup {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-wrap: wrap;
  .header {
    color: #ebebeb;
    width: 100%;
  }
  .form {
    width: 500px;
    margin: 80px 20px 0;
    padding: 40px 30px;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    border-radius: 5px;
    box-shadow: rgb(36 36 36) 0px 10px 50px;
  
  
    .input-block {
      width: 100%;
      &:not(:last-child) {
        margin-bottom: 60px;
      }
      label {
        color: #ebebeb;
        display: flex;
        justify-content: flex-start;
        align-items: center;
        flex-wrap: wrap;
        .title {
          margin: 0;
          font-size: 22px;
          font-weight: bold;
        }
        .description {
          margin: 10px 5px;
          font-size: 12px;
          width: 100%;
          text-align: left;
          color: #999999;
          a {
            color: #ebebeb;
          }
        }
      }
      input {
        margin: 30px 0 0;
        padding: 5px;
        width: 100%;
        background: #161616;
        box-shadow: rgb(14 14 14) 3px 3px 6px 0px inset, rgb(44 44 44 / 50%) -3px -3px 6px 1px inset;
        border: none;
        border-bottom: solid 1px #ebebeb;
        color: #ebebeb;
      }
    }
    .submit-block {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100%;
      .submit {
        width: 100%;
        margin-top: 20px;
        padding: 10px;
        background: rgb(104, 104, 255);
        color: #ebebeb;
        font-weight: bold;
        border-radius: 5px;
        cursor: pointer;
        transition: .3s;
        &:hover {
          opacity: .7;
        }
      }
    }
  }
}
</style>