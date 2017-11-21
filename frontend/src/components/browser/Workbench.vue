<template>
<div id='workbentch'>

    <el-row>
  <el-col :span="3"><div class="grid-content bg-purple">


  </div></el-col>

  <el-col :span="21"><div class="grid-content bg-purple-light">
<el-form :model="Form"  ref="Form">
<el-row>

    <el-col :span="3">
        <!--Method选项-->

        <el-select v-model="Form.method" placeholder="请选择">

            <el-option
                    v-for="item in methods"
                    :key="item.index"
                    :value="item">
            </el-option>
        </el-select>

    </el-col>

  <el-col :span="14"><div class="grid-content bg-purple-light">
    <el-input v-model="Form.request_uri" placeholder="URI" @change="checkUriParams()">

        <!--host-->

        <el-select slot="prepend" class="host-select"
                   v-model="Form.request_host"
                   filterable
                   allow-create
        >
            <el-option
                    v-for="item in project_hosts"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
            </el-option>
        </el-select>


        <!--param按钮-->
        <el-button slot="append" v-model="query_params_show" @click="checkUriParams();query_params_show = !query_params_show">params</el-button>

    </el-input>


  </div></el-col>

  <el-col :span="2"><div class="grid-content bg-purple-light">
  <el-button type="primary" @click="submitForm('Form')">Send</el-button>
  </div></el-col>
    <el-col :span="2"><el-button @click="createDocument()">生成文档</el-button></el-col>
</el-row>

    <query-param v-if="query_params_show" v-for="(param_info,index) in uriParamsArr"
            :keyState="param_info.keyStateSearch"
                :k="param_info.key"
                :v="param_info.value"
                :required="param_info.required"
                :key_state="param_info.keyState"
                :key="index"
                :index="index"
                @selfUpdate="changeParam"
    ></query-param>

<div v-for="param in getParams">
<el-row>
  <el-col :span="3">
    <el-input v-model="param.k"></el-input>
  </el-col>
  <el-col :span="5">
  <el-input></el-input>
  </el-col>

  <el-col :span="6">
  <el-input v-model="param.v"></el-input>
  </el-col>
</el-row>
</div>

<el-form-item
  v-for="(domain, index) in Form.dataList"
  :key="domain.key"
  :prop="'dataList.' + index + '.v'"
>
  <el-row>
    <el-col :span="2"><div class="grid-content bg-purple"></div>
    <el-switch
    v-model="domain.use"
    on-color="#13ce66"
    off-color="#ff4949">
  </el-switch>
    </el-col>
    <el-col :span="2"><div class="grid-content bg-purple"></div>
  <el-checkbox v-model="domain.required">必填</el-checkbox>
    </el-col>
    <el-col :span="5"><div class="grid-content bg-purple"></div>
    <el-input v-model="domain.k"></el-input>

    </el-col>
    <!--<el-col :span="3"><div class="grid-content bg-purple"></div>-->

    <!--<el-select-->
  <!--v-model="domain.keyState"-->
  <!--filterable-->
  <!--allow-create-->
  <!--&gt;-->
  <!--<el-option-->
    <!--v-for="item in getKeyState()"-->
    <!--:key="item.index"-->
    <!--:value="item">-->
  <!--</el-option>-->
<!--</el-select>-->

    <!--</el-col>-->
    <el-col :span="9"><div class="grid-content bg-purple"></div>

    <el-date-picker
      v-model="domain.v"
      v-if="domain.value_type=='time'"
      type="datetime"
      placeholder="选择日期时间"
      format="yyyy-MM-dd HH:mm:ss"
      @change="setTime(domain)"
      >
    </el-date-picker>
<el-input v-model="domain.v" v-if="domain.value_type=='text'" ></el-input>
<input type="file" @change.prevent="getFile($event, domain)"  multiple="multiple" v-if="domain.value_type=='file'">

    </el-col>
    <el-col :span="2"><div class="grid-content bg-purple"></div>

    <el-select v-model="domain.value_type" @change="clearValue(domain)">
        <el-option
          v-for="item in value_type"
          :key="item.index"
          :value="item">
        </el-option>
      </el-select>

    </el-col>
    <el-col :span="1"><div class="grid-content bg-purple"></div>
    <el-button @click.prevent="removeDomain(domain)"><i class="el-icon-minus"></i></el-button>

    </el-col>
  </el-row>

</el-form-item>


  <el-form-item>
    <el-button @click="addDomain('text')"><i class="el-icon-plus"></i> text</el-button>
    <el-button @click="addDomain('file')"><i class="el-icon-plus"></i> file</el-button>
    <el-button @click="addDomain('time')"><i class="el-icon-plus"></i> time</el-button>


  </el-form-item>
