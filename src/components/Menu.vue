<script setup lang="ts">
import logo from '../assets/logo.svg';
</script>

<template>
    <router-link style="text-decoration: none; color: inherit;" to="/">
        <div class="n-text logo">
            <img :src="logo" />
            <span>Teaset</span>
        </div>
    </router-link>
    <div></div>
</template>

<style lang="scss">
.menu {
    align-items: left;
}
.logo {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    margin-left: 0.5rem;
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
</style>
