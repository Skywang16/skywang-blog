<script setup>

// 只缓存 home 路由
</script>

<template>
  <router-view v-slot="{ Component }">
    <keep-alive include="HomeView">
      <component :is="Component" />
    </keep-alive>
  </router-view>
</template>

<style scoped>
.logo {
  display: block;
  margin: 0 auto;
}
</style>
