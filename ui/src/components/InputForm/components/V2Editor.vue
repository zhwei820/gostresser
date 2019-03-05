<template>
  <div id="app">
    <Button id="btn" type='primary' @click="preview">预览</Button>

    <vue-editor id="editor"
                useCustomImageHandler
                @imageAdded="handleImageAdded"
                :editorOptions="editorSettings"
                :editorToolbar="customToolbar"
                v-model="htmlForEditor"
                style="width: 100%">
    </vue-editor>
  </div>
</template>

<script>
  import Quill from 'quill'
  import { VueEditor } from 'vue2-editor'
  import {uploadImg} from '@/apis/misc'
  import ImageResize from 'quill-image-resize-module'
  Quill.register('modules/imageResize', ImageResize)
  var Size = Quill.import('attributors/style/size');
  Quill.register(Size, true);

  const sizeList = []
  for( let ii = 12; ii <= 32; ii += 2)
  {
    sizeList.push(`${ii}px`)
  }
  Size.whitelist = sizeList;
  export default {
    components: {
      VueEditor,
    },
    props: {
      value: {
        type: [String],
        default: () => (''),
      },  // url 对象列表
    },
    data() {
      return {
        htmlForEditor: this.value,
        editorSettings: {
          modules: {
            // imageDrop: true,
            imageResize: {},
          },
        },
        customToolbar: [
          [{ size: sizeList }],

          ['bold', 'italic', 'underline', 'strike'],
          [{align: ''}, {align: 'center'}, {align: 'right'}, {align: 'justify'}],
          ['blockquote', 'code-block'],
          [{ list: 'ordered'}, { list: 'bullet' }, { list: 'check' }],
          [{ script: 'sub'}, { script: 'super' }],
          [{ indent: '-1'}, { indent: '+1' }],
          [{ color: [] }, { background: [] }],
          ['link', 'image', 'video'],
          ['clean'],
        ],
      }
    },
    watch: {
      htmlForEditor(nextVal) {
        this.$emit('update:value', nextVal)
        this.$emit('on-change')
      },
      value() {
        this.htmlForEditor = this.value
      },

    },

    methods: {
      preview() {
        const content = `<div class="ql-editor">${this.htmlForEditor}</div>`
        this.$Modal.confirm({
          title: '预览',
          content: content,
          width: 1200,
          // closable: true,
        })
      },
      handleImageAdded: async (file, Editor, cursorLocation, resetUploader) => {
        // An example of using FormData
        // NOTE: Your key could be different such as:
        // formData.append('file', file)

        const formData = new FormData()
        formData.append('upload_file', file)

        try {
          const res = await uploadImg(formData)

          const url = res.preview_url // Get url from response
          Editor.insertEmbed(cursorLocation, 'image', url)
          resetUploader()
        } catch (e) {
          console.log(e)
        }

      },
      setEditorContent() {
        this.htmlForEditor = '<h1>Html For Editor</h1>'
      },
    },
  }
</script>

<style lang="scss" scoped>
  .quillWrapper {
    height: 550px;
    width: 750px;
  }
  #app {
    height: 680px;
  }
  #btn {
    margin-right: 60px;
    float: right;
    margin-top: 10px;
  }
</style>