</el-form>
</div></el-col>

</el-row>
<el-col v-loading="request_loading" v-if="jsonData">
    <html-json  :jsonObj="jsonData" class="json-data" ></html-json>
</el-col>


    <el-dialog  title="文档预览" :visible.sync="previewBox">

        <el-row :gutter="20" style="text-align: left ;font-size: large">
            <el-col :span="2" size="large"><el-tag>{{ Form.method }}</el-tag> </el-col>
            <el-col :span="22" ><span style="color:red">{host}</span> {{ Form.request_uri }}</el-col>
        </el-row>

        <!--参数说明-->
        <el-table
                :data="document.params" border stripe
                style="width: 100%">

            <el-table-column scope="scope"
                    label="参数"
                    width="180">
                <template scope="scope">
                    {{ scope.row.key }}
                </template>
            </el-table-column>

            <el-table-column scope="scope"
                             label="说明"
                             width="180">
                <template scope="scope">

                    <el-select
                            v-model="scope.row.statement"
                            filterable
                            allow-create>
                        <el-option
                                v-for="item in scope.row.statement_options"
                                :label="item.statement"
                                :value="item.statement">
                        </el-option>
                    </el-select>


                </template>
            </el-table-column>

            <el-table-column scope="scope"
                             label="类型"
                             width="100">
                <template scope="scope">
                    {{ scope.row.type }}
                </template>
            </el-table-column>
            <el-table-column scope="scope"
                             label="必填"
                             width="80">
                <template scope="scope" >
                    <span style="cursor:pointer" @click="scope.row.required = !scope.row.required"> {{ scope.row.required ? "是":"否" }}</span>

                </template>
            </el-table-column>

            <el-table-column scope="scope"
                             label="备注"
                             width="180">
                <template scope="scope">
                    <el-input v-model="scope.row.remark" style="border:none" type="textarea" autosize></el-input>

                </template>
            </el-table-column>
        </el-table>

<!--返回值-->
        <el-table
                :data="document.response" border stripe
                style="width: 100%">
            <el-table-column scope="scope"
                             label="标识"
                             width="180">
                <template scope="scope">
                   {{ scope.row.key }}

                </template>
            </el-table-column>
<!--类型-->
            <el-table-column scope="scope"
                             label="类型"
                             width="180">
                <template scope="scope">
                    {{ scope.row.type }}

                </template>
            </el-table-column>

            <!--说明-->
            <el-table-column scope="scope"
                             label="说明"
                             width="180">
                <template scope="scope">

                    <el-select
                            v-model="scope.row.statement"
                            filterable
                            allow-create>
                        <el-option
                                v-for="item in scope.row.statement_options"
                                :label="item.statement"
                                :value="item.statement">
                        </el-option>
                    </el-select>


                </template>
            </el-table-column>

        </el-table>

<!--注释-->
        <el-row>
                <el-col :span="8" :offset="16">

                    <el-dropdown>
  <span class="el-dropdown-link">
    下载<i class="el-icon-caret-bottom el-icon--right"></i>
  </span>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item>markdown</el-dropdown-item>
                            <el-dropdown-item>swagger</el-dropdown-item>

                        </el-dropdown-menu>
                    </el-dropdown>

                    <el-button  type="primary">保存</el-button>
                </el-col>
            </el-row>
        <el-row >

            <el-col :span="4" :offset="20" style=" background: #d3dce6; on-hover:pointer" >
                <span v-clipboard:copy="commentaries"
                      v-clipboard:success="onCopy" style="cursor:pointer"
            >复制到剪切板</span></el-col>
        </el-row>
    <el-row >
        <el-col style="text-align: left;font-size: large">
            <el-input
                    type="textarea"
                    autosize
                    placeholder="请输入内容"
                    v-model="commentaries">
            </el-input>
        </el-col>
    </el-row>


    </el-dialog>


</div>
</template>



