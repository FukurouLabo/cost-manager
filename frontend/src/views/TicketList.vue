<template>
  <div class="about">
    <h1>Ticket List</h1>
    <div class="ticket-list">
      <div v-for="ticket in ticketList" :key="ticket.key" class="ticket" @click="selectTicket(ticket.key)" v-bind:class="judgementTicket(ticket.key)">
        <div class="header">
          <p class="key">{{ticket.key}}</p>
          <div class="fields">
            <p class="summary">{{ticket.fields.summary}}</p>
            <p class="self">{{ticket.self}}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
  return {
      ticketList: null,
      selectedTicket: null,
    }
  },
  mounted() {
    window.backend
      .fetchIssueList()
      .then((res) => {
        this.ticketList = res
      })
      .catch((err) => {
        console.log(err);
      });
  },
  methods: {
    selectTicket(key) {
      if (this.selectedTicket === key) {
        this.selectedTicket = null;
      } else {
        this.selectedTicket = key
      }
      
    },
    judgementTicket(key) {
      return (
        (this.selectedTicket === key) ? "active" : ""
      );
    }
  }
};
</script>

<style lang="scss" scoped>
.ticket-list * {
  color: #344563;
}
.ticket-list {
  display: flex;
  justify-content: flex-start;
  align-items: flex-start;
  flex-wrap: wrap;
  .ticket {
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
      content: '記録開始！';
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
  .ticket.active {
    border: solid #FF719A 1px;
    &::after {
      content: '記録終了！';
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