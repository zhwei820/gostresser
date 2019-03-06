<style scoped>
    .layout {
        max-width: 1000px;
        margin: 0 auto;
        margin-top: 50px;
    }
</style>

<template>
    <div class="layout">

        <Layout>
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

