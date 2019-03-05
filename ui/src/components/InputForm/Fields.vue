<template>
  <FormItem
          :label="schema.title"
          :prop="schema.key"
          :error='errorStr'
  >

    <p v-html="currentValue" v-if="schema.type === 'html'"></p>

    <Input
            v-else-if="schema.type === 'input'"
            v-model="currentValue"
            v-bind="schema.props"
            @on-change="handleChange">
    </Input>

    <i-switch
            v-else-if="schema.type === 'switch'"
            v-model="currentValue"
            v-bind="schema.props"
            @on-change="handleChange">
    </i-switch>

    <Input
            v-else-if="schema.type === 'textarea'"
            v-model="currentValue"
            type="textarea"
            v-bind="schema.props"
            :rows="4"
            @on-change="handleChange">
    </Input>

    <DatePicker
            v-else-if="schema.type === 'date'"
            v-model="currentValue"
            v-bind="schema.props"
            style="width:100%;"
            type="date"
            @on-change="handleChange">
    </DatePicker>

    <DatePicker
            v-else-if="schema.type === 'datetimerange'"
            v-model="currentValue"
            v-bind="schema.props"
            style="width:100%;"
            type="datetimerange"
            @on-change="handleChange">
    </DatePicker>

    <Uploader
            v-else-if="schema.type === 'upload'"
            :value.sync="currentValue"
            v-bind="schema.props"
            style="width:100%;"
            @on-change="handleChange">
    </Uploader>

    <!--<V2Editor-->
            <!--v-else-if="schema.type === 'editor'"-->
            <!--:value.sync="currentValue"-->
            <!--v-bind="schema.props"-->
            <!--style="width:100%;"-->
            <!--@on-change="handleChange">-->
    <!--</V2Editor>-->

    <RadioGroup v-model="currentValue"
                @on-change="handleChange"
                v-else-if="schema.type === 'radio'">
      <Radio
              v-for="option in schema.options"
              :label="option.value"
              :key="option.value"
      >
        <span>{{ option.text }}</span>
      </Radio>
    </RadioGroup>

    <Select
            v-else-if="schema.type === 'select'"
            v-model="currentValue"
            v-bind="schema.props"
            :multiple="schema.multiple"
            :filterable="schema.filterable"
            :clearable="schema.clearable"
            style="width:100%"
            placeholder="请选择"
            @on-change="handleChange"
    >
      <Option
              v-for="option in schema.options"
              :key="option.value"
              :value="option.value"
      >{{ option.text }}
      </Option>
    </Select>

  </FormItem>
</template>

<script>

  // import Uploader from './components/Uploader.vue'
  // import V2Editor from './components/V2Editor.vue'

  export default {
    name: 'normalField',

    data() {
      return {
        currentValue: '',
      }
    },

    watch: {
      value: {
        immediate: true,
        handler(nextVal) {
          // debugger
          this.currentValue = nextVal

        },
      },
    },

    props: {
      value: null,
      error: null,
      schema: Object,
    },

    computed: {
      errorStr() {
        return JSON.stringify(this.error) || ''
      },
    },

    methods: {
      handleChange() {
        this.$emit('update:value', this.currentValue)
      },
    },

    components: {
      // Uploader,
      // V2Editor,
    },
  }
</script>
