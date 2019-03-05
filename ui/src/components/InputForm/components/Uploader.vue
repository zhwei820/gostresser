<template>
  <div>
    <div class="demo-upload-list" v-for="item in uploadList">
      <img :src="item.url">
      <div class="demo-upload-list-cover">
        <Icon type="ios-eye-outline" @click.native="handleView(item.url)"></Icon>
        <Icon type="ios-trash-outline" @click.native="handleRemove(item)"></Icon>
      </div>
    </div>
    <Upload
            ref="upload"
            :show-upload-list="false"
            :default-file-list="defaultList"
            :format="['jpg','jpeg','png']"
            :max-size="maxFileSize"
            :on-format-error="handleFormatError"
            :on-exceeded-size="handleMaxSize"
            :before-upload="handleBeforeUpload"
            multiple
            type="drag"
            action="//jsonplaceholder.typicode.com/posts/"
            style="display: inline-block;width:58px;">
      <div style="width: 58px;height:58px;line-height: 58px;">
        <Icon type="ios-camera" size="20"></Icon>
      </div>
    </Upload>
    <Modal title="View Image" v-model="visible">
      <img :src="imageUrl" v-if="visible" style="width: 100%">
    </Modal>
  </div>
</template>
<script>
  import {uploadImg} from '@/apis/misc'

  export default {
    props: {
      value: {
        type: [Array, String],
        default: () => ([]),
      },  // url 对象列表
      maxNum: {
        type: [Number],
        default: () => (1),
      },  // 图片最大数量
      maxFileSize: {
        type: [Number],
        default: () => (100 * 1024),
      } ,  // 图片最大文件大小
    },
    data() {
      const imgList = Array.isArray(this.value) ? this.value : (this.value ? [{url: this.value}] : [])
      return {
        defaultList: imgList,
        imageUrl: '',
        visible: false,
        uploadList: imgList,
      }
    },
    watch: {
      value() {
        const imgList = Array.isArray(this.value) ? this.value : (this.value ? [{url: this.value}] : [])
        this.defaultList = imgList
        this.uploadList = imgList
      },
    },
    methods: {
      handleView(url) {
        this.imageUrl = url
        this.visible = true
      },
      handleRemove(file) {
        this.uploadList.splice(this.uploadList.indexOf(file), 1)
        this.handleChanage()
      },
      handleFormatError(file) {
        this.$Notice.warning({
          title: 'The file format is incorrect',
          desc: 'File format of ' + file.name + ' is incorrect, please select jpg or png.',
        })
      },
      handleMaxSize(file) {
        this.$Notice.warning({
          title: 'Exceeding file size limit',
          desc: 'File  ' + file.name + ' is too large, no more than 2M.',
        })
      },
      handleBeforeUpload(file) {
        const check = this.uploadList.length < this.maxNum
        if (!check) {
          this.$Notice.warning({
            title: 'Up to ' + this.maxNum + ' pictures can be uploaded.',
          })
          return false
        }
        this.doUpload(file)
        return false
      },
      async doUpload(file) {
        const formData = new FormData()
        formData.append('upload_file', file)

        const res = await uploadImg(formData)
        this.uploadList.push({
          url: res.preview_url,
        })
        this.handleChanage()
      },
      handleChanage() {
        if (this.maxNum === 1) {
          this.$emit('update:value', this.uploadList.length > 0 ? this.uploadList[0].url : '')
          this.$emit('on-change')
        } else {
          let data = []
          this.uploadList.forEach((item) => {
            data.push({url: item.url})
          })
          this.$emit('update:value', data)
          this.$emit('on-change')
        }

      }
    },
  }
</script>
<style>
  .demo-upload-list {
    display: inline-block;
    width: 60px;
    height: 60px;
    text-align: center;
    line-height: 60px;
    border: 1px solid transparent;
    border-radius: 4px;
    overflow: hidden;
    background: #fff;
    position: relative;
    box-shadow: 0 1px 1px rgba(0, 0, 0, .2);
    margin-right: 4px;
  }

  .demo-upload-list img {
    width: 100%;
    height: 100%;
  }

  .demo-upload-list-cover {
    display: none;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: rgba(0, 0, 0, .6);
  }

  .demo-upload-list:hover .demo-upload-list-cover {
    display: block;
  }

  .demo-upload-list-cover i {
    color: #fff;
    font-size: 20px;
    cursor: pointer;
    margin: 0 2px;
  }
</style>
