<template>
    <div>
        <h2 style="font-size: 20px; margin: 10px">{{ okMsg }}</h2>

        <Card v-for="(item, key) in failures" :key="key" style="margin: 10px">
            <p slot="title"><span style="color: red">API Name: </span> {{ key }}</p>
            <p><span style="color: red">Fail Reason</span>: {{ item }}</p>
        </Card>
    </div>
</template>

<script>

  import {fetchHealthCheck} from '@/apis/healthcheck'

  export default {
    name: "HealthCheckList",
    data () {
      return {
        okMsg: 'Everything is ok!',
        failures: {},
      }
    },
    computed: {
    },
    methods: {

      async fetchHealthCheck () {
        // debugger
        const response = await fetchHealthCheck()

        if (response.failures !== undefined) {
          this.okMsg = 'Some apis not reachable!'
          this.failures = response.failures
        }

      },
    },
    mounted () {
      this.fetchHealthCheck()
    }

  }
</script>

<style scoped>

</style>
