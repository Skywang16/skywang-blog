<script setup>

console.log(import.meta.env.BASE_URL)
//import { RouterLink, RouterView } from 'vue-router'
// import HelloWorld from './components/HelloWorld.vue'


// import { onMounted } from 'vue'
// function savetoken() {
//   // 保存token
//   localStorage.setItem('token', 'hdhsdsdsiweiwjijwei@sada')
//   localStorage.setItem('admin', 'admin')
//   // console.log('获取token',localStorage.getItem('token'))
//   // console.log('获取token',localStorage.getItem('admin'))
// }

// onMounted(savetoken);
</script>

<template>
  <!-- <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header> -->
  <RouterView />
</template>
<style scoped>
.logo{
  display: block;
  margin: 0 auto;
}

</style>
