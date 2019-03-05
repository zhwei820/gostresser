<template>
    <div>
        <Card>
            <Form ref="formDynamic" :model="formDynamic" :label-width="100" style="width: 850px">

                <FormItem label="Name" required>
                    <Col span="12">
                        <Input v-model="formDynamic.name" placeholder="Enter your name"></Input>
                    </Col>
                </FormItem>


                <Card>
                    <p slot="title">token strategy</p>

                    <FormItem label="name" prop="token_strategy.name" required>
                        <Col span="12">
                            <Select v-model="formDynamic.token_strategy.name">
                                <Option v-for="item in token_strategy_names" :value="item" :key="item">{{ item }}
                                </Option>
                            </Select>
                        </Col>
                    </FormItem>


                    <FormItem label="settings"
                              v-for="(item, ii) in formDynamic.token_strategy.settings"
                              :key="ii"
                              :prop="'token_strategy.settings.' + ii + '.alg'"
                    >
                        <Row>
                            <Col span="20">
                                <Col span="6">
                                    <Select v-model="item.alg">
                                        <Option v-for="item in token_strategy_algs" :value="item" :key="item">
                                            {{ item }}
                                        </Option>
                                    </Select>
                                </Col>
                                <Col span="2" offset="1">
                                    <span style="float: right; margin-right: 10px"> key </span>
                                </Col>
                                <Col span="12">
                                    <Input type="text" v-model="item.key" placeholder="Enter something..."></Input>
                                </Col>
                                <Col span="2" offset="1">
                                    <Button type="warning" @click="handleRemovetokenSettings(ii)">del</Button>
                                </Col>
                            </Col>
                        </Row>
                    </FormItem>
                    <FormItem>
                        <Row>
                            <Col span="6">
                                <Button type="dashed" long @click="handleAddtokenSettings" icon="md-add">Add target
                                </Button>
                            </Col>
                        </Row>
                    </FormItem>

                </Card>


                <Proxy :value.sync="formDynamic.oauth_endpoints.authorize" :title="'authorize'"/>
                <Proxy :value.sync="formDynamic.oauth_endpoints.introspect" :title="'introspect'"/>
                <Proxy :value.sync="formDynamic.oauth_endpoints.token" :title="'token'"/>


                <FormItem>
                    <Button type="primary" @click="handleSubmit('formDynamic')">Submit</Button>
                </FormItem>
            </Form>
        </Card>
    </div>
</template>
<script lang="jsx">

  import {createAuth, putAuth, fetchAuthByName} from '@/apis/auths'
  import InputForm from '@/components/InputForm'
  import Proxy from '@/components/Proxy'

  export default {
    name: "AuthDetail",
    components: {
      // eslint-disable-next-line vue/no-unused-components
      InputForm: InputForm,
      Proxy: Proxy,
    },
    data () {
      return {
        token_strategy_names: [
          'jwt'
        ],
        token_strategy_algs: [
          'HS256',
          'HS384',
          'HS512',
          'RS256',
          'RS384',
          'RS512',
        ],
        copy: '',
        formDynamic: {
          "name": "",
          "oauth_endpoints": {
            "authorize": {
              "listen_path": "/auth/github/authorize",
              "methods": [
                "ALL"
              ],
              "preserve_host": false,
              "strip_path": false,
              "upstreams": {
                "balancing": "roundrobin",
                "targets": [
                  {
                    "target": "https://github.com/login/oauth/authorize"
                  }
                ]
              }
            },
            "token": {
              "listen_path": "/auth/github/token",
              "methods": [
                "GET",
                "POST"
              ],
              "preserve_host": false,
              "strip_path": false,
              "upstreams": {
                "balancing": "roundrobin",
                "targets": [
                  {
                    "target": "https://github.com/login/oauth/access_token"
                  }
                ]
              }
            },
            "introspect": {
              "listen_path": "/auth/github/introspect",
              "methods": [
                "GET"
              ],
              "preserve_host": false,
              "strip_path": false,
              "upstreams": {
                "balancing": "roundrobin",
                "targets": [
                  {
                    "target": "https://api.github.com/user"
                  }
                ]
              }
            }
          },
          "cors_meta": {
            "domains": [
              "*"
            ],
            "methods": [
              "GET",
              "POST",
              "PUT",
              "PATCH",
              "DELETE"
            ],
            "request_headers": [
              "Origin",
              "Authorization",
              "Content-Type"
            ],
            "exposed_headers": [
              "X-Debug-Token",
              "X-Debug-Token-Link"
            ],
            "enabled": true
          },
          "rate_limit": {
            "limit": "200-S",
            "enabled": true
          },
          "token_strategy": {
            "name": "jwt",
            "settings": [
              {"alg": "HS512", "key": "dsfdsfdsfsdf"}
            ]
          },
        },
        authData: {},
      }
    },
    computed: {
      name () {
        if (!this.copy) {
          const name = this.$route.params.name
          return name === 'null' ? null : name
        }
        return null
      }
    },
    methods: {

      handleSubmit (name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.$Message.success('Success!');
            this.createOrUpdateAuth(this.formDynamic)
          } else {
            this.$Message.error('Fail!');
          }
        })
      },
      handleAddtokenSettings () {
        this.formDynamic.token_strategy.settings.push({
          target: '',
          weight: 0,
        });
      },

      handleRemovetokenSettings (ii) {
        if (this.formDynamic.token_strategy.settings.length > 1) {
          this.formDynamic.token_strategy.settings.splice(ii, 1)
        } else {
          this.$Message.warning('Can not delete last target!')
        }
      },

      async createOrUpdateAuth (data) {
        this.errors = {}  //
                          // https://github.com/iview/iview/blob/master/src/components/form/form-item.vue#L81
        try {
          if (!this.name) {
            await createAuth(data)
          } else {
            await putAuth(this.name, data)
          }
          this.$router.push({name: 'authList'})
        } catch (e) {
          console.log(e)
          this.errors = e.data || {}
        }
      },

      async fetchAuthByName (name) {
        const response = await fetchAuthByName(name)
        // @ts-ignore
        this.formDynamic = response
        if (!this.formDynamic.plugins) {
          this.formDynamic.plugins = []
        }
        const urlParams = new URLSearchParams(window.location.search);
        this.copy = urlParams.get('copy');
        if (this.copy) {
          this.formDynamic.name = ''
        }
      },

    },
    created () {
      if (this.name) {
        this.fetchAuthByName(this.name)
      }
    }
  }

</script>

<style>
    .ivu-card .ivu-card-bordered {
        margin-bottom: 15px;
    }
</style>
