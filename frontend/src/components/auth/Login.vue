<template>
    <div class="login">


        <el-row type="flex" class="row-bg" justify="center">


            <el-col :span="8" class="form">
                <el-form label-width="80px" :model="form">

                    <el-form-item label="邮箱">
                        <el-input v-model="form.username"></el-input>
                    </el-form-item>
                    <el-form-item label="密码">
                        <el-input v-model="form.password" type="password"></el-input>
                    </el-form-item>
                    <el-button type="primary" @click="onSubmit()">登录</el-button>

                </el-form>
            </el-col>


        </el-row>

    </div>
</template>

<script type="text/ecmascript-6">
    export default {
        data() {
            return {
                form: {
                    username: 'zhurkai@163.com',
                    password: '123456',
                    grant_type: 'password',
                    client_id: 2,
                    client_secret: 'jaRRCCNvPujoAM8cgLv4rjbpnmjmGfETze6CE1Ir',
                }
            }
        },
        methods: {
            onSubmit() {
                this.axios.post('/api/oauth/token',this.form).then( (response)=> {
                    // 记录token
                    localStorage.setItem('access_token',response.data.access_token)
                    localStorage.setItem('token_type',response.data.token_type)
                    localStorage.setItem('refresh_token',response.data.refresh_token)
                    let expire_time =Math.round(new  Date().getTime()/1000) + response.data.expires_in
                    localStorage.setItem('expire_time',expire_time)
                    // 记录用户信息
                    this.axios.get('api/api/user-info').then((response)=>{
                       let  user = response.data.data
                        localStorage.setItem('user_id',user.id)
                        localStorage.setItem('user_name',user.name)
                        localStorage.setItem('user_email',user.email)
                      this.$store.commit('userInfo', user);
                        let redirect = decodeURIComponent(this.$route.query.redirect || '/');
                          this.$router.push({path: redirect})
                    }).catch((response)=>{
                        this.$message.error('获取用户信息失败')
                    })
                },(response)=>{
                    console.log(response)
                    this.$message.error(response.body.message)
                })

            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .form {
        margin-top: 200px;
        /*border:1px #BFCBD9 solid;*/

    }
</style>
