import { ref } from "vue"

const store = {
  inputValue: ref(''),
  rightClickMenu: null,
  user: {},
  quoteValue: null,
  inputRef: ref(null),
  currentPage: ref(''),
}

export default store
