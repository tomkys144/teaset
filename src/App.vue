<script setup lang="ts">
  import { NLayout, NLayoutContent, NLayoutHeader } from 'naive-ui'
  import 'vfonts/FiraSans.css'
  import Menu from './components/Menu.vue'
</script>

<template>
  <NLayout>
    <NLayoutHeader class="header">
      <Menu />
    </NLayoutHeader>
    <NLayoutContent class="content">
      <router-view />
    </NLayoutContent>
  </NLayout>
</template>

<style lang="scss">
  #app {
    font-family: v-sans, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    font-weight: 400;
  }

  .content {
    margin-top: 5rem;
  }

  .header {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    max-height: 5rem;
    z-index: 1;
  }
</style>
