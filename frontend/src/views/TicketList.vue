<template>
  <div class="about">
    <h1>チケット一覧</h1>
    <div v-if="display">
      <p>{{message}}</p>
    </div>

    <div v-for="ticket in ticketList" :key="ticket.name">
      <p>{{ticket.name}}</p>
      <p>{{ticket.url}}</p>
      <p>{{ticket.number}}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
  return {
      display: false,
      message: '',
      ticketList: null,
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
    
    window.backend
      .testTicketList()
      .then((res) => {
        this.ticketList = res
        console.log(res)
      })
      .catch((err) => {
        console.log(err);
      });
  },
};
</script>