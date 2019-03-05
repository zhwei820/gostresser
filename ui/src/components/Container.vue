<style scoped>
  .layout {
    max-width: 1000px;
    margin: 0 auto;
  }
</style>

<template>
  <div class="layout">



      <div style="display: inline-block">
        <Menu mode="horizontal" theme="light" width="auto" :active-name="activeMenu"
              @on-select="onSelectMenuChange">

          <MenuItem name="service" to="/service">GoStresser Table</MenuItem>
          <MenuItem name="authservers" to="/auths">GoStresser Chart</MenuItem>
          <MenuItem name="health" to="/health">Health Error</MenuItem>

        </Menu>
      </div>

    <div class="layout-nav" style="float: right; margin-right: 50px; margin-top: 25px">
      <Dropdown v-if="user.access_token">

        <a href="javascript:void(0)">
          <p>ADMIN</p>
          <Icon type="ios-arrow-down"></Icon>
        </a>
        <DropdownMenu slot="list">
          <DropdownItem>
            <a href="/logout">Login</a>
          </DropdownItem>

        </DropdownMenu>
      </Dropdown>

      <a v-else class="gb_header_user_login" href="/login">登录系统</a>

    </div>

    <Layout >
      <Content :style="{padding: '0 16px 16px'}">
          <router-view/>
      </Content>
    </Layout>
  </div>
</template>

<script lang="js">
  import {USER_INFO} from '@/config'
  import storage from '@/utils/storage'

  const ACTIVEMENU = 'ACTIVEMENU'

  export default {
    name: "Container",
    data() {
      return {
        formInline: {
          username: '',
          password: ''
        },

      }
    },
    computed: {
      user() {
        const userInfo = storage.getItem(USER_INFO)
        console.log(userInfo)
        return userInfo
      },
      activeMenu() {
        return storage.getItem(ACTIVEMENU)
      },
    },
    methods: {
      onSelectMenuChange(val) {
        storage.setItem(ACTIVEMENU, val)
      }
    },
  }

</script>

