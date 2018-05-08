<template>
    <div id="markdown">
        <mavon-editor :ishljs = "true" :value="c" @change="changed"></mavon-editor>
    </div>
</template>

<script  type="text/ecmascript-6">
    import { mavonEditor } from 'mavon-editor'
    import 'mavon-editor/dist/css/index.css'
    export default {
        name: 'Markdown',
        data () {
            return {
              c: this.content
            }
        },
        props:{
            content: String
        },
      methods:{
          getContent () {
            return this.c
          },
        changed (content) {
            this.c = content
        }
      },
        components:{mavonEditor}
    }
</script>