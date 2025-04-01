import router from '@/router'

export const leave = (address) => {
  const routeMask = document.querySelector('#routeMask')
  routeMask.classList.add('leave')
  routeMask.classList.remove('waiting')
  setTimeout(() => {
    router.push(address)
  }, 200)
}