<template>
  <div class="login-container">
    <Row>
      <Col span="6" offset="9">
        <Card :bordered="true" style="margin-top: 40%">
          <p slot="title">GoStresser Admin</p>
          <Form ref="formInline" :model="formInline" :rules="ruleInline" inline>
            <FormItem prop="username">
              <Input type="text" v-model="formInline.username" placeholder="Username">
                <Icon type="ios-person-outline" slot="prepend"></Icon>
              </Input>
            </FormItem>
            <FormItem prop="password">
              <Input type="password" v-model="formInline.password" placeholder="Password">
                <Icon type="ios-lock-outline" slot="prepend"></Icon>
              </Input>
            </FormItem>
            <FormItem>
              <Button type="primary" @click="handleSubmit('formInline')">Login</Button>
            </FormItem>
          </Form>
        </Card>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {Login} from '@/apis/auth'
  import * as config from '@/config'
  import router from '@/router'

  export default {
    name: "Login",
    data() {
      return {
        formInline: {
          username: '',
          password: ''
        },
        ruleInline: {
          username: [
            {required: true, message: 'Please fill in the username name', trigger: 'blur'}
          ],
          password: [
            {required: true, message: 'Please fill in the password.', trigger: 'blur'},
            {type: 'string', min: 5, message: 'The password length cannot be less than 6 bits', trigger: 'blur'}
          ]
        }
      }
    },
    methods: {
      handleSubmit(name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const res = await Login(this.formInline)

            const {access_token} = res
            localStorage.setItem(config.TOKEN_NAME, JSON.stringify(access_token || ''))
            localStorage.setItem(config.USER_INFO, JSON.stringify(res))
            router.push({ name: 'serviceList' })

          } else {
            this.$Message.error('Fail!');
          }
        })
      }
    }
  }
</script>

<style scoped>
  .login-container {
    min-height: 100%;
    width: 100%;
    background-color: #2d3a4b;
    overflow: hidden;
  }
</style>
