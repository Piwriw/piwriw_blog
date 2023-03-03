<template>
  <div>
    <mavon-editor
      v-model='data.content'
      ref='md'
      @change='change'
      style='min-height: 600px'
    />

    <!--    <button @click='submit'>提交</button>-->
  </div>
</template>

<script>
export default {
  name: 'edit',
  props: ['content'],
  watch:{
    content(val){
      this.data.content=this.content
    }
  },
  components: {
    mavonEditor
  }, data() {
    return {
      data: {
        content: this.content, // 输入的markdown
        html: ''    // 及时转的html
      }
    }
  },
  methods: {
    // 所有操作都会被解析重新渲染
    async change(value, render) {
      // render 为 markdown 解析后的结果[html]
      this.data.html = render
      this.$emit('article', this.data)
    }
    // 提交
    // submit() {
    //   console.log(this.content)
    //   console.log(this.html)
    // }
  },
  mounted() {
  }
}
// 导入组件 及 组件样式
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
</script>

<style scoped>

</style>
