<template>
    <div id='workbench'>

        <el-row>
            <el-col :span="3">
                <div class="grid-content bg-purple">


                </div>
            </el-col>

            <el-col :span="21">
                <div class="grid-content bg-purple-light">
                    <el-form :model="Form" ref="Form">
                        <el-row>

                            <el-col :span="3">
                                <!--Method选项-->

                                <el-select v-model="currentApi.method" placeholder="请选择">

                                    <el-option
                                            v-for="item in methods"
                                            :key="item.index"
                                            :value="item">
                                    </el-option>
                                </el-select>

                            </el-col>

                            <el-col :span="14">
                                <div class="grid-content bg-purple-light">
                                    <el-input v-model="currentApi.path" placeholder="URI" @change="checkUriParams()">

                                        <!--host-->

                                        <el-select slot="prepend" class="host-select"
                                                   v-model="currentApi.host"
                                                   filterable
                                                   allow-create
                                                   @change="changeHost"
                                        >
                                            <el-option
                                                    v-for="item in currentApi.hosts"
                                                    :key="item.name"
                                                    :label="item.label"
                                                    :value="item.name">
                                            </el-option>
                                        </el-select>


                                        <!--param按钮-->
                                        <el-button slot="append" v-model="query_params_show"
                                                   @click="checkUriParams();query_params_show = !query_params_show">
                                            params
                                        </el-button>

                                    </el-input>


                                </div>
                            </el-col>

                            <el-col :span="2">
                                <div class="grid-content bg-purple-light">
                                    <el-button type="primary" @click="submitForm('Form')">Send</el-button>
                                </div>
                            </el-col>
                            <el-col :span="2">
                                <el-button @click="saveApi">save</el-button>
                            </el-col>
                            <el-col :span="2">
                                <el-button @click="createDocument()">生成文档</el-button>
                            </el-col>
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

                        <el-row>
                            <el-col :span="24">

                                <el-tabs v-model="tabSelected">
                                    <el-tab-pane label="form" name="form" :disabled="currentApi.method ==='GET'">
                                <span slot="label"><el-radio :disabled="currentApi.method ==='GET'" v-model="bodyType" label="form">form</el-radio></span>
                                        <el-form-item
                                                v-for="(item, index) in currentApi.params"
                                                :key="item.id"

                                        >
                                            <el-row>
                                                <el-col :span="2">
                                                    <div class="grid-content bg-purple"></div>
                                                    <el-switch
                                                            v-model="item.is_use"
                                                            on-color="#13ce66"
                                                            off-color="#ff4949">
                                                    </el-switch>
                                                </el-col>
                                                <el-col :span="2">
                                                    <div class="grid-content bg-purple"></div>
                                                    <el-checkbox v-model="item.required">必填</el-checkbox>
                                                </el-col>
                                                <el-col :span="5">
                                                    <div class="grid-content bg-purple"></div>
                                                    <el-input v-model="item.name"></el-input>

                                                </el-col>

                                                <el-col :span="9">
                                                    <div class="grid-content bg-purple"></div>

                                                    <el-date-picker
                                                            v-model="item.value"
                                                            v-if="item.type=='time'"
                                                            type="datetime"
                                                            placeholder="选择日期时间"
                                                            format="yyyy-MM-dd HH:mm:ss"
                                                            @change="setTime(item)"
                                                    >
                                                    </el-date-picker>
                                                    <el-input v-model="item.value"
                                                              v-if="item.type!='file' && item.type!='time'"></el-input>
                                                    <input type="file" @change.prevent="getFile($event, item)"
                                                           multiple="multiple" v-if="item.type=='file'">

                                                </el-col>
                                                <el-col :span="2">
                                                    <div class="grid-content bg-purple"></div>

                                                    <el-select v-model="item.type" @change="clearValue(item)">
                                                        <el-option
                                                                v-for="t in type"
                                                                :key="t.index"
                                                                :value="t">
                                                        </el-option>
                                                    </el-select>

                                                </el-col>
                                                <el-col :span="1">
                                                    <div class="grid-content bg-purple"></div>
                                                    <el-button @click.prevent="removeParam(item)"><i
                                                            class="el-icon-minus"></i></el-button>

                                                </el-col>
                                            </el-row>

                                        </el-form-item>
                                        <el-form-item>
                                            <el-button @click="addParam('text')"><i class="el-icon-plus"></i> text
                                            </el-button>
                                            <el-button @click="addParam('file')"><i class="el-icon-plus"></i> file
                                            </el-button>
                                            <el-button @click="addParam('time')"><i class="el-icon-plus"></i> time
                                            </el-button>

                                        </el-form-item>

                                    </el-tab-pane>

                                    <el-tab-pane label="json" name="json" :disabled="currentApi.method ==='GET'">
                                       <span slot="label">
                                           <el-radio :disabled="currentApi.method ==='GET'" v-model="bodyType" label="json" @change="bodyTypeChange">json</el-radio>
                                       </span>
                                        <el-input
                                                type="textarea"
                                                :rows="12"
                                                v-model="currentApi.json_input"
                                                @blur="prettyInput"
                                        >
                                        </el-input>

                                    </el-tab-pane>
                                    <el-tab-pane  name="headers">
                                        <span slot="label">headers({{headers.length}})</span>
                                        <el-row v-for="header,index in headers" :key="index">
                                            <el-col :span="10">
                                                <el-input type="text" v-model="header.key"></el-input>
                                            </el-col>
                                            <el-col :span="10">
                                                <el-input type="text" v-model="header.value"></el-input>
                                            </el-col>
                                            <el-col :span="4">
                                                <el-button @click.preventDefault="headers.splice(index,1)">-</el-button>
                                            </el-col>

                                        </el-row>

                                        <el-button @click.preventDefault="headers.push({key:'',value:''})">+</el-button>
                                    </el-tab-pane>
                                </el-tabs>


                            </el-col>
                        </el-row>


                    </el-form>
                </div>
            </el-col>

        </el-row>


        <el-col v-if="jsonData">
            <vue-json-pretty id="vue-json-pretty"
                             :data="jsonData"
            >
            </vue-json-pretty>
        </el-col>


        <el-dialog title="文档预览" :visible.sync="previewBox">

            <div v-if="!docShowType">
                <el-row :gutter="20" style="text-align: left ;font-size: large">
                    <el-col :span="2" size="large">
                        <el-tag>{{ currentApi.method }}</el-tag>
                    </el-col>
                    <el-col :span="22"><span style="color:red">{host}</span> {{ currentApi.path }}</el-col>
                </el-row>

                <!--参数说明-->
                <el-table
                        :data="document.params" border stripe
                        style="width: 100%">


                    <el-table-column
                            label="参数"
                            width="180">
                        <template slot-scope="scope">

                            <span style="margin-left: 10px">{{ scope.row.key }}</span>
                        </template>
                    </el-table-column>

                    <el-table-column
                            label="说明"
                            width="180">
                        <template slot-scope="scope">
                            <el-select
                                    v-model="scope.row.statement"
                                    filterable
                                    allow-create>
                                <el-option
                                        v-for="item,index in scope.row.statement_options"
                                        :label="item.statement"
                                        :key="item.id"
                                        :value="item.statement">
                                </el-option>
                            </el-select>
                        </template>
                    </el-table-column>

                    <el-table-column
                            label="必填"
                            width="180">
                        <template slot-scope="scope">
                            <span style="cursor:pointer" @click="scope.row.required = !scope.row.required"> {{ scope.row.required ? "是":"否" }}</span>
                            <!--<span style="margin-left: 10px">{{ scope.row.required }}</span>-->
                        </template>
                    </el-table-column>
                    <el-table-column
                            label="类型"
                            width="180">
                        <template slot-scope="scope">

                            <span style="margin-left: 10px">{{ scope.row.type }}</span>
                        </template>
                    </el-table-column>


                    <el-table-column
                            label="备注"
                            width="180">
                        <template slot-scope="scope">
                            <el-input v-model="scope.row.remark" style="border:none" type="textarea"
                                      autosize></el-input>

                        </template>
                    </el-table-column>
                </el-table>

                <!--返回值-->
                <el-table
                        :data="document.response" border stripe
                        style="width: 100%">
                    <el-table-column
                            label="标识"
                            width="180">
                        <template slot-scope="scope">
                            {{ scope.row.key }}

                        </template>
                    </el-table-column>
                    <!--类型-->
                    <el-table-column
                            label="类型"
                            width="180">
                        <template slot-scope="scope">
                            {{ scope.row.type }}

                        </template>
                    </el-table-column>

                    <!--说明-->
                    <el-table-column
                            label="说明"
                            width="180">
                        <template slot-scope="scope">

                            <el-select
                                    v-model="scope.row.statement"
                                    filterable
                                    allow-create>
                                <el-option
                                        v-for="item in scope.row.statement_options"
                                        :label="item.statement"
                                        :value="item.statement"
                                        :key="item.id"
                                >
                                </el-option>
                            </el-select>


                        </template>
                    </el-table-column>

                </el-table>
            </div>

            <markdown
                    v-if="docShowType=='markdown'"
                    :content="markdownDoc"
            ></markdown>

            <!--注释-->
            <el-row>
                <el-col :span="8" :offset="16">

                    <el-select v-model="docShowType" placeholder="格式" @change="generateDoc">
                        <el-option
                                key="默认"
                                label="默认"
                                value="">
                        </el-option>
                        <el-option
                                key="markdown"
                                value="markdown">
                        </el-option>

                        <el-option
                                key="swagger"
                                value="swagger">
                        </el-option>
                    </el-select>

                    <el-button type="primary" @click="saveDoc">保存</el-button>
                </el-col>
            </el-row>
            <!--<el-row>-->

                <!--<el-col :span="4" :offset="20" style=" background: #d3dce6; on-hover:pointer">-->
                <!--<span v-clipboard:copy="commentaries"-->
                      <!--v-clipboard:success="onCopy" style="cursor:pointer"-->
                <!--&gt;复制到剪切板</span></el-col>-->
            <!--</el-row>-->
            <!--<el-row>-->
                <!--<el-col style="text-align: left;font-size: large">-->
                    <!--<el-input-->
                            <!--type="textarea"-->
                            <!--autosize-->
                            <!--placeholder="请输入内容"-->
                            <!--v-model="commentaries">-->
                    <!--</el-input>-->
                <!--</el-col>-->
            <!--</el-row>-->


        </el-dialog>


    </div>
