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
        </template>
    </div>
</template>

<script type="text/ecmascript-6">
  import { getUri } from '../../config'

  export default{
    data () {
      return {
        projects: [],
        project_id: ''
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
        console.log(id)
      }
    }
  }
</script>