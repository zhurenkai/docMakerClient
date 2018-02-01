<template>
    <div id="topMenu">

    <el-row>
        <el-col :span="10" :offset="5"><span class="title">文档制造机</span></el-col>
        <el-col :span="2" :offset="5"> <i class="el-icon-setting title"></i></el-col>
        <el-col :span="2" ><span class="user">朱仁凯</span></el-col>
    </el-row>

    </div>
</template>

<script type="text/ecmascript-6">
    export default{

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