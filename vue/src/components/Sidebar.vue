<template>
  <SpSidebar v-on:sidebar-open="sidebarOpen = true" v-on:sidebar-close="sidebarOpen = false">
    <template v-slot:default>
      <SpLinkIcon link="/" text="Dashboard" icon="Dashboard" />
      
      <SpLinkIcon link="/relayers" text="Relayers" icon="Transactions" />
      <br>
      <SpLinkIcon href="https://www.arkhadian.com/arkhadian" target="_blank" text="Buy" icon="Token" />
    </template>
    <template v-slot:footer>
      <SpStatusAPI :showText="sidebarOpen" />
      <SpStatusRPC :showText="sidebarOpen" />
      <SpStatusWS :showText="sidebarOpen" />
      <div class="sp-text">Build: v0.3.8</div>
    </template>
  </SpSidebar>
</template>

<script>
export default {
  name: 'Sidebar',
  data() {
    return {
      sidebarOpen: true,
    }
  },
  computed: {
    hasWallet() {
      return this.$store.hasModule(['common', 'wallet'])
    },
  },
}
</script>
