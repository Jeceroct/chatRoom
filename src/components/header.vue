<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="container">
    <h1>{{ roomName }}</h1>
    <div class="more">
      <el-icon><MoreFilled /></el-icon>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const roomName = ref('望子成龙小学')
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px;
  background-color: #1f1f1f;
  position: absolute;
  top: 0;
  left: 0;
  width: calc(100% - 20px);
  z-index: 1000;
  height: 4em;

  h1 {
    color: #cecece;
    font-size: 1.5em;
    font-weight: bold;
  }

  .more {
    position: absolute;
    right: 0.5em;
    color: #cecece;
    font-size: 1.5em;
    padding: 0.5em;
    height: 1em;
    width: 1em;
    cursor: pointer;
    border-radius: 50%;
    transition: all 0.2s ease-in-out;

    &:hover {
      background-color: #cecece;
      color: #1f1f1f; 
    }
  }
}
</style>