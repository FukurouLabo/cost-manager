<template>
  <div>
    <h1 v-if="!selectedIssueId">Issue List</h1>
    <h1 v-if="selectedIssueId" class="active">Recoding...</h1>
    <div class="refresh">
      <button @click="fetchIssues">
        <img
          alt="refresh"
          src="../assets/refresh.png"
        />
      </button>
    </div>
    <div class="issue-list">
      <div v-if="!issueList" class="loading">
        <Loading />
      </div>
      <div 
        v-for="issue in issueList"
        :key="issue.id"
        class="issue"
        @click="selectIssue(issue.id, issue.fields.summary)"
        v-bind:class="judgementIssue(issue.id)"
      >
        <div class="header">
          <p class="key">{{issue.key}} | {{issue.id}}</p>
          <div class="fields">
            <p class="summary">{{issue.fields.summary}}</p>
            <p class="self">{{issue.self}}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Loading from "@/components/Loading.vue";

export default {
  name: "IssueList",
  components: {
    Loading,
  },
  data() {
  return {
    issueList: null,
    selectedIssueId: null,
    selectedIssueSummary: null,
    }
  },
  mounted() {
    // 起動時に前回記録中になっているIssueが無いか確認
    window.backend
      .fetchRecordingIssue()
      .then((id) => {
        this.selectedIssueId = id;

        // Issueの一覧取得
        window.backend
          .fetchIssues()
          .then((res) => {
            if (id) {
              this.issueList = res.filter(issue => issue.id === id);
            } else {
              this.issueList = res;
            }
          })
          .catch((err) => {
            console.log(err);
          });
      })
      .catch((err) => {
        console.log(err);
      })
  },
  methods: {
    fetchIssues() {
      this.issueList = null;
      window.backend
      .fetchRecordingIssue()
      .then((id) => {
        this.selectedIssueId = id;

          // Issueの一覧取得
          window.backend
            .fetchIssues()
            .then((res) => {
              if (id) {
                this.issueList = res.filter(issue => issue.id === id);
              } else {
                this.issueList = res;
              }
            })
            .catch((err) => {
              console.log(err);
            });
        })
        .catch((err) => {
          console.log(err);
        });
    },
    selectIssue(id, summary) {
      if (this.selectedIssueId === id) {
        this.issueList =null;
        this.selectedIssueId = null;
        this.selectedIssueSummary = null;
        // 計測停止
        window.backend
          .finishMeasurement()
          .then((res) => {
            console.log(res);
          })
          .catch((err) => {
            console.log(err);
          });
        // Issueの一覧取得
        window.backend
          .fetchIssues()
          .then((res) => {
            this.issueList = res;
          })
          .catch((err) => {
            console.log(err);
          });
      } else {
        this.issueList = this.issueList.filter(issue => issue.id === id);
        this.selectedIssueId = id;
        this.selectedIssueSummary = summary;
        // 計測開始
        window.backend
          .startMeasurement(id)
          .then((res) => {
            console.log(res);
          })
          .catch((err) => {
            console.log(err);
          });
      }
      
    },
    judgementIssue(id) {
      return (
        (this.selectedIssueId === id) ? "active" : ""
      );
    }
  }
};
</script>

<style lang="scss" scoped>
.issue-list * {
  color: #5571a2;
}
h1 {
  color: #ebebeb;
}
h1.active {
  color: #FF719A;
}
.refresh {
  display: flex;
  justify-content: flex-end;
  button {
    width: 74px;
    padding: 5px 20px 0;
    cursor: pointer;
    background: #ebebeb;
    border: solid 1px #ebebeb;
    border-radius: 3px;
    margin-bottom: 10px;
    transition: .2s;
    &:hover {
      opacity: .7;
    }
    img {
      width: 80%;
    }
  }
}
.loading {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}
.issue-list {
  display: flex;
  justify-content: flex-start;
  align-items: flex-start;
  flex-wrap: wrap;
  .issue {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    width: 100%;
    margin: 3px;
    border: solid #344563 1px;
    cursor: pointer;
    transition: .3s;
    &::after {
      content: 'Record Start';
      width: 100%;
      height: 100%;
      background: rgba(0, 0, 0, .7);
      color: #fff;
      display: none;
      position: absolute;
      top: 0;
      left: 0;
    }
    &:hover {
      opacity: .7;
      &::after {
        display: flex;
        justify-content: center;
        align-items: center;
      }
    }
    .header {
      display: flex;
      justify-content: flex-start;
      width: 100%;
      .key {
        display: flex;
        justify-content: center;
        align-items: center;
        background: #344563;
        height: auto;
        color: #fff;
        font-weight: bold;
        margin: 0;
        padding: 20px;
        width: 20%;
      }
      .fields {
        display: flex;
        justify-content: flex-start;
        flex-wrap: wrap;
        width: 80%;
        .summary {
          width: 100%;
          font-size: 20px;
          font-weight: bold;
          text-align: left;
          margin: 0;
          padding: 20px 20px 0;
        }
        .self {
          width: 100%;
          text-align: left;
          margin: 0;
          padding: 5px 20px 20px;
        }
      }
    }
  }
  .issue.active {
    border: solid #FF719A 1px;
    &::after {
      content: 'Record End';
      width: 100%;
      height: 100%;
      background: rgba(0, 0, 0, .7);
      color: #fff;
      display: none;
      position: absolute;
      top: 0;
      left: 0;
    }
    &:hover {
      opacity: .7;
      &::after {
        display: flex;
        justify-content: center;
        align-items: center;
      }
    }
    .header {
      .key {
        background-image: linear-gradient(-225deg, #FFE29F 0%, #FFA99F 48%, #FF719A 100%);
        background-size: 200% 200%;
        animation: bggradient 5s ease infinite;
      }
      .fields {
        .summary, .self {
          color: #FF719A;
        }
      }
    }
  }
}

@keyframes bggradient{
	0% { background-position: 0% 50% }
	50% { background-position: 100% 50% }
	100% { background-position: 0% 50% }
}
</style>
