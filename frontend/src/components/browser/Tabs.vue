<template>
    <div id="tabs">
<el-row>
    <el-col :span="24">

        <el-tabs :value="activeApi" type="border-card" closable @tab-remove="removeTab">
            <el-tab-pane
                    v-for="(item, index) in apiList"
                    :key="item.name"
                    :label="item.name"
                    :name="String(index)"
            >
                <workbench :api="item"></workbench>
            </el-tab-pane>

        </el-tabs>
    </el-col>
</el-row>


    </div>
</template>

<script type="text/ecmascript-6">
    import Workbench from './Workbench.vue'
    import {mapState} from 'vuex'

    export default{
    name:'Tabs',
    data(){
        return {

        }
    },
     methods:{
         removeTab(index){
            this.$store.commit('removeTab',index)
             console.log(index)
         }
     },
    computed:{
        ...mapState({
            activeApi: state => state.activeApi
        }),
        apiList(){
            return this.$store.state.apiTabs
        },
        currentApi(){
            return this.$store.state.selectedApi
        }
    },
    components:{Workbench}
}

</script>