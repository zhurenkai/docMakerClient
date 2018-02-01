<template>
<div id="queryParam">
  <el-row>
    <el-col :span="2">
  <el-checkbox v-model="required_pointer" @change="param_bind_change('required_pointer', 'required')">必填</el-checkbox>
    </el-col>
    <el-col :span="5">
  <el-input  placeholder="请输入内容" v-model="k_pointer" @change="param_bind_change('k_pointer', 'key')"></el-input>
    </el-col>

    <el-col :span="9">

    <el-input  placeholder="请输入内容" v-model="v_pointer" @change="param_bind_change('v_pointer', 'value')"></el-input>
    </el-col>

  </el-row>

</div>

</template>


<script type="text/ecmascript-6">
  export default {
    name: 'QueryParam',
    props: {
      k: String,
      v: String,
      required: Boolean,
      index:Number
    },
    data () {
      return {
        k_pointer: this.k,
        v_pointer: this.v,
        required_pointer: this.required,
        keyState_pointer: this.keyState
      }
    },
    watch: {
      k () {
        this.k_pointer = this.k
      },
      v () {
        this.v_pointer = this.v
      },
      required () {
        this.required_pointer = this.required
      }
    },
    methods: {
      param_bind_change: function (attr, prop) {
        this.$emit('param_bind_change')
        let change = {'index': this.index, 'attr': prop, 'value': this[attr]}
        this.$emit('selfUpdate', change)
      }
    }
  }
</script>