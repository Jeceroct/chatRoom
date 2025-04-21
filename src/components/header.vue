<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <Setting />
  <div class="container">
    <h1>{{ roomName }}</h1>
    <div class="more" @click="showSetting">
      <el-icon><MoreFilled /></el-icon>
    </div>
  </div>
</template>

<script setup>
// import { leave } from '@/utils/leave'
import { ref } from 'vue'
import Setting from './setting.vue';

const roomName = ref('聊天室')

const showSetting = () => {
  const settingContainer = document.querySelector('.settingContainer');
  settingContainer.classList.add('show');
}
</script>

<style scoped>
.container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background-color: var(--color-background);
  position: absolute;
  top: 0;
  left: 0;
  width: calc(100% - 20px);
  z-index: 1000;
  height: 2em;

  h1 {
    color: var(--color-text);
    font-size: 1.2em;
    font-weight: bold;
  }

  .more {
    position: absolute;
    right: 0.5em;
    color: var(--color-text);
    font-size: 1.5em;
    padding: 0.5em;
    height: 1em;
    width: 1em;
    cursor: pointer;
    border-radius: 50%;
    transition: all 0.2s ease-in-out;

    &:hover {
      background-color: var(--color-hover);
      color: var(--color-hover-text);
    }
  }
}
</style>