<template>
  <div class="sp-component sp-welcome">
    <div class="sp-box sp-shadow sp-welcome__main">
      <div class="sp-welcome__header">
        <h2>Welcome to Arkhadian smart chain wallet</h2>
      </div>
      <p>
       Manage your ASC assets<br />
        
      </p>
      
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue'
import SpButton from '../SpButton'

export default defineComponent({
  name: 'SpWelcome',
  components: {
    SpButton,
  },
})
</script>
