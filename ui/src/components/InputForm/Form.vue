<template>
  <div class="manual-input-form">
    <Form :model="form" ref="theForm" label-position="right" :label-width='120' :rules="rules">
      <template v-for="scm in formList">
        <normal-field
                :key="scm.key"
                :type="scm.type"
                :value.sync="form[scm.key]"
                :schema="scm"
                :error="errors[scm.key]"
        ></normal-field>
      </template>
      <FormItem v-if="showSubmit">
        <Button type="primary" @click="handleSubmit()">Submit</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
  import Fields from './Fields'

  const noneType = [null, undefined]

  export default {
    name: 'InputForm',

    data() {
      return {
        formInit: {},
        form: {},
      }
    },

    props: {
      formList: {
        type: [Array, Object],
        default: () => ([]),
      },
      errors: {
        type: Object,
        default: () => ({}),
      },
      showSubmit: {
        type: Boolean,
        default: () => (true),
      },
    },

    watch: {
      formList: {
        immediate: true,
        handler(nextSchema) {
          const formInit = {}
          nextSchema.forEach((ns) => {
            formInit[ns.key] = !noneType.includes(ns.defaultValue) ? ns.defaultValue : ''

          })
          this.formInit = formInit
          this.form = {...formInit}
        },
      },

      form: {
        deep: true,
        handler(nextForm) {
          const _nextForm = {}
          Object.keys(nextForm).forEach((v) => {
            const tmp = v.split('.')
            if (tmp.length > 1) {
              if (tmp.length == 2) {
                if (_nextForm[tmp[0]] == undefined) {
                  _nextForm[tmp[0]] = {}
                }
                  _nextForm[tmp[0]][tmp[1]] = nextForm[v]
              } else if (tmp.length == 3) {
                if (_nextForm[tmp[0]] == undefined) {
                  _nextForm[tmp[0]] = {}
                }
                if (_nextForm[tmp[0]][tmp[1]] == undefined) {
                  _nextForm[tmp[0]][tmp[1]] = {}
                }
                _nextForm[tmp[0]][tmp[1]][tmp[2]] = nextForm[v]
              }
            } else {
              _nextForm[v] = nextForm[v]
            }
          })
          this.$emit('change', {..._nextForm})
        },
      },
    },

    computed: {
      rules() {
        const rules = {}
        this.formList.forEach((item) => {
          if (item.rule !== undefined) {
            rules[item.key] = item.rule
          }
        })
        return rules
      },
    },

    components: {
      'normal-field': Fields,
    },
    methods: {
      handleSubmit() {
        this.$refs.theForm.validate((valid) => {
          this.$emit('submit', this.getForm(), valid)
        })
      },
      // 获取整个 form
      getForm() {
        return {
          ...this.form,
        }
      },
    },
  }
</script>

<style lang="scss">
  .manual-input-form {
    position: relative;
    padding: 0 0 20px 20px;
    width: 80%;
  }
</style>
