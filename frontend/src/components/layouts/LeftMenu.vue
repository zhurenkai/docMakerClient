<template>
    <div id="leftMenu">
        <!--菜单栏-->

        <!--<el-menu default-active="2" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose">-->
        <el-menu default-active="1" class="el-menu-vertical-demo">

            <el-submenu index="1">
                <template slot="title"><i class="el-icon-message"></i>项目</template>


                <el-submenu :index="project.name" v-for="project,index in projects" :key="index">
                    <template slot="title"> {{ project.name }}</template>
                    <!--<el-menu-item index="1-4-1" v-for="module in project.modules" >{{ module.title }}</el-menu-item>-->

                    <el-submenu :index="module.name" v-for="module,index in project.modules" :key="index">

                        <template slot="title">{{ module.name }}</template>

                        <el-menu-item style="padding-left: 0px; " :index="api.name" v-for="api,index in module.apis" :key="index"  @click="addTab(api.id)">
                            <el-row>
                                <el-col :span="4">
                                    <span class="request_method" :style="{ color: activeColor(api.method)}">{{ api.method }} </span>
                                </el-col>
                                <el-col :span="19">
                                    {{ api.name }}
                                </el-col>
                                <el-col :span="1"><i class="el-icon-circle-close" @click="removeApiConfirm()"></i></el-col>

                                <el-dialog title="提示" :visible.sync="removeApiBox">
                                    <span>确定删除?</span>
                                    <span slot="footer" class="dialog-footer">
    <el-button @click="removeApiBox = false">取 消</el-button>
    <el-button type="primary" @click="removeApi(project,module,api)">确 定</el-button>
  </span>
                                </el-dialog>

                            </el-row>
                        </el-menu-item>
                        <el-menu-item index="0" style="padding-left: 50px" @click="addApi(project,module)"><i class="el-icon-plus"></i>接口</el-menu-item>
                    </el-submenu>
                    <el-menu-item index="0" style="padding-left: 50px" @click="addModule(project)"><i class="el-icon-plus"></i>模块</el-menu-item>

                </el-submenu>
                <el-menu-item index="0" style="padding-left: 50px" @click="addProject(projects)"><i class="el-icon-plus"></i>项目</el-menu-item>

            </el-submenu>

            <el-menu-item index="2"><i class="el-icon-menu"></i>导航二</el-menu-item>
            <el-menu-item index="3"><i class="el-icon-setting"></i>导航三</el-menu-item>
        </el-menu>

 <!--添加-->



        <el-dialog title="添加" :visible.sync="addBox">
            <el-form :model="form">
                <el-form-item label="名称" >
                    <el-input v-model="form.name"  auto-complete="off"></el-input>
                </el-form-item>
                <el-form-item label="简介" >
                    <el-input v-model="form.description"  auto-complete="off"></el-input>

                </el-form-item>


            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button >取 消</el-button>
                <el-button type="primary" @click="submitForm">确 定</el-button>
            </div>
        </el-dialog>

    </div>


</template>

<script type="text/ecmascript-6">
    import {getUri} from '../../config/config.js'
    export default{
        data:function () {
            return {
                projects: [],
                removeApiBox: false,
                pushIndex: {
                    project:false,
                    module:false
                },
                addBox: false,
                form:{
                    name:'',
                    description:''
                }
            }
        },
        methods:{
            getProjects:function () {
                let uri = getUri('project','user_project')
                this.axios.get(uri).then((response)=>{
                    let projects = response.data.data
                    this.projects = projects
                })
            },
            activeColor:function (method) {
                let color
               switch (method){
                   case 'POST' :
                       color = 'orange'
                       break
                   case 'GET' :
                       color = 'green'
                       break;
                   case 'DELETE':
                       color = 'red'
                       break;
                   case 'PUT':
                       color = 'blue'
                       break;
               }
               return color
            },
            removeApi: function(project,module,api){
                this.removeApiBox =false
                let project_index = this.projects.indexOf(project)
                let module_index = project.modules.indexOf(module)
                let api_index = module.apis.indexOf(api)
                let uri = getUri('api','resource')
                this.axios.delete(uri + api.id).then((response)=>{
                    if(!response.data.code){
                        this.projects[project_index].modules[module_index].apis.splice(api_index,1)
                    }else{
                        this.$message.error(response.data.message)
                    }
                })
            },
            removeApiConfirm(){
                    this.removeApiBox = true
            },
            addProject(){
                this.addBox = true
                this.uri = getUri('project','user_project')
            },
            addModule(project){
                this.form.project_id = project.id
                this.addBox = true
                this.uri = getUri('module','resource')
                this.pushIndex.project = this.projects.indexOf(project)
            },
            addApi(project,module){
                this.form.project_id = project.id
                this.form.module_id = module.id
                this.addBox = true
                this.uri = getUri('api','resource')
                this.pushIndex.project = this.projects.indexOf(project)
                this.pushIndex.module = project.modules.indexOf(module)
            },
            submitForm(){
                console.log(this.form)
                this.axios.post(this.uri,this.form).then((response)=>{
                    let data = this.getApiData(response)
                    if(data){
                        let proIndex = this.pushIndex.project
                        let modIdex  = this.pushIndex.module
                        if( modIdex !==false){
                            this.projects[proIndex].modules[modIdex].apis.push(data)
                        }else if(proIndex!==false){
                            this.projects[proIndex].modules.push(data)
                        }else{
                            this.projects.push(data)
                        }
                       this.pushIndex=  {
                            project:false,
                                module:false
                        }
                        this.addBox=  false

                    }
                },(response)=>{
                    this.$message.error('创建失败')
                })
            },
            addTab(apiId){
                let uri = getUri('api','resource')
                this.axios.get(uri+'/'+apiId).then((response)=>{
                    let data = this.getApiData(response)
                    this.$store.commit('addTab',data);
                },(response)=>{})
            }
        },
        created:function () {
            this.getProjects()
        }
    }
</script>

<style>
    .request_method{
        font-size: 7px;
        font-weight: 800;

    }
    .el-icon-circle-close{color: red}
</style>