<template>
    <div id="markdown">
        <mavon-editor :ishljs = "true" :value="content" ></mavon-editor>
    </div>
</template>

<script  type="text/ecmascript-6">
    import { mavonEditor } from 'mavon-editor'
    import 'mavon-editor/dist/css/index.css'
    export default {
        name: 'Markdown',
        data () {
            return {
            }
        },
        props:{
            content: String
        },
        components:{mavonEditor}
    }
</script>