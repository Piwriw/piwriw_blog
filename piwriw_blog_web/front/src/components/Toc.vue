<template>
 <div>
   <VMdPreview
     ref="previewRef"
     :text="data.content"
     lg:mx-20
   />
   <Catalogue v-if="!loading" :preview="previewRef" />
 </div>
</template>

<script>
import VMdPreview  from '@kangc/v-md-editor/lib/preview'
import '@kangc/v-md-editor/lib/style/base-editor.css'
import '@kangc/v-md-editor/lib/theme/style/vuepress.css'
import Catalogue from './Catalogue.vue'

export default {
  name: 'Toc',
  props:['content'],
  watch:{
    content(){
    this.data.content=String(this.content)
    }
  },
  components:{VMdPreview},
  data(){
    return {
      data:{
        content:'',
      }
    }
  }
}

</script>

<style scoped>

</style>
