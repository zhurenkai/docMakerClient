<template>
  <div id="app">

<el-row>

    <router-view name="top_menu"></router-view>
</el-row>

<el-row>
    <el-col :span="4">

<router-view  name="left_menu"></router-view>
    </el-col>

<el-col :span="20">
    <router-view></router-view>

</el-col>
  <!--  <img src="./assets/logo.png"> -->


</el-row>


  </div>
</template>

<script>
export default {
  name: 'app',
  data(){
      return {
          user: false
      }
  },
    mouted(){
      this.getUserInfo
    },
    methods:{
      getUserInfo:function () {
          let user = {}
          user.id = localStorage.getItem('user_id')
          user.name = localStorage.getItem('user_name')
          user.email = localStorage.getItem('user_email')
          this.user = user
          console.log(user)
      }
    }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 10px;
}
</style>
