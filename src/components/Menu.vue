<script setup lang="ts">
  import logo from '@/assets/logo.svg'
  import { NDivider } from 'naive-ui'
</script>

<template>
  <div class="nav">
    <router-link
      style="text-decoration: none; color: inherit"
      :to="{ name: 'Home' }"
    >
      <div class="n-text logo">
        <img :src="logo" />
        <span>Teaset</span>
      </div>
    </router-link>
    <div />
  </div>
  <NDivider class="nav-divider" />
</template>

<style lang="scss">
  .nav {
    padding: 0.5rem 1rem;
  }
  .logo {
    padding: 0;
    display: flex;
    align-items: center;
    justify-content: flex-start;
    img {
      height: 2.5rem;
      margin-top: auto;
      margin-bottom: auto;
    }
    span {
      margin-left: 0.5rem;
      font-size: 2rem;
      font-weight: 400;
    }
  }
  .nav-divider {
    margin: 0 !important;
  }
</style>
