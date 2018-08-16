<template>
    <div id="importStateSettings">
        <template>
            <el-select v-model="project_id" placeholder="请选择项目" @change="changeProject">
                <el-option
                        v-for="item in projects"
                        :key="item.id"
                        :label="item.name"
                        :value="item.id">
                </el-option>
            </el-select>
    <el-row v-if="project_id">
            <el-form>
                <el-form-item label="MYSQL HOST">
                    <el-col :span="8"> <el-input v-model="form.host"></el-input></el-col>
                </el-form-item>

                <el-form-item label="PORT">
                    <el-col :span="8"> <el-input v-model="form.port"></el-input></el-col>

                </el-form-item>

                <el-form-item label="USERNAME">
                    <el-col :span="8"> <el-input v-model="form.username"></el-input></el-col>

                </el-form-item>

                <el-form-item label="PASSWORD">
                    <el-col :span="8"> <el-input v-model="form.password"></el-input></el-col>

                </el-form-item>
                <el-form-item label="DATABASES">
                    <el-col :span="8"> <el-input placeholder="项目数据库名称，用`,`隔开" v-model="form.databases"></el-input></el-col>

                </el-form-item>


                <el-form-item>
                    <el-button type="primary" @click="onSubmit">保存信息</el-button>
                    <el-button @click="readComment">读取注释</el-button>
                    <el-button @click="submitComment" :disabled="!new_comment_count">提交注释({{new_comment_count}})</el-button>
                </el-form-item>
            </el-form>

    </el-row>
        </template>
    </div>
</template>

<script type="text/ecmascript-6">
  import { getUri } from '../../config'

  export default{
    data () {
      return {
        projects: [],
        project_id: '',
        form: {
          host: '127.0.0.1',
          port: 3306,
          username: 'root',
          password:'',
          project_id: '',
          databases:''
        },
        comments: [],
        new_comments: [],
        new_comment_count: 0
      }
    },
    created () {
      this.loadData()
    },
    methods: {
      loadData(){
        let uri = getUri('project', 'user_project')
        this.axios.get(uri).then((response) => {
          let data = this.getApiData(response)
          if (data) {
            this.projects = data
          }
        }, (response) => {
        })
      },
      changeProject(id) {
        this.form.project_id = id
        let uri = getUri('db_config','resource')
        let query = {project_id : id}
        this.axios.get(uri,{params:query}).then(response=>{
          let data = this.getApiData(response)
          if (data) {
            this.form = data
          }else{
            this.form = {
              host: '127.0.0.1',
              port: 3306,
              username: 'root',
              password:'',
              databases:'',
              project_id: id
            }
          }
        },response=>{})
      },
      onSubmit(){
        let uri = getUri('db_config','resource')
        // 更新
        if(!this.form.id){
          this.axios.post(uri,this.form).then(response=>{
            this.getApiData(response,'保存成功')
          })
        } // 创建
        else{
          this.axios.put(uri + '/' + this.form.id,this.form).then(response=>{
            this.getApiData(response,'保存成功')
          })
        }
      },
      readComment() {
        let db_config = this.form
        let uri = getUri('db_comment','import')
        this.axios.post(uri,db_config).then(response=>{
          let data = this.getApiData(response)
          this.comments = data
          let uri2 = getUri('db_comment','getHash')
          let query = {
            project_id: this.project_id
          }
          this.axios.get(uri2,{params:query}).then(res=>{
            let data2 = this.getApiData(res)
            this.new_comments = this.comments.filter(v=>{
              return data2.indexOf(v.hash) === -1
            })
            this.new_comment_count = this.new_comments.length
            this.$message('读取完毕,新增字段注释' + this.new_comment_count + '条' )
          })
        },response=>{})
      },
      submitComment(){
        let uri = getUri('db_comment','storeMany')
        let data = {
          list: this.new_comments,
          project_id: this.project_id
        }
        this.axios.post(uri,data).then(response=>{
         let data = this.getApiData(response)
          if(data){
           this.new_comment_count = 0
            this.new_comments = []
            this.$message('提交成功')
          }
        },response=>{})
      }
    }
  }
</script>