<script type="text/ecmascript-6">
  import QueryParam from './QueryParam.vue'
  import VueHtmlJson from 'vue-html-json'
  import {getUri} from '../../config/config.js'
  export default {
    data () {
      return {
        methods: ['GET', 'POST', 'PUT', 'DELETE'],
        value_type: ['text', 'file', 'time'],
        getParams: [],
        previewBox:false,
        previewLoading:false,
        Form: {
          dataList: [{
            k: '',
            v: '',
            use: true,
            required: true,
            value_type: 'file',
            keyState: '',
            file: ''
          }],
          method: 'GET',
          request_uri: '/party?id=aa&title=bb',
          request_host: 'https://www.randqun.com'
        },
        uriParamsArr: {},
        searchQuery: '',
        jsonData: false,
        request_loading: false,
        query_params_show: false,
        project_id: 1,
        project_hosts: [],
        document:{},
        commentaries:'',
          api_name:'测试API'
      }
    },
    methods: {
        onCopy: function (e) {
           this.$message({ message: '复制成功',
               type: 'success'})
        },
      submitForm (event) {
        this.request_loading = true
        this.$refs[event].validate((valid) => {
          if (valid) {
            let formData = new FormData()
            let form = this.Form
            for (let index in form.dataList) {
              let obj = form.dataList[index]
              if (obj.value_type === 'file') {
                for (let i = 0; i < obj.v.length; i++) {
                  formData.append(obj.k, obj.v[i])
                }
              } else {
                formData.append(obj.k, obj.v)
              }
            }
            formData.append('request_uri', this.Form.request_uri)
            formData.append('request_host', this.Form.request_host)
            formData.append('request_method', this.Form.method)
            this.$http.post('client-api/client/request', formData).then((response) => {
              this.request_loading = false
              this.jsonData = response.body
            }, (response) => {
                 // error callback
            })
          } else {
            console.log('error submit!!')
            return false
          }
        })
      },
      getProjectHosts () {
          let project_id = this.project_id
          let host_list = []
         this.axios.get('api/api/project-host/'+project_id).then((response)=>{
             if(response.data.code){
                 this.$message.error(response.data.error)
             }
            let hosts = response.data.data
            hosts.forEach((host)=>{
                let project_host = {
                    value: host.name,
                    label: host.description
                }
                this.project_hosts.push(project_host)
            })

         })
      },
      resetForm (formName) {
        this.$refs[formName].resetFields()
      },
      removeDomain (item) {
        let index = this.Form.dataList.indexOf(item)
        if (index !== -1) {
          this.Form.dataList.splice(index, 1)
        }
      },
      getFile (event, domain) {
        let index = this.Form.dataList.indexOf(domain)
        this.Form.dataList[index].v = event.target.files
      },
      setTime (event) {
        let index = this.Form.dataList.indexOf(event)
        this.Form.dataList[index].v = event.v.toJSON()
      },
      clearValue (domain) {
        console.log(domain)
        let index = this.Form.dataList.indexOf(domain)
        this.Form.dataList[index].v = ''
      },
      checkUriParams () {
        let uri = this.Form.request_uri
        let paramArr = []
        let start = uri.indexOf('?')
        if (start !== -1) {
          let str = uri.substr(start + 1)
          let strs = []
          strs = str.split('&')
          for (let i = 0; i < strs.length; i++) {
            let key = strs[i].split('=')[0]
            let value = ''
            value = strs[i].split('=')[1]
            paramArr.push({
              key: key,
              value: value,
              required: true
            })
          }
        }
        this.uriParamsArr = paramArr
      },
      addDomain (type) {
        this.Form.dataList.push({
          v: '',
          k: '',
          use: true,
          required: true,
          value_type: type,
          keyState: ''
        })
//        console.log(this.Form)
      },
      changeParam (val) {
        this.uriParamsArr[val.index][val.attr] = val.value
        let uriParamsObjArr = this.uriParamsArr
        console.log(uriParamsObjArr)
        let paramArr = []
        for (let index in uriParamsObjArr) {
          paramArr.push(uriParamsObjArr[index].key + '=' + uriParamsObjArr[index].value)
        }
        let newUriParam = paramArr.join('&')
        console.log(newUriParam)
        let start = this.Form.request_uri.indexOf('?')
        let str = this.Form.request_uri.substr(0, start)
        let newUri = str + '?' + newUriParam
        console.log(newUri)
        this.Form.request_uri = newUri
      },
      createDocument () {
            this.checkUriParams()
          this.previewBox = true
            let data = {
                uri:this.Form.request_uri,
                method: this.Form.method,
                queryParam:this.uriParamsArr,
                formParam: this.Form.dataList,
                response:this.jsonData
          }
          let uri = getUri('doc','generate')
          this.previewLoading = true
          this.axios.post(uri,data).then((response)=>{
                this.document = response.data.data
              this.commentaries = "/**\n *"+this.api_name + "\n"
               this.document.params.forEach((param)=>{
                    this.commentaries +=' *@param '+ param.type+ ' '+param.key + ' ' +param.statement+'\n'
               })
              this.commentaries +='*/'
              this.previewLoading = false
          },(response)=>{
              this.$message.error('生成出错');
          })
      }
    },
    components: {
      QueryParam, [VueHtmlJson.name]: VueHtmlJson
    }
  }
</script>
<style scoped>
    .json-data{text-align: left}
    .host-select{
        width: 210px;
    }
    </style>