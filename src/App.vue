<script setup lang="ts">
import { NLayout, NLayoutContent, NLayoutHeader } from 'naive-ui';
import 'vfonts/FiraSans.css';
import Menu from './components/Menu.vue';
</script>

<template>
  <NLayout>
    <NLayoutHeader>
      <Menu />
    </NLayoutHeader>
    <NLayoutContent>
      <router-view></router-view>
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
</style>
