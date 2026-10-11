<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
  <div class="not-found">
    <div class="glass-card card">
      <h1>404</h1>
      <p>Page not found</p>
      <RouterLink to="/" class="btn-primary link">Back home</RouterLink>
    </div>
  </div>
</template>

<style scoped>
.not-found {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}
.card {
  padding: 48px;
  text-align: center;
  display: flex;
  flex-direction: column;
  gap: 12px;
  align-items: center;
}
h1 { font-size: 4rem; margin: 0; }
.link { text-decoration: none; }
</style>