</template>


<script type="text/ecmascript-6">
  import QueryParam from './QueryParam.vue'
  import { getUri } from '../../config/config.js'
  import Markdown from '../doc/Markdown.vue'
  import { mapState } from 'vuex'
  import VueJsonPretty from 'vue-json-pretty'
  var stringify = require('json-stable-stringify')
  export default {
    components: {
      QueryParam, Markdown, VueJsonPretty
    },
    computed: {},
    props: {
      api: Object
    },
    data () {
      return {
        methods: ['GET', 'POST', 'PUT', 'DELETE'],
        type: ['text', 'file', 'time', 'varchar'],
        previewBox: false,
        previewLoading: false,
        Form: {
          dataList: []
        },
        headers: [],
        tabSelected: 'form',
        bodyType:'form',
        currentApi: this.api,
        uriParamsArr: {},
        searchQuery: '',
        jsonData: false,
        request_loading: false,
        query_params_show: true,
        document: {},
        commentaries: '',
        api_name: '测试API',
        docShowType: '',
        markdownDoc: ''
      }
    },
    methods: {
      onCopy: function (e) {
        this.$message({
          message: '复制成功',
          type: 'success'
        })
      },
      bodyTypeChange (type){
        if(type !=='json'){
          return
        }
       let  exists = this.headers.filter((item)=>{
          return item.key == 'Content-Type' && item.value =='application/json'
        })
        if(exists.length ===0){
         this.headers.push({key:'Content-Type',value:'application/json'})
        }
      },
      changeHost (name) {
        let host = this.currentApi.hosts.filter((v)=>{
          return v.name === name
        })
        if(host.length === 1)
        {
          this.headers = host[0].headers
        }
      },
      submitForm (event) {
        this.request_loading = true
        this.$refs[event].validate((valid) => {
          if (valid) {
            let request_headers = this.headerToStr()
            if (this.bodyType == 'form') {
              var data = new FormData()
              let form = this.Form
              for (let index in this.currentApi.params) {
                let obj = this.currentApi.params[index]
                if (obj.type === 'file') {
                  for (let i = 0; i < obj.value.length; i++) {
                    data.append(obj.name, obj.value[i])
                  }
                }
                else {
                  data.append(obj.name, obj.value)
                }
              }
              data.append('request_uri', this.currentApi.path)
              data.append('request_host', this.currentApi.host)
              data.append('request_method', this.currentApi.method)
              data.append('request_headers', request_headers)
            }
            else {
                var data = {
               request_uri : this.currentApi.path,
              request_host : this.currentApi.host,
               request_method : this.currentApi.method,
               request_headers : request_headers,
              json_input : this.currentApi.json_input }
            }
            let option = {headers:{'Content-Type': 'application/json'}}
            this.axios.post('client-api/client/request', data,option).then((response) => {
              this.request_loading = false
              this.jsonData = response.data
            }, (response) => {
              this.$message.error('请求出错')
            })
          }
          else {
            return false
          }
        })
      },
      resetForm (formName) {
        this.$refs[formName].resetFields()
      },
      removeParam (item) {
        let index = this.currentApi.params.indexOf(item)
        if (index !== -1) {
          this.currentApi.params.splice(index, 1)
        }
      },
      getFile (event, item) {
        let index = this.currentApi.params.indexOf(item)
        this.currentApi.params[index].v = event.target.files
      },
      setTime (event) {
        let index = this.currentApi.params.indexOf(event)
        this.currentApi.params[index].v = event.v.toJSON()
      },
      clearValue (item) {
        let index = this.currentApi.params.indexOf(item)
        this.currentApi.params[index].v = ''
      },
      checkUriParams () {
        let uri = this.currentApi.path
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
      addParam (type) {
        this.currentApi.params.push({
          value: '',
          name: '',
          is_use: true,
          required: true,
          api_id: this.currentApi.id,
          type: type,
        })
      },
      changeParam (val) {
        this.uriParamsArr[val.index][val.attr] = val.value
        let uriParamsObjArr = this.uriParamsArr
        let paramArr = []
        for (let index in uriParamsObjArr) {
          paramArr.push(uriParamsObjArr[index].key + '=' + uriParamsObjArr[index].value)
        }
        let newUriParam = paramArr.join('&')
        let start = this.currentApi.path.indexOf('?')
        let str = this.currentApi.path.substr(0, start)
        let newUri = str + '?' + newUriParam
        this.currentApi.path = newUri
      },
      createDocument () {
        this.checkUriParams()
        this.previewBox = true
        try {
          var jsonParse = JSON.parse(this.currentApi.json_input)
        }
        catch (e) {
        }
        let data = {
          uri: this.currentApi.path,
          method: this.currentApi.method,
          queryParam: this.uriParamsArr,
          formParam: this.currentApi.params,
          response: this.jsonData,
          jsonParam: jsonParse
        }
        let uri = getUri('doc', 'generate')
        this.previewLoading = true
        this.axios.post(uri, data).then((response) => {
          this.document = this.getApiData(response)
          this.commentaries = "/**\n *" + this.currentApi.name + "\n"
          this.document.params.forEach((param) => {
            this.commentaries += ' *@param ' + param.type + ' ' + param.key + ' ' + param.statement + '\n'
          })
          this.commentaries += '*/'
          this.previewLoading = false
        }, (response) => {
          this.$message.error('生成出错')
        })
      },
      saveDoc(){
        let uri = getUri('doc', 'resource')
        let data = this.document
        this.axios.post(uri, data).then((response) => {
          this.$message.success(response.data.data)
          this.getApiData(response)
        }, (response) => {
          this.$message.error('保存失败！')
        })
      },
      generateDoc(value){
        if (value == 'markdown') {
          this.markdownDoc = this.generateMarkdown()
        }
      },
      generateMarkdown () {
        let text = '### ' + this.currentApi.name + '\n\n'
          + '  [**' + this.currentApi.method + '**] {host}' + this.currentApi.path + '\n\n'
          + '**请求参数说明**\n\n'
          + '|参数|说明|类型|必填|备注| \n |---|---|---|---|---| \n'
        this.document.params.forEach((param) => {
          text += '|' + param.key + '|' + param.statement + '|' + param.type + '|' + param.required + '|\n'
        })
        text += '\n **返回参数说明**\n\n'
          + '|参数|说明|类型|备注| \n |---|---|---|---| \n'
        if (this.document.response) {
          this.document.response.forEach((param) => {
            text += '|' + param.key + '|' + param.statement + '|' + param.type + '|\n'
          })
        }
        text += '\n``` \n'
        text += stringify(this.jsonData, {space: ' '})
        text += '\n``` \n'
        return text
      },
      saveApi(){
        let api = this.currentApi
        let uri = getUri('api', 'resource') + api.id
        this.axios.put(uri, api).then((response) => {
          this.getApiData(response, '保存成功！')
        }, (response) => {
          this.$message.error('保存失败')
        })

        let host = this.currentApi.hosts.filter((v)=>{
          return v.name === this.currentApi.host
        })
        if(host.length ===1) {
          let h = host.pop()
          h.headers = JSON.stringify(this.headers)
          let uri = getUri('host','resource') + h.id
          this.axios.patch(uri,h).then((response)=>{
            this.getApiData(response)
          },(response)=>{

          })
        }
      },
      headerToStr () {
        let headers = {}
        this.headers.forEach((v)=>{
          headers[v.key] = v.value
        })
        return  JSON.stringify(headers)
      },
      prettyInput(){
        let jsonObj
        if(this.currentApi.json_input.trim() === ''){
          return
        }
        try {
           jsonObj = JSON.parse(this.currentApi.json_input)
        }
        catch (e) {
          this.$message.error('json字符串格式有误')
          return
        }
        let jsonPretty = stringify(jsonObj, {space: ' '})
        if(jsonPretty){
          this.currentApi.json_input =jsonPretty
        }else{
          this.currentApi.json_input = ''
        }
      }
    },
    created(){
      if (!this.currentApi.method) {
        this.currentApi.method = 'GET'
      }
      if (this.currentApi.json_input) {
        this.tabSelected = 'json'
        this.bodyType = 'json'
      }
      if(this.currentApi.method === 'GET'){
        this.tabSelected = 'headers'
      }
      this.changeHost(this.currentApi.host)
    }
  }
</script>
<style scoped>
    #vue-json-pretty {
        text-align: left
    }

    .host-select {
        width: 210px;
    }
</style>