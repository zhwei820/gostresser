<template>
  <div>
    <Card>
      <Form ref="formDynamic" :model="formDynamic" :label-width="100" style="width: 850px">

        <FormItem label="Name" prop="name" required>
          <Col span="12">
            <Input v-model="formDynamic.name" placeholder="Enter your name"></Input>
          </Col>
        </FormItem>

        <FormItem label="Active" prop="active">

          <i-switch size="large" v-model="formDynamic.active">
            <span slot="open">ON</span>
            <span slot="close">OFF</span>
          </i-switch>
        </FormItem>


        <Proxy :value.sync="formDynamic.proxy" :title="'Proxy'"/>

        <Card>
          <p slot="title">Plugins</p>

          <FormItem label="Plugin"
                    v-for="(item, ii) in formDynamic.plugins"
                    :key="ii">
            <Row>
              <Col span="10">
                <Button id="btn" type='primary'
                        @click="showPluginModal(ii)">
                  {{ item.name }}
                </Button>

              </Col>
              <Col span="2" offset="1">
                <Button type="warning" @click="handleRemovePlugin(ii)">del</Button>
              </Col>
            </Row>
          </FormItem>
          <FormItem>
            <Row>
              <Col span="6">
                <Button type="dashed" long @click="showSelectPluginModal" icon="md-add">Add Plugin</Button>
              </Col>
            </Row>
          </FormItem>
        </Card>


        <FormItem label="health_check" prop="health_check" required>
          <Col span="12">
            <Input v-model="formDynamic.health_check.url" placeholder="health_check url"></Input>
          </Col>
          <Col span="2" offset="1">
            <span style="float: right; margin-right: 10px"> timeout </span>
          </Col>
          <Col span="2">
            <Input type="text" v-model="formDynamic.health_check.timeout" placeholder="timeout"></Input>
          </Col>
        </FormItem>


        <FormItem>
          <Button type="primary" @click="handleSubmit('formDynamic')">Submit</Button>
        </FormItem>
      </Form>
    </Card>
  </div>
</template>
<script lang="jsx">

  import {createService, putService, fetchServiceByName} from '@/apis/service'
  import InputForm from '@/components/InputForm'
  import Proxy from '@/components/Proxy'
  import {getDefaultFormData, getFormList} from './helper.jsx'
  import {getSelectPluginForm} from './plugins/select'
  import {fetchAuth} from '@/apis/auths'

  export default {
    name: "ServiceDetail",
    components: {
      // eslint-disable-next-line vue/no-unused-components
      InputForm: InputForm,
      Proxy: Proxy,
    },
    data() {
      return {
        pluginName: 'body_limit',
        pluginIndex: -1,

        copy: '',
        OauthList: [],
        formDynamic: {
          "name": "",
          "active": true,
          "proxy": {
            "preserve_host": false,
            "listen_path": "",
            "upstreams": {
              "balancing": "roundrobin",
              "targets": [{
                "target": "",
                "weight": 0
              }]
            },
            "insecure_skip_verify": false,
            "strip_path": false,
            "append_path": false,
            "methods": ["GET"],
            "hosts": null,
            "forwarding_timeouts": {
              "dial_timeout": "0s",
              "response_header_timeout": "0s"
            }
          },
          "plugins": [],
          "health_check": {
            "url": "",
            "timeout": 3
          }
        },
        serviceData: {},
      }
    },
    computed: {
      name() {
        if (!this.copy) {
          const name = this.$route.params.name
          return name === 'null' ? null : name
        }
        return null
      }
    },
    methods: {

      handleSubmit(name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.$Message.success('Success!');
            this.createOrUpdateService(this.formDynamic)
          } else {
            this.$Message.error('Fail!');
          }
        })
      },


      async createOrUpdateService(data) {
        this.errors = {}  //
                          // https://github.com/iview/iview/blob/master/src/components/form/form-item.vue#L81
        try {
          if (!this.name) {
            await createService(data)
          } else {
            await putService(this.name, data)
          }
          this.$router.push({name: 'serviceList'})
        } catch (e) {
          console.log(e)
          this.errors = e.data || {}
        }
      },

      async fetchServiceByName(name) {
        const response = await fetchServiceByName(name)
        // @ts-ignore
        this.formDynamic = response
        // this.$set(this.formDynamic, response)
        // console.log(this.formDynamic)
        if (!this.formDynamic.plugins) {
          this.formDynamic.plugins = []
        }
        const urlParams = new URLSearchParams(window.location.search);
        this.copy = urlParams.get('copy');
        if (this.copy) {
          this.formDynamic.name = ''
        }

        const response2 = await fetchAuth(name)
        this.OauthList = []
        response2.forEach((val) => {
          this.OauthList.push(
            {value: val.name, text: val.name},
          )
        })

      },

      showPluginModal(ii) {
        this.pluginIndex = ii
        console.log(this.formDynamic.plugins)
        console.log(ii)
        this.formList = getFormList(this.formDynamic.plugins[ii], this.OauthList)
        this.$Modal.confirm({
          title: 'Edit Plugin',
          width: 600,
          // maskClosable: true,  这个参数现在还不支持...
          // closable: true,  这个参数有bug...
          onOk: async () => {
          },
          onCancel: () => {
            // this.$Message.error()
          },
          render: () => {
            return (
              <InputForm showSubmit={false} onChange={this.handleFormDataChange} formList={this.formList}
                         errors={{}}></InputForm>
            )
          },
        })
      },
      showSelectPluginModal() {
        this.formList = getSelectPluginForm(this.pluginName)
        this.$Modal.confirm({
          title: 'Select Plugin',
          width: 600,
          // maskClosable: true,  这个参数现在还不支持...
          // closable: true,  这个参数有bug...
          onOk: async () => {
            this.formDynamic.plugins.push(getDefaultFormData(this.pluginName))
            // this.showPluginModal(this.formDynamic.plugins.length - 1)
          },
          onCancel: () => {
            // this.$Message.error()
          },
          render: () => {
            return (
              <InputForm showSubmit={false} onChange={this.handleSelectPluginChange} formList={this.formList}
                         errors={{}}></InputForm>
            )
          },
        })
      },
      handleSelectPluginChange(newVal) {
        this.pluginName = newVal.pluginName
      },
      handleFormDataChange(newVal) {
        this.formDynamic.plugins[this.pluginIndex] = newVal
      },
      handleRemovePlugin(ii) {
        this.formDynamic.plugins.splice(ii, 1)
      },

    },
    created() {
      if (this.name) {
        this.fetchServiceByName(this.name)
      }
    }
  }

</script>

<style>
  .ivu-card .ivu-card-bordered {
    margin-bottom: 15px;
  }
</style>
