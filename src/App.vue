<script setup lang="ts">
  import * as themeJson from '@/assets/theme.json'
  import Menu from '@/components/Menu.vue'
  import {
    NConfigProvider,
    NLayout,
    NLayoutContent,
    NLayoutHeader,
  } from 'naive-ui'
  import 'vfonts/FiraSans.css'
</script>

<template>
  <NConfigProvider :theme-overrides="themeJson">
    <NLayout>
      <NLayoutHeader class="header">
        <Menu />
      </NLayoutHeader>
      <NLayoutContent class="content">
        <router-view />
      </NLayoutContent>
    </NLayout>
  </NConfigProvider>
</template>

<style lang="scss">
  #app {
    text-align: center;
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
