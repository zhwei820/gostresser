<template>
  <div :class="$style.not_found">
    <h1 :class="$style.title">404</h1>
    <h1 :class="[$style.title, $style._sub]">No Found</h1>
  </div>
</template>

<script lang="ts">
export default {
  data() {
    return {
      style: { top: 0, left: 0 },
    }
  },

  methods: {
  },
}
</script>


<style lang="scss" module>
.not_found {
  position: absolute;
  top: 0;
  right: 0ex;
  bottom: 0px;
  left: 0px;
  margin: auto;
  margin-top: 15vw;
  width: 60vw;
}

.title {
  text-align: center;
  font-size: 56px;
  color: #0064ec;

  &._sub {
    font-size: 40px;
  }
}
</style>
