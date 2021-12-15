<script setup lang="ts">
  import logo from '@/assets/logo.svg'
  import { NH1 } from 'naive-ui'
</script>

<template>
  <div class="hero">
    <img :src="logo" alt="teaset logo" class="hero-logo" />
    <NH1>Teaset</NH1>
  </div>
</template>

<style lang="scss">
  .hero {
    margin: 4rem 0;
    display: flex;
    justify-content: center;
    padding: 0 1rem;
    flex-wrap: wrap;
    &-logo {
      max-height: 15vw;
      width: 100%;
    }
  }
</style>
