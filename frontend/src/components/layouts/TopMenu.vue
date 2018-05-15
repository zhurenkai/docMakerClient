<template>
    <div id="topMenu">

        <el-menu
                class="el-menu-demo"
                mode="horizontal"
                background-color="#545c64"
                text-color="#fff"
                menu-trigger="click"
                active-text-color="#ffd04b">
            <el-row>
                <el-col :offset="20">
                    <el-submenu index="2">
                        <template slot="title">{{user.name}}</template>
                        <el-menu-item index="2-1" @click="logout">退出</el-menu-item>
                    </el-submenu>

                </el-col>
            </el-row>

            </el-menu>

    </div>
</template>

<script type="text/ecmascript-6">
  import {mapState} from 'vuex'
  import { getUri } from '../../config'
    export default{
      computed:{
        ...mapState({
          userInfo: state => state.userInfo
        }),
      },
      data () {
        return {
          user: {}
        }
      },
      methods: {
        logout() {
          localStorage.clear()
          this.$router.push('login')
        },
        getUserInfo(){
          if(this.$store.state.userInfo.name){
            this.user = this.$store.state.userInfo
          }else{
            let uri = getUri('user','info')
            this.axios.get(uri).then(response=>{
              let  user = response.data.data
              this.$store.commit('userInfo', user);
              this.user = user
            })
          }
        }
      },
      created () {
        this.getUserInfo()
      }
    }
</script>
<style>
    #topMenu{
        background-color: #1f2d3d;
        width:100%;
        height:100%;
        line-height: 60px;
    }
    .title{
        color: white;
        font-size: 30px;
    }
    .user{
        color: white;
        font-size: 24;
    }
</style>