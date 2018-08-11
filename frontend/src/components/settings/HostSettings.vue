<template>
    <div id="settings">
        <el-collapse v-model="activeNames" @change="handleChange" :accordion="true">
            <el-collapse-item :title="item.name" :name="index" v-for="item,index in project_hosts" :key="item.id">

                <el-form  :model="item" class="demo-form-inline">
                    <el-row>
                        <el-col :span="8">HOST</el-col>
                        <el-col :span="8">说明 </el-col>
                    </el-row>
                        <el-form-item
                                v-for="(host, index) in item.hosts"
                                :key="host.id"
                        >
                            <el-row>
                                <el-col :span="8"> <el-input v-model="host.name"></el-input></el-col>
                                <el-col :span="8"> <el-input v-model="host.description"></el-input></el-col>
                                <el-col :span="3">  <el-button @click.prevent="removeHost(host)">-</el-button></el-col>
                            </el-row>

                        </el-form-item>

                    <el-form-item>
                    <el-button @click="addHost">+</el-button>
                    <el-button @click="saveHost(item)" type="primary">保存</el-button>
                    </el-form-item>

                </el-form>

            </el-collapse-item>

        </el-collapse>
    </div>
</template>

<script type="text/ecmascript-6">
  import { getUri } from '../../config'
  export default {
    name: 'HostSettings',
    data () {
      return {
        project_hosts: [],
        activeNames: 0
      }
    },
    methods: {
      removeHost (item) {
        let projectIndex = this.activeNames
        let index = this.project_hosts[projectIndex].hosts.indexOf(item)
        this.project_hosts[projectIndex].hosts.splice(index, 1)
      },
      addHost () {
        let projectIndex = this.activeNames
        let project =  this.project_hosts[projectIndex]
        let hostItem = {
          name: '',
          description: '',
          is_default: 0,
          project_id:project.id
        }
        this.project_hosts[projectIndex].hosts.push(hostItem)
      },
      saveHost(item){
        let uri = getUri('project_host','resource')
        let update_uri = uri+ item.id
        let validata = item.hosts.filter((item,index)=>{
          return item.name.trim() === '' || item.description.trim() ==='';
        })
        if(validata.length > 0){
          this.$message.error('不能为空')
          return
        }
        this.axios.put(update_uri,item).then((res)=>{
          let data = this.getApiData(res,'保存成功')
        },(res)=>{
          this.$message.error('保存失败')
        })
      },
      handleChange (name) {
        this.activeNames = name
      }
    },
    created(){
      let uri = getUri('project_host', 'resource')
      this.axios.get(uri).then((response) => {
        let data = this.getApiData(response)
        this.project_hosts = data
      }, (response) => {
      })
    }
  }
</script>

<style>


</style>