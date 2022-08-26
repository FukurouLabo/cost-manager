<template>
  <div class="about">
    <h1>チケット一覧</h1>
    <div class="ticket-list">
      <div v-for="ticket in ticketList" :key="ticket.name" class="ticket" @click="selectTicket(ticket.name)">
        <div class="header">
          <p class="id">{{ticket.number}}</p>
          <div class="title">
            <p class="name">{{ticket.name}}</p>
            <p class="url">{{ticket.url}}</p>
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
    }
  },
  mounted() {
    window.backend
      .testTicketList()
      .then((res) => {
        this.ticketList = res
      })
      .catch((err) => {
        console.log(err);
      });
  },
  methods: {
    selectTicket(name) {
      alert(name);
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
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    width: 100%;
    margin: 3px;
    border: solid #344563 1px;
    cursor: pointer;
    transition: .3s;
    &:hover {
      opacity: .7;
    }
    .header {
      display: flex;
      justify-content: flex-start;
      width: 100%;
      .id {
        display: flex;
        justify-content: center;
        align-items: center;
        background: #344563;
        height: auto;
        color: #fff;
        font-weight: bold;
        margin: 0;
        padding: 20px;
      }
      .title {
        display: flex;
        justify-content: flex-start;
        flex-wrap: wrap;
        .name {
          width: 100%;
          font-size: 20px;
          font-weight: bold;
          text-align: left;
          margin: 0;
          padding: 20px 20px 0;
        }
        .url {
          width: 100%;
          text-align: left;
          margin: 0;
          padding: 5px 20px 20px;
        }
      }
    }
  }
}
</style>