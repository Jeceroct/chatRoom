<template>
  <div class="myInput">
    <slot name="pre" id="pre" class="pre"></slot>
    <textarea v-model="value" class="input"></textarea>
    <span style="visibility: hidden; height: 0; position: fixed; top: 0;">{{ value }}</span>
    <slot name="end" id="end" class="end"></slot>
  </div>
</template>

<script setup>
import { defineProps, onMounted, defineExpose, defineEmits, computed } from 'vue';

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
})

const value = computed({
  get() {
    return props.modelValue;
  },
  set(val) {
    emit('update:modelValue', val);
  }
})

const emit = defineEmits(['update:modelValue', 'enter', 'paste', 'submit'])

var myInput
var textarea
var myInputHeight
var p

const focus = () => {
  textarea.focus();
}

let blankWidth = 0;
let lines = 1;
// let oldLines = lines;
// let oldPWidth;
function getRowCount() {
  lines = Math.floor(p.clientWidth / (textarea.clientWidth - lines * blankWidth)) + 1
  lines += textarea.value.split(/\r?\n/).length - 1;
  // 消除行尾误差
  // if (lines != oldLines) {
  //   blankWidth = textarea.clientWidth * oldLines - oldPWidth;
  //   oldLines = lines;
  //   console.log(lines, oldLines, blankWidth);
  // }
  // oldPWidth = p.clientWidth;
  return lines;
}

const updateHeight = () => {
  textarea.style.height = getRowCount() * parseFloat(getComputedStyle(textarea).lineHeight) + 'px';
  myInput.style.height = getRowCount() == 1 ? myInputHeight + 'px' : textarea.clientHeight + 10 + 'px';
}

defineExpose({
  focus,
  updateHeight
})

onMounted(() => {
  myInput = document.querySelector('.myInput');
  myInputHeight = myInput.clientHeight;
  textarea = myInput.querySelector('textarea');
  p = myInput.querySelector('span');
  textarea.addEventListener('input', updateHeight)
  textarea.addEventListener('focus', updateHeight)
  // watch(value, updateHeight)

  // 监听回车事件
  textarea.addEventListener('keydown', (e) => {
    if (e.key == 'Enter' && e.ctrlKey) {
      e.preventDefault();
      textarea.value += '\n';
      setTimeout(() => {
        updateHeight()
      }, 0);
    }
    else if (e.key == 'Enter') {
      e.preventDefault();
      document.querySelector('#inputForm').submit();
      emit('enter', e);
      emit('submit', e);
    }
  })
  // 监听粘贴事件
  textarea.addEventListener('paste', (e) => {
    e.preventDefault();
    emit('paste', e);
  })
})
</script>

<style scoped>
.myInput {

  display: flex;
  align-items: center;
  justify-content: space-around;

  background-color: #fff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);

  transition: all 0.1s;

  padding: 0 10px;

  .pre,
  .end {
    width: 5%;
  }

  p,
  textarea {
    font-size: 1.2em;
    font-family: '微软雅黑', '黑体';
  }

  .input {
    width: 90%;
    height: 1.5em;
    background-color: transparent;
    border: none;
    outline: none;
    resize: none;

    margin: 5px 0;

    display: flex;
    align-items: center;

    line-height: 1.5em;

    &::-webkit-scrollbar {
      display: none;
    }
  }
}
</style